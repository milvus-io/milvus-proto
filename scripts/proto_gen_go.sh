#!/usr/bin/env bash

# Licensed to the LF AI & Data foundation under one
# or more contributor license agreements. See the NOTICE file
# distributed with this work for additional information
# regarding copyright ownership. The ASF licenses this file
# to you under the Apache License, Version 2.0 (the
# "License"); you may not use this file except in compliance
# with the License. You may obtain a copy of the License at
#
#     http://www.apache.org/licenses/LICENSE-2.0
#
# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS,
# WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
# See the License for the specific language governing permissions and
# limitations under the License.

SCRIPTS_DIR=$(dirname "$0")
PROTO_DIR=$SCRIPTS_DIR/../proto/
PROGRAM=$(basename "$0")
GOPATH=$(go env GOPATH)
GOOGLE_PROTO_DIR=$SCRIPTS_DIR/../cmake-build/protobuf/protobuf-src/src/
INSTALL_PATH="$1"

if [ -z $GOPATH ]; then
    printf "Error: the environment variable GOPATH is not set, please set it before running %s\n" $PROGRAM > /dev/stderr
    exit 1
fi

case ":$PATH:" in
    *":$GOPATH/bin:"*) ;;
    *) export PATH="$GOPATH/bin:$PATH";;
esac

# make sure protoc-gen-go and protoc-gen-go-grpc came from $INSTALL_PATH
export PATH=$INSTALL_PATH:${PATH}
echo "using protoc-gen-go: $(which protoc-gen-go)"
echo "using protoc-gen-go-grpc: $(which protoc-gen-go-grpc)"


pushd ${PROTO_DIR}
mkdir -p ../go-api/commonpb
mkdir -p ../go-api/schemapb
mkdir -p ../go-api/milvuspb
mkdir -p ../go-api/msgpb
mkdir -p ../go-api/federpb
mkdir -p ../go-api/tokenizerpb

echo "$(pwd)"

$protoc --version

$protoc --proto_path="${GOOGLE_PROTO_DIR}" --proto_path=. \
    --go_opt="Mmilvus.proto=github.com/milvus-io/milvus-proto/go-api/v2/milvuspb;milvuspb" \
    --go_opt=Mcommon.proto=github.com/milvus-io/milvus-proto/go-api/v2/commonpb \
    --go_opt=Mschema.proto=github.com/milvus-io/milvus-proto/go-api/v2/schemapb \
    --go_out=paths=source_relative:./../go-api/milvuspb \
    --go-grpc_out=require_unimplemented_servers=false,paths=source_relative:./../go-api/milvuspb milvus.proto

$protoc --proto_path="${GOOGLE_PROTO_DIR}" --proto_path=. \
    --go_opt=Mmilvus.proto=github.com/milvus-io/milvus-proto/go-api/v2/milvuspb \
    --go_opt=Mcommon.proto=github.com/milvus-io/milvus-proto/go-api/v2/commonpb \
    --go_opt="Mschema.proto=github.com/milvus-io/milvus-proto/go-api/v2/schemapb;schemapb" \
    --go_out=paths=source_relative:./../go-api/schemapb \
    --go-grpc_out=require_unimplemented_servers=false,paths=source_relative:./../go-api/schemapb schema.proto

$protoc --proto_path="${GOOGLE_PROTO_DIR}" --proto_path=. \
    --go_opt=Mmilvus.proto=github.com/milvus-io/milvus-proto/go-api/v2/milvuspb \
    --go_opt="Mcommon.proto=github.com/milvus-io/milvus-proto/go-api/v2/commonpb;commonpb" \
    --go_opt=Mschema.proto=github.com/milvus-io/milvus-proto/go-api/v2/schemapb \
    --go_out=paths=source_relative:./../go-api/commonpb \
    --go-grpc_out=require_unimplemented_servers=false,paths=source_relative:./../go-api/commonpb common.proto

$protoc --proto_path="${GOOGLE_PROTO_DIR}" --proto_path=. \
    --go_opt=Mschema.proto=github.com/milvus-io/milvus-proto/go-api/v2/schemapb \
    --go_opt=Mcommon.proto=github.com/milvus-io/milvus-proto/go-api/v2/commonpb \
    --go_opt="Mmessage.proto=github.com/milvus-io/milvus-proto/go-api/v2/msgapb;msgpb" \
    --go_out=paths=source_relative:./../go-api/msgpb \
    --go-grpc_out=require_unimplemented_servers=false,paths=source_relative:./../go-api/msgpb msg.proto

$protoc --proto_path="${GOOGLE_PROTO_DIR}" --proto_path=. \
    --go_opt=Mschema.proto=github.com/milvus-io/milvus-proto/go-api/v2/schemapb \
    --go_opt=Mcommon.proto=github.com/milvus-io/milvus-proto/go-api/v2/commonpb \
    --go_opt="Mmessage.proto=github.com/milvus-io/milvus-proto/go-api/v2/federpb;federpb" \
    --go_out=paths=source_relative:./../go-api/federpb \
    --go-grpc_out=require_unimplemented_servers=false,paths=source_relative:./../go-api/federpb feder.proto

$protoc --proto_path="${GOOGLE_PROTO_DIR}" --proto_path=. \
    --go_out=paths=source_relative:./../go-api/rgpb rg.proto

$protoc --proto_path="${GOOGLE_PROTO_DIR}" --proto_path=. \
    --go_out=paths=source_relative:./../go-api/tokenizerpb \
    --go-grpc_out=require_unimplemented_servers=false,paths=source_relative:./../go-api/tokenizerpb tokenizer.proto

popd
