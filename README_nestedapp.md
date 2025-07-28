# nestedappinstancemetrics

nestedappinstancemetrics defines an interface for an http webserver inside of a
'runtime' on EVE which returns metric data to EVE for containers or groups
of containers which reside in the runtime.  This runtime can be a VM and
is expected to be running inside an EVE AppInstance and identified by the
eve-api config where AppInstanceConfig.runtime_type!=APP_RUNTIME_TYPE_UNSPECIFIED.
The server listens on the address defined by CollectStatsIPAddr in app
instance config, this address will be on an airgapped local network
instance with no external access outside of EVE.  The consumer of this
data is in the pillar service container in EVE-OS.

## Endpoints

The following are the API endpoints which must be implemented by the runtime
metric server. All endpoints will start with an api prefix path "/api/v1/"
until further versions are defined. All endpoints are accessible over
http at tcp port 57475.

### Inventory

GET /api/v1/inventory/nested-app-id

Return codes:

* Inventory Returned: `200`
* Non-GET type request: `405`

Response:

The body will contain the NestedAppInventory structure in json form.

### Metrics

GET /api/v1/metrics/nested-app-id/\<app-id\>

Return codes:

* Metrics returned: `200`
* App Id not in uuid format: `400`
* App Id not defined in runtime: `404`
* Non-GET type request: `405`

Response:

The body will contain the NestedAppMetrics structure in json form.

### Storage Metric

To retrieve runtime level metrics containing usage for all nested applications contained
inside the runtime this will be returned from a single endpoint below.
Each NestedAppRuntimeDiskMetric is associated with a runtime persistent filesystem
and provides the usage snapshot at a runtime scope which
contains 1 or more nested app ids sharing the fs.

GET /api/v1/storagemetrics/runtime

Return codes:

* Metric Object Returned: `200`
* Non-Get type request: `405`

Response:

The body will contain the NestedAppRuntimeDiskMetric structure in protobuf form.
