# EVE API — Rust Bindings

Rust protobuf bindings for the [EVE Device API v2](https://github.com/lf-edge/eve-api),
auto-generated at build time from `.proto` definitions using
[prost](https://github.com/tokio-rs/prost).

This is a **pure data types** crate — it contains only generated protobuf message
structs and enums. No business logic, no crypto, no HTTP client. Higher-level
functionality belongs in companion crates (e.g., `crypto-provider`, `eve-api-client`).

## Modules

Each module corresponds to a protobuf package from the EVE API:

| Module | Protobuf package | Key types |
|---|---|---|
| `common` | `org.lfedge.eve.common` | `HashAlgorithm`, `CipherBlock`, `CipherContext` |
| `auth` | `org.lfedge.eve.auth` | `AuthBody`, `AuthContainer` |
| `register` | `org.lfedge.eve.register` | `ZRegisterMsg` |
| `certs` | `org.lfedge.eve.certs` | `ZControllerCert`, `ZCert` |
| `config` | `org.lfedge.eve.config` | `EdgeDevConfig`, `ConfigRequest`, `ConfigResponse` |
| `attest` | `org.lfedge.eve.attest` | `ZAttestReq`, `ZAttestResp` |
| `info` | `org.lfedge.eve.info` | `ZInfoMsg` |
| `metrics` | `org.lfedge.eve.metrics` | `ZMetricMsg` |
| `logs` | `org.lfedge.eve.logs` | `LogBundle`, `LogEntry` |
| `flowlog` | `org.lfedge.eve.flowlog` | `FlowMessage` |
| `uuid` | `org.lfedge.eve.uuid` | `UuidRequest`, `UuidResponse` |
| `hardwarehealth` | `org.lfedge.eve.hardwarehealth` | `ZHardwareHealth` |
| `profile` | `org.lfedge.eve.profile` | `LocalProfile` |
| `nestedappinstancemetrics` | `org.lfedge.eve.nestedappinstancemetrics` | nested app inventory, logs, metrics |

## Usage

```rust
use eve_api::register::ZRegisterMsg;
use prost::Message;

let msg = ZRegisterMsg {
    pem_cert: bytes::Bytes::from_static(b"my-cert"),
    serial: "DEVICE-001".to_string(),
    soft_serial: "SOFT-001".to_string(),
    ..Default::default()
};

// Serialize to protobuf wire format
let encoded = msg.encode_to_vec();

// Deserialize
let decoded = ZRegisterMsg::decode(encoded.as_slice()).unwrap();
assert_eq!(decoded.serial, "DEVICE-001");
```

## Convenience re-exports

Commonly used types are re-exported from the crate root:

```rust
use eve_api::{AuthBody, AuthContainer, HashAlgorithm, EdgeDevConfig, ZRegisterMsg};
```

Constants:

```rust
assert_eq!(eve_api::API_VERSION, "v2");
assert_eq!(eve_api::API_PATH_PREFIX, "/api/v2/edgedevice");
```

## How bindings are generated

There are no pre-generated `.rs` files. The `build.rs` script compiles all
`.proto` files from `../proto/` via `prost-build` during `cargo build`. The
generated code lands in `$OUT_DIR` and is included via `include!()` macros
in `src/lib.rs`.

If upstream adds a new `.proto` file or package, update `build.rs` and
`src/lib.rs` accordingly.

## Building

```sh
cargo build          # generates and compiles bindings
cargo test           # roundtrip and smoke tests
cargo clippy         # lint (generated code warnings are suppressed in Cargo.toml)
```

Or from the repository root:

```sh
make rust            # build release
make rust-test       # run tests
make rust-check      # clippy + fmt check
```

## Dependencies

| Crate | Purpose |
|---|---|
| `prost` | Protobuf runtime (derive `Message`) |
| `prost-types` | Well-known protobuf types |
| `bytes` | Zero-copy `Bytes` for binary fields |
| `prost-build` | Build-time proto compilation |

## License

Apache-2.0 — see [LICENSE](../LICENSE).
