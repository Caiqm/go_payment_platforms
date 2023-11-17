package logic

import (
	"context"
	"github.com/Caiqm/go_payment_platforms/pkg/alipay"
	pb "github.com/Caiqm/go_payment_platforms/protoc/aliyun_pb"
	"github.com/goccy/go-json"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type AliPay struct {
	pb.UnimplementedAliPayServer
}

// 支付宝APP支付
func (ap *AliPay) TradeAppPay(ctx context.Context, req *pb.PayRequest) (rpy *pb.PayReply, err error) {
	var p = alipay.TradeAppPay{}
	p.NotifyURL = req.NotifyURL
	p.Subject = req.Subject
	p.OutTradeNo = req.OutTradeNo
	p.TotalAmount = req.TotalAmount
	p.AppAuthToken = req.AppAuthToken
	if req.ProductCode != "" {
		p.ProductCode = "QUICK_MSECURITY_PAY"
	}
	param, err := NewPayClient(req.AppId, false).TradeAppPay(p)
	if err != nil {
		return
	}
	rpy.PayURL = param
	return
}

// 二维码支付
func (ap *AliPay) TradePagePay(ctx context.Context, req *pb.PayRequest) (rpy *pb.PayReply, err error) {
	var p = alipay.TradePagePay{}
	p.NotifyURL = req.NotifyURL
	p.Subject = req.Subject
	p.OutTradeNo = req.OutTradeNo
	p.TotalAmount = req.TotalAmount
	p.AppAuthToken = req.AppAuthToken
	if req.ProductCode != "" {
		p.ProductCode = "FAST_INSTANT_TRADE_PAY"
	}
	param, err := NewPayClient(req.AppId, false).TradePagePay(p)
	if err != nil {
		return
	}
	r, _ := json.Marshal(param)
	rpy.PayURL = string(r)
	return
}

// 支付宝小程序支付
func (ap *AliPay) TradeCreate(ctx context.Context, req *pb.PayRequest) (rpy *pb.PayReply, err error) {
	var p = alipay.TradeCreate{}
	p.NotifyURL = req.NotifyURL
	p.Subject = req.Subject
	p.OutTradeNo = req.OutTradeNo
	p.TotalAmount = req.TotalAmount
	p.BuyerId = req.BuyerId
	p.AppAuthToken = req.AppAuthToken
	if req.ProductCode != "" {
		p.ProductCode = "FACE_TO_FACE_PAYMENT"
	}
	param, err := NewPayClient(req.AppId, false).TradeCreate(p)
	if err != nil {
		return
	}
	rpy.TradeNo = param.TradeNo
	rpy.OutTradeNo = param.OutTradeNo
	return
}

// 支付宝手机网站支付接口2.0
func (ap *AliPay) TradeWapPay(ctx context.Context, req *pb.PayRequest) (rpy *pb.PayReply, err error) {
	var p = alipay.TradeWapPay{}
	p.NotifyURL = req.NotifyURL
	p.Subject = req.Subject
	p.OutTradeNo = req.OutTradeNo
	p.TotalAmount = req.TotalAmount
	p.AppAuthToken = req.AppAuthToken
	if req.ProductCode != "" {
		p.ProductCode = "QUICK_WAP_WAY"
	}
	param, err := NewPayClient(req.AppId, false).TradeWapPay(p)
	if err != nil {
		return
	}
	r, _ := json.Marshal(param)
	rpy.PayURL = string(r)
	return nil, status.Errorf(codes.Unimplemented, "method TradeWapPay not implemented")
}

// 支付查询
func (ap *AliPay) TradeQuery(ctx context.Context, req *pb.PayQueryRequest) (rpy *pb.PayQueryReply, err error) {
	var p = alipay.TradeQuery{}
	p.TradeNo = req.TradeNo
	p.OutTradeNo = req.OutTradeNo
	p.AppAuthToken = req.AppAuthToken
	param, err := NewPayClient(req.AppId, false).TradeQuery(p)
	if err != nil {
		return
	}
	r, _ := json.Marshal(param)
	rpy.QueryResult = r
	return
}

// 执行退款
func (ap *AliPay) TradeRefund(ctx context.Context, req *pb.RefundRequest) (rpy *pb.RefundReply, err error) {
	var p = alipay.TradeRefund{}
	p.TradeNo = req.TradeNo
	p.OutTradeNo = req.OutTradeNo
	p.OutRequestNo = req.OutRequestNo
	p.RefundAmount = req.RefundAmount
	p.AppAuthToken = req.AppAuthToken
	p.RefundReason = req.RefundReason
	param, err := NewPayClient(req.AppId, false).TradeRefund(p)
	if err != nil {
		return
	}
	r, _ := json.Marshal(param)
	rpy.RefundResult = r
	return
}

// 退款查询
func (ap *AliPay) TradeFastPayRefundQuery(ctx context.Context, req *pb.RefundQueryRequest) (rpy *pb.RefundQueryReply, err error) {
	var p = alipay.TradeFastPayRefundQuery{}
	p.TradeNo = req.TradeNo
	p.OutTradeNo = req.OutTradeNo
	p.OutRequestNo = req.OutRequestNo
	p.AppAuthToken = req.AppAuthToken
	param, err := NewPayClient(req.AppId, false).TradeFastPayRefundQuery(p)
	if err != nil {
		return
	}
	r, _ := json.Marshal(param)
	rpy.RefundQueryResult = r
	return
}
