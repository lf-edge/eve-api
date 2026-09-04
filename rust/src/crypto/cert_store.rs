// SPDX-License-Identifier: Apache-2.0

//! Certificate storage trait for the EVE OBJECT-SIGNING protocol.
//!
//! The [`CertStore`] trait provides access to the certificates needed by
//! [`Authenticatable`](super::Authenticatable) to populate `AuthContainer`
//! fields (`sender_cert_hash`, `sender_cert`) and by `EveClient` to include
//! the device certificate in `ZRegisterMsg.pem_cert`.
//!
//! # Design
//!
//! `CertStore` is intentionally **separate from [`Signer`](super::Signer)**:
//! crypto providers generate certs and sign data; certificate storage holds
//! them. Different `CertStore` implementations serve different operational
//! phases:
//!
//! - [`OnboardCertStore`] — onboarding phase: `signing_cert` = onboard cert,
//!   `device_cert` = device cert
//! - [`DeviceCertStore`] — operational phase: `signing_cert` = device cert,
//!   `device_cert` = device cert
//!
//! # Infallible Access
//!
//! All methods return references (`&[u8]`), not `Result`. Certificates are
//! loaded/generated during provider construction — once a `CertStore` is
//! built, cert access is infallible. Any I/O or TPM NV reads happen before
//! the store is constructed.
//!
//! # Concrete Implementations
//!
//! `CertStore` is a trait (not a struct) to allow platform-specific backends:
//! in-memory (tests), file-based PEM (embedded Linux), TPM NV-backed (device
//! cert persisted in NV RAM), PKCS#11/HSM, etc. The concrete
//! `OnboardCertStore` and `DeviceCertStore` structs below are the simplest
//! impls — the trait keeps the door open for any backend.

use base64::Engine;

/// Certificate storage for a specific operational phase.
///
/// During onboarding: `signing_cert` = onboard cert, `device_cert` = device cert.
/// During operation:  `signing_cert` = device cert,  `device_cert` = device cert.
///
/// # Thread Safety
///
/// Implementations must be `Send + Sync` so the provider can be shared
/// across async tasks.
pub trait CertStore: Send + Sync {
    /// The certificate whose private key is used for signing.
    ///
    /// This determines what goes into `AuthContainer.sender_cert_hash`.
    /// During onboarding this is the onboard cert; during operation this
    /// is the device cert.
    fn signing_cert_der(&self) -> &[u8];

    /// SHA-256 hash of the signing certificate (all 32 bytes).
    ///
    /// Used as `AuthContainer.sender_cert_hash` with
    /// `algo = SHA256_32BYTES`.
    fn signing_cert_hash(&self) -> &[u8];

    /// The signing certificate in PEM encoding.
    ///
    /// Default implementation PEM-wraps the DER bytes.
    fn signing_cert_pem(&self) -> Vec<u8> {
        pem_encode("CERTIFICATE", self.signing_cert_der())
    }

    /// The device certificate in DER encoding (always the device cert,
    /// regardless of phase).
    ///
    /// Needed for `ZRegisterMsg.pem_cert` during onboarding.
    fn device_cert_der(&self) -> &[u8];

    /// The device certificate in PEM encoding.
    ///
    /// Default implementation PEM-wraps the DER bytes.
    fn device_cert_pem(&self) -> Vec<u8> {
        pem_encode("CERTIFICATE", self.device_cert_der())
    }

    /// SHA-256 hash of the device certificate (all 32 bytes).
    fn device_cert_hash(&self) -> &[u8];
}

/// Encode DER bytes into PEM format with the given label.
///
/// Produces standard PEM with 64-character base64 lines.
pub fn pem_encode(label: &str, der: &[u8]) -> Vec<u8> {
    let b64 = base64::engine::general_purpose::STANDARD.encode(der);

    let mut pem = format!("-----BEGIN {label}-----\n");
    for chunk in b64.as_bytes().chunks(64) {
        // SAFETY: base64 output is always valid UTF-8
        pem.push_str(std::str::from_utf8(chunk).unwrap());
        pem.push('\n');
    }
    pem.push_str(&format!("-----END {label}-----\n"));
    pem.into_bytes()
}

// ── Concrete implementations ───────────────────────────────────────────

/// Onboarding phase: signs with onboard key, device cert is separate.
///
/// `signing_cert` returns the onboard certificate.
/// `device_cert` returns the device certificate.
#[derive(Debug, Clone)]
pub struct OnboardCertStore {
    onboard_cert_der: Vec<u8>,
    onboard_cert_hash: Vec<u8>,
    device_cert_der: Vec<u8>,
    device_cert_hash: Vec<u8>,
}

impl OnboardCertStore {
    /// Create a new `OnboardCertStore` from raw DER bytes and pre-computed
    /// SHA-256 hashes (32 bytes each).
    pub fn new(
        onboard_cert_der: Vec<u8>,
        onboard_cert_hash: Vec<u8>,
        device_cert_der: Vec<u8>,
        device_cert_hash: Vec<u8>,
    ) -> Self {
        Self {
            onboard_cert_der,
            onboard_cert_hash,
            device_cert_der,
            device_cert_hash,
        }
    }
}

impl CertStore for OnboardCertStore {
    fn signing_cert_der(&self) -> &[u8] {
        &self.onboard_cert_der
    }

    fn signing_cert_hash(&self) -> &[u8] {
        &self.onboard_cert_hash
    }

    fn device_cert_der(&self) -> &[u8] {
        &self.device_cert_der
    }

    fn device_cert_hash(&self) -> &[u8] {
        &self.device_cert_hash
    }
}

/// Operational phase: signs with device key, signing cert = device cert.
///
/// Both `signing_cert` and `device_cert` return the same device certificate.
#[derive(Debug, Clone)]
pub struct DeviceCertStore {
    device_cert_der: Vec<u8>,
    device_cert_hash: Vec<u8>,
}

impl DeviceCertStore {
    /// Create a new `DeviceCertStore` from raw DER bytes and pre-computed
    /// SHA-256 hash (32 bytes).
    pub fn new(device_cert_der: Vec<u8>, device_cert_hash: Vec<u8>) -> Self {
        Self {
            device_cert_der,
            device_cert_hash,
        }
    }
}

impl CertStore for DeviceCertStore {
    fn signing_cert_der(&self) -> &[u8] {
        &self.device_cert_der
    }

    fn signing_cert_hash(&self) -> &[u8] {
        &self.device_cert_hash
    }

    fn device_cert_der(&self) -> &[u8] {
        &self.device_cert_der
    }

    fn device_cert_hash(&self) -> &[u8] {
        &self.device_cert_hash
    }
}

#[cfg(test)]
mod tests {
    use super::*;

    fn fake_hash(prefix: u8) -> Vec<u8> {
        vec![prefix; 32]
    }

    fn fake_cert(id: u8) -> Vec<u8> {
        vec![id; 100]
    }

    #[test]
    fn test_onboard_cert_store_signing_is_onboard() {
        let store = OnboardCertStore::new(
            fake_cert(0xAA),
            fake_hash(0x11),
            fake_cert(0xBB),
            fake_hash(0x22),
        );

        // signing cert = onboard cert
        assert_eq!(store.signing_cert_der(), &fake_cert(0xAA));
        assert_eq!(store.signing_cert_hash(), &fake_hash(0x11));

        // device cert = device cert
        assert_eq!(store.device_cert_der(), &fake_cert(0xBB));
        assert_eq!(store.device_cert_hash(), &fake_hash(0x22));
    }

    #[test]
    fn test_device_cert_store_signing_is_device() {
        let store = DeviceCertStore::new(fake_cert(0xCC), fake_hash(0x33));

        // both signing and device cert are the same
        assert_eq!(store.signing_cert_der(), store.device_cert_der());
        assert_eq!(store.signing_cert_hash(), store.device_cert_hash());
        assert_eq!(store.signing_cert_der(), &fake_cert(0xCC));
        assert_eq!(store.signing_cert_hash(), &fake_hash(0x33));
    }

    #[test]
    fn test_pem_encode_roundtrip() {
        let der = b"hello world this is some fake DER data for testing purposes";
        let pem = pem_encode("CERTIFICATE", der);
        let pem_str = std::str::from_utf8(&pem).unwrap();

        assert!(pem_str.starts_with("-----BEGIN CERTIFICATE-----\n"));
        assert!(pem_str.ends_with("-----END CERTIFICATE-----\n"));

        // Decode the base64 content back
        let lines: Vec<&str> = pem_str.lines().collect();
        let b64: String = lines[1..lines.len() - 1].join("");
        let decoded = base64::engine::general_purpose::STANDARD
            .decode(&b64)
            .unwrap();
        assert_eq!(decoded, der);
    }
}
