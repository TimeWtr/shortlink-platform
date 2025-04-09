.PHONY: gen_http
gen_http:
	@cd api/proto && protoc --proto_path=$(GOPATH)/pkg/protobuf/src \
					   --proto_path=$(GOPATH)/pkg/googleapis \
                       --proto_path=./intr/v1 \
                       --plugin=protoc-gen-go_grpc=$(GOPATH)/bin/protoc-gen-go-grpc \
                       --go_out=./gen/ \
                       --go_grpc_out=./gen intr/v1/generate_admin.proto

.PHONY: gen_rpc
gen_rpc:
	@cd api/proto && protoc --proto_path=./intr/v1 \
					   --plugin=protoc-gen-go_grpc=$(GOPATH)/bin/protoc-gen-go-grpc \
					   --go_out=./gen/ \
					   --go_grpc_out=./gen intr/v1/generate.proto