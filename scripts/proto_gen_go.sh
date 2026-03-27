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

SOURCE="${BASH_SOURCE[0]}"
while [ -h "$SOURCE" ]; do # resolve $SOURCE until the file is no longer a symlink
  DIR="$( cd -P "$( dirname "$SOURCE" )" && pwd )"
  SOURCE="$(readlink "$SOURCE")"
  [[ $SOURCE != /* ]] && SOURCE="$DIR/$SOURCE" # if $SOURCE was a relative symlink, we need to resolve it relative to the path where the symlink file was located
done
ROOT_DIR="$( cd -P "$( dirname "$SOURCE" )/.." && pwd )"

PROTOC_BIN=$ROOT_DIR/cmake-build/protobuf/protobuf-build/protoc

PROTO_DIR=$ROOT_DIR/proto
PROGRAM=$(basename "$0")
GOPATH=$(go env GOPATH)
GOOGLE_PROTO_DIR=$ROOT_DIR/cmake-build/protobuf/protobuf-src/src/
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
echo "using protoc-gen-go-vtproto: $(which protoc-gen-go-vtproto)"


pushd ${PROTO_DIR}
mkdir -p ${ROOT_DIR}/go-api/commonpb
mkdir -p ${ROOT_DIR}/go-api/schemapb
mkdir -p ${ROOT_DIR}/go-api/milvuspb
mkdir -p ${ROOT_DIR}/go-api/msgpb
mkdir -p ${ROOT_DIR}/go-api/federpb
mkdir -p ${ROOT_DIR}/go-api/tokenizerpb
mkdir -p ${ROOT_DIR}/go-api/rgpb

echo "protoc path: $PROTOC_BIN"
export protoc=$PROTOC_BIN

$protoc --version

# Common vtproto M-flags used across all proto files
VTPROTO_COMMON_OPTS="--go-vtproto_opt=Mcommon.proto=github.com/milvus-io/milvus-proto/go-api/v2/commonpb"
VTPROTO_SCHEMA_OPTS="--go-vtproto_opt=Mschema.proto=github.com/milvus-io/milvus-proto/go-api/v2/schemapb"
VTPROTO_MILVUS_OPTS="--go-vtproto_opt=Mmilvus.proto=github.com/milvus-io/milvus-proto/go-api/v2/milvuspb"
VTPROTO_MSG_OPTS="--go-vtproto_opt=Mmsg.proto=github.com/milvus-io/milvus-proto/go-api/v2/msgpb"
VTPROTO_FEDER_OPTS="--go-vtproto_opt=Mfeder.proto=github.com/milvus-io/milvus-proto/go-api/v2/federpb"
VTPROTO_RG_OPTS="--go-vtproto_opt=Mrg.proto=github.com/milvus-io/milvus-proto/go-api/v2/rgpb"
VTPROTO_FEATURES="--go-vtproto_out=features=marshal+unmarshal+size,paths=source_relative"

$protoc --proto_path="${GOOGLE_PROTO_DIR}" --proto_path=. \
    --go_opt="Mmilvus.proto=github.com/milvus-io/milvus-proto/go-api/v2/milvuspb;milvuspb" \
    --go_opt=Mcommon.proto=github.com/milvus-io/milvus-proto/go-api/v2/commonpb \
    --go_opt=Mschema.proto=github.com/milvus-io/milvus-proto/go-api/v2/schemapb \
    --go_out=paths=source_relative:${ROOT_DIR}/go-api/milvuspb \
    --go-grpc_out=require_unimplemented_servers=false,paths=source_relative:${ROOT_DIR}/go-api/milvuspb \
    ${VTPROTO_COMMON_OPTS} ${VTPROTO_SCHEMA_OPTS} ${VTPROTO_MILVUS_OPTS} ${VTPROTO_RG_OPTS} ${VTPROTO_FEDER_OPTS} ${VTPROTO_MSG_OPTS} \
    ${VTPROTO_FEATURES}:${ROOT_DIR}/go-api/milvuspb \
    milvus.proto

$protoc --proto_path="${GOOGLE_PROTO_DIR}" --proto_path=. \
    --go_opt=Mmilvus.proto=github.com/milvus-io/milvus-proto/go-api/v2/milvuspb \
    --go_opt=Mcommon.proto=github.com/milvus-io/milvus-proto/go-api/v2/commonpb \
    --go_opt="Mschema.proto=github.com/milvus-io/milvus-proto/go-api/v2/schemapb;schemapb" \
    --go_out=paths=source_relative:${ROOT_DIR}/go-api/schemapb \
    --go-grpc_out=require_unimplemented_servers=false,paths=source_relative:${ROOT_DIR}/go-api/schemapb \
    ${VTPROTO_COMMON_OPTS} ${VTPROTO_SCHEMA_OPTS} ${VTPROTO_MILVUS_OPTS} \
    ${VTPROTO_FEATURES}:${ROOT_DIR}/go-api/schemapb \
    schema.proto

$protoc --proto_path="${GOOGLE_PROTO_DIR}" --proto_path=. \
    --go_opt=Mmilvus.proto=github.com/milvus-io/milvus-proto/go-api/v2/milvuspb \
    --go_opt="Mcommon.proto=github.com/milvus-io/milvus-proto/go-api/v2/commonpb;commonpb" \
    --go_opt=Mschema.proto=github.com/milvus-io/milvus-proto/go-api/v2/schemapb \
    --go_out=paths=source_relative:${ROOT_DIR}/go-api/commonpb \
    --go-grpc_out=require_unimplemented_servers=false,paths=source_relative:${ROOT_DIR}/go-api/commonpb \
    ${VTPROTO_COMMON_OPTS} ${VTPROTO_SCHEMA_OPTS} \
    ${VTPROTO_FEATURES}:${ROOT_DIR}/go-api/commonpb \
    common.proto

$protoc --proto_path="${GOOGLE_PROTO_DIR}" --proto_path=. \
    --go_opt=Mschema.proto=github.com/milvus-io/milvus-proto/go-api/v2/schemapb \
    --go_opt=Mcommon.proto=github.com/milvus-io/milvus-proto/go-api/v2/commonpb \
    --go_opt="Mmessage.proto=github.com/milvus-io/milvus-proto/go-api/v2/msgapb;msgpb" \
    --go_out=paths=source_relative:${ROOT_DIR}/go-api/msgpb \
    --go-grpc_out=require_unimplemented_servers=false,paths=source_relative:${ROOT_DIR}/go-api/msgpb \
    ${VTPROTO_COMMON_OPTS} ${VTPROTO_SCHEMA_OPTS} \
    ${VTPROTO_FEATURES}:${ROOT_DIR}/go-api/msgpb \
    msg.proto

$protoc --proto_path="${GOOGLE_PROTO_DIR}" --proto_path=. \
    --go_opt=Mschema.proto=github.com/milvus-io/milvus-proto/go-api/v2/schemapb \
    --go_opt=Mcommon.proto=github.com/milvus-io/milvus-proto/go-api/v2/commonpb \
    --go_opt="Mmessage.proto=github.com/milvus-io/milvus-proto/go-api/v2/federpb;federpb" \
    --go_out=paths=source_relative:${ROOT_DIR}/go-api/federpb \
    --go-grpc_out=require_unimplemented_servers=false,paths=source_relative:${ROOT_DIR}/go-api/federpb \
    ${VTPROTO_COMMON_OPTS} ${VTPROTO_SCHEMA_OPTS} \
    ${VTPROTO_FEATURES}:${ROOT_DIR}/go-api/federpb \
    feder.proto

$protoc --proto_path="${GOOGLE_PROTO_DIR}" --proto_path=. \
    --go_out=paths=source_relative:${ROOT_DIR}/go-api/rgpb \
    ${VTPROTO_FEATURES}:${ROOT_DIR}/go-api/rgpb \
    rg.proto

$protoc --proto_path="${GOOGLE_PROTO_DIR}" --proto_path=. \
    --go_out=paths=source_relative:${ROOT_DIR}/go-api/tokenizerpb \
    --go-grpc_out=require_unimplemented_servers=false,paths=source_relative:${ROOT_DIR}/go-api/tokenizerpb \
    ${VTPROTO_COMMON_OPTS} ${VTPROTO_SCHEMA_OPTS} \
    ${VTPROTO_FEATURES}:${ROOT_DIR}/go-api/tokenizerpb \
    tokenizer.proto
popd
