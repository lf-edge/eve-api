// SPDX-License-Identifier: Apache-2.0

//! Cryptographic protocol traits for the EVE Device API.
//!
//! This module defines the traits that codify the
//! [OBJECT-SIGNING](https://github.com/lf-edge/eve-api/blob/main/OBJECT-SIGNING.md)
//! specification:
//!
//! - [`CryptoError`] — protocol-level error type (no backend-specific variants)
//! - [`Signer`] — sign and verify with ECDSA-P256-SHA256
//! - [`CertStore`] — certificate access (signing cert, device cert, hashes)
//! - [`OnboardCertStore`] — `CertStore` for the onboarding phase
//! - [`DeviceCertStore`] — `CertStore` for the operational phase
//! - [`pem_encode`] — utility to PEM-wrap DER bytes
//!
//! - [`Authenticatable`] — blanket impl on `prost::Message + Default` for
//!   `AuthContainer` wrapping/unwrapping

pub mod authenticatable;
pub mod cert_store;
pub mod error;
pub mod provider;
pub mod signer;

pub use authenticatable::Authenticatable;
pub use cert_store::{pem_encode, CertStore, DeviceCertStore, OnboardCertStore};
pub use error::CryptoError;
pub use provider::CryptoProvider;
pub use signer::Signer;
