module validation-example

go 1.24.0

replace github.com/lf-edge/eve-api/go => ../go

require github.com/lf-edge/eve-api/go v0.0.0-00010101000000-000000000000

require (
	github.com/envoyproxy/protoc-gen-validate v1.2.1 // indirect
	github.com/grpc-ecosystem/grpc-gateway/v2 v2.27.2 // indirect
	google.golang.org/genproto/googleapis/api v0.0.0-20250826171959-ef028d996bc1 // indirect
	google.golang.org/protobuf v1.36.8 // indirect
)
