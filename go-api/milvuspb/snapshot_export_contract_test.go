// Licensed to the LF AI & Data foundation under one
// or more contributor license agreements. See the NOTICE file
// distributed with this work for additional information
// regarding copyright ownership. The ASF licenses this file
// to you under the Apache License, Version 2.0 (the
// "License"); you may not use this file except in compliance
// with the License. You may obtain a copy of the License at
//
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package milvuspb_test

import (
	"testing"

	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/reflect/protoreflect"
	"google.golang.org/protobuf/types/descriptorpb"

	"github.com/milvus-io/milvus-proto/go-api/v3/commonpb"
	"github.com/milvus-io/milvus-proto/go-api/v3/milvuspb"
)

func TestExportSnapshotResponseContract(t *testing.T) {
	fields := (&milvuspb.ExportSnapshotResponse{}).ProtoReflect().Descriptor().Fields()
	metadataURI := fields.ByNumber(protoreflect.FieldNumber(2))
	if metadataURI == nil || metadataURI.Name() != "snapshot_metadata_uri" {
		t.Fatalf("field 2 must remain snapshot_metadata_uri, got %v", metadataURI)
	}
	options, ok := metadataURI.Options().(*descriptorpb.FieldOptions)
	if !ok || !options.GetDeprecated() {
		t.Fatal("snapshot_metadata_uri must remain deprecated")
	}
	jobID := fields.ByNumber(protoreflect.FieldNumber(3))
	if jobID == nil || jobID.Name() != "job_id" {
		t.Fatalf("field 3 must be job_id, got %v", jobID)
	}
}

func TestGetExportSnapshotStateContract(t *testing.T) {
	service := milvuspb.File_milvus_proto.Services().ByName("MilvusService")
	if service == nil || service.Methods().ByName("GetExportSnapshotState") == nil {
		t.Fatal("MilvusService.GetExportSnapshotState is not registered")
	}

	options, ok := (&milvuspb.GetExportSnapshotStateRequest{}).
		ProtoReflect().Descriptor().Options().(*descriptorpb.MessageOptions)
	if !ok || !proto.HasExtension(options, commonpb.E_PrivilegeExtObj) {
		t.Fatal("GetExportSnapshotStateRequest is missing privilege options")
	}
	extension := proto.GetExtension(options, commonpb.E_PrivilegeExtObj)
	privilege, ok := extension.(*commonpb.PrivilegeExt)
	if !ok {
		t.Fatalf("unexpected privilege extension type %T", extension)
	}
	if privilege.GetObjectType() != commonpb.ObjectType_Global ||
		privilege.GetObjectPrivilege() != commonpb.ObjectPrivilege_PrivilegeExportSnapshot ||
		privilege.GetObjectNameIndex() != -1 {
		t.Fatalf("unexpected export state privilege option: %v", privilege)
	}
}
