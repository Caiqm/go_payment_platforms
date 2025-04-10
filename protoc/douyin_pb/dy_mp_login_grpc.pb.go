// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v4.25.0
// source: dy_mp_login.proto

package douyin_pb

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
	DyMpLogin_Code2Session_FullMethodName   = "/douyin_pb.DyMpLogin/Code2Session"
	DyMpLogin_GetPhoneNumber_FullMethodName = "/douyin_pb.DyMpLogin/GetPhoneNumber"
)

// DyMpLoginClient is the client API for DyMpLogin service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type DyMpLoginClient interface {
	Code2Session(ctx context.Context, in *MpLoginRequest, opts ...grpc.CallOption) (*MpLoginReply, error)
	GetPhoneNumber(ctx context.Context, in *MpPhoneNumberRequest, opts ...grpc.CallOption) (*MpPhoneNumberReply, error)
}

type dyMpLoginClient struct {
	cc grpc.ClientConnInterface
}

func NewDyMpLoginClient(cc grpc.ClientConnInterface) DyMpLoginClient {
	return &dyMpLoginClient{cc}
}

func (c *dyMpLoginClient) Code2Session(ctx context.Context, in *MpLoginRequest, opts ...grpc.CallOption) (*MpLoginReply, error) {
	out := new(MpLoginReply)
	err := c.cc.Invoke(ctx, DyMpLogin_Code2Session_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dyMpLoginClient) GetPhoneNumber(ctx context.Context, in *MpPhoneNumberRequest, opts ...grpc.CallOption) (*MpPhoneNumberReply, error) {
	out := new(MpPhoneNumberReply)
	err := c.cc.Invoke(ctx, DyMpLogin_GetPhoneNumber_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// DyMpLoginServer is the server API for DyMpLogin service.
// All implementations must embed UnimplementedDyMpLoginServer
// for forward compatibility
type DyMpLoginServer interface {
	Code2Session(context.Context, *MpLoginRequest) (*MpLoginReply, error)
	GetPhoneNumber(context.Context, *MpPhoneNumberRequest) (*MpPhoneNumberReply, error)
	mustEmbedUnimplementedDyMpLoginServer()
}

// UnimplementedDyMpLoginServer must be embedded to have forward compatible implementations.
type UnimplementedDyMpLoginServer struct {
}

func (UnimplementedDyMpLoginServer) Code2Session(context.Context, *MpLoginRequest) (*MpLoginReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Code2Session not implemented")
}
func (UnimplementedDyMpLoginServer) GetPhoneNumber(context.Context, *MpPhoneNumberRequest) (*MpPhoneNumberReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetPhoneNumber not implemented")
}
func (UnimplementedDyMpLoginServer) mustEmbedUnimplementedDyMpLoginServer() {}

// UnsafeDyMpLoginServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to DyMpLoginServer will
// result in compilation errors.
type UnsafeDyMpLoginServer interface {
	mustEmbedUnimplementedDyMpLoginServer()
}

func RegisterDyMpLoginServer(s grpc.ServiceRegistrar, srv DyMpLoginServer) {
	s.RegisterService(&DyMpLogin_ServiceDesc, srv)
}

func _DyMpLogin_Code2Session_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MpLoginRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DyMpLoginServer).Code2Session(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: DyMpLogin_Code2Session_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DyMpLoginServer).Code2Session(ctx, req.(*MpLoginRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DyMpLogin_GetPhoneNumber_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MpPhoneNumberRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DyMpLoginServer).GetPhoneNumber(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: DyMpLogin_GetPhoneNumber_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DyMpLoginServer).GetPhoneNumber(ctx, req.(*MpPhoneNumberRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// DyMpLogin_ServiceDesc is the grpc.ServiceDesc for DyMpLogin service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var DyMpLogin_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "douyin_pb.DyMpLogin",
	HandlerType: (*DyMpLoginServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Code2Session",
			Handler:    _DyMpLogin_Code2Session_Handler,
		},
		{
			MethodName: "GetPhoneNumber",
			Handler:    _DyMpLogin_GetPhoneNumber_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "dy_mp_login.proto",
}
