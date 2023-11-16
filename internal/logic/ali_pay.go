package logic

import (
	"context"
	"github.com/Caiqm/go_payment_platforms/pkg/alipay"
	pb "github.com/Caiqm/go_payment_platforms/protoc/aliyun_pb"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type AliPay struct {
	pb.UnimplementedAliPayServer
}

func (ap *AliPay) TradeAppPay(ctx context.Context, req *pb.PayRequest) (rpy *pb.PayReply, err error) {
	var p = alipay.TradeAppPay{}
	p.NotifyURL = req.NotifyURL
	p.Subject = req.Subject
	p.OutTradeNo = req.OutTradeNo
	p.TotalAmount = req.TotalAmount
	p.ProductCode = req.ProductCode
	client := alipay.New(req.AppId, "", false)
	param, err := client.TradeAppPay(p)
	if err != nil {
		return
	}
	rpy.PayURL = param
	return
}

func (ap *AliPay) TradePagePay(ctx context.Context, req *pb.PayRequest) (rpy *pb.PayReply, err error) {
	return nil, status.Errorf(codes.Unimplemented, "method TradePagePay not implemented")
}

func (ap *AliPay) TradeCreate(ctx context.Context, req *pb.PayRequest) (rpy *pb.PayReply, err error) {
	return nil, status.Errorf(codes.Unimplemented, "method TradeCreate not implemented")
}

func (ap *AliPay) TradeQuery(ctx context.Context, req *pb.PayQueryRequest) (rpy *pb.PayQueryReply, err error) {
	return nil, status.Errorf(codes.Unimplemented, "method TradeQuery not implemented")
}

func (ap *AliPay) TradeRefund(ctx context.Context, req *pb.RefundRequest) (rpy *pb.RefundReply, err error) {
	return nil, status.Errorf(codes.Unimplemented, "method TradeRefund not implemented")
}
