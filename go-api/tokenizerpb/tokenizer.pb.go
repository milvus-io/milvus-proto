// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.33.0
// 	protoc        v3.21.4
// source: tokenizer.proto

package schemapb

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type TokenizationRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Text       string                           `protobuf:"bytes,1,opt,name=text,proto3" json:"text,omitempty"`
	Parameters []*TokenizationRequest_Parameter `protobuf:"bytes,2,rep,name=parameters,proto3" json:"parameters,omitempty"`
}

func (x *TokenizationRequest) Reset() {
	*x = TokenizationRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tokenizer_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TokenizationRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TokenizationRequest) ProtoMessage() {}

func (x *TokenizationRequest) ProtoReflect() protoreflect.Message {
	mi := &file_tokenizer_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TokenizationRequest.ProtoReflect.Descriptor instead.
func (*TokenizationRequest) Descriptor() ([]byte, []int) {
	return file_tokenizer_proto_rawDescGZIP(), []int{0}
}

func (x *TokenizationRequest) GetText() string {
	if x != nil {
		return x.Text
	}
	return ""
}

func (x *TokenizationRequest) GetParameters() []*TokenizationRequest_Parameter {
	if x != nil {
		return x.Parameters
	}
	return nil
}

type TokenizationResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Tokens []*Token `protobuf:"bytes,1,rep,name=tokens,proto3" json:"tokens,omitempty"`
}

func (x *TokenizationResponse) Reset() {
	*x = TokenizationResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tokenizer_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TokenizationResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TokenizationResponse) ProtoMessage() {}

func (x *TokenizationResponse) ProtoReflect() protoreflect.Message {
	mi := &file_tokenizer_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TokenizationResponse.ProtoReflect.Descriptor instead.
func (*TokenizationResponse) Descriptor() ([]byte, []int) {
	return file_tokenizer_proto_rawDescGZIP(), []int{1}
}

func (x *TokenizationResponse) GetTokens() []*Token {
	if x != nil {
		return x.Tokens
	}
	return nil
}

type Token struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Text           string `protobuf:"bytes,1,opt,name=text,proto3" json:"text,omitempty"`
	OffsetFrom     int32  `protobuf:"varint,2,opt,name=offset_from,json=offsetFrom,proto3" json:"offset_from,omitempty"`
	OffsetTo       int32  `protobuf:"varint,3,opt,name=offset_to,json=offsetTo,proto3" json:"offset_to,omitempty"`
	Position       int32  `protobuf:"varint,4,opt,name=position,proto3" json:"position,omitempty"`
	PositionLength int32  `protobuf:"varint,5,opt,name=position_length,json=positionLength,proto3" json:"position_length,omitempty"`
}

func (x *Token) Reset() {
	*x = Token{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tokenizer_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Token) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Token) ProtoMessage() {}

func (x *Token) ProtoReflect() protoreflect.Message {
	mi := &file_tokenizer_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Token.ProtoReflect.Descriptor instead.
func (*Token) Descriptor() ([]byte, []int) {
	return file_tokenizer_proto_rawDescGZIP(), []int{2}
}

func (x *Token) GetText() string {
	if x != nil {
		return x.Text
	}
	return ""
}

func (x *Token) GetOffsetFrom() int32 {
	if x != nil {
		return x.OffsetFrom
	}
	return 0
}

func (x *Token) GetOffsetTo() int32 {
	if x != nil {
		return x.OffsetTo
	}
	return 0
}

func (x *Token) GetPosition() int32 {
	if x != nil {
		return x.Position
	}
	return 0
}

func (x *Token) GetPositionLength() int32 {
	if x != nil {
		return x.PositionLength
	}
	return 0
}

type TokenizationRequest_Parameter struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Key    string   `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
	Values []string `protobuf:"bytes,2,rep,name=values,proto3" json:"values,omitempty"`
}

func (x *TokenizationRequest_Parameter) Reset() {
	*x = TokenizationRequest_Parameter{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tokenizer_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TokenizationRequest_Parameter) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TokenizationRequest_Parameter) ProtoMessage() {}

func (x *TokenizationRequest_Parameter) ProtoReflect() protoreflect.Message {
	mi := &file_tokenizer_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TokenizationRequest_Parameter.ProtoReflect.Descriptor instead.
func (*TokenizationRequest_Parameter) Descriptor() ([]byte, []int) {
	return file_tokenizer_proto_rawDescGZIP(), []int{0, 0}
}

func (x *TokenizationRequest_Parameter) GetKey() string {
	if x != nil {
		return x.Key
	}
	return ""
}

func (x *TokenizationRequest_Parameter) GetValues() []string {
	if x != nil {
		return x.Values
	}
	return nil
}

var File_tokenizer_proto protoreflect.FileDescriptor

var file_tokenizer_proto_rawDesc = []byte{
	0x0a, 0x0f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x69, 0x7a, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x16, 0x6d, 0x69, 0x6c, 0x76, 0x75, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e,
	0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x69, 0x7a, 0x65, 0x72, 0x22, 0xb7, 0x01, 0x0a, 0x13, 0x54, 0x6f,
	0x6b, 0x65, 0x6e, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x65, 0x78, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x04, 0x74, 0x65, 0x78, 0x74, 0x12, 0x55, 0x0a, 0x0a, 0x70, 0x61, 0x72, 0x61, 0x6d, 0x65, 0x74,
	0x65, 0x72, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x35, 0x2e, 0x6d, 0x69, 0x6c, 0x76,
	0x75, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x69, 0x7a,
	0x65, 0x72, 0x2e, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x2e, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x65, 0x74, 0x65, 0x72,
	0x52, 0x0a, 0x70, 0x61, 0x72, 0x61, 0x6d, 0x65, 0x74, 0x65, 0x72, 0x73, 0x1a, 0x35, 0x0a, 0x09,
	0x50, 0x61, 0x72, 0x61, 0x6d, 0x65, 0x74, 0x65, 0x72, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x16, 0x0a, 0x06, 0x76,
	0x61, 0x6c, 0x75, 0x65, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x09, 0x52, 0x06, 0x76, 0x61, 0x6c,
	0x75, 0x65, 0x73, 0x22, 0x4d, 0x0a, 0x14, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x69, 0x7a, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x35, 0x0a, 0x06, 0x74,
	0x6f, 0x6b, 0x65, 0x6e, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1d, 0x2e, 0x6d, 0x69,
	0x6c, 0x76, 0x75, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x74, 0x6f, 0x6b, 0x65, 0x6e,
	0x69, 0x7a, 0x65, 0x72, 0x2e, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x52, 0x06, 0x74, 0x6f, 0x6b, 0x65,
	0x6e, 0x73, 0x22, 0x9e, 0x01, 0x0a, 0x05, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x12, 0x0a, 0x04,
	0x74, 0x65, 0x78, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x74, 0x65, 0x78, 0x74,
	0x12, 0x1f, 0x0a, 0x0b, 0x6f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x5f, 0x66, 0x72, 0x6f, 0x6d, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0a, 0x6f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x46, 0x72, 0x6f,
	0x6d, 0x12, 0x1b, 0x0a, 0x09, 0x6f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x5f, 0x74, 0x6f, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x6f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x54, 0x6f, 0x12, 0x1a,
	0x0a, 0x08, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x04, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x08, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x27, 0x0a, 0x0f, 0x70, 0x6f,
	0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x6c, 0x65, 0x6e, 0x67, 0x74, 0x68, 0x18, 0x05, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x0e, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x4c, 0x65, 0x6e,
	0x67, 0x74, 0x68, 0x32, 0x74, 0x0a, 0x09, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x69, 0x7a, 0x65, 0x72,
	0x12, 0x67, 0x0a, 0x08, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x69, 0x7a, 0x65, 0x12, 0x2b, 0x2e, 0x6d,
	0x69, 0x6c, 0x76, 0x75, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x74, 0x6f, 0x6b, 0x65,
	0x6e, 0x69, 0x7a, 0x65, 0x72, 0x2e, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x69, 0x7a, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x2c, 0x2e, 0x6d, 0x69, 0x6c, 0x76,
	0x75, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x69, 0x7a,
	0x65, 0x72, 0x2e, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x36, 0x5a, 0x34, 0x67, 0x69, 0x74,
	0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6d, 0x69, 0x6c, 0x76, 0x75, 0x73, 0x2d, 0x69,
	0x6f, 0x2f, 0x6d, 0x69, 0x6c, 0x76, 0x75, 0x73, 0x2d, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x67,
	0x6f, 0x2d, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x32, 0x2f, 0x73, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x70,
	0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_tokenizer_proto_rawDescOnce sync.Once
	file_tokenizer_proto_rawDescData = file_tokenizer_proto_rawDesc
)

func file_tokenizer_proto_rawDescGZIP() []byte {
	file_tokenizer_proto_rawDescOnce.Do(func() {
		file_tokenizer_proto_rawDescData = protoimpl.X.CompressGZIP(file_tokenizer_proto_rawDescData)
	})
	return file_tokenizer_proto_rawDescData
}

var file_tokenizer_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_tokenizer_proto_goTypes = []interface{}{
	(*TokenizationRequest)(nil),           // 0: milvus.proto.tokenizer.TokenizationRequest
	(*TokenizationResponse)(nil),          // 1: milvus.proto.tokenizer.TokenizationResponse
	(*Token)(nil),                         // 2: milvus.proto.tokenizer.Token
	(*TokenizationRequest_Parameter)(nil), // 3: milvus.proto.tokenizer.TokenizationRequest.Parameter
}
var file_tokenizer_proto_depIdxs = []int32{
	3, // 0: milvus.proto.tokenizer.TokenizationRequest.parameters:type_name -> milvus.proto.tokenizer.TokenizationRequest.Parameter
	2, // 1: milvus.proto.tokenizer.TokenizationResponse.tokens:type_name -> milvus.proto.tokenizer.Token
	0, // 2: milvus.proto.tokenizer.Tokenizer.Tokenize:input_type -> milvus.proto.tokenizer.TokenizationRequest
	1, // 3: milvus.proto.tokenizer.Tokenizer.Tokenize:output_type -> milvus.proto.tokenizer.TokenizationResponse
	3, // [3:4] is the sub-list for method output_type
	2, // [2:3] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_tokenizer_proto_init() }
func file_tokenizer_proto_init() {
	if File_tokenizer_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_tokenizer_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TokenizationRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_tokenizer_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TokenizationResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_tokenizer_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Token); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_tokenizer_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TokenizationRequest_Parameter); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_tokenizer_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_tokenizer_proto_goTypes,
		DependencyIndexes: file_tokenizer_proto_depIdxs,
		MessageInfos:      file_tokenizer_proto_msgTypes,
	}.Build()
	File_tokenizer_proto = out.File
	file_tokenizer_proto_rawDesc = nil
	file_tokenizer_proto_goTypes = nil
	file_tokenizer_proto_depIdxs = nil
}
