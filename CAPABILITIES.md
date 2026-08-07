# Capabilities

A Device reports what it supports so a Controller can avoid sending
configuration the Device would ignore, and avoid waiting for messages the Device
will never send. There are **three independent mechanisms**, all in `ZInfoDevice`
(see [info.proto](./proto/info/info.proto)), with **different semantics**.
Conflating them is a correctness bug.

| mechanism | field | kind |
| --- | --- | --- |
| `APICapability` | `api_capability` | monotonic level |
| `OptionalCapabilities` | `optional_capabilities` | independent booleans |
| `Capabilities` | `capabilities` | independent booleans (hardware) |

## APICapability — a level, not a set

`APICapability` covers two kinds of support:

1. **`EdgeDevConfig` fields the Device parses.** Set one on a Device below the
   level and it is silently ignored.
2. **Messages or fields the Device sends.** A Controller expecting one from a
   Device below the level waits for it indefinitely.

**A larger value implies every smaller one.** A Controller MUST test with `>=`:

```text
if device.api_capability >= API_CAPABILITY_MTU {
    // safe to set NetworkConfig.mtu and NetworkInstanceConfig.mtu
}
```

Equality and set-membership tests are wrong. A Device reporting
`API_CAPABILITY_SMART_REPORT` (20) also supports 1 through 19, and EVE-OS
reports a single top value rather than a set.

`API_CAPABILITY_UNSPECIFIED` (0) also covers Devices predating the field, so it
means "none of the below", not "unknown, try anyway".

### What each value covers

`cfg` = an `EdgeDevConfig` field the Device parses. `rpt` = something the Device
sends. Rows marked **(?)** are inferred from the enum comment and the commit that
introduced the value, not stated anywhere authoritative — corrections welcome.

| value | | covers |
| --- | --- | --- |
| `RETRY_UPDATE` = 1 | cfg | `BaseOS.retry_update` |
| `SHUTDOWN` = 2 | cfg | `EdgeDevConfig.shutdown` |
| `START_DELAY_IN_SECONDS` = 3 | cfg | `AppInstanceConfig.start_delay_in_seconds` |
| `EDGEVIEW` = 4 | cfg | `EdgeDevConfig.edgeview`, `EdgeViewConfig.token` |
| `VOLUME_SNAPSHOTS` = 5 | cfg | `AppInstanceConfig.snapshot`, `SnapshotConfig`; snapshots taken during an app instance update (`SNAPSHOT_TYPE_APP_UPDATE`) **(?)** |
| `NETWORK_INSTANCE_ROUTING` = 6 | cfg | `NetworkInstanceConfig.static_routes`, `.propagate_connected_routes`, and `IPRoute` |
| `BOOT_MODE` = 7 | cfg | `VmConfig.boot_mode` |
| `MTU` = 8 | cfg | `NetworkConfig.mtu`, `NetworkInstanceConfig.mtu` |
| `ADAPTER_USER_LABELS` = 9 | cfg | `SystemAdapter.shared_labels` |
| `ENFORCED_NET_INTERFACE_ORDER` = 10 | cfg | `VmConfig.enforce_network_interface_order`, and hence `NetworkAdapter.interface_order` and `Adapter.interface_order` |
| `NTPS_FQDN` = 11 | cfg | NTP servers as FQDN, and more than one — `ipspec.ntp`, `ipspec.more_ntp` **(?)** |
| `WIN_LIC_PASSTHROUGH` = 12 | cfg | `VmConfig.enable_oem_win_license_key` |
| `VOLUME_SNAPSHOTS_IMMEDIATE` = 13 | cfg | `SnapshotType.SNAPSHOT_TYPE_IMMEDIATE`; snapshots on demand, which restart the app **(?)** |
| `ENCRYPTED_PATCH_ENVELOPE` = 14 | cfg | `EveBinaryArtifact.encrypted_inline`, `.encrypted_volumeref`, `.metadata_cipher_data` |
| `SINGLE_STACK_IP_NETWORK` = 15 | cfg | `NetworkType.V4Only`, `NetworkType.V6Only` |
| `CELLULAR_ATTACH_CONFIG` = 16 | cfg | `CellularAccessPoint.attach_apn`, `.attach_ip_type`, `.attach_auth_protocol` |
| `EDGEVIEW_AUTHENTICATION` = 17 | cfg | EdgeView command authentication; no single field — see [EDGEVIEW-CONTAINER-API.md](https://github.com/lf-edge/eve/blob/master/docs/EDGEVIEW-CONTAINER-API.md) **(?)** |
| `DISABLE_VTPM` = 18 | cfg | `VmConfig.disable_vtpm` |
| `LOC_REBOOT_COLLECT_INFO` = 19 | cfg | LOC-initiated reboot and collect-info; `LOCConfig.datastore_collect_info_id` **(?)** |
| `SMART_REPORT` = 20 | rpt | S.M.A.R.T. data in `ZHardwareHealth.disks`, superseding the deprecated `ZInfoHardware.disks` |
| `REPORT_TPM_EVENTLOG` = 21 | rpt | `ZAttestQuote.tpm_binary_event_log`, superseding the deprecated `ZAttestQuote.event_log` |

### Adding a value

Appending value *N* asserts that a Device reporting it also supports everything
below. So:

1. Append at the end. Never insert, never renumber.
2. Name the field(s) or message(s) it covers, in the enum comment and in the
   table above.
3. If the capability depends on build flavor or hardware rather than on version,
   it belongs in `OptionalCapabilities` instead.

## OptionalCapabilities — independent booleans

Not monotonic; test each separately. These describe properties that vary by
EVE-OS build flavor rather than by version.

| field | meaning | matters because |
| --- | --- | --- |
| `hv_type_kubevirt` | Device runs the Kubevirt hypervisor | required before sending `EdgeNodeCluster`; the KVM flavor cannot join a cluster |
| `hw_inventory_support` | Device can produce `HardwareInventory` | distinguishes "found no hardware" from "cannot report" when `ZInfoHardware.inventory` is empty |
| `etcd_snapshot` | Device supports etcd snapshots | EVE-k cluster operations |

An absent boolean means "not supported or not reported". Do not infer support
from a Device predating the field.

## Capabilities — hardware

Independent booleans describing the platform, not the software:

| field | meaning |
| --- | --- |
| `HWAssistedVirtualization` | VMX/SVM on amd64, virtualization extensions on arm64 |
| `IOVirtualization` | IOMMU / I/O virtualization support |

These bound what a Device can ever run: device passthrough
(`AppInstanceConfig.adapters`, SR-IOV VFs) will not work without
`IOVirtualization`, whatever `APICapability` says.

## Checklist for Controller implementers

1. Test `APICapability` with `>=`; treat `0` as "none".
2. Test each `OptionalCapabilities` and `Capabilities` boolean individually.
3. Check the gate before sending any field in the table above.
4. A Device does not report being sent a gated field it does not understand — it
   just ignores it. That silence is why these gates exist.
