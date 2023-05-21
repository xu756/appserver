// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.21.12
// source: login.proto

package pb

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

// LoginRpcClient is the client API for LoginRpc service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type LoginRpcClient interface {
	MiniLoginByMobile(ctx context.Context, in *LoginReq, opts ...grpc.CallOption) (*LoginResp, error)
	MiniLoginByAuth(ctx context.Context, in *MiniAuthReq, opts ...grpc.CallOption) (*LoginResp, error)
}

type loginRpcClient struct {
	cc grpc.ClientConnInterface
}

func NewLoginRpcClient(cc grpc.ClientConnInterface) LoginRpcClient {
	return &loginRpcClient{cc}
}

func (c *loginRpcClient) MiniLoginByMobile(ctx context.Context, in *LoginReq, opts ...grpc.CallOption) (*LoginResp, error) {
	out := new(LoginResp)
	err := c.cc.Invoke(ctx, "/pb.LoginRpc/MiniLoginByMobile", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *loginRpcClient) MiniLoginByAuth(ctx context.Context, in *MiniAuthReq, opts ...grpc.CallOption) (*LoginResp, error) {
	out := new(LoginResp)
	err := c.cc.Invoke(ctx, "/pb.LoginRpc/MiniLoginByAuth", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// LoginRpcServer is the server API for LoginRpc service.
// All implementations must embed UnimplementedLoginRpcServer
// for forward compatibility
type LoginRpcServer interface {
	MiniLoginByMobile(context.Context, *LoginReq) (*LoginResp, error)
	MiniLoginByAuth(context.Context, *MiniAuthReq) (*LoginResp, error)
	mustEmbedUnimplementedLoginRpcServer()
}

// UnimplementedLoginRpcServer must be embedded to have forward compatible implementations.
type UnimplementedLoginRpcServer struct {
}

func (UnimplementedLoginRpcServer) MiniLoginByMobile(context.Context, *LoginReq) (*LoginResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method MiniLoginByMobile not implemented")
}
func (UnimplementedLoginRpcServer) MiniLoginByAuth(context.Context, *MiniAuthReq) (*LoginResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method MiniLoginByAuth not implemented")
}
func (UnimplementedLoginRpcServer) mustEmbedUnimplementedLoginRpcServer() {}

// UnsafeLoginRpcServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to LoginRpcServer will
// result in compilation errors.
type UnsafeLoginRpcServer interface {
	mustEmbedUnimplementedLoginRpcServer()
}

func RegisterLoginRpcServer(s grpc.ServiceRegistrar, srv LoginRpcServer) {
	s.RegisterService(&LoginRpc_ServiceDesc, srv)
}

func _LoginRpc_MiniLoginByMobile_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LoginReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LoginRpcServer).MiniLoginByMobile(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.LoginRpc/MiniLoginByMobile",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LoginRpcServer).MiniLoginByMobile(ctx, req.(*LoginReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _LoginRpc_MiniLoginByAuth_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MiniAuthReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LoginRpcServer).MiniLoginByAuth(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.LoginRpc/MiniLoginByAuth",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LoginRpcServer).MiniLoginByAuth(ctx, req.(*MiniAuthReq))
	}
	return interceptor(ctx, in, info, handler)
}

// LoginRpc_ServiceDesc is the grpc.ServiceDesc for LoginRpc service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var LoginRpc_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "pb.LoginRpc",
	HandlerType: (*LoginRpcServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "MiniLoginByMobile",
			Handler:    _LoginRpc_MiniLoginByMobile_Handler,
		},
		{
			MethodName: "MiniLoginByAuth",
			Handler:    _LoginRpc_MiniLoginByAuth_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "login.proto",
}
