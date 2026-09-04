// SPDX-License-Identifier: Apache-2.0

//! Blanket `Authenticatable` trait for AuthContainer wrapping/unwrapping.
//!
//! Every `prost::Message + Default` automatically gets `AuthContainer`
//! construction and extraction via the blanket impl below. This codifies
//! the EVE [OBJECT-SIGNING](https://github.com/lf-edge/eve-api/blob/main/OBJECT-SIGNING.md)
//! protocol — any Rust consumer of the EVE API gets this for free.
//!
//! # Usage
//!
//! ```rust,ignore
//! use eve_api::crypto::Authenticatable;
//! use eve_api::register::ZRegisterMsg;
//!
//! // Wrap: message → signed AuthContainer
//! let msg = ZRegisterMsg { serial: "DEV-001".into(), ..Default::default() };
//! let container = msg.to_auth_container(&crypto)?;
//!
//! // Extract without verification
//! let msg = ZRegisterMsg::from_auth_container(&container)?;
//!
//! // Extract with signature verification
//! let msg = ZRegisterMsg::from_auth_container_verified(&container, cert_der, &crypto)?;
//! ```
//!
//! # Signing Convention
//!
//! The signature is computed over the raw `AuthBody.payload` bytes (the
//! serialized inner message), **NOT** the serialized `AuthBody` protobuf
//! wrapper. This matches the proven working behaviour of the existing
//! micro-eve client and EVE Go's `signAuthData()`.
//!
//! # `sender_cert` Field
//!
//! `to_auth_container()` leaves `sender_cert` empty. Only `register()`
//! needs the full cert — it mutates the container after construction:
//!
//! ```rust,ignore
//! let mut container = msg.to_auth_container(&crypto)?;
//! container.sender_cert = base64_encode(&crypto.signing_cert_pem()).into();
//! ```

use bytes::Bytes;
use prost::Message;

use crate::auth::{AuthBody, AuthContainer};
use crate::common::HashAlgorithm;

use super::{CryptoError, CryptoProvider};

/// Trait for fluent AuthContainer wrapping/unwrapping of protobuf messages.
///
/// Automatically implemented for every `prost::Message + Default` type via
/// blanket impl — no per-type boilerplate needed.
pub trait Authenticatable: Message + Default + Sized {
    /// Wrap this message in a signed `AuthContainer`.
    ///
    /// Signs with whatever key the provider holds. The "which key" question
    /// is answered by which `CryptoProvider` instance you pass — no
    /// `SigningIdentity` enum, no runtime branching.
    ///
    /// # Arguments
    ///
    /// * `crypto` — provides signing and certificate access
    ///
    /// # Returns
    ///
    /// A fully-assembled `AuthContainer` ready to be serialized and sent.
    /// The `sender_cert` field is left empty — `register()` fills it in
    /// for the registration endpoint specifically.
    fn to_auth_container<C: CryptoProvider>(
        &self,
        crypto: &C,
    ) -> Result<AuthContainer, CryptoError> {
        // 1. Serialize the inner message → payload bytes
        let payload = self.encode_to_vec();

        // 2. Sign the raw payload bytes (NOT the serialized AuthBody wrapper)
        let signature = crypto.sign(&payload)?;

        // 3. Wrap in AuthBody
        let auth_body = AuthBody {
            payload: Bytes::from(payload),
        };

        // 4. Assemble the AuthContainer
        Ok(AuthContainer {
            protected_payload: Some(auth_body),
            algo: HashAlgorithm::Sha25632bytes as i32,
            sender_cert_hash: Bytes::copy_from_slice(crypto.signing_cert_hash()),
            signature_hash: Bytes::from(signature),
            sender_cert: Bytes::new(), // register() fills this in
            cipher_context: None,
            cipher_data: None,
        })
    }

    /// Extract this message type from an `AuthContainer` without verification.
    ///
    /// Use this only when the transport layer (TLS) already guarantees
    /// authenticity, or when verification is handled separately.
    fn from_auth_container(container: &AuthContainer) -> Result<Self, CryptoError> {
        let payload =
            container
                .protected_payload
                .as_ref()
                .ok_or(CryptoError::AuthMissingField {
                    field: "protected_payload",
                })?;
        Self::decode(payload.payload.as_ref())
            .map_err(|e| CryptoError::parse("protobuf", format!("{e}")))
    }

    /// Extract this message type from a verified `AuthContainer`.
    ///
    /// Verifies the ECDSA signature over the payload using the public key
    /// from `sender_cert_der` before extracting the message.
    ///
    /// # Arguments
    ///
    /// * `container` — the received `AuthContainer`
    /// * `sender_cert_der` — DER-encoded X.509 certificate of the sender
    ///   (e.g., the controller's signing cert)
    /// * `crypto` — provides the `verify()` implementation
    fn from_auth_container_verified<C: CryptoProvider>(
        container: &AuthContainer,
        sender_cert_der: &[u8],
        crypto: &C,
    ) -> Result<Self, CryptoError> {
        let protected =
            container
                .protected_payload
                .as_ref()
                .ok_or(CryptoError::AuthMissingField {
                    field: "protected_payload",
                })?;

        let valid = crypto.verify(
            &protected.payload,
            &container.signature_hash,
            sender_cert_der,
        )?;

        if !valid {
            return Err(CryptoError::SignatureMismatch);
        }

        Self::decode(protected.payload.as_ref())
            .map_err(|e| CryptoError::parse("protobuf", format!("{e}")))
    }
}

/// Zero-cost blanket impl — every protobuf message gets this for free.
impl<T> Authenticatable for T where T: Message + Default {}

#[cfg(test)]
mod tests {
    use super::*;
    use crate::crypto::{CertStore, DeviceCertStore, Signer};
    use crate::register::ZRegisterMsg;
    use crate::uuid::UuidResponse;
    use std::fmt;

    // ── Mock crypto provider ───────────────────────────────────────────

    struct MockCrypto {
        cert_store: DeviceCertStore,
    }

    impl MockCrypto {
        fn new() -> Self {
            Self {
                cert_store: DeviceCertStore::new(vec![0xAA; 100], vec![0x11; 32]),
            }
        }
    }

    impl fmt::Debug for MockCrypto {
        fn fmt(&self, f: &mut fmt::Formatter<'_>) -> fmt::Result {
            f.debug_struct("MockCrypto").finish()
        }
    }

    impl Signer for MockCrypto {
        fn sign(&self, data: &[u8]) -> Result<Vec<u8>, CryptoError> {
            // Simple "signature": SHA-256-ish mock — just use data length
            Ok(vec![data.len() as u8; 64])
        }

        fn verify(
            &self,
            data: &[u8],
            signature: &[u8],
            _cert_der: &[u8],
        ) -> Result<bool, CryptoError> {
            // Verify: signature should be 64 bytes of (data.len() as u8)
            Ok(signature.len() == 64 && signature[0] == data.len() as u8)
        }
    }

    impl CertStore for MockCrypto {
        fn signing_cert_der(&self) -> &[u8] {
            self.cert_store.signing_cert_der()
        }
        fn signing_cert_hash(&self) -> &[u8] {
            self.cert_store.signing_cert_hash()
        }
        fn device_cert_der(&self) -> &[u8] {
            self.cert_store.device_cert_der()
        }
        fn device_cert_hash(&self) -> &[u8] {
            self.cert_store.device_cert_hash()
        }
    }

    // ── Tests ──────────────────────────────────────────────────────────

    #[test]
    fn test_to_auth_container_basic() {
        let crypto = MockCrypto::new();
        let msg = ZRegisterMsg {
            serial: "DEV-001".to_string(),
            ..Default::default()
        };

        let container = msg.to_auth_container(&crypto).unwrap();

        // Check structure
        assert!(container.protected_payload.is_some());
        assert_eq!(container.algo, HashAlgorithm::Sha25632bytes as i32);
        assert_eq!(container.sender_cert_hash.len(), 32);
        assert_eq!(container.signature_hash.len(), 64);
        assert!(container.sender_cert.is_empty()); // not filled by to_auth_container
    }

    #[test]
    fn test_roundtrip_without_verify() {
        let crypto = MockCrypto::new();
        let original = ZRegisterMsg {
            serial: "SN-12345".to_string(),
            soft_serial: "SOFT-67890".to_string(),
            ..Default::default()
        };

        let container = original.to_auth_container(&crypto).unwrap();
        let extracted = ZRegisterMsg::from_auth_container(&container).unwrap();

        assert_eq!(extracted.serial, "SN-12345");
        assert_eq!(extracted.soft_serial, "SOFT-67890");
    }

    #[test]
    fn test_roundtrip_with_verify() {
        let crypto = MockCrypto::new();
        let original = ZRegisterMsg {
            serial: "DEV-002".to_string(),
            ..Default::default()
        };

        let container = original.to_auth_container(&crypto).unwrap();
        let extracted =
            ZRegisterMsg::from_auth_container_verified(&container, b"sender-cert", &crypto)
                .unwrap();

        assert_eq!(extracted.serial, "DEV-002");
    }

    #[test]
    fn test_verify_fails_with_wrong_signature() {
        let crypto = MockCrypto::new();
        let msg = ZRegisterMsg {
            serial: "DEV-003".to_string(),
            ..Default::default()
        };

        let mut container = msg.to_auth_container(&crypto).unwrap();
        // Corrupt the signature
        container.signature_hash = Bytes::from(vec![0xFF; 64]);

        let result =
            ZRegisterMsg::from_auth_container_verified(&container, b"sender-cert", &crypto);

        assert!(result.is_err());
        let err = result.unwrap_err();
        assert!(matches!(err, CryptoError::SignatureMismatch));
    }

    #[test]
    fn test_from_auth_container_missing_payload() {
        let container = AuthContainer {
            protected_payload: None,
            ..Default::default()
        };

        let result = ZRegisterMsg::from_auth_container(&container);
        assert!(result.is_err());
        assert!(matches!(
            result.unwrap_err(),
            CryptoError::AuthMissingField {
                field: "protected_payload"
            }
        ));
    }

    #[test]
    fn test_from_auth_container_verified_missing_payload() {
        let crypto = MockCrypto::new();
        let container = AuthContainer {
            protected_payload: None,
            ..Default::default()
        };

        let result = ZRegisterMsg::from_auth_container_verified(&container, b"cert", &crypto);
        assert!(result.is_err());
        assert!(matches!(
            result.unwrap_err(),
            CryptoError::AuthMissingField {
                field: "protected_payload"
            }
        ));
    }

    #[test]
    fn test_sender_cert_hash_matches_provider() {
        let crypto = MockCrypto::new();
        let msg = ZRegisterMsg::default();

        let container = msg.to_auth_container(&crypto).unwrap();
        assert_eq!(
            container.sender_cert_hash.as_ref(),
            crypto.signing_cert_hash()
        );
    }

    #[test]
    fn test_blanket_impl_works_for_any_message() {
        // UuidResponse is a different message type — blanket impl should work
        let crypto = MockCrypto::new();
        let msg = UuidResponse {
            uuid: "test-uuid-1234".to_string(),
            ..Default::default()
        };

        let container = msg.to_auth_container(&crypto).unwrap();
        let extracted = UuidResponse::from_auth_container(&container).unwrap();
        assert_eq!(extracted.uuid, "test-uuid-1234");
    }

    #[test]
    fn test_sender_cert_empty_by_default() {
        let crypto = MockCrypto::new();
        let msg = ZRegisterMsg::default();

        let container = msg.to_auth_container(&crypto).unwrap();
        assert!(
            container.sender_cert.is_empty(),
            "sender_cert should be empty — register() fills it in"
        );
    }

    #[test]
    fn test_register_pattern_with_sender_cert() {
        // Simulate the register() pattern: to_auth_container + mutate sender_cert
        let crypto = MockCrypto::new();
        let msg = ZRegisterMsg {
            serial: "DEV-REG".to_string(),
            ..Default::default()
        };

        let mut container = msg.to_auth_container(&crypto).unwrap();

        // register() would do this:
        use base64::Engine;
        let pem = crypto.signing_cert_pem();
        let b64 = base64::engine::general_purpose::STANDARD.encode(&pem);
        container.sender_cert = Bytes::from(b64.into_bytes());

        assert!(!container.sender_cert.is_empty());

        // The message should still be extractable
        let extracted = ZRegisterMsg::from_auth_container(&container).unwrap();
        assert_eq!(extracted.serial, "DEV-REG");
    }
}
