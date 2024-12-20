// protos/MultipleIssuer.proto

// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             (unknown)
// source: protos/multiple_issuer.proto

package protos

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.64.0 or later.
const _ = grpc.SupportPackageIsVersion9

const (
	MultipleIssuer_IssueMultipleVC_FullMethodName = "/MultipleIssuer.MultipleIssuer/IssueMultipleVC"
)

// MultipleIssuerClient is the client API for MultipleIssuer service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type MultipleIssuerClient interface {
	IssueMultipleVC(ctx context.Context, in *MsgRequestMultipleVC, opts ...grpc.CallOption) (*MsgResponseMultipleVC, error)
}

type multipleIssuerClient struct {
	cc grpc.ClientConnInterface
}

func NewMultipleIssuerClient(cc grpc.ClientConnInterface) MultipleIssuerClient {
	return &multipleIssuerClient{cc}
}

func (c *multipleIssuerClient) IssueMultipleVC(ctx context.Context, in *MsgRequestMultipleVC, opts ...grpc.CallOption) (*MsgResponseMultipleVC, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(MsgResponseMultipleVC)
	err := c.cc.Invoke(ctx, MultipleIssuer_IssueMultipleVC_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MultipleIssuerServer is the server API for MultipleIssuer service.
// All implementations must embed UnimplementedMultipleIssuerServer
// for forward compatibility.
type MultipleIssuerServer interface {
	IssueMultipleVC(context.Context, *MsgRequestMultipleVC) (*MsgResponseMultipleVC, error)
	mustEmbedUnimplementedMultipleIssuerServer()
}

// UnimplementedMultipleIssuerServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedMultipleIssuerServer struct{}

func (UnimplementedMultipleIssuerServer) IssueMultipleVC(context.Context, *MsgRequestMultipleVC) (*MsgResponseMultipleVC, error) {
	return nil, status.Errorf(codes.Unimplemented, "method IssueMultipleVC not implemented")
}
func (UnimplementedMultipleIssuerServer) mustEmbedUnimplementedMultipleIssuerServer() {}
func (UnimplementedMultipleIssuerServer) testEmbeddedByValue()                        {}

// UnsafeMultipleIssuerServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to MultipleIssuerServer will
// result in compilation errors.
type UnsafeMultipleIssuerServer interface {
	mustEmbedUnimplementedMultipleIssuerServer()
}

func RegisterMultipleIssuerServer(s grpc.ServiceRegistrar, srv MultipleIssuerServer) {
	// If the following call pancis, it indicates UnimplementedMultipleIssuerServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&MultipleIssuer_ServiceDesc, srv)
}

func _MultipleIssuer_IssueMultipleVC_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgRequestMultipleVC)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MultipleIssuerServer).IssueMultipleVC(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: MultipleIssuer_IssueMultipleVC_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MultipleIssuerServer).IssueMultipleVC(ctx, req.(*MsgRequestMultipleVC))
	}
	return interceptor(ctx, in, info, handler)
}

// MultipleIssuer_ServiceDesc is the grpc.ServiceDesc for MultipleIssuer service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var MultipleIssuer_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "MultipleIssuer.MultipleIssuer",
	HandlerType: (*MultipleIssuerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "IssueMultipleVC",
			Handler:    _MultipleIssuer_IssueMultipleVC_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "protos/multiple_issuer.proto",
}
