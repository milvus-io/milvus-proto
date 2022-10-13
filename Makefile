GO		?= go
PWD 	:= $(shell pwd)
GOPATH 	:= $(shell $(GO) env GOPATH)
PROTOC	:= $(shell which protoc)
PROTOC_VER := $(shell protoc --version)

all: generate-proto

check-protoc:
ifeq (, $(shell which protoc))
	$(error "No protoc in PATH, consider doing apt-get install protoc")
else
	@echo "using $(shell which protoc)"
endif

check-protoc-version: check-protoc
	@(env bash $(PWD)/scripts/check_protoc_version.sh)

build:
	@(env bash $(PWD)/scripts/core_build.sh)

generate-proto: check-protoc-version build
	@which protoc-gen-go 1>/dev/null || (echo "Installing protoc-gen-go" && go get github.com/golang/protobuf/protoc-gen-go@v1.5.2)
	@(env bash $(PWD)/scripts/proto_gen_go.sh)