use prost_build::Config;
use std::path::Path;
use std::process::Command;

fn main() -> Result<(), Box<dyn std::error::Error>> {
    let proto_dir = Path::new(env!("CARGO_MANIFEST_DIR")).join("../proto");

    if !proto_dir.exists() {
        return Err(format!(
            "Proto directory not found at {}. Make sure this crate lives inside the eve-api repo.",
            proto_dir.display()
        )
        .into());
    }

    let proto_dir_str = proto_dir.to_str().unwrap();

    // Extra include path for third-party proto deps (e.g., validate/validate.proto
    // from protoc-gen-validate, imported by register.proto).
    //
    // Resolution order:
    // 1. PROTO_DEPS_DIR env var (set inside devcontainer → /proto/.proto_deps)
    // 2. Local proto-deps/ directory (auto-fetched for local dev)
    let proto_deps_dir = if let Ok(d) = std::env::var("PROTO_DEPS_DIR") {
        Path::new(&d).join("protoc-gen-validate").to_path_buf()
    } else {
        let local = Path::new(env!("CARGO_MANIFEST_DIR")).join("proto-deps");
        let validate = local.join("validate").join("validate.proto");
        if !validate.exists() {
            eprintln!("Fetching validate.proto for local build...");
            std::fs::create_dir_all(local.join("validate")).expect("create proto-deps/validate/");
            let status = Command::new("curl")
                .args([
                    "-sSfL",
                    "-o",
                    validate.to_str().unwrap(),
                    "https://raw.githubusercontent.com/bufbuild/protoc-gen-validate/v1.2.1/validate/validate.proto",
                ])
                .status()
                .expect("failed to run curl — install curl or use the devcontainer");
            if !status.success() {
                panic!("failed to download validate.proto (exit {})", status);
            }
        }
        local
    };
    let proto_deps_str = proto_deps_dir.to_str().unwrap().to_owned();

    let mut config = Config::new();

    // Use Bytes for all bytes fields for zero-copy efficiency
    config.bytes(["."]);

    // AuthContainer contains large binary payloads, but we still need Debug
    // for prost::Message trait bound. We let prost derive Debug normally and
    // accept the verbose output — users can use the {:?} alternate form or
    // wrap in a custom Display if needed.

    // Proto files grouped by package, ordered so dependencies come first.
    //
    // The include path is the proto root directory; file paths passed to
    // compile_protos are relative to that root.
    let proto_files: Vec<String> = [
        // --- evecommon (no deps on other eve protos) ---
        "evecommon/evecommon.proto",
        "evecommon/devmodelcommon.proto",
        "evecommon/acipherinfo.proto",
        "evecommon/netcmn.proto",
        // --- eveuuid ---
        "eveuuid/eveuuid.proto",
        // --- auth (depends on evecommon) ---
        "auth/auth.proto",
        // --- certs (depends on evecommon) ---
        "certs/certs.proto",
        // --- register ---
        "register/register.proto",
        // --- attest (depends on certs) ---
        "attest/attest.proto",
        // --- config (depends on evecommon, certs, auth) ---
        "config/devcommon.proto",
        "config/storage.proto",
        "config/vm.proto",
        "config/fw.proto",
        "config/scep.proto",
        "config/netconfig.proto",
        "config/netinst.proto",
        "config/baseosconfig.proto",
        "config/appconfig.proto",
        "config/edgeview.proto",
        "config/edge_node_cluster.proto",
        "config/patch_envelope.proto",
        "config/devmodel.proto",
        "config/compound_devconfig.proto",
        "config/devconfig.proto",
        // --- info (depends on evecommon, config) ---
        "info/cert.proto",
        "info/hardware.proto",
        "info/pnac.proto",
        "info/ntpsources.proto",
        "info/patch_envelope.proto",
        "info/edge_node_cluster.proto",
        "info/info.proto",
        // --- metrics (depends on evecommon) ---
        "metrics/nestedappruntimemetrics.proto",
        "metrics/metrics.proto",
        // --- logs ---
        "logs/log.proto",
        // --- flowlog ---
        "flowlog/flowlog.proto",
        // --- hardwarehealth ---
        "hardwarehealth/hardware_health.proto",
        // --- profile (depends on info, metrics, config) ---
        "profile/local_profile.proto",
        "profile/network.proto",
        // --- proxy (depends on evecommon) ---
        "proxy/scep.proto",
        // --- nestedappinstancemetrics ---
        "nestedappinstancemetrics/nestedappinstanceinventory.proto",
        "nestedappinstancemetrics/nestedappinstancelog.proto",
        "nestedappinstancemetrics/nestedappinstancemetrics.proto",
    ]
    .iter()
    .map(|f| format!("{}/{}", proto_dir_str, f))
    .collect();

    config.compile_protos(&proto_files, &[proto_dir_str, &proto_deps_str])?;

    // Rerun if any proto file changes
    println!("cargo:rerun-if-changed={}", proto_dir_str);
    println!("cargo:rerun-if-changed={}", proto_deps_str);
    println!("cargo:rerun-if-env-changed=PROTO_DEPS_DIR");
    println!("cargo:rerun-if-changed=build.rs");

    Ok(())
}
