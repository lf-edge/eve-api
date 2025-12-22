# Local Profile Server (LPS) API

This document defines the API for Local Profile Server (LPS).

This document is version 1, and all endpoints will begin with `/api/v1`.

## Server endpoint

EVE MUST use the server endpoint specified using the `local_profile_server`
in [EdgeDevConfig](./proto/config/devconfig.proto) and use the associated
`profile_server_token` to validate the responses. If no port number is specified
in `local_profile_server` EVE MUST default to `8888`.

If the `local_profile_server` is empty, then EVE MUST NOT invoke this API.
If the `local_profile_server` is cleared, then EVE MUST forget any configuration
it had received from the `local_profile_server`.

## Mime Types

All `GET` requests MUST have no mime type set.
All `POST` requests MUST have the mime type set of `application/x-proto-binary`.
All responses with a body MUST have the mime type set of
`application/x-proto-binary`.

## Endpoints

The following are the API endpoints that MAY be implemented by a profile server.

### Local Profile

Retrieve the local profile, which will override any global profile.

```http
   GET /api/v1/local_profile
```

Return codes:

* Valid: `200`
* Not implemented: `404`

Request:

The request MUST use HTTP for this request.
The request MUST NOT contain any body content.

Response:

The response mime type MUST be `application/x-proto-binary`. The response MUST
contain a single protobuf message of type [LocalProfile](./proto/profile/local_profile.proto).

The requester MUST verify that the response payload has the correct
`server_token`.
If the profile is empty, it will reset any saved local profile but otherwise
have no effect.

A non-empty profile will override the `global_profile` specified in
[EdgeDevConfig](./proto/config/devconfig.proto). The resulting current
profile will be used to determine which app instances are started and
stopped by matching against the [profile_list in AppInstanceConfig](./proto/config/appconfig.proto).

### Radio

Publishes the current state of all wireless network adapters.
The response may optionally include radio configuration to apply,
which is currently limited to toggling the Radio-Silence mode.
Radio-Silence mode disables or enables all wireless radio devices simultaneously,
turning wireless communications OFF or ON.
To request changes related to IP settings or cellular network port
configurations, the Local Profile Server should use the Network endpoint instead
(`api/v1/network`).

```http
   POST /api/v1/radio
```

Return codes:

* Success; with new radio configuration in the response: `200`
* Success; without radio configuration in the response: `204`
* Not implemented: `404`

Request:

The request mime type MUST be `application/x-proto-binary`.
The request MUST have the body of a single protobuf message of type
[RadioStatus](./proto/profile/local_profile.proto).

Response:

The response MAY contain the body of a single protobuf message of type
[RadioConfig](./proto/profile/local_profile.proto) encoded as
`application/x-proto-binary`.

The requester MUST verify that the response payload (if provided)
has the correct `server_token`.
If the verification succeeds, it will apply the received radio configuration.

Device MUST stop publishing radio status until all changes in the received radio
configuration are fully applied, without any ongoing or pending operations left
behind.
When device fails to apply the configuration, it SHOULD eventually stop retrying
and publish the new radio status afterwards, indicating the error condition
inside the `RadioStatus.config_error` field.

### AppInfo

Publish the current state of app instances on the device to LPS and optionally
obtain a list of app commands to execute.

```http
   POST /api/v1/appinfo
```

Return codes:

* Success; with commands to execute as defined in the response body: `200`
* Success; without commands to execute: `204`
* Not implemented: `404`

Request:

The request mime type MUST be `application/x-proto-binary`.
The request MUST have the body of a single protobuf message
of type [LocalAppInfoList](./proto/profile/local_profile.proto).
Device publishes information repeatedly to keep the local server updated and
to allow the server to submit application commands for execution.
Local server MAY throttle or cancel this communication stream by returning
the `404` code.

Response:

The response MAY contain the body of a single protobuf message of type
[LocalAppCmdList](./proto/profile/local_profile.proto), encoded as
`application/x-proto-binary`.

The requester MUST verify that the response payload (if provided)
has the correct `server_token`.
If the verification succeeds, all entries of `app_commands` are iterated,
and those that successfully match a running application instance (by `id`
and/or `displayname`) are applied.

Currently, the method allows to request a locally running application instance
to be *restarted* or *purged* by EVE. This may help to resolve a case of
an application being in a broken state, and the user not being able to fix it
(remotely) due to a lack of connectivity between the device and the controller.
Rather than rebooting the entire device (locally), it is possible to
restart/purge only a selected application.

A command request, as defined by `AppCommand` protobuf message, includes
an important field `timestamp` (`uint64`), which should record the time when
the request was made by the user. The format of the timestamp is not defined.
It can be a Unix timestamp or a different time representation. It is not even
required for the timestamp to match the real time or to be in-sync with
the device clock.

What is required, however, is that two successive but distinct requests made
for the same application will have different timestamps attached.
This requirement applies even between restarts of the Local profile server.
A request made after a restart should not have the same timestamp attached
as the previous request made for the same application before the restart.

EVE guarantees that a newly added command request (into `LocalAppCmdList.app_commands`),
or a change of the `timestamp` field, will result in the command being triggered
ASAP. Even if the execution of a command is interrupted by a device reboot/crash,
the eventuality of the command completion is still guaranteed.
The only exception is if Local Profile Server restarts/crashes shortly after
a request is made, in which case it can get lost before EVE is able to receive
it. For this scenario to be avoided, a persistence of command requests
on the side of the Local Profile server is necessary.

It is not required for the Local profile server to stop submitting command
requests that have been already processed by EVE. Using the `timestamp` field,
EVE is able to determine if a given command request has been already handled
or not.
To check if the last requested command has completed, compare its timestamp with
`last_cmd_timestamp` field from `LocalAppInfo` message, submitted by EVE in
the request body of the API.

### App Boot Info

Publish the effective boot order and its source for all application instances
to LPS, and optionally receive boot configuration updates in response.
This endpoint allows bidirectional communication: EVE reports current boot
status, and LPS can respond with new boot configurations.

```http
   POST /api/v1/appbootinfo
```

Return codes:

* Success; with boot configuration in the response: `200`
* Success; no configuration changes needed: `204`
* Not implemented or no config for this device: `404`

Request:

The request mime type MUST be `application/x-proto-binary`.
The request MUST have the body of a single protobuf message of type
[AppBootInfoList](./proto/profile/local_profile.proto).
Device publishes information repeatedly to keep the local server updated and
to allow the server to submit boot configuration changes.
Local server MAY throttle or cancel this communication stream by returning
the `404` code.

The `AppBootInfoList` contains boot information for all application instances.
Each `AppBootInfo` entry includes:

* `id` - Application UUID
* `displayname` - User-friendly application name
* `boot_order` - The effective boot order applied to this application
  (BootOrder enum: `BOOT_ORDER_UNSPECIFIED`, `BOOT_ORDER_USB`, `BOOT_ORDER_NOUSB`)
* `source` - Which configuration source determined the effective boot order
  (BootOrderSource enum):
  * `BOOT_ORDER_SOURCE_UNSPECIFIED` (0): No explicit boot order was configured
  * `BOOT_ORDER_SOURCE_LPS` (1): Set by LPS via `/api/v1/appbootinfo`
  * `BOOT_ORDER_SOURCE_CONTROLLER` (2): Set by Controller API via
    `VmConfig.boot_order`
  * `BOOT_ORDER_SOURCE_DEVICE_PROPERTY` (3): Set by device property
    `app.boot.order`

Response:

The response MAY contain the body of a single protobuf message of type
[AppBootConfigList](./proto/profile/local_profile.proto), encoded as
`application/x-proto-binary`.

The requester MUST verify that the response payload (if provided) has the
correct `server_token`. If the verification succeeds, the boot configuration
is applied.

The `AppBootConfigList` contains boot configurations for one or more
applications. Each `AppBootConfig` entry specifies:

* `id` - Application UUID (at least one of id/displayname must be set)
* `displayname` - User-friendly application name
* `usb_boot` - USB boot mode for this application (BootOrder enum):
  * `BOOT_ORDER_UNSPECIFIED` (0): No override - use the next priority level
    (Controller API setting, then Device Property, then OVMF default).
    This allows selectively overriding only specific apps via LPS.
  * `BOOT_ORDER_USB` (1): Enable USB boot priority - OVMF will prioritize USB
    devices in boot order
  * `BOOT_ORDER_NOUSB` (2): Disable USB boot - OVMF will remove USB devices from
    boot order completely, ensuring the VM boots from disk

The setting is passed to OVMF via fw_cfg `opt/eve.bootorder` when the VM starts.

**Precedence**: Boot order can be configured from multiple sources (highest to
lowest priority):

1. **LPS API** - via this `/api/v1/appbootinfo` endpoint (per-app)
2. **Controller API** - via `VmConfig.boot_order` in `AppInstanceConfig` (per-app)
3. **Device Property** - via `app.boot.order` device configuration (device-wide)

EVE uses the highest priority source that provides a specified value. A source
"provides" a boot order when:

* **LPS**: EVE has a cached LPS config for the app with `usb_boot` set to
  `BOOT_ORDER_USB` or `BOOT_ORDER_NOUSB` (not `BOOT_ORDER_UNSPECIFIED`)
  (see "LPS Response Handling" below for how the cache is populated/cleared)
* **Controller API**: `VmConfig.boot_order` is `BOOT_ORDER_USB` or `BOOT_ORDER_NOUSB`
  (not `BOOT_ORDER_UNSPECIFIED`)
* **Device Property**: `app.boot.order` is `"usb"` or `"nousb"` (not empty `""`)

`BOOT_ORDER_UNSPECIFIED` (for API fields) or empty string `""` (for device property)
means "no override" - fall back to the next priority level. If all sources are
unspecified/empty, OVMF default behavior applies.

**LPS Response Handling and Persistence**:

EVE caches LPS boot configurations and persists them to disk. This is primarily
needed to handle EVE reboots: since LPS typically runs as a VM on the same
EVE device, when EVE reboots, LPS is unavailable until EVE and the LPS VM come
back up. The persisted cache ensures boot order settings remain active during
this window.

| LPS Response | Effect on Cache | Behavior |
|--------------|-----------------|----------|
| **LPS unavailable** | Cache preserved | Persisted boot order applied |
| **404 Not Found** | Cache cleared | Fall back to Controller/Device |
| **204 No Content** | Cache preserved | No changes; keep current order |
| **200 OK with apps** | Cache updated | Update listed apps; reset others |
| **200 OK with empty list** | Cache cleared | ALL apps reset to default |

**Key behavior during EVE reboot**:

1. EVE starts, LPS VM is not yet running
2. EVE loads persisted LPS boot configs from disk and applies them
3. VMs start with the previously configured boot order
4. LPS VM starts and begins responding to requests
5. EVE fetches fresh config from LPS; if unchanged, no action needed

**Key distinction for 404**: When LPS returns 404, it means "LPS is running but
has no config for this device" - the cache is **cleared** and boot order falls
back to Controller/Device settings. This is different from LPS being unavailable
(e.g., during EVE boot), where the cache is **preserved**.

**Priority Resolution Examples**:

| Device | Controller | LPS Cache | Effective |
|--------|------------|-----------|-----------|
| `"usb"` | `UNSPECIFIED` | *(empty)* | `usb` |
| `"usb"` | `NOUSB` | *(empty)* | `nousb` |
| `"usb"` | `NOUSB` | `USB` | `usb` |
| `"nousb"` | `USB` | `NOUSB` | `nousb` |
| `""` | `UNSPECIFIED` | *(empty)* | *(default)* |
| `""` | `USB` | *(empty)* | `usb` |

**LPS Response Scenarios (detailed)**:

| LPS Response | App in List | `usb_boot` | Effect |
|--------------|-------------|------------|--------|
| Unavailable | N/A | N/A | Cache preserved; use persisted config |
| 404 | N/A | N/A | Cache cleared; reset to default |
| 204 | N/A | N/A | Cache preserved; no changes |
| 200 | No (empty) | N/A | Cache cleared; ALL apps reset |
| 200 | No (others) | N/A | This app reset to default |
| 200 | Yes | `BOOT_ORDER_UNSPECIFIED` | Unspecified; falls back |
| 200 | Yes | `BOOT_ORDER_USB` | LPS wins; use USB priority |
| 200 | Yes | `BOOT_ORDER_NOUSB` | LPS wins; deprioritize USB |

**Behavior (for 200 response with config):**

* Each response from LPS is treated as the complete desired state for boot
  configuration.
* Applications included in `app_configs` will have their boot order set to the
  specified `usb_boot` value.
* Applications NOT included in the response will use the default boot order.
* To explicitly set default boot order for an app, include it with `usb_boot`
  set to `BOOT_ORDER_UNSPECIFIED`.
* Changes take effect on the next VM restart.

**Note:** A `204` response means "no changes" - EVE will preserve the current
boot configuration without modification.

**Posting Behavior:**

EVE posts boot info immediately when boot order decisions are finalized, including:

* After processing configuration from the controller
* When LPS updates boot order via this endpoint
* When the device property `app.boot.order` changes

Additionally, EVE posts periodically approximately once per minute. When LPS
returns `404`, EVE will throttle periodic posting to approximately once per hour.

### DevInfo

Publish the current state of the device to LPS and optionally obtain a command
to execute.

```http
   POST /api/v1/devinfo
```

Return codes:

* Success; with a command to execute as defined in the response body: `200`
* Success; without a command to execute: `204`
* Not implemented: `404`

Request:

The request mime type MUST be `application/x-proto-binary`.
The request MUST have the body of a single protobuf message of type
[LocalDevInfo](./proto/profile/local_profile.proto).
Device publishes information repeatedly to keep the local server updated and
to allow the server to submit commands for execution.
Local server MAY throttle or cancel this communication stream by returning
the `404` code.

Response:

The response MAY contain the body of a single protobuf message of type
[LocalDevCmd](./proto/profile/local_profile.proto), encoded as
`application/x-proto-binary`.

The requester MUST verify that the response payload (if provided) has
the correct `server_token`.
If the verification succeeds, then the timestamp is checked to determine whether
or not the command has already been executed, and it not it is applied.

Currently, the method allows to request a graceful Shutdown (of all app
instances) or such a Shutdown followed by a Poweroff of EVE. This allows for
graceful shutdown of applications and optionally a poweroff whether triggered
by a user on the local profile server or a UPS interfacing with the local
profile server.

The command request includes an important field `timestamp` (`uint64`), which
should record the time when the request was made
by the user. The format of the timestamp is not defined. It can be a Unix
timestamp or a different time representation. It is not even required for the
timestamp to match the real time or to be in-sync with the device clock.

What is required, however, is that two successive but distinct requests made for
the device will have different timestamps attached.
This requirement applies even between restarts of the Local profile server.
A request made after a restart should not have the same timestamp attached
as the previous request made before the restart.

EVE guarantees that a newly added command request,
or a change of the `timestamp` field, will result in the command being triggered
ASAP.
Even if the execution of a command is interrupted by a device reboot/crash,
the eventuality of the command completion is still guaranteed.
The only exception is if Local Profile Server restarts/crashes shortly after
a request is made, in which case it can get lost before EVE is able to receive
it. For this scenario to be avoided, a persistence of command requests
on the side of the Local Profile server is necessary.

It is not required for the Local profile server to stop submitting command
requests that have been already processed by EVE. Using the `timestamp` field,
EVE is able to determine if a given command request has been already handled
or not.
To check if the last requested command has completed, compare its timestamp with
`last_cmd_timestamp` field from `LocalDevInfo` message, submitted by EVE in
the request body of the API.

### Device Location Info (GNSS)

Publish the current location of the device as obtained from a GNSS receiver
to the local server.

```http
   POST /api/v1/location
```

Return codes:

* Success: `200`
* Not implemented: `404`

Request:

The request mime type MUST be `application/x-proto-binary`.
The request MUST have the body of a single protobuf message of type
[ZInfoLocation](./proto/info/info.proto).
Device publishes information repeatedly with a (default) period of 20 seconds
to keep the local server updated (configurable using
[timer.location.app.interval](https://github.com/lf-edge/eve/blob/master/docs/CONFIG-PROPERTIES.md).
Local server MAY throttle or cancel this communication stream by returning
the `404` code.

### Network

Publishes the current IP configuration of all network adapters (excluding those
directly assigned to applications).
The response may optionally include a locally-declared desired configuration
for one or more adapters, which EVE will validate and apply if permitted.

```http
   POST /api/v1/network
```

Return codes:

* Success, with locally-declared network configuration included in the response
  for EVE to apply: `200`
* Success, no local network configuration updates requested; any existing
  local configuration remains in effect: `204`
* Not implemented, or intentionally used by the local server to throttle
  the periodic network information updates: `404`
  When `404` is returned, any previously submitted local network configuration
  is reverted.

Request:

The request MIME type MUST be `application/x-proto-binary`.
The request MUST contain the body of a single protobuf message of type
[NetworkInfo](./proto/profile/network.proto).
Device publishes network information repeatedly to keep LPS updated and
to allow the server to submit local configuration updates.
Local server MAY throttle or cancel this communication stream by returning
the `404` code.

`NetworkInfo` includes:

* The latest network configuration received from the controller.
* Status of controller connectivity
* A fallback network configuration, used when the latest configuration
  fails to provide working controller connectivity.
* Status of the local network configuration submitted previously
  by the Local Profile Server (LPS), indicating whether it was successfully
  applied or if errors occurred.

Response:

The response MAY contain the body of a single protobuf message of type
[LocalNetworkConfig](./proto/profile/network.proto), encoded as
`application/x-proto-binary`.
If no further updates are needed to the local configuration, the server MAY return
HTTP 204 (`No Content`), and EVE will continue using the most recently submitted
local configuration.

`LocalNetworkConfig` contains:

* An authorization token (`server_token`) to verify the request against
  the controller-provisioned secret.
* Declarative network configuration for ports managed locally.

Behavior:

* EVE validates the received local configuration to ensure it is well-formed
  and that the submitted `server_token` matches the value provisioned
  by the controller.
* The controller may specify, on a per-port basis, whether local modifications
  are allowed (see `SystemAdapter.allow_local_modifications`).
  Ports that are not permitted to receive local configuration are skipped,
  and an error is reported for each such port in `LocalNetworkConfigInfo`.
* If the controller revokes local modification permissions or un-configures LPS,
  EVE reverts affected adapters to the controller configuration.
* Locally-manageable fields include the IP configuration, wireless settings,
  and proxy configuration — the attributes defined in [NetworkPortConfig](proto/profile/network.proto).
  Fields that affect overall network topology, such as interface usage, cost,
  or assigned labels, are not locally-manageable and remain under controller
  management.
* Local port configuration overrides the entire set of locally-manageable
  fields. Partial updates are not supported — fields omitted or set to empty/zero
  values are applied as such, rather than inheriting values from the controller
  configuration. The locally-manageable portion is treated as a single unit,
  so the controller cannot manage some of these fields while the local user
  manages others.
* Fields that are not included in the set of locally-manageable attributes
  (e.g., interface usage, cost) continue to follow the controller-provided
  configuration, either the latest or the active fallback configuration,
  as appropriate.
* When LPS (temporarily or indefinitely) throttles/cancels the communication
  stream by returning `404`, previously submitted local network configuration
  is reverted.
* If LPS becomes inaccessible or unresponsive, a previously submitted network
  configuration remains in effect. To revert local config received from a crashed
  or misbehaving LPS, the controller user must explicitly revoke permissions
  or disable the LPS.
* EVE disables fallback mechanism only for adapters with a local configuration
  with respect to fields that can be locally managed. All other adapters, or fields
  that are not locally-manageable, such as interface usage or cost, continue using
  either the latest or fallback controller configuration, depending on connectivity
  status.
* The published controller connectivity status helps the local operator create
  a local configuration that ensures proper connectivity.
* When local configuration fails validation or application, EVE reports errors
  back in `NetworkInfo.local_config`.

## Security

In addition to using a `server_token` it is recommended that ACLs/firewall rules
are deployed so that the traffic to/from the local profile server can not be
directed to non-local destinations.
