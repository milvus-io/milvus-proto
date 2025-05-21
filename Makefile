GO		?= go
PWD 	:= $(shell pwd)
GOPATH 	:= $(shell $(GO) env GOPATH)
PROTOC	:= $(shell which protoc)
PROTOC_VER := $(shell protoc --version)
INSTALL_PATH := $(PWD)/bin

all: generate-proto

build:
	@(env bash $(PWD)/scripts/core_build.sh)

generate-proto: export protoc=$(PROTOC)
generate-proto: build
	@echo "Generate proto files"
	@echo "Installing protoc-gen-go to ./bin" && GOBIN=$(INSTALL_PATH) go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.33.0
	@echo "Installing protoc-gen-go-grpc to ./bin" && GOBIN=$(INSTALL_PATH)  go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.3.0
	@(env bash $(PWD)/scripts/proto_gen_go.sh $(INSTALL_PATH))

clean:
	@echo "Cleaning up all the generated files"
	@rm -rf cmake-build