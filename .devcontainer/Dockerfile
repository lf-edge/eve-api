# syntax=docker/dockerfile:1.6.0
FROM alpine:3.18
SHELL ["/bin/ash", "-eo", "pipefail", "-c"]

ARG TARGETOS=linux
ARG TARGETARCH
ARG PROTOC_GEN_GO_VERSION=1.31.0
ARG PROTOC_VERSION=23.4
ARG OPENAPI_VERSION=91958df0371da5c71794adc92e21cf8fed58df97 # v2.27.2
ARG GRPC_GATEWAY_VERSION=91958df0371da5c71794adc92e21cf8fed58df97 # 2.27.2
ARG PROTO_GEN_VALIDATE_VERSION=7b06248484ceeaa947e93ca2747eccf336a88ecc # v1.2.1
ARG GOOGLEAPIS_VERSION=376467058c288ad34dd7aafa892a95883e4acd0c # master
ARG PROTO_DEPS_DIR=/proto/.proto_deps

RUN apk add --no-cache bash curl git make graphviz ttf-freefont


RUN arch="$(uname -m)" && \
    case "$arch" in \
    aarch64|arm64) arch="arm64" ;; \
    riscv64) arch="riscv_64" ;; \
    x86_64) arch="amd64";; \
    *) echo "Unsupported architecture: $arch" && exit 1 ;; \
    esac && \
    \
    echo "Installing Go..." && \
    echo "https://go.dev/dl/go1.23.4.${TARGETOS}-${arch}.tar.gz" && \
    curl -fsSL -o go1.23.4.${TARGETOS}-${arch}.tar.gz https://go.dev/dl/go1.23.4.${TARGETOS}-${arch}.tar.gz && \
    tar -C /usr/local -xzf go1.23.4.${TARGETOS}-${arch}.tar.gz && \
    rm go1.23.4.${TARGETOS}-${arch}.tar.gz && \
    ln -sf /usr/local/go/bin/go /usr/bin/go && \
    ln -sf /usr/local/go/bin/gofmt /usr/bin/gofmt

RUN arch="$(uname -m)" && \
    case "$arch" in \
    aarch64|arm64) arch="aarch_64" ;; \
    riscv64) arch="riscv_64" ;; \
    x86_64) ;; \
    *) echo "Unsupported architecture: $arch" && exit 1 ;; \
    esac && \
    \
    echo "Installing protoc..." && \
    curl -fsSL -o protobuf.zip https://github.com/protocolbuffers/protobuf/releases/download/v${PROTOC_VERSION}/protoc-${PROTOC_VERSION}-${TARGETOS}-${arch}.zip; \
    unzip protobuf.zip -d /usr/local; \
    rm -f protobuf.zip

# Download googleapis proto
RUN mkdir -p "${PROTO_DEPS_DIR}" && \
    curl -L https://github.com/googleapis/googleapis/archive/${GOOGLEAPIS_VERSION}.tar.gz | \
    tar -xz -C "${PROTO_DEPS_DIR}" && \
    mv "${PROTO_DEPS_DIR}/googleapis-${GOOGLEAPIS_VERSION}" "${PROTO_DEPS_DIR}/googleapis"


# Download protoc-gen-validate proto
RUN mkdir -p "${PROTO_DEPS_DIR}" && \
    curl -L https://github.com/bufbuild/protoc-gen-validate/archive/${PROTO_GEN_VALIDATE_VERSION}.tar.gz | \
    tar -xz -C "${PROTO_DEPS_DIR}" && \
    mv "${PROTO_DEPS_DIR}/protoc-gen-validate-${PROTO_GEN_VALIDATE_VERSION}" "${PROTO_DEPS_DIR}/protoc-gen-validate"

# Download grpc-gateway
RUN mkdir -p "${PROTO_DEPS_DIR}" && \
    curl -L https://github.com/grpc-ecosystem/grpc-gateway/archive/${GRPC_GATEWAY_VERSION}.tar.gz | \
    tar -xz -C "${PROTO_DEPS_DIR}" && \
    mv "${PROTO_DEPS_DIR}/grpc-gateway-${GRPC_GATEWAY_VERSION}" "${PROTO_DEPS_DIR}/grpc-gateway"

RUN git clone https://github.com/seamia/protodot.git /protodot
RUN go build -C /protodot -o /usr/local/bin/protodot .

RUN mkdir -p $(go env GOPATH)
RUN go install google.golang.org/protobuf/cmd/protoc-gen-go@v${PROTOC_GEN_GO_VERSION} && mv /root/go/bin/protoc-gen-go /usr/local/bin/protoc-gen-go
RUN go install github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2@${OPENAPI_VERSION} && mv /root/go/bin/protoc-gen-openapiv2 /usr/local/bin/protoc-gen-openapiv2
RUN go install github.com/envoyproxy/protoc-gen-validate@${PROTO_GEN_VALIDATE_VERSION} && mv /root/go/bin/protoc-gen-validate /usr/local/bin/protoc-gen-validate
