// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v4.23.1
// source: handler_gRPC.proto

package proto

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
	MyService_GetFuncRPC_FullMethodName     = "/proto.MyService/GetFuncRPC"
	MyService_PostFuncRPC_FullMethodName    = "/proto.MyService/PostFuncRPC"
	MyService_GetFuncPingRPC_FullMethodName = "/proto.MyService/GetFuncPingRPC"
)

// MyServiceClient is the client API for MyService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type MyServiceClient interface {
	GetFuncRPC(ctx context.Context, in *GetFuncRequest, opts ...grpc.CallOption) (*GetFuncResponse, error)
	PostFuncRPC(ctx context.Context, in *PostFuncRequest, opts ...grpc.CallOption) (*PostFuncResponse, error)
	GetFuncPingRPC(ctx context.Context, in *GetFuncPingRequest, opts ...grpc.CallOption) (*GetFuncPingResponse, error)
}

type myServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewMyServiceClient(cc grpc.ClientConnInterface) MyServiceClient {
	return &myServiceClient{cc}
}

func (c *myServiceClient) GetFuncRPC(ctx context.Context, in *GetFuncRequest, opts ...grpc.CallOption) (*GetFuncResponse, error) {
	out := new(GetFuncResponse)
	err := c.cc.Invoke(ctx, MyService_GetFuncRPC_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *myServiceClient) PostFuncRPC(ctx context.Context, in *PostFuncRequest, opts ...grpc.CallOption) (*PostFuncResponse, error) {
	out := new(PostFuncResponse)
	err := c.cc.Invoke(ctx, MyService_PostFuncRPC_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *myServiceClient) GetFuncPingRPC(ctx context.Context, in *GetFuncPingRequest, opts ...grpc.CallOption) (*GetFuncPingResponse, error) {
	out := new(GetFuncPingResponse)
	err := c.cc.Invoke(ctx, MyService_GetFuncPingRPC_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MyServiceServer is the server API for MyService service.
// All implementations must embed UnimplementedMyServiceServer
// for forward compatibility
type MyServiceServer interface {
	GetFuncRPC(context.Context, *GetFuncRequest) (*GetFuncResponse, error)
	PostFuncRPC(context.Context, *PostFuncRequest) (*PostFuncResponse, error)
	GetFuncPingRPC(context.Context, *GetFuncPingRequest) (*GetFuncPingResponse, error)
	mustEmbedUnimplementedMyServiceServer()
}

// UnimplementedMyServiceServer must be embedded to have forward compatible implementations.
type UnimplementedMyServiceServer struct {
}

func (UnimplementedMyServiceServer) GetFuncRPC(context.Context, *GetFuncRequest) (*GetFuncResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetFuncRPC not implemented")
}
func (UnimplementedMyServiceServer) PostFuncRPC(context.Context, *PostFuncRequest) (*PostFuncResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PostFuncRPC not implemented")
}
func (UnimplementedMyServiceServer) GetFuncPingRPC(context.Context, *GetFuncPingRequest) (*GetFuncPingResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetFuncPingRPC not implemented")
}
func (UnimplementedMyServiceServer) mustEmbedUnimplementedMyServiceServer() {}

// UnsafeMyServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to MyServiceServer will
// result in compilation errors.
type UnsafeMyServiceServer interface {
	mustEmbedUnimplementedMyServiceServer()
}

func RegisterMyServiceServer(s grpc.ServiceRegistrar, srv MyServiceServer) {
	s.RegisterService(&MyService_ServiceDesc, srv)
}

func _MyService_GetFuncRPC_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetFuncRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MyServiceServer).GetFuncRPC(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: MyService_GetFuncRPC_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MyServiceServer).GetFuncRPC(ctx, req.(*GetFuncRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MyService_PostFuncRPC_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PostFuncRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MyServiceServer).PostFuncRPC(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: MyService_PostFuncRPC_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MyServiceServer).PostFuncRPC(ctx, req.(*PostFuncRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MyService_GetFuncPingRPC_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetFuncPingRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MyServiceServer).GetFuncPingRPC(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: MyService_GetFuncPingRPC_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MyServiceServer).GetFuncPingRPC(ctx, req.(*GetFuncPingRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// MyService_ServiceDesc is the grpc.ServiceDesc for MyService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var MyService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "proto.MyService",
	HandlerType: (*MyServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetFuncRPC",
			Handler:    _MyService_GetFuncRPC_Handler,
		},
		{
			MethodName: "PostFuncRPC",
			Handler:    _MyService_PostFuncRPC_Handler,
		},
		{
			MethodName: "GetFuncPingRPC",
			Handler:    _MyService_GetFuncPingRPC_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "handler_gRPC.proto",
}
