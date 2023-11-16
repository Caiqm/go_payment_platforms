// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v4.25.0
// source: ali_mp_login.proto

package aliyun_pb

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
	AliMpLogin_SystemOauthToken_FullMethodName = "/aliyun_pb.AliMpLogin/SystemOauthToken"
	AliMpLogin_UserInfoShare_FullMethodName    = "/aliyun_pb.AliMpLogin/UserInfoShare"
)

// AliMpLoginClient is the client API for AliMpLogin service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type AliMpLoginClient interface {
	SystemOauthToken(ctx context.Context, in *MpLoginRequest, opts ...grpc.CallOption) (*MpLoginReply, error)
	UserInfoShare(ctx context.Context, in *MpUserInfoRequest, opts ...grpc.CallOption) (*MpUserInfoReply, error)
}

type aliMpLoginClient struct {
	cc grpc.ClientConnInterface
}

func NewAliMpLoginClient(cc grpc.ClientConnInterface) AliMpLoginClient {
	return &aliMpLoginClient{cc}
}

func (c *aliMpLoginClient) SystemOauthToken(ctx context.Context, in *MpLoginRequest, opts ...grpc.CallOption) (*MpLoginReply, error) {
	out := new(MpLoginReply)
	err := c.cc.Invoke(ctx, AliMpLogin_SystemOauthToken_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *aliMpLoginClient) UserInfoShare(ctx context.Context, in *MpUserInfoRequest, opts ...grpc.CallOption) (*MpUserInfoReply, error) {
	out := new(MpUserInfoReply)
	err := c.cc.Invoke(ctx, AliMpLogin_UserInfoShare_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AliMpLoginServer is the server API for AliMpLogin service.
// All implementations must embed UnimplementedAliMpLoginServer
// for forward compatibility
type AliMpLoginServer interface {
	SystemOauthToken(context.Context, *MpLoginRequest) (*MpLoginReply, error)
	UserInfoShare(context.Context, *MpUserInfoRequest) (*MpUserInfoReply, error)
	mustEmbedUnimplementedAliMpLoginServer()
}

// UnimplementedAliMpLoginServer must be embedded to have forward compatible implementations.
type UnimplementedAliMpLoginServer struct {
}

func (UnimplementedAliMpLoginServer) SystemOauthToken(context.Context, *MpLoginRequest) (*MpLoginReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SystemOauthToken not implemented")
}
func (UnimplementedAliMpLoginServer) UserInfoShare(context.Context, *MpUserInfoRequest) (*MpUserInfoReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UserInfoShare not implemented")
}
func (UnimplementedAliMpLoginServer) mustEmbedUnimplementedAliMpLoginServer() {}

// UnsafeAliMpLoginServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to AliMpLoginServer will
// result in compilation errors.
type UnsafeAliMpLoginServer interface {
	mustEmbedUnimplementedAliMpLoginServer()
}

func RegisterAliMpLoginServer(s grpc.ServiceRegistrar, srv AliMpLoginServer) {
	s.RegisterService(&AliMpLogin_ServiceDesc, srv)
}

func _AliMpLogin_SystemOauthToken_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MpLoginRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AliMpLoginServer).SystemOauthToken(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AliMpLogin_SystemOauthToken_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AliMpLoginServer).SystemOauthToken(ctx, req.(*MpLoginRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AliMpLogin_UserInfoShare_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MpUserInfoRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AliMpLoginServer).UserInfoShare(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AliMpLogin_UserInfoShare_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AliMpLoginServer).UserInfoShare(ctx, req.(*MpUserInfoRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// AliMpLogin_ServiceDesc is the grpc.ServiceDesc for AliMpLogin service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var AliMpLogin_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "aliyun_pb.AliMpLogin",
	HandlerType: (*AliMpLoginServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SystemOauthToken",
			Handler:    _AliMpLogin_SystemOauthToken_Handler,
		},
		{
			MethodName: "UserInfoShare",
			Handler:    _AliMpLogin_UserInfoShare_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "ali_mp_login.proto",
}
