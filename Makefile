proto-diagram:
	rm -rf images/*
	protodot -inc /usr/include -src ./proto/config/devconfig.proto -output devconfig -generated ./images && dot ./images/devconfig.dot -Tpng -o ./images/devconfig.png && echo generated ./images/devconfig.*

.PHONY: proto-api-%

proto: go python proto-diagram
	@echo Done building protobuf, you may want to vendor it into your packages, e.g. pkg/pillar.
	@echo See ./go/README.md for more information.

go: PROTOC_OUT_OPTS=paths=source_relative:
go: proto-api-go

python: proto-api-python

proto-api-%:
	rm -rf $*/*/; mkdir -p $* # building $@
	protoc -I./proto --$(*)_out=$(PROTOC_OUT_OPTS)./$* proto/*/*.proto

