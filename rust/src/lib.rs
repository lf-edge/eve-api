//! # EVE API Rust Bindings
//!
//! This crate provides auto-generated Rust protobuf bindings for the
//! [EVE Device API](https://github.com/lf-edge/eve-api), enabling communication
//! between EVE-compatible edge devices and controllers.
//!
//! ## Overview
//!
//! This is a **pure data types** crate — it contains only the generated protobuf
//! message structs and enums with no business logic, no crypto, and no HTTP client.
//! Higher-level functionality is provided by companion crates:
//!
//! - `crypto-provider` — Traits and implementations for signing/verification
//! - `eve-api-client` — Async HTTP client for the EVE controller API
//!
//! ## Modules
//!
//! Each module corresponds to a protobuf package from the EVE API:
//!
//! - [`common`] — Shared types: `HashAlgorithm`, `CipherBlock`, `CipherContext`, etc.
//! - [`auth`] — `AuthBody`, `AuthContainer` for message signing envelopes
//! - [`register`] — `ZRegisterMsg` for device onboarding
//! - [`config`] — `EdgeDevConfig`, `ConfigRequest`, `ConfigResponse`
//! - [`certs`] — `ZControllerCert`, `ZCert` for certificate management
//! - [`attest`] — `ZAttestReq`, `ZAttestResp` for TPM attestation
//! - [`info`] — `ZInfoMsg` for device/app status reporting
//! - [`metrics`] — `ZMetricMsg` for resource usage reporting
//! - [`logs`] — `LogBundle`, `LogEntry` for device logging
//! - [`flowlog`] — `FlowMessage` for network flow statistics
//! - [`uuid`] — `UuidRequest`, `UuidResponse`
//! - [`hardwarehealth`] — `ZHardwareHealth` for hardware health reports
//! - [`profile`] — `LocalProfile` for local profile server
//!
//! ## Example
//!
//! ```rust
//! use eve_api::register::ZRegisterMsg;
//! use prost::Message;
//!
//! let msg = ZRegisterMsg {
//!     pem_cert: bytes::Bytes::from_static(b"my-cert"),
//!     serial: "DEVICE-001".to_string(),
//!     soft_serial: "SOFT-001".to_string(),
//!     ..Default::default()
//! };
//!
//! // Serialize to protobuf wire format
//! let encoded = msg.encode_to_vec();
//! assert!(!encoded.is_empty());
//!
//! // Deserialize back
//! let decoded = ZRegisterMsg::decode(encoded.as_slice()).unwrap();
//! assert_eq!(decoded.serial, "DEVICE-001");
//! ```

// ── evecommon ──────────────────────────────────────────────────────────
/// Common types shared across EVE API packages.
///
/// Includes `HashAlgorithm`, `CipherBlock`, `CipherContext`,
/// `DiskDescription`, `RadioAccessTechnology`, and more.
pub mod common {
    include!(concat!(env!("OUT_DIR"), "/org.lfedge.eve.common.rs"));
}

// ── auth ───────────────────────────────────────────────────────────────
/// Authentication envelope messages.
///
/// The [`AuthContainer`] wraps every API request/response to provide
/// end-to-end integrity via ECDSA signatures, even through TLS MitM proxies.
pub mod auth {
    include!(concat!(env!("OUT_DIR"), "/org.lfedge.eve.auth.rs"));
}

// ── register ───────────────────────────────────────────────────────────
/// Device registration (onboarding) messages.
pub mod register {
    include!(concat!(env!("OUT_DIR"), "/org.lfedge.eve.register.rs"));
}

// ── certs ──────────────────────────────────────────────────────────────
/// Controller certificate management messages.
pub mod certs {
    include!(concat!(env!("OUT_DIR"), "/org.lfedge.eve.certs.rs"));
}

// ── attest ─────────────────────────────────────────────────────────────
/// Attestation and trust-anchoring messages (TPM quotes, nonces, keys).
pub mod attest {
    include!(concat!(env!("OUT_DIR"), "/org.lfedge.eve.attest.rs"));
}

// ── config ─────────────────────────────────────────────────────────────
/// Device and application configuration messages.
///
/// The central type is [`EdgeDevConfig`] which carries the complete
/// desired state for a device.
pub mod config {
    include!(concat!(env!("OUT_DIR"), "/org.lfedge.eve.config.rs"));
}

// ── info ───────────────────────────────────────────────────────────────
/// Device and application status/information messages.
pub mod info {
    include!(concat!(env!("OUT_DIR"), "/org.lfedge.eve.info.rs"));
}

// ── metrics ────────────────────────────────────────────────────────────
/// Device and application metrics messages.
pub mod metrics {
    include!(concat!(env!("OUT_DIR"), "/org.lfedge.eve.metrics.rs"));
}

// ── logs ───────────────────────────────────────────────────────────────
/// Device and application log messages.
pub mod logs {
    include!(concat!(env!("OUT_DIR"), "/org.lfedge.eve.logs.rs"));
}

// ── flowlog ────────────────────────────────────────────────────────────
/// Network flow logging messages (TCP/UDP flows, DNS lookups).
pub mod flowlog {
    include!(concat!(env!("OUT_DIR"), "/org.lfedge.eve.flowlog.rs"));
}

// ── eveuuid ────────────────────────────────────────────────────────────
/// UUID request/response messages for device identity.
pub mod uuid {
    include!(concat!(env!("OUT_DIR"), "/org.lfedge.eve.uuid.rs"));
}

// ── hardwarehealth ─────────────────────────────────────────────────────
/// Hardware health reporting messages (ECC memory, storage, etc.).
pub mod hardwarehealth {
    include!(concat!(
        env!("OUT_DIR"),
        "/org.lfedge.eve.hardwarehealth.rs"
    ));
}

// ── profile ────────────────────────────────────────────────────────────
/// Local profile and network status messages for Local Operator Console.
pub mod profile {
    include!(concat!(env!("OUT_DIR"), "/org.lfedge.eve.profile.rs"));
}

// ── proxy ──────────────────────────────────────────────────────────────
/// SCEP proxy messages for certificate enrollment.
pub mod proxy {
    include!(concat!(env!("OUT_DIR"), "/org.lfedge.eve.proxy.rs"));
}

// ── nestedappinstancemetrics ───────────────────────────────────────────
/// Nested (child) application instance metrics, logs, and inventory.
pub mod nestedappinstancemetrics {
    include!(concat!(
        env!("OUT_DIR"),
        "/org.lfedge.eve.nestedappinstancemetrics.rs"
    ));
}

// ── Convenience re-exports ─────────────────────────────────────────────

pub use auth::{AuthBody, AuthContainer};
pub use common::{CipherBlock, CipherContext, HashAlgorithm};
pub use config::EdgeDevConfig;
pub use register::ZRegisterMsg;

/// API version implemented by this crate.
pub const API_VERSION: &str = "v2";

/// Base path prefix for all API endpoints.
pub const API_PATH_PREFIX: &str = "/api/v2/edgedevice";
