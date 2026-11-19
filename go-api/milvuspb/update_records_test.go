package milvuspb_test

import (
	"testing"

	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/structpb"

	"github.com/milvus-io/milvus-proto/go-api/v3/commonpb"
	"github.com/milvus-io/milvus-proto/go-api/v3/milvuspb"
	"github.com/milvus-io/milvus-proto/go-api/v3/schemapb"
)

func TestUpdateRecordsSharedLiteralsRoundTrip(t *testing.T) {
	req := &milvuspb.UpdateRecordsRequest{CollectionName: "people", PrimaryKeys: &schemapb.IDs{
		IdField: &schemapb.IDs_IntId{IntId: &schemapb.LongArray{Data: []int64{100, 200, 300}}},
	}, Assignments: []*milvuspb.UpdateAssignment{
		{FieldName: "age", Literal: &schemapb.UpdateLiteral{Value: &schemapb.UpdateLiteral_ScalarValue{ScalarValue: &schemapb.ValueField{Data: &schemapb.ValueField_LongData{LongData: 0}}}}},
		{FieldName: "name", Literal: &schemapb.UpdateLiteral{Value: &schemapb.UpdateLiteral_NullValue{NullValue: structpb.NullValue_NULL_VALUE}}},
		{FieldName: "tags", Literal: &schemapb.UpdateLiteral{Value: &schemapb.UpdateLiteral_ArrayValue{ArrayValue: &schemapb.ScalarField{Data: &schemapb.ScalarField_StringData{StringData: &schemapb.StringArray{}}}}}},
		{FieldName: "note", Literal: &schemapb.UpdateLiteral{Value: &schemapb.UpdateLiteral_ScalarValue{ScalarValue: &schemapb.ValueField{Data: &schemapb.ValueField_StringData{StringData: ""}}}}},
	}}
	encoded, err := proto.Marshal(req)
	if err != nil {
		t.Fatal(err)
	}
	decoded := &milvuspb.UpdateRecordsRequest{}
	if err := proto.Unmarshal(encoded, decoded); err != nil {
		t.Fatal(err)
	}
	if !proto.Equal(req, decoded) {
		t.Fatal("PKs or shared assignments changed")
	}
	if len(decoded.Assignments) != 4 {
		t.Fatal("assignments were expanded per PK")
	}
	if _, ok := decoded.Assignments[1].Literal.Value.(*schemapb.UpdateLiteral_NullValue); !ok {
		t.Fatal("explicit NULL lost")
	}
	if decoded.Assignments[2].Literal.GetArrayValue().GetStringData() == nil {
		t.Fatal("empty array lost its type")
	}
	privilege := proto.GetExtension(req.ProtoReflect().Descriptor().Options(), commonpb.E_PrivilegeExtObj).(*commonpb.PrivilegeExt)
	if privilege.GetObjectPrivilege() != commonpb.ObjectPrivilege_PrivilegeUpsert {
		t.Fatal("update must require write authorization")
	}
}
