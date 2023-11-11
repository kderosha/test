// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v4.24.4
// source: test.proto

package test

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// TesterClient is the client API for Tester service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type TesterClient interface {
	SayHello(ctx context.Context, in *Test, opts ...grpc.CallOption) (Tester_SayHelloClient, error)
}

type testerClient struct {
	cc grpc.ClientConnInterface
}

func NewTesterClient(cc grpc.ClientConnInterface) TesterClient {
	return &testerClient{cc}
}

func (c *testerClient) SayHello(ctx context.Context, in *Test, opts ...grpc.CallOption) (Tester_SayHelloClient, error) {
	stream, err := c.cc.NewStream(ctx, &Tester_ServiceDesc.Streams[0], "/proto.Tester/SayHello", opts...)
	if err != nil {
		return nil, err
	}
	x := &testerSayHelloClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type Tester_SayHelloClient interface {
	Recv() (*Test, error)
	grpc.ClientStream
}

type testerSayHelloClient struct {
	grpc.ClientStream
}

func (x *testerSayHelloClient) Recv() (*Test, error) {
	m := new(Test)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// TesterServer is the server API for Tester service.
// All implementations must embed UnimplementedTesterServer
// for forward compatibility
type TesterServer interface {
	SayHello(*Test, Tester_SayHelloServer) error
	mustEmbedUnimplementedTesterServer()
}

// UnimplementedTesterServer must be embedded to have forward compatible implementations.
type UnimplementedTesterServer struct {
}

func (UnimplementedTesterServer) SayHello(*Test, Tester_SayHelloServer) error {
	return status.Errorf(codes.Unimplemented, "method SayHello not implemented")
}
func (UnimplementedTesterServer) mustEmbedUnimplementedTesterServer() {}

// UnsafeTesterServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to TesterServer will
// result in compilation errors.
type UnsafeTesterServer interface {
	mustEmbedUnimplementedTesterServer()
}

func RegisterTesterServer(s grpc.ServiceRegistrar, srv TesterServer) {
	s.RegisterService(&Tester_ServiceDesc, srv)
}

func _Tester_SayHello_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(Test)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(TesterServer).SayHello(m, &testerSayHelloServer{stream})
}

type Tester_SayHelloServer interface {
	Send(*Test) error
	grpc.ServerStream
}

type testerSayHelloServer struct {
	grpc.ServerStream
}

func (x *testerSayHelloServer) Send(m *Test) error {
	return x.ServerStream.SendMsg(m)
}

// Tester_ServiceDesc is the grpc.ServiceDesc for Tester service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Tester_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "proto.Tester",
	HandlerType: (*TesterServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "SayHello",
			Handler:       _Tester_SayHello_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "test.proto",
}
