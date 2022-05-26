// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             (unknown)
// source: proto-public/pbserverdiscovery/serverdiscovery.proto

package pbserverdiscovery

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

// ServerDiscoveryServiceClient is the client API for ServerDiscoveryService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ServerDiscoveryServiceClient interface {
	// WatchServers will stream back sets of ready servers as they change such as
	// when new servers are added or older ones removed. A ready server is one that
	// should be considered ready for sending general RPC requests towards that would
	// catalog queries, xDS proxy configurations and similar services.
	WatchServers(ctx context.Context, in *WatchServersRequest, opts ...grpc.CallOption) (ServerDiscoveryService_WatchServersClient, error)
}

type serverDiscoveryServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewServerDiscoveryServiceClient(cc grpc.ClientConnInterface) ServerDiscoveryServiceClient {
	return &serverDiscoveryServiceClient{cc}
}

func (c *serverDiscoveryServiceClient) WatchServers(ctx context.Context, in *WatchServersRequest, opts ...grpc.CallOption) (ServerDiscoveryService_WatchServersClient, error) {
	stream, err := c.cc.NewStream(ctx, &ServerDiscoveryService_ServiceDesc.Streams[0], "/serverdiscovery.ServerDiscoveryService/WatchServers", opts...)
	if err != nil {
		return nil, err
	}
	x := &serverDiscoveryServiceWatchServersClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type ServerDiscoveryService_WatchServersClient interface {
	Recv() (*WatchServersResponse, error)
	grpc.ClientStream
}

type serverDiscoveryServiceWatchServersClient struct {
	grpc.ClientStream
}

func (x *serverDiscoveryServiceWatchServersClient) Recv() (*WatchServersResponse, error) {
	m := new(WatchServersResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// ServerDiscoveryServiceServer is the server API for ServerDiscoveryService service.
// All implementations should embed UnimplementedServerDiscoveryServiceServer
// for forward compatibility
type ServerDiscoveryServiceServer interface {
	// WatchServers will stream back sets of ready servers as they change such as
	// when new servers are added or older ones removed. A ready server is one that
	// should be considered ready for sending general RPC requests towards that would
	// catalog queries, xDS proxy configurations and similar services.
	WatchServers(*WatchServersRequest, ServerDiscoveryService_WatchServersServer) error
}

// UnimplementedServerDiscoveryServiceServer should be embedded to have forward compatible implementations.
type UnimplementedServerDiscoveryServiceServer struct {
}

func (UnimplementedServerDiscoveryServiceServer) WatchServers(*WatchServersRequest, ServerDiscoveryService_WatchServersServer) error {
	return status.Errorf(codes.Unimplemented, "method WatchServers not implemented")
}

// UnsafeServerDiscoveryServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ServerDiscoveryServiceServer will
// result in compilation errors.
type UnsafeServerDiscoveryServiceServer interface {
	mustEmbedUnimplementedServerDiscoveryServiceServer()
}

func RegisterServerDiscoveryServiceServer(s grpc.ServiceRegistrar, srv ServerDiscoveryServiceServer) {
	s.RegisterService(&ServerDiscoveryService_ServiceDesc, srv)
}

func _ServerDiscoveryService_WatchServers_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(WatchServersRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(ServerDiscoveryServiceServer).WatchServers(m, &serverDiscoveryServiceWatchServersServer{stream})
}

type ServerDiscoveryService_WatchServersServer interface {
	Send(*WatchServersResponse) error
	grpc.ServerStream
}

type serverDiscoveryServiceWatchServersServer struct {
	grpc.ServerStream
}

func (x *serverDiscoveryServiceWatchServersServer) Send(m *WatchServersResponse) error {
	return x.ServerStream.SendMsg(m)
}

// ServerDiscoveryService_ServiceDesc is the grpc.ServiceDesc for ServerDiscoveryService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ServerDiscoveryService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "serverdiscovery.ServerDiscoveryService",
	HandlerType: (*ServerDiscoveryServiceServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "WatchServers",
			Handler:       _ServerDiscoveryService_WatchServers_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "proto-public/pbserverdiscovery/serverdiscovery.proto",
}