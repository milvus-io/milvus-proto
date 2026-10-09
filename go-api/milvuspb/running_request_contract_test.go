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

// The two running-request RPCs are guarded by different privileges: listing is
// read-only, canceling stops other users' requests. A request message pointing
// at the wrong one would still compile, so the binding is checked here.
func TestRunningRequestPrivilegeContract(t *testing.T) {
	cases := []struct {
		method    string
		request   proto.Message
		privilege commonpb.ObjectPrivilege
	}{
		{"ListRunningRequests", &milvuspb.ListRunningRequestsRequest{}, commonpb.ObjectPrivilege_PrivilegeListRunningRequests},
		{"CancelRequests", &milvuspb.CancelRequestsRequest{}, commonpb.ObjectPrivilege_PrivilegeCancelRequests},
	}
	service := milvuspb.File_milvus_proto.Services().ByName("MilvusService")
	if service == nil {
		t.Fatal("MilvusService is not registered")
	}
	for _, c := range cases {
		if service.Methods().ByName(protoreflect.Name(c.method)) == nil {
			t.Fatalf("MilvusService.%s is not registered", c.method)
		}
		options, ok := c.request.ProtoReflect().Descriptor().Options().(*descriptorpb.MessageOptions)
		if !ok || !proto.HasExtension(options, commonpb.E_PrivilegeExtObj) {
			t.Fatalf("%s request is missing privilege options", c.method)
		}
		privilege, ok := proto.GetExtension(options, commonpb.E_PrivilegeExtObj).(*commonpb.PrivilegeExt)
		if !ok {
			t.Fatalf("%s request has an unexpected privilege extension type", c.method)
		}
		if privilege.GetObjectType() != commonpb.ObjectType_Global ||
			privilege.GetObjectPrivilege() != c.privilege ||
			privilege.GetObjectNameIndex() != -1 {
			t.Fatalf("unexpected %s privilege option: %v", c.method, privilege)
		}
	}
}
