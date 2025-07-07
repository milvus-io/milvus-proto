GO		?= go
PWD 	:= $(shell pwd)
GOPATH 	:= $(shell $(GO) env GOPATH)
PROTOC	:= $(shell which protoc)
PROTOC_VER := $(shell protoc --version)
INSTALL_PATH := $(PWD)/bin

# protoc-gen-go
PROTOC_GEN_GO_VERSION := 1.33.0
PROTOC_GEN_GO_OUTPUT := $(shell echo | $(INSTALL_PATH)/protoc-gen-go --version 2>/dev/null)
INSTALL_PROTOC_GEN_GO := $(findstring $(PROTOC_GEN_GO_VERSION),$(PROTOC_GEN_GO_OUTPUT))
# protoc-gen-go-grpc
PROTOC_GEN_GO_GRPC_VERSION := 1.3.0
PROTOC_GEN_GO_GRPC_OUTPUT := $(shell echo | $(INSTALL_PATH)/protoc-gen-go-grpc  --version 2>/dev/null)
INSTALL_PROTOC_GEN_GO_GRPC := $(findstring $(PROTOC_GEN_GO_GRPC_VERSION),$(PROTOC_GEN_GO_GRPC_OUTPUT))

all: generate-proto

build:
	@(env bash $(PWD)/scripts/core_build.sh)

generate-proto: get-proto-deps
generate-proto: build
	@echo "Generate proto files"
	@(env bash $(PWD)/scripts/proto_gen_go.sh $(INSTALL_PATH))

clean:
	@echo "Cleaning up all the generated files"
	@rm -rf cmake-build

get-proto-deps:
	@mkdir -p $(INSTALL_PATH) # make sure directory exists
	@if [ -z "$(INSTALL_PROTOC_GEN_GO)" ]; then \
		echo "install protoc-gen-go $(PROTOC_GEN_GO_VERSION) to $(INSTALL_PATH)" && GOBIN=$(INSTALL_PATH) go install google.golang.org/protobuf/cmd/protoc-gen-go@v$(PROTOC_GEN_GO_VERSION); \
	else \
		echo "protoc-gen-go@v$(PROTOC_GEN_GO_VERSION) already installed";\
	fi
	@if [ -z "$(INSTALL_PROTOC_GEN_GO_GRPC)" ]; then \
		echo "install protoc-gen-go-grpc $(PROTOC_GEN_GO_GRPC_VERSION) to $(INSTALL_PATH)" && GOBIN=$(INSTALL_PATH) go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v$(PROTOC_GEN_GO_GRPC_VERSION); \
	else \
		echo "protoc-gen-go-grpc@v$(PROTOC_GEN_GO_GRPC_VERSION) already installed";\
	fi
