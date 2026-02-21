// SPDX-License-Identifier: Apache-2.0

//! Protocol-level error types for cryptographic operations.
//!
//! [`CryptoError`] is the canonical error type used by the protocol traits
//! (`Signer`, `CertStore`, `Authenticatable`) defined in this crate. It
//! covers all failure modes relevant to the EVE OBJECT-SIGNING protocol
//! without depending on any specific crypto backend.
//!
//! Backend-specific errors (e.g., `ring::error::Unspecified`, `TpmError`)
//! should be mapped to `CryptoError` variants (typically `SigningFailed`
//! or `VerificationFailed`) in their respective crates via `From` impls.

use thiserror::Error;

/// Errors that can occur during cryptographic operations.
///
/// This enum is intentionally free of backend-specific variants — it lives
/// in `eve-api-rs` alongside the protocol traits and must not pull in
/// dependencies like `ring` or `tss-esapi`. Backend crates convert their
/// native errors into these variants.
#[derive(Error, Debug)]
pub enum CryptoError {
    // ── Key / Certificate loading ──────────────────────────────────────
    /// Failed to read a certificate or key file from disk.
    #[error("failed to read {kind} from {path}: {source}")]
    FileRead {
        kind: &'static str,
        path: String,
        source: std::io::Error,
    },

    /// The PEM or DER data could not be parsed.
    #[error("failed to parse {kind}: {reason}")]
    Parse { kind: &'static str, reason: String },

    /// The certificate or key uses an unsupported algorithm.
    ///
    /// The EVE API requires ECDSA with P-256 (prime256v1 / secp256r1).
    #[error("unsupported algorithm: {algorithm} (expected ECDSA P-256)")]
    UnsupportedAlgorithm { algorithm: String },

    /// A required certificate or key is missing.
    #[error("missing {kind}: {detail}")]
    Missing { kind: &'static str, detail: String },

    // ── Signing ────────────────────────────────────────────────────────
    /// The signing operation failed.
    ///
    /// Backend crates map their native errors to this variant, e.g.:
    /// - `ring::error::Unspecified` → `SigningFailed { reason: "ring: ..." }`
    /// - `TpmError` → `SigningFailed { reason: "tpm: ..." }`
    #[error("signing failed: {reason}")]
    SigningFailed { reason: String },

    // ── Verification ───────────────────────────────────────────────────
    /// The signature verification operation itself failed (not "invalid
    /// signature", but rather an operational failure like a malformed
    /// input).
    #[error("verification failed: {reason}")]
    VerificationFailed { reason: String },

    /// The signature is structurally valid but does not match.
    #[error("signature mismatch")]
    SignatureMismatch,

    // ── Hashing ────────────────────────────────────────────────────────
    /// Hash computation failed.
    #[error("hash computation failed: {reason}")]
    HashFailed { reason: String },

    // ── Certificate validation ─────────────────────────────────────────
    /// Certificate chain validation failed.
    #[error("certificate chain validation failed: {reason}")]
    CertificateChainInvalid { reason: String },

    /// Certificate has expired or is not yet valid.
    #[error("certificate validity error: {reason}")]
    CertificateValidity { reason: String },

    /// Certificate hash mismatch (e.g., `sender_cert_hash` lookup failed).
    #[error("certificate hash mismatch: expected {expected}, got {actual}")]
    CertificateHashMismatch { expected: String, actual: String },

    // ── AuthContainer ──────────────────────────────────────────────────
    /// A required field in the `AuthContainer` is missing or empty.
    #[error("auth container missing field: {field}")]
    AuthMissingField { field: &'static str },

    /// The `AuthContainer` uses an unsupported hash algorithm.
    #[error("unsupported hash algorithm in auth container: {algo}")]
    AuthUnsupportedAlgo { algo: i32 },

    // ── Generic / catch-all ────────────────────────────────────────────
    /// Base64 decoding error.
    #[error("base64 decode error: {0}")]
    Base64Decode(#[from] base64::DecodeError),

    /// Any other error that doesn't fit the categories above.
    #[error("{0}")]
    Other(String),
}

// ── Convenience constructors ───────────────────────────────────────────

impl CryptoError {
    /// Create a [`CryptoError::FileRead`] from an `io::Error`.
    pub fn file_read(kind: &'static str, path: impl Into<String>, source: std::io::Error) -> Self {
        Self::FileRead {
            kind,
            path: path.into(),
            source,
        }
    }

    /// Create a [`CryptoError::Parse`] error.
    pub fn parse(kind: &'static str, reason: impl Into<String>) -> Self {
        Self::Parse {
            kind,
            reason: reason.into(),
        }
    }

    /// Create a [`CryptoError::Missing`] error.
    pub fn missing(kind: &'static str, detail: impl Into<String>) -> Self {
        Self::Missing {
            kind,
            detail: detail.into(),
        }
    }

    /// Create a [`CryptoError::SigningFailed`] error.
    pub fn signing(reason: impl Into<String>) -> Self {
        Self::SigningFailed {
            reason: reason.into(),
        }
    }

    /// Create a [`CryptoError::VerificationFailed`] error.
    pub fn verification(reason: impl Into<String>) -> Self {
        Self::VerificationFailed {
            reason: reason.into(),
        }
    }
}
