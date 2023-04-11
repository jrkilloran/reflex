// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v4.22.2
// source: reflex.proto

package reflexpb

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

const (
	Reflex_Stream_FullMethodName = "/reflexpb.Reflex/Stream"
)

// ReflexClient is the client API for Reflex service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ReflexClient interface {
	Stream(ctx context.Context, in *StreamRequest, opts ...grpc.CallOption) (Reflex_StreamClient, error)
}

type reflexClient struct {
	cc grpc.ClientConnInterface
}

func NewReflexClient(cc grpc.ClientConnInterface) ReflexClient {
	return &reflexClient{cc}
}

func (c *reflexClient) Stream(ctx context.Context, in *StreamRequest, opts ...grpc.CallOption) (Reflex_StreamClient, error) {
	stream, err := c.cc.NewStream(ctx, &Reflex_ServiceDesc.Streams[0], Reflex_Stream_FullMethodName, opts...)
	if err != nil {
		return nil, err
	}
	x := &reflexStreamClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type Reflex_StreamClient interface {
	Recv() (*Event, error)
	grpc.ClientStream
}

type reflexStreamClient struct {
	grpc.ClientStream
}

func (x *reflexStreamClient) Recv() (*Event, error) {
	m := new(Event)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// ReflexServer is the server API for Reflex service.
// All implementations must embed UnimplementedReflexServer
// for forward compatibility
type ReflexServer interface {
	Stream(*StreamRequest, Reflex_StreamServer) error
	mustEmbedUnimplementedReflexServer()
}

// UnimplementedReflexServer must be embedded to have forward compatible implementations.
type UnimplementedReflexServer struct {
}

func (UnimplementedReflexServer) Stream(*StreamRequest, Reflex_StreamServer) error {
	return status.Errorf(codes.Unimplemented, "method Stream not implemented")
}
func (UnimplementedReflexServer) mustEmbedUnimplementedReflexServer() {}

// UnsafeReflexServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ReflexServer will
// result in compilation errors.
type UnsafeReflexServer interface {
	mustEmbedUnimplementedReflexServer()
}

func RegisterReflexServer(s grpc.ServiceRegistrar, srv ReflexServer) {
	s.RegisterService(&Reflex_ServiceDesc, srv)
}

func _Reflex_Stream_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(StreamRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(ReflexServer).Stream(m, &reflexStreamServer{stream})
}

type Reflex_StreamServer interface {
	Send(*Event) error
	grpc.ServerStream
}

type reflexStreamServer struct {
	grpc.ServerStream
}

func (x *reflexStreamServer) Send(m *Event) error {
	return x.ServerStream.SendMsg(m)
}

// Reflex_ServiceDesc is the grpc.ServiceDesc for Reflex service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Reflex_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "reflexpb.Reflex",
	HandlerType: (*ReflexServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Stream",
			Handler:       _Reflex_Stream_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "reflex.proto",
}
