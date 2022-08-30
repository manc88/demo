// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.21.3
// source: user_service.proto

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

// UserServiceClient is the client API for UserService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type UserServiceClient interface {
	GetAllV1(ctx context.Context, in *GetAllV1Request, opts ...grpc.CallOption) (*GetAllV1Response, error)
	CreateV1(ctx context.Context, in *CreateV1Request, opts ...grpc.CallOption) (*CreateV1Response, error)
	DeleteV1(ctx context.Context, in *DeleteV1Request, opts ...grpc.CallOption) (*DeleteV1Response, error)
}

type userServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewUserServiceClient(cc grpc.ClientConnInterface) UserServiceClient {
	return &userServiceClient{cc}
}

func (c *userServiceClient) GetAllV1(ctx context.Context, in *GetAllV1Request, opts ...grpc.CallOption) (*GetAllV1Response, error) {
	out := new(GetAllV1Response)
	err := c.cc.Invoke(ctx, "/user_service.UserService/GetAllV1", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) CreateV1(ctx context.Context, in *CreateV1Request, opts ...grpc.CallOption) (*CreateV1Response, error) {
	out := new(CreateV1Response)
	err := c.cc.Invoke(ctx, "/user_service.UserService/CreateV1", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) DeleteV1(ctx context.Context, in *DeleteV1Request, opts ...grpc.CallOption) (*DeleteV1Response, error) {
	out := new(DeleteV1Response)
	err := c.cc.Invoke(ctx, "/user_service.UserService/DeleteV1", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// UserServiceServer is the server API for UserService service.
// All implementations must embed UnimplementedUserServiceServer
// for forward compatibility
type UserServiceServer interface {
	GetAllV1(context.Context, *GetAllV1Request) (*GetAllV1Response, error)
	CreateV1(context.Context, *CreateV1Request) (*CreateV1Response, error)
	DeleteV1(context.Context, *DeleteV1Request) (*DeleteV1Response, error)
	mustEmbedUnimplementedUserServiceServer()
}

// UnimplementedUserServiceServer must be embedded to have forward compatible implementations.
type UnimplementedUserServiceServer struct {
}

func (UnimplementedUserServiceServer) GetAllV1(context.Context, *GetAllV1Request) (*GetAllV1Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAllV1 not implemented")
}
func (UnimplementedUserServiceServer) CreateV1(context.Context, *CreateV1Request) (*CreateV1Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateV1 not implemented")
}
func (UnimplementedUserServiceServer) DeleteV1(context.Context, *DeleteV1Request) (*DeleteV1Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteV1 not implemented")
}
func (UnimplementedUserServiceServer) mustEmbedUnimplementedUserServiceServer() {}

// UnsafeUserServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to UserServiceServer will
// result in compilation errors.
type UnsafeUserServiceServer interface {
	mustEmbedUnimplementedUserServiceServer()
}

func RegisterUserServiceServer(s grpc.ServiceRegistrar, srv UserServiceServer) {
	s.RegisterService(&UserService_ServiceDesc, srv)
}

func _UserService_GetAllV1_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetAllV1Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).GetAllV1(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/user_service.UserService/GetAllV1",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).GetAllV1(ctx, req.(*GetAllV1Request))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_CreateV1_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateV1Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).CreateV1(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/user_service.UserService/CreateV1",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).CreateV1(ctx, req.(*CreateV1Request))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_DeleteV1_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteV1Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).DeleteV1(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/user_service.UserService/DeleteV1",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).DeleteV1(ctx, req.(*DeleteV1Request))
	}
	return interceptor(ctx, in, info, handler)
}

// UserService_ServiceDesc is the grpc.ServiceDesc for UserService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var UserService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "user_service.UserService",
	HandlerType: (*UserServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetAllV1",
			Handler:    _UserService_GetAllV1_Handler,
		},
		{
			MethodName: "CreateV1",
			Handler:    _UserService_CreateV1_Handler,
		},
		{
			MethodName: "DeleteV1",
			Handler:    _UserService_DeleteV1_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "user_service.proto",
}
