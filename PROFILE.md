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
