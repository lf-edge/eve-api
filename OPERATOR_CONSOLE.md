# Local Operator Console (LOC) API

This document explains the Local Operator Console(LOC) from the API point of view.

This document is similarly structured as PROFILE.md, because the topic is very similar.

## Server endpoint

EVE periodically sends a `CompoundEdgeDevConfigRequest` to the Local Operator Console.
EVE can find the URL to the LOC in `loc_url` which is part of the
`EdgeDevConfig` that EVE received from the controller.

## Endpoints

In the following `$UUID` is the uuid of the edge node contacting LOC.

### Certs
```http
       GET /api/v2/edgedevice/certs
```

Returns the controller certificates.

### Config

```http
    POST /api/v2/edgedevice/id/$UUID/config
```

Returns the signed config from the controller.

### Compound Config

```http
    POST /api/v2/edgedevice/id/$UUID/compound-config
```

Returns edge device config as part of the compound config. It also includes:

* current Unix timestamp in nano seconds precision
* device command
* app commands
* radio config

The device command can be:

* shutdown of the device
* shutdown and poweroff of the device
* reboot of the device
* run collect-info and upload the tarball to the configured datastore via POST/PUT

The app commands can be:

* restart of the app
* purge of the app

### Edge Metrics

```http
    POST /api/v2/edgedevice/id/$UUID/metrics
```

This is where the metrics are sent. Similar to the controller.

### Edge Info

```http
    POST /api/v2/edgedevice/id/$UUID/info
```

This is where the info messages are sent. Similar to the controller.

