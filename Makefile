GO		?= go
PWD 	:= $(shell pwd)
GOPATH 	:= $(shell $(GO) env GOPATH)
PROTOC	:= $(shell which protoc)
PROTOC_VER := $(shell protoc --version)

all: generate-proto

build:
	@(env bash $(PWD)/scripts/core_build.sh)

generate-proto: export protoc:=${PWD}/cmake-build/protobuf/protobuf-build/protoc
generate-proto: build
	@which protoc-gen-go 1>/dev/null || (echo "Installing protoc-gen-go" && go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.33.0)
	@which protoc-gen-go-grpc 1>/dev/null || (echo "Installing protoc-gen-go-grpc" && go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.3.0)
	@(env bash $(PWD)/scripts/proto_gen_go.sh)

clean:
	@echo "Cleaning up all the generated files"
	@rm -rf cmake-build