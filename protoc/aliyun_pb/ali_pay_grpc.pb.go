// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v4.25.0
// source: ali_pay.proto

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
	AliPay_TradeAppPay_FullMethodName             = "/aliyun_pb.AliPay/TradeAppPay"
	AliPay_TradePagePay_FullMethodName            = "/aliyun_pb.AliPay/TradePagePay"
	AliPay_TradeCreate_FullMethodName             = "/aliyun_pb.AliPay/TradeCreate"
	AliPay_TradeWapPay_FullMethodName             = "/aliyun_pb.AliPay/TradeWapPay"
	AliPay_TradeQuery_FullMethodName              = "/aliyun_pb.AliPay/TradeQuery"
	AliPay_TradeRefund_FullMethodName             = "/aliyun_pb.AliPay/TradeRefund"
	AliPay_TradeFastPayRefundQuery_FullMethodName = "/aliyun_pb.AliPay/TradeFastPayRefundQuery"
)

// AliPayClient is the client API for AliPay service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type AliPayClient interface {
	TradeAppPay(ctx context.Context, in *PayRequest, opts ...grpc.CallOption) (*PayReply, error)
	TradePagePay(ctx context.Context, in *PayRequest, opts ...grpc.CallOption) (*PayReply, error)
	TradeCreate(ctx context.Context, in *PayRequest, opts ...grpc.CallOption) (*PayReply, error)
	TradeWapPay(ctx context.Context, in *PayRequest, opts ...grpc.CallOption) (*PayReply, error)
	TradeQuery(ctx context.Context, in *PayQueryRequest, opts ...grpc.CallOption) (*PayQueryReply, error)
	TradeRefund(ctx context.Context, in *RefundRequest, opts ...grpc.CallOption) (*RefundReply, error)
	TradeFastPayRefundQuery(ctx context.Context, in *RefundQueryRequest, opts ...grpc.CallOption) (*RefundQueryReply, error)
}

type aliPayClient struct {
	cc grpc.ClientConnInterface
}

func NewAliPayClient(cc grpc.ClientConnInterface) AliPayClient {
	return &aliPayClient{cc}
}

func (c *aliPayClient) TradeAppPay(ctx context.Context, in *PayRequest, opts ...grpc.CallOption) (*PayReply, error) {
	out := new(PayReply)
	err := c.cc.Invoke(ctx, AliPay_TradeAppPay_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *aliPayClient) TradePagePay(ctx context.Context, in *PayRequest, opts ...grpc.CallOption) (*PayReply, error) {
	out := new(PayReply)
	err := c.cc.Invoke(ctx, AliPay_TradePagePay_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *aliPayClient) TradeCreate(ctx context.Context, in *PayRequest, opts ...grpc.CallOption) (*PayReply, error) {
	out := new(PayReply)
	err := c.cc.Invoke(ctx, AliPay_TradeCreate_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *aliPayClient) TradeWapPay(ctx context.Context, in *PayRequest, opts ...grpc.CallOption) (*PayReply, error) {
	out := new(PayReply)
	err := c.cc.Invoke(ctx, AliPay_TradeWapPay_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *aliPayClient) TradeQuery(ctx context.Context, in *PayQueryRequest, opts ...grpc.CallOption) (*PayQueryReply, error) {
	out := new(PayQueryReply)
	err := c.cc.Invoke(ctx, AliPay_TradeQuery_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *aliPayClient) TradeRefund(ctx context.Context, in *RefundRequest, opts ...grpc.CallOption) (*RefundReply, error) {
	out := new(RefundReply)
	err := c.cc.Invoke(ctx, AliPay_TradeRefund_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *aliPayClient) TradeFastPayRefundQuery(ctx context.Context, in *RefundQueryRequest, opts ...grpc.CallOption) (*RefundQueryReply, error) {
	out := new(RefundQueryReply)
	err := c.cc.Invoke(ctx, AliPay_TradeFastPayRefundQuery_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AliPayServer is the server API for AliPay service.
// All implementations must embed UnimplementedAliPayServer
// for forward compatibility
type AliPayServer interface {
	TradeAppPay(context.Context, *PayRequest) (*PayReply, error)
	TradePagePay(context.Context, *PayRequest) (*PayReply, error)
	TradeCreate(context.Context, *PayRequest) (*PayReply, error)
	TradeWapPay(context.Context, *PayRequest) (*PayReply, error)
	TradeQuery(context.Context, *PayQueryRequest) (*PayQueryReply, error)
	TradeRefund(context.Context, *RefundRequest) (*RefundReply, error)
	TradeFastPayRefundQuery(context.Context, *RefundQueryRequest) (*RefundQueryReply, error)
	mustEmbedUnimplementedAliPayServer()
}

// UnimplementedAliPayServer must be embedded to have forward compatible implementations.
type UnimplementedAliPayServer struct {
}

func (UnimplementedAliPayServer) TradeAppPay(context.Context, *PayRequest) (*PayReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method TradeAppPay not implemented")
}
func (UnimplementedAliPayServer) TradePagePay(context.Context, *PayRequest) (*PayReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method TradePagePay not implemented")
}
func (UnimplementedAliPayServer) TradeCreate(context.Context, *PayRequest) (*PayReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method TradeCreate not implemented")
}
func (UnimplementedAliPayServer) TradeWapPay(context.Context, *PayRequest) (*PayReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method TradeWapPay not implemented")
}
func (UnimplementedAliPayServer) TradeQuery(context.Context, *PayQueryRequest) (*PayQueryReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method TradeQuery not implemented")
}
func (UnimplementedAliPayServer) TradeRefund(context.Context, *RefundRequest) (*RefundReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method TradeRefund not implemented")
}
func (UnimplementedAliPayServer) TradeFastPayRefundQuery(context.Context, *RefundQueryRequest) (*RefundQueryReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method TradeFastPayRefundQuery not implemented")
}
func (UnimplementedAliPayServer) mustEmbedUnimplementedAliPayServer() {}

// UnsafeAliPayServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to AliPayServer will
// result in compilation errors.
type UnsafeAliPayServer interface {
	mustEmbedUnimplementedAliPayServer()
}

func RegisterAliPayServer(s grpc.ServiceRegistrar, srv AliPayServer) {
	s.RegisterService(&AliPay_ServiceDesc, srv)
}

func _AliPay_TradeAppPay_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PayRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AliPayServer).TradeAppPay(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AliPay_TradeAppPay_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AliPayServer).TradeAppPay(ctx, req.(*PayRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AliPay_TradePagePay_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PayRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AliPayServer).TradePagePay(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AliPay_TradePagePay_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AliPayServer).TradePagePay(ctx, req.(*PayRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AliPay_TradeCreate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PayRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AliPayServer).TradeCreate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AliPay_TradeCreate_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AliPayServer).TradeCreate(ctx, req.(*PayRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AliPay_TradeWapPay_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PayRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AliPayServer).TradeWapPay(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AliPay_TradeWapPay_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AliPayServer).TradeWapPay(ctx, req.(*PayRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AliPay_TradeQuery_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PayQueryRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AliPayServer).TradeQuery(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AliPay_TradeQuery_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AliPayServer).TradeQuery(ctx, req.(*PayQueryRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AliPay_TradeRefund_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RefundRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AliPayServer).TradeRefund(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AliPay_TradeRefund_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AliPayServer).TradeRefund(ctx, req.(*RefundRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AliPay_TradeFastPayRefundQuery_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RefundQueryRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AliPayServer).TradeFastPayRefundQuery(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AliPay_TradeFastPayRefundQuery_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AliPayServer).TradeFastPayRefundQuery(ctx, req.(*RefundQueryRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// AliPay_ServiceDesc is the grpc.ServiceDesc for AliPay service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var AliPay_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "aliyun_pb.AliPay",
	HandlerType: (*AliPayServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "TradeAppPay",
			Handler:    _AliPay_TradeAppPay_Handler,
		},
		{
			MethodName: "TradePagePay",
			Handler:    _AliPay_TradePagePay_Handler,
		},
		{
			MethodName: "TradeCreate",
			Handler:    _AliPay_TradeCreate_Handler,
		},
		{
			MethodName: "TradeWapPay",
			Handler:    _AliPay_TradeWapPay_Handler,
		},
		{
			MethodName: "TradeQuery",
			Handler:    _AliPay_TradeQuery_Handler,
		},
		{
			MethodName: "TradeRefund",
			Handler:    _AliPay_TradeRefund_Handler,
		},
		{
			MethodName: "TradeFastPayRefundQuery",
			Handler:    _AliPay_TradeFastPayRefundQuery_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "ali_pay.proto",
}
