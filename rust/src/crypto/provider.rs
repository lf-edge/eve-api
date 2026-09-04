// SPDX-License-Identifier: Apache-2.0

//! Combined `CryptoProvider` supertrait for the EVE OBJECT-SIGNING protocol.
//!
//! [`CryptoProvider`] is the convenience bound used by
//! [`Authenticatable`](super::Authenticatable) and `EveClient` — it combines
//! [`Signer`] (signing/verification) with [`CertStore`] (certificate access)
//! into a single trait bound.
//!
//! # Blanket Implementation
//!
//! Any type that implements both `Signer` and `CertStore` is automatically a
//! `CryptoProvider` — no manual impl needed:
//!
//! ```rust,ignore
//! use eve_api::crypto::{Signer, CertStore, CryptoProvider};
//!
//! // SoftwareCryptoProvider implements Signer + CertStore,
//! // so it's automatically a CryptoProvider.
//! fn send_request(crypto: &impl CryptoProvider) {
//!     let sig = crypto.sign(b"payload").unwrap();
//!     let hash = crypto.signing_cert_hash();
//!     // ...
//! }
//! ```
//!
//! # Design Rationale
//!
//! `Signer` and `CertStore` are independently useful — you can pass a bare
//! `Signer` to lower-level code that doesn't need certs. `CryptoProvider` is
//! the convenient combined bound for protocol-level code that needs both.

use super::{CertStore, Signer};

/// Combined interface: signing + certificate access.
///
/// This is what [`Authenticatable::to_auth_container()`](super::Authenticatable)
/// and `EveClient` are generic over.
///
/// Implementations live in micro-eve crates (`SoftwareCryptoProvider`,
/// `TpmCryptoProvider`). Each holds a signing backend + certificate data,
/// implementing both [`Signer`] and [`CertStore`]. The blanket impl below
/// makes them automatically a `CryptoProvider`.
pub trait CryptoProvider: Signer + CertStore {}

/// Blanket implementation: anything that is both a `Signer` and a `CertStore`
/// is automatically a `CryptoProvider`.
impl<T: Signer + CertStore> CryptoProvider for T {}

#[cfg(test)]
mod tests {
    use super::*;
    use crate::crypto::{CryptoError, DeviceCertStore};
    use std::fmt;

    /// A minimal mock that implements both Signer and CertStore.
    struct MockProvider {
        cert_store: DeviceCertStore,
    }

    impl MockProvider {
        fn new() -> Self {
            Self {
                cert_store: DeviceCertStore::new(vec![0xAA; 100], vec![0x11; 32]),
            }
        }
    }

    impl fmt::Debug for MockProvider {
        fn fmt(&self, f: &mut fmt::Formatter<'_>) -> fmt::Result {
            f.debug_struct("MockProvider").finish()
        }
    }

    impl Signer for MockProvider {
        fn sign(&self, data: &[u8]) -> Result<Vec<u8>, CryptoError> {
            Ok(vec![data.len() as u8; 64])
        }

        fn verify(
            &self,
            _data: &[u8],
            signature: &[u8],
            _cert_der: &[u8],
        ) -> Result<bool, CryptoError> {
            Ok(signature.len() == 64)
        }
    }

    impl CertStore for MockProvider {
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

    #[test]
    fn test_blanket_impl_works() {
        // MockProvider implements Signer + CertStore, so it should
        // automatically be a CryptoProvider.
        fn requires_crypto_provider(p: &impl CryptoProvider) -> Vec<u8> {
            let sig = p.sign(b"test").unwrap();
            let _hash = p.signing_cert_hash();
            sig
        }

        let provider = MockProvider::new();
        let sig = requires_crypto_provider(&provider);
        assert_eq!(sig.len(), 64);
    }
}
