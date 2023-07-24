# EVE Device API

This is the "Device API", for communications between an [lf-edge eve](https://github.com/lf-edge/eve) edge device
and a controller.

See [https://www.lfedge.org/projects/eve/](https://www.lfedge.org/projects/eve/)

This directory defines the API itself as message definitions, documentation and language-specific bindings.

## Contents

### Documentation

Documentation is in markdown files in this directory, specifically:

* General overview and how to use and build in this [README.md](./README.md)
* The protocol in [APIv2.md](./APIv2.md)
* Local profile overrides in [PROFILE.md](./PROFILE.md)
* Object signing in [OBJECT-SIGNING.md](./OBJECT-SIGNING.md)

### Message definitions

The message definitions are in [protobufs](https://developers.google.com/protocol-buffers/) in
the [proto](./proto) subdirectory.

### Language bindings

Language bindings are generated from the protobufs, in language-specific directories:

* [go](./go)
* [python](./python)

## Using the language-specific bindings

To use the language-specific bindings, import them as libraries to your appropriate language.
For example, in go:

```go
import (
    "github.com/lf-edge/eve-api/go/config"
)
```

## Generating the language-specific bindings

To compile the protobufs into the target language, run:

```bash
make proto
```

To generate just a single language, run:

```bash
make go
# or
make python
```

### Visualizations

In addition to the language-specific libraries, `make proto` generates visualizations of the protobuf structure,
beginning with the root of an edge device config. These are
available as `.svg`, `.dot` and `.png` as below. Click to zoom in.

* [![png](./images/devconfig.png "protobuf structure")](./images/devconfig.png)
* [svg](./images/devconfig.dot.svg)
* [dot](./images/devconfig.dot)

### Prerequisites

To generate the protobufs, you need to have the following prerequisites installed:

* [protoc](https://grpc.io/docs/protoc-installation/)
* [protodot](https://github.com/seamia/protodot)
* [dot](https://graphviz.org/docs/layouts/dot/)

## Updating Messages

To update the messages:

1. Edit the specific `.proto` files in the [proto](./proto) directory.
1. Run `make proto` to generate the language-specific bindings and visualizations.
1. Commit the changes to the `.proto` files and the generated files.

The Pull Request process for this repository will regenerate the language-specific bindings and visualizations,
and check that no files are missing from the repository.
