// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v6.30.2
// source: proto/sync.proto

package proto

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
	SyncService_SendChange_FullMethodName = "/sync.SyncService/SendChange"
)

// SyncServiceClient is the client API for SyncService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type SyncServiceClient interface {
	SendChange(ctx context.Context, in *FileChange, opts ...grpc.CallOption) (*Ack, error)
}

type syncServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewSyncServiceClient(cc grpc.ClientConnInterface) SyncServiceClient {
	return &syncServiceClient{cc}
}

func (c *syncServiceClient) SendChange(ctx context.Context, in *FileChange, opts ...grpc.CallOption) (*Ack, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(Ack)
	err := c.cc.Invoke(ctx, SyncService_SendChange_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// SyncServiceServer is the server API for SyncService service.
// All implementations must embed UnimplementedSyncServiceServer
// for forward compatibility.
type SyncServiceServer interface {
	SendChange(context.Context, *FileChange) (*Ack, error)
	mustEmbedUnimplementedSyncServiceServer()
}

// UnimplementedSyncServiceServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedSyncServiceServer struct{}

func (UnimplementedSyncServiceServer) SendChange(context.Context, *FileChange) (*Ack, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SendChange not implemented")
}
func (UnimplementedSyncServiceServer) mustEmbedUnimplementedSyncServiceServer() {}
func (UnimplementedSyncServiceServer) testEmbeddedByValue()                     {}

// UnsafeSyncServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to SyncServiceServer will
// result in compilation errors.
type UnsafeSyncServiceServer interface {
	mustEmbedUnimplementedSyncServiceServer()
}

func RegisterSyncServiceServer(s grpc.ServiceRegistrar, srv SyncServiceServer) {
	// If the following call pancis, it indicates UnimplementedSyncServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&SyncService_ServiceDesc, srv)
}

func _SyncService_SendChange_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FileChange)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SyncServiceServer).SendChange(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: SyncService_SendChange_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SyncServiceServer).SendChange(ctx, req.(*FileChange))
	}
	return interceptor(ctx, in, info, handler)
}

// SyncService_ServiceDesc is the grpc.ServiceDesc for SyncService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var SyncService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "sync.SyncService",
	HandlerType: (*SyncServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SendChange",
			Handler:    _SyncService_SendChange_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/sync.proto",
}
