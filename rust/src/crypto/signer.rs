// SPDX-License-Identifier: Apache-2.0

//! Core signing and verification trait for the EVE OBJECT-SIGNING protocol.
//!
//! The [`Signer`] trait is the minimal interface needed by
//! [`Authenticatable`](super::Authenticatable) and `EveClient` for day-to-day
//! `AuthContainer` construction and verification.
//!
//! # Implementations
//!
//! | Implementation     | Backend      | Crate            |
//! |--------------------|--------------|------------------|
//! | `SoftwareSigner`   | `ring`       | `crypto-software`|
//! | `TpmSigner`        | `tss-esapi`  | `tpm-provider`   |
//!
//! # Signature Format
//!
//! `sign()` returns a **fixed-size R‖S** signature (64 bytes for P-256),
//! suitable for `AuthContainer.signature_hash`. This is distinct from the
//! ASN.1 DER encoding used by TLS — the conversion happens at a lower layer
//! (e.g., `rustls_signer.rs` in `tpm-provider`).
//!
//! # Verification
//!
//! `verify()` is always performed in software using the public key extracted
//! from the provided DER certificate. There is no security benefit to using
//! the TPM for verification — it only involves public keys (no secrets).
//! This matches EVE Go's `verifyAuthSig()` in `controllerconn/authen.go`,
//! which uses Go's standard `ecdsa.Verify()` regardless of TPM presence.

use std::fmt::Debug;

use super::CryptoError;

/// Core signing and verification operations for the EVE OBJECT-SIGNING protocol.
///
/// Uses concrete [`CryptoError`] (not an associated error type) so that
/// [`Authenticatable`](super::Authenticatable) and `EveClient` can use `?`
/// without `From` bounds at every call site.
///
/// # Thread Safety
///
/// Implementations must be `Send + Sync` so the provider can be shared
/// across async tasks (e.g., held inside an `Arc` in the API client).
///
/// # Example
///
/// ```rust,ignore
/// use eve_api::crypto::{Signer, CryptoError};
///
/// fn sign_payload(signer: &impl Signer, data: &[u8]) -> Result<Vec<u8>, CryptoError> {
///     signer.sign(data)
/// }
/// ```
pub trait Signer: Send + Sync + Debug {
    /// Sign `data` with ECDSA-P256-SHA256.
    ///
    /// Returns a fixed-size R‖S signature (64 bytes for P-256) suitable
    /// for `AuthContainer.signature_hash`.
    ///
    /// The implementation hashes `data` with SHA-256 internally before
    /// signing — callers pass the raw payload bytes, not a pre-computed
    /// digest.
    fn sign(&self, data: &[u8]) -> Result<Vec<u8>, CryptoError>;

    /// Verify an ECDSA-P256-SHA256 signature using the public key from
    /// the given DER-encoded X.509 certificate.
    ///
    /// Returns `Ok(true)` if the signature is valid, `Ok(false)` if it is
    /// structurally well-formed but does not match, or `Err` if verification
    /// cannot be performed (e.g., unsupported key type, malformed cert).
    ///
    /// # Software-only verification
    ///
    /// This method is **always performed in software** — verification uses
    /// only the public key (extracted from the cert), so there is no security
    /// benefit to using the TPM. The TPM protects *private* keys (signing,
    /// ECDH, quote); public-key verification is pure math with no secrets
    /// involved.
    ///
    /// This matches EVE Go's `verifyAuthSig()` which uses Go's standard
    /// `ecdsa.Verify()` / `rsa.VerifyPKCS1v15()` regardless of whether the
    /// device has a TPM.
    fn verify(&self, data: &[u8], signature: &[u8], cert_der: &[u8]) -> Result<bool, CryptoError>;
}

#[cfg(test)]
mod tests {
    use super::*;

    /// A minimal mock signer for testing trait bounds and usage patterns.
    #[derive(Debug)]
    struct MockSigner;

    impl Signer for MockSigner {
        fn sign(&self, data: &[u8]) -> Result<Vec<u8>, CryptoError> {
            // Return data length as a fake "signature" for testing
            Ok(vec![data.len() as u8; 64])
        }

        fn verify(
            &self,
            _data: &[u8],
            signature: &[u8],
            _cert_der: &[u8],
        ) -> Result<bool, CryptoError> {
            // Accept any 64-byte signature
            Ok(signature.len() == 64)
        }
    }

    #[test]
    fn test_signer_generic_usage() {
        fn sign_and_verify(s: &impl Signer, data: &[u8]) -> Result<bool, CryptoError> {
            let sig = s.sign(data)?;
            s.verify(data, &sig, b"cert")
        }

        let signer = MockSigner;
        assert!(sign_and_verify(&signer, b"payload").unwrap());
    }
}
