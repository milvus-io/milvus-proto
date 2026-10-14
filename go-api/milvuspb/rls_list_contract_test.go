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

func TestListRLSPrincipalsCountContract(t *testing.T) {
	request := &milvuspb.ListRLSPrincipalsRequest{}
	if request.GetCount() {
		t.Fatal("an omitted count option must preserve names enumeration")
	}
	for _, c := range []struct {
		message proto.Message
		name    protoreflect.Name
		number  protoreflect.FieldNumber
		kind    protoreflect.Kind
	}{
		{request, "base", 1, protoreflect.MessageKind},
		{request, "db_name", 2, protoreflect.StringKind},
		{request, "collection_name", 3, protoreflect.StringKind},
		{request, "count", 4, protoreflect.BoolKind},
		{&milvuspb.ListRLSPrincipalsResponse{}, "status", 1, protoreflect.MessageKind},
		{&milvuspb.ListRLSPrincipalsResponse{}, "principal_names", 2, protoreflect.StringKind},
		{&milvuspb.ListRLSPrincipalsResponse{}, "db_name", 3, protoreflect.StringKind},
		{&milvuspb.ListRLSPrincipalsResponse{}, "collection_name", 4, protoreflect.StringKind},
		{&milvuspb.ListRLSPrincipalsResponse{}, "principal_cnt", 5, protoreflect.Int64Kind},
	} {
		field := c.message.ProtoReflect().Descriptor().Fields().ByNumber(c.number)
		if field == nil || field.Name() != c.name || field.Kind() != c.kind {
			t.Fatalf("%T field %d must remain %s (%s), got %v", c.message, c.number, c.name, c.kind, field)
		}
	}

	options := request.ProtoReflect().Descriptor().Options().(*descriptorpb.MessageOptions)
	privilege := proto.GetExtension(options, commonpb.E_PrivilegeExtObj).(*commonpb.PrivilegeExt)
	if privilege.GetObjectType() != commonpb.ObjectType_Collection ||
		privilege.GetObjectPrivilege() != commonpb.ObjectPrivilege_PrivilegeViewRLS ||
		privilege.GetObjectNameIndex() != 3 {
		t.Fatalf("count-only requests must retain collection ViewRLS authorization: %v", privilege)
	}

	for _, message := range []proto.Message{
		&milvuspb.ListRLSPrincipalsRequest{DbName: "db", CollectionName: "collection", Count: true},
		&milvuspb.ListRLSPrincipalsResponse{PrincipalNames: []string{"alice", "bob"}, PrincipalCnt: 2},
		&milvuspb.ListRLSPrincipalsResponse{PrincipalCnt: 1 << 33},
	} {
		wire, err := proto.Marshal(message)
		if err != nil {
			t.Fatal(err)
		}
		decoded := message.ProtoReflect().New().Interface()
		if err := proto.Unmarshal(wire, decoded); err != nil {
			t.Fatal(err)
		}
		if !proto.Equal(message, decoded) {
			t.Fatalf("round trip changed %v to %v", message, decoded)
		}
	}
}
