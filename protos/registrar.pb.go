// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.35.2
// 	protoc        (unknown)
// source: protos/registrar.proto

package protos

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

type RegistrarRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Did         string `protobuf:"bytes,1,opt,name=did,proto3" json:"did,omitempty"`
	DidDocument string `protobuf:"bytes,2,opt,name=didDocument,proto3" json:"didDocument,omitempty"`
}

func (x *RegistrarRequest) Reset() {
	*x = RegistrarRequest{}
	mi := &file_protos_registrar_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *RegistrarRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RegistrarRequest) ProtoMessage() {}

func (x *RegistrarRequest) ProtoReflect() protoreflect.Message {
	mi := &file_protos_registrar_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RegistrarRequest.ProtoReflect.Descriptor instead.
func (*RegistrarRequest) Descriptor() ([]byte, []int) {
	return file_protos_registrar_proto_rawDescGZIP(), []int{0}
}

func (x *RegistrarRequest) GetDid() string {
	if x != nil {
		return x.Did
	}
	return ""
}

func (x *RegistrarRequest) GetDidDocument() string {
	if x != nil {
		return x.DidDocument
	}
	return ""
}

type RegistrarResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Result string `protobuf:"bytes,1,opt,name=result,proto3" json:"result,omitempty"`
}

func (x *RegistrarResponse) Reset() {
	*x = RegistrarResponse{}
	mi := &file_protos_registrar_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *RegistrarResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RegistrarResponse) ProtoMessage() {}

func (x *RegistrarResponse) ProtoReflect() protoreflect.Message {
	mi := &file_protos_registrar_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RegistrarResponse.ProtoReflect.Descriptor instead.
func (*RegistrarResponse) Descriptor() ([]byte, []int) {
	return file_protos_registrar_proto_rawDescGZIP(), []int{1}
}

func (x *RegistrarResponse) GetResult() string {
	if x != nil {
		return x.Result
	}
	return ""
}

var File_protos_registrar_proto protoreflect.FileDescriptor

var file_protos_registrar_proto_rawDesc = []byte{
	0x0a, 0x16, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2f, 0x72, 0x65, 0x67, 0x69, 0x73, 0x74, 0x72,
	0x61, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x09, 0x72, 0x65, 0x67, 0x69, 0x73, 0x74,
	0x72, 0x61, 0x72, 0x22, 0x46, 0x0a, 0x10, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x72, 0x61, 0x72,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x10, 0x0a, 0x03, 0x64, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x64, 0x69, 0x64, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x69, 0x64,
	0x44, 0x6f, 0x63, 0x75, 0x6d, 0x65, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b,
	0x64, 0x69, 0x64, 0x44, 0x6f, 0x63, 0x75, 0x6d, 0x65, 0x6e, 0x74, 0x22, 0x2b, 0x0a, 0x11, 0x52,
	0x65, 0x67, 0x69, 0x73, 0x74, 0x72, 0x61, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x16, 0x0a, 0x06, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x06, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x32, 0x57, 0x0a, 0x09, 0x52, 0x65, 0x67, 0x69,
	0x73, 0x74, 0x72, 0x61, 0x72, 0x12, 0x4a, 0x0a, 0x0b, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65,
	0x72, 0x44, 0x69, 0x64, 0x12, 0x1b, 0x2e, 0x72, 0x65, 0x67, 0x69, 0x73, 0x74, 0x72, 0x61, 0x72,
	0x2e, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x72, 0x61, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x1c, 0x2e, 0x72, 0x65, 0x67, 0x69, 0x73, 0x74, 0x72, 0x61, 0x72, 0x2e, 0x52, 0x65,
	0x67, 0x69, 0x73, 0x74, 0x72, 0x61, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22,
	0x00, 0x42, 0x11, 0x5a, 0x0f, 0x73, 0x73, 0x69, 0x2d, 0x62, 0x6f, 0x6f, 0x6b, 0x2f, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_protos_registrar_proto_rawDescOnce sync.Once
	file_protos_registrar_proto_rawDescData = file_protos_registrar_proto_rawDesc
)

func file_protos_registrar_proto_rawDescGZIP() []byte {
	file_protos_registrar_proto_rawDescOnce.Do(func() {
		file_protos_registrar_proto_rawDescData = protoimpl.X.CompressGZIP(file_protos_registrar_proto_rawDescData)
	})
	return file_protos_registrar_proto_rawDescData
}

var file_protos_registrar_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_protos_registrar_proto_goTypes = []any{
	(*RegistrarRequest)(nil),  // 0: registrar.RegistrarRequest
	(*RegistrarResponse)(nil), // 1: registrar.RegistrarResponse
}
var file_protos_registrar_proto_depIdxs = []int32{
	0, // 0: registrar.Registrar.RegisterDid:input_type -> registrar.RegistrarRequest
	1, // 1: registrar.Registrar.RegisterDid:output_type -> registrar.RegistrarResponse
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_protos_registrar_proto_init() }
func file_protos_registrar_proto_init() {
	if File_protos_registrar_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_protos_registrar_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_protos_registrar_proto_goTypes,
		DependencyIndexes: file_protos_registrar_proto_depIdxs,
		MessageInfos:      file_protos_registrar_proto_msgTypes,
	}.Build()
	File_protos_registrar_proto = out.File
	file_protos_registrar_proto_rawDesc = nil
	file_protos_registrar_proto_goTypes = nil
	file_protos_registrar_proto_depIdxs = nil
}
