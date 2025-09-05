.DEFAULT_GOAL := proto

PROTO_DEPS_DIR := /proto/.proto_deps
PROTOC_COMMON_FLAGS := --proto_path=$(PROTO_DEPS_DIR)/googleapis \
	--proto_path=$(PROTO_DEPS_DIR)/protoc-gen-validate \
	--proto_path=$(PROTO_DEPS_DIR)/grpc-gateway

help:
	@echo "Available targets:"
	@echo "  help              Show this help message"
	@echo "  proto             Build protobuf files (default target)"
	@echo "  proto-validate-go Generate Go validation rules for protos (should be run after proto)."
	@echo "  swagger           Generate Swagger/OpenAPI documentation"

proto-container:
	docker build -f .devcontainer/Dockerfile --build-arg PROTO_DEPS_DIR=$(PROTO_DEPS_DIR) -t eve-api-builder .

proto-diagram:
	protodot -inc /usr/local/include -src ./proto/config/devconfig.proto -output devconfig -generated ./images
	dot ./images/devconfig.dot -Tpng -o ./images/devconfig.png
	dot ./images/devconfig.dot -Tsvg -o ./images/devconfig.svg
	echo generated ./images/devconfig.*

.PHONY: proto-api-% proto proto-container proto-local swagger swagger-local

proto: proto-container
	docker run --rm --env HOME=/tmp -v $(PWD):/src -w /src -u $$(id -u) eve-api-builder make proto-local

proto-validate-go: proto-container
	docker run --rm --env HOME=/tmp -v $(PWD):/src -w /src -u $$(id -u) eve-api-builder make proto-validate-go-local

proto-local: go go-vet python proto-diagram
	@echo Done building protobuf, you may want to vendor it into your packages, e.g. pkg/pillar.
	@echo See ./go/README.md for more information.

go: PROTOC_OUT_OPTS=paths=source_relative:
go: proto-api-go
go-vet:
	go mod tidy -C go/
	go vet -C go ./...

python: proto-api-python

proto-api-%:
	rm -rf $*/*/; mkdir -p $* # building $@
	protoc -I./proto --$(*)_out=$(PROTOC_OUT_OPTS)./$* \
	$(PROTOC_COMMON_FLAGS) \
    proto/*/*.proto

proto-validate-python-local:
	@echo "Not implemented"

# For now, only register.proto is annotated with validation rules,
# later we will make this to generate validation for all protos.
proto-validate-go-local:
	protoc \
	$(PROTOC_COMMON_FLAGS) \
	--validate_out="lang=go,paths=source_relative:./go" \
	-I./proto \
	proto/register/register.proto

swagger: proto-container
	docker run --rm --env HOME=/tmp -v $(PWD):/src -w /src -u $$(id -u) eve-api-builder make swagger-local

swagger-local:
	rm -rf swagger/ && mkdir -p swagger
	protoc -I./proto \
	$(PROTOC_COMMON_FLAGS) \
    --openapiv2_opt=logtostderr=true \
    --openapiv2_opt=disable_default_errors=true \
	--openapiv2_out=swagger/ \
	proto/*/*.proto
