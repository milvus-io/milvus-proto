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
	"google.golang.org/protobuf/reflect/protodesc"
	"google.golang.org/protobuf/reflect/protoreflect"
	"google.golang.org/protobuf/types/descriptorpb"
	"google.golang.org/protobuf/types/dynamicpb"

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

func TestSnapshotSummaryUTCTimestampContract(t *testing.T) {
	const hybridTimestamp int64 = 445644800032243754
	const unixMillis int64 = 1700000000123
	original := &milvuspb.SnapshotSummary{Name: "backup", CreateTs: hybridTimestamp, CreateUtcTs: unixMillis}
	descriptor := original.ProtoReflect().Descriptor()
	fields := descriptor.Fields()
	utcField := fields.ByName("create_utc_ts")
	ttField := fields.ByName("create_ts")
	if utcField == nil || utcField.Number() != 4 || utcField.Kind() != protoreflect.Int64Kind || ttField.Number() != 2 {
		t.Fatalf("unexpected timestamp fields: %v, %v", ttField, utcField)
	}
	wire, err := proto.Marshal(original)
	if err != nil {
		t.Fatal(err)
	}
	decoded := original.ProtoReflect().New().Interface()
	if err := proto.Unmarshal(wire, decoded); err != nil || !proto.Equal(original, decoded) {
		t.Fatalf("timestamp roundtrip failed: %v, %v", err, decoded)
	}

	// Model an older client that knows only the original TT field.
	legacyFile, err := protodesc.NewFile(&descriptorpb.FileDescriptorProto{
		Name: proto.String("legacy_snapshot.proto"), Syntax: proto.String("proto3"),
		MessageType: []*descriptorpb.DescriptorProto{{
			Name:  proto.String("LegacySnapshot"),
			Field: []*descriptorpb.FieldDescriptorProto{protodesc.ToFieldDescriptorProto(ttField)},
		}},
	}, nil)
	if err != nil {
		t.Fatal(err)
	}
	legacy := dynamicpb.NewMessage(legacyFile.Messages().Get(0))
	if err := (proto.UnmarshalOptions{DiscardUnknown: true}).Unmarshal(wire, legacy); err != nil {
		t.Fatal(err)
	}
	if got := legacy.Get(legacy.Descriptor().Fields().ByName("create_ts")).Int(); got != hybridTimestamp {
		t.Fatalf("legacy TT changed: %d", got)
	}
	legacyWire, err := proto.Marshal(legacy)
	if err != nil {
		t.Fatal(err)
	}
	if err := proto.Unmarshal(legacyWire, decoded); err != nil || decoded.ProtoReflect().Get(utcField).Int() != 0 {
		t.Fatalf("older server must default UTC to zero: %v, %v", err, decoded)
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
