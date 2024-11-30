// protos/AtomicIssuer.proto

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.35.2
// 	protoc        (unknown)
// source: protos/atomic_issuer.proto

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

type MsgRequestAtomicVC struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Did string `protobuf:"bytes,1,opt,name=did,proto3" json:"did,omitempty"`
	Vp  string `protobuf:"bytes,2,opt,name=vp,proto3" json:"vp,omitempty"`
}

func (x *MsgRequestAtomicVC) Reset() {
	*x = MsgRequestAtomicVC{}
	mi := &file_protos_atomic_issuer_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *MsgRequestAtomicVC) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MsgRequestAtomicVC) ProtoMessage() {}

func (x *MsgRequestAtomicVC) ProtoReflect() protoreflect.Message {
	mi := &file_protos_atomic_issuer_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MsgRequestAtomicVC.ProtoReflect.Descriptor instead.
func (*MsgRequestAtomicVC) Descriptor() ([]byte, []int) {
	return file_protos_atomic_issuer_proto_rawDescGZIP(), []int{0}
}

func (x *MsgRequestAtomicVC) GetDid() string {
	if x != nil {
		return x.Did
	}
	return ""
}

func (x *MsgRequestAtomicVC) GetVp() string {
	if x != nil {
		return x.Vp
	}
	return ""
}

type MsgResponseAtomicVC struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Result string `protobuf:"bytes,1,opt,name=result,proto3" json:"result,omitempty"`
	Msg    string `protobuf:"bytes,2,opt,name=msg,proto3" json:"msg,omitempty"`
	Vcs    []*VC  `protobuf:"bytes,3,rep,name=vcs,proto3" json:"vcs,omitempty"`
}

func (x *MsgResponseAtomicVC) Reset() {
	*x = MsgResponseAtomicVC{}
	mi := &file_protos_atomic_issuer_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *MsgResponseAtomicVC) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MsgResponseAtomicVC) ProtoMessage() {}

func (x *MsgResponseAtomicVC) ProtoReflect() protoreflect.Message {
	mi := &file_protos_atomic_issuer_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MsgResponseAtomicVC.ProtoReflect.Descriptor instead.
func (*MsgResponseAtomicVC) Descriptor() ([]byte, []int) {
	return file_protos_atomic_issuer_proto_rawDescGZIP(), []int{1}
}

func (x *MsgResponseAtomicVC) GetResult() string {
	if x != nil {
		return x.Result
	}
	return ""
}

func (x *MsgResponseAtomicVC) GetMsg() string {
	if x != nil {
		return x.Msg
	}
	return ""
}

func (x *MsgResponseAtomicVC) GetVcs() []*VC {
	if x != nil {
		return x.Vcs
	}
	return nil
}

type VC struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name  string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Token string `protobuf:"bytes,2,opt,name=token,proto3" json:"token,omitempty"`
}

func (x *VC) Reset() {
	*x = VC{}
	mi := &file_protos_atomic_issuer_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *VC) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*VC) ProtoMessage() {}

func (x *VC) ProtoReflect() protoreflect.Message {
	mi := &file_protos_atomic_issuer_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use VC.ProtoReflect.Descriptor instead.
func (*VC) Descriptor() ([]byte, []int) {
	return file_protos_atomic_issuer_proto_rawDescGZIP(), []int{2}
}

func (x *VC) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *VC) GetToken() string {
	if x != nil {
		return x.Token
	}
	return ""
}

var File_protos_atomic_issuer_proto protoreflect.FileDescriptor

var file_protos_atomic_issuer_proto_rawDesc = []byte{
	0x0a, 0x1a, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2f, 0x61, 0x74, 0x6f, 0x6d, 0x69, 0x63, 0x5f,
	0x69, 0x73, 0x73, 0x75, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0c, 0x41, 0x74,
	0x6f, 0x6d, 0x69, 0x63, 0x49, 0x73, 0x73, 0x75, 0x65, 0x72, 0x22, 0x36, 0x0a, 0x12, 0x4d, 0x73,
	0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x41, 0x74, 0x6f, 0x6d, 0x69, 0x63, 0x56, 0x43,
	0x12, 0x10, 0x0a, 0x03, 0x64, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x64,
	0x69, 0x64, 0x12, 0x0e, 0x0a, 0x02, 0x76, 0x70, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02,
	0x76, 0x70, 0x22, 0x63, 0x0a, 0x13, 0x4d, 0x73, 0x67, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x41, 0x74, 0x6f, 0x6d, 0x69, 0x63, 0x56, 0x43, 0x12, 0x16, 0x0a, 0x06, 0x72, 0x65, 0x73,
	0x75, 0x6c, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x72, 0x65, 0x73, 0x75, 0x6c,
	0x74, 0x12, 0x10, 0x0a, 0x03, 0x6d, 0x73, 0x67, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03,
	0x6d, 0x73, 0x67, 0x12, 0x22, 0x0a, 0x03, 0x76, 0x63, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x10, 0x2e, 0x41, 0x74, 0x6f, 0x6d, 0x69, 0x63, 0x49, 0x73, 0x73, 0x75, 0x65, 0x72, 0x2e,
	0x56, 0x43, 0x52, 0x03, 0x76, 0x63, 0x73, 0x22, 0x2e, 0x0a, 0x02, 0x56, 0x43, 0x12, 0x12, 0x0a,
	0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x32, 0x66, 0x0a, 0x0c, 0x41, 0x74, 0x6f, 0x6d, 0x69,
	0x63, 0x49, 0x73, 0x73, 0x75, 0x65, 0x72, 0x12, 0x56, 0x0a, 0x0d, 0x49, 0x73, 0x73, 0x75, 0x65,
	0x41, 0x74, 0x6f, 0x6d, 0x69, 0x63, 0x56, 0x43, 0x12, 0x20, 0x2e, 0x41, 0x74, 0x6f, 0x6d, 0x69,
	0x63, 0x49, 0x73, 0x73, 0x75, 0x65, 0x72, 0x2e, 0x4d, 0x73, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x41, 0x74, 0x6f, 0x6d, 0x69, 0x63, 0x56, 0x43, 0x1a, 0x21, 0x2e, 0x41, 0x74, 0x6f,
	0x6d, 0x69, 0x63, 0x49, 0x73, 0x73, 0x75, 0x65, 0x72, 0x2e, 0x4d, 0x73, 0x67, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x41, 0x74, 0x6f, 0x6d, 0x69, 0x63, 0x56, 0x43, 0x22, 0x00, 0x42,
	0x11, 0x5a, 0x0f, 0x73, 0x73, 0x69, 0x2d, 0x62, 0x6f, 0x6f, 0x6b, 0x2f, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_protos_atomic_issuer_proto_rawDescOnce sync.Once
	file_protos_atomic_issuer_proto_rawDescData = file_protos_atomic_issuer_proto_rawDesc
)

func file_protos_atomic_issuer_proto_rawDescGZIP() []byte {
	file_protos_atomic_issuer_proto_rawDescOnce.Do(func() {
		file_protos_atomic_issuer_proto_rawDescData = protoimpl.X.CompressGZIP(file_protos_atomic_issuer_proto_rawDescData)
	})
	return file_protos_atomic_issuer_proto_rawDescData
}

var file_protos_atomic_issuer_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_protos_atomic_issuer_proto_goTypes = []any{
	(*MsgRequestAtomicVC)(nil),  // 0: AtomicIssuer.MsgRequestAtomicVC
	(*MsgResponseAtomicVC)(nil), // 1: AtomicIssuer.MsgResponseAtomicVC
	(*VC)(nil),                  // 2: AtomicIssuer.VC
}
var file_protos_atomic_issuer_proto_depIdxs = []int32{
	2, // 0: AtomicIssuer.MsgResponseAtomicVC.vcs:type_name -> AtomicIssuer.VC
	0, // 1: AtomicIssuer.AtomicIssuer.IssueAtomicVC:input_type -> AtomicIssuer.MsgRequestAtomicVC
	1, // 2: AtomicIssuer.AtomicIssuer.IssueAtomicVC:output_type -> AtomicIssuer.MsgResponseAtomicVC
	2, // [2:3] is the sub-list for method output_type
	1, // [1:2] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_protos_atomic_issuer_proto_init() }
func file_protos_atomic_issuer_proto_init() {
	if File_protos_atomic_issuer_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_protos_atomic_issuer_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_protos_atomic_issuer_proto_goTypes,
		DependencyIndexes: file_protos_atomic_issuer_proto_depIdxs,
		MessageInfos:      file_protos_atomic_issuer_proto_msgTypes,
	}.Build()
	File_protos_atomic_issuer_proto = out.File
	file_protos_atomic_issuer_proto_rawDesc = nil
	file_protos_atomic_issuer_proto_goTypes = nil
	file_protos_atomic_issuer_proto_depIdxs = nil
}
