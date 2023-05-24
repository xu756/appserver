// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.21.12
// source: im.proto

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

// ImLoginClient is the client API for ImLogin service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ImLoginClient interface {
	// 元事件 连接 断开 状态更新 解密错误
	Meta(ctx context.Context, in *ImMeta, opts ...grpc.CallOption) (*ImResp, error)
	// 服务端主动推送
	GetImData(ctx context.Context, in *ImData, opts ...grpc.CallOption) (*ImResp, error)
}

type imLoginClient struct {
	cc grpc.ClientConnInterface
}

func NewImLoginClient(cc grpc.ClientConnInterface) ImLoginClient {
	return &imLoginClient{cc}
}

func (c *imLoginClient) Meta(ctx context.Context, in *ImMeta, opts ...grpc.CallOption) (*ImResp, error) {
	out := new(ImResp)
	err := c.cc.Invoke(ctx, "/pb.ImLogin/Meta", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *imLoginClient) GetImData(ctx context.Context, in *ImData, opts ...grpc.CallOption) (*ImResp, error) {
	out := new(ImResp)
	err := c.cc.Invoke(ctx, "/pb.ImLogin/GetImData", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ImLoginServer is the server API for ImLogin service.
// All implementations must embed UnimplementedImLoginServer
// for forward compatibility
type ImLoginServer interface {
	// 元事件 连接 断开 状态更新 解密错误
	Meta(context.Context, *ImMeta) (*ImResp, error)
	// 服务端主动推送
	GetImData(context.Context, *ImData) (*ImResp, error)
	mustEmbedUnimplementedImLoginServer()
}

// UnimplementedImLoginServer must be embedded to have forward compatible implementations.
type UnimplementedImLoginServer struct {
}

func (UnimplementedImLoginServer) Meta(context.Context, *ImMeta) (*ImResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Meta not implemented")
}
func (UnimplementedImLoginServer) GetImData(context.Context, *ImData) (*ImResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetImData not implemented")
}
func (UnimplementedImLoginServer) mustEmbedUnimplementedImLoginServer() {}

// UnsafeImLoginServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ImLoginServer will
// result in compilation errors.
type UnsafeImLoginServer interface {
	mustEmbedUnimplementedImLoginServer()
}

func RegisterImLoginServer(s grpc.ServiceRegistrar, srv ImLoginServer) {
	s.RegisterService(&ImLogin_ServiceDesc, srv)
}

func _ImLogin_Meta_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ImMeta)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ImLoginServer).Meta(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.ImLogin/Meta",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ImLoginServer).Meta(ctx, req.(*ImMeta))
	}
	return interceptor(ctx, in, info, handler)
}

func _ImLogin_GetImData_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ImData)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ImLoginServer).GetImData(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.ImLogin/GetImData",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ImLoginServer).GetImData(ctx, req.(*ImData))
	}
	return interceptor(ctx, in, info, handler)
}

// ImLogin_ServiceDesc is the grpc.ServiceDesc for ImLogin service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ImLogin_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "pb.ImLogin",
	HandlerType: (*ImLoginServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Meta",
			Handler:    _ImLogin_Meta_Handler,
		},
		{
			MethodName: "GetImData",
			Handler:    _ImLogin_GetImData_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "im.proto",
}
