package logic

import (
	"context"
	"fmt"
	pb "github.com/Caiqm/go_payment_platforms/protoc/weixin_pb"
	wechat "github.com/Caiqm/wxpay-v2"
	"github.com/goccy/go-json"
	"time"
)

type WxPay struct {
	pb.UnimplementedWxPayServer
}

// APP支付
func (wp *WxPay) TradeAppPay(ctx context.Context, req *pb.PayRequest) (rpy *pb.PayReply, err error) {
	client, err := NewWeChatClient(req.AppId, req.Secret)
	if err != nil {
		err = fmt.Errorf("初始化失败：%s", err.Error())
		return
	}
	client.LoadOptionFunc(wechat.WithMchInformation(req.MchId, req.ApiKey), wechat.WithPayHost())
	var p = wechat.TradeApp{}
	p.Body = req.Body
	p.OutTradeNo = req.OutTradeNo
	p.TotalFee = fmt.Sprintf("%f", req.TotalFee*100)
	p.SpbillCreateIp = req.SpbillCreateIp
	p.NotifyUrl = req.NotifyUrl
	payRsp, err := client.TradeApp(p)
	if err != nil {
		err = fmt.Errorf("支付失败：%s", err.Error())
		return
	}
	rpy.MchId = payRsp.MchID
	rpy.NonceStr = payRsp.NonceStr
	rpy.PrepayId = payRsp.PrepayId
	rpy.Sign = payRsp.Sign
	rpy.Timestamp = fmt.Sprintf("%d", time.Now().Unix())
	rpy.AppId = payRsp.AppID
	rpy.TradeType = payRsp.TradeType
	return
}

// 小程序支付
func (wp *WxPay) TradeMpPay(ctx context.Context, req *pb.PayRequest) (rpy *pb.MpPayReply, err error) {
	client, err := NewWeChatClient(req.AppId, req.Secret)
	if err != nil {
		err = fmt.Errorf("初始化失败：%s", err.Error())
		return
	}
	client.LoadOptionFunc(wechat.WithMchInformation(req.MchId, req.ApiKey), wechat.WithPayHost())
	var p = wechat.TradeApplet{}
	p.Body = req.Body
	p.OutTradeNo = req.OutTradeNo
	p.TotalFee = fmt.Sprintf("%f", req.TotalFee*100)
	p.SpbillCreateIp = req.SpbillCreateIp
	p.NotifyUrl = req.NotifyUrl
	p.OpenId = req.OpenId
	payRsp, err := client.TradeApplet(p)
	if err != nil {
		err = fmt.Errorf("支付失败：%s", err.Error())
		return
	}
	rpy.AppId = payRsp.AppID
	rpy.NonceStr = payRsp.NonceStr
	rpy.Package = payRsp.Package
	rpy.Sign = payRsp.PaySign
	rpy.Timestamp = fmt.Sprintf("%d", time.Now().Unix())
	rpy.SignType = payRsp.SignType
	return
}

// JSAPI支付
func (wp *WxPay) TradeJsApiPay(ctx context.Context, req *pb.PayRequest) (rpy *pb.PayReply, err error) {
	client, err := NewWeChatClient(req.AppId, req.Secret)
	if err != nil {
		err = fmt.Errorf("初始化失败：%s", err.Error())
		return
	}
	client.LoadOptionFunc(wechat.WithMchInformation(req.MchId, req.ApiKey), wechat.WithPayHost())
	var p = wechat.TradeJSAPI{}
	p.Body = req.Body
	p.OutTradeNo = req.OutTradeNo
	p.TotalFee = fmt.Sprintf("%f", req.TotalFee*100)
	p.SpbillCreateIp = req.SpbillCreateIp
	p.NotifyUrl = req.NotifyUrl
	p.OpenId = req.OpenId
	payRsp, err := client.TradeJSAPI(p)
	if err != nil {
		err = fmt.Errorf("支付失败：%s", err.Error())
		return
	}
	rpy.AppId = payRsp.AppID
	rpy.NonceStr = payRsp.NonceStr
	rpy.PrepayId = payRsp.PrepayId
	rpy.Sign = payRsp.Sign
	rpy.Timestamp = fmt.Sprintf("%d", time.Now().Unix())
	rpy.TradeType = payRsp.TradeType
	rpy.MchId = payRsp.MchID
	return
}

// H5支付
func (wp *WxPay) TradeWapPay(ctx context.Context, req *pb.PayRequest) (rpy *pb.WapPayReply, err error) {
	client, err := NewWeChatClient(req.AppId, req.Secret)
	if err != nil {
		err = fmt.Errorf("初始化失败：%s", err.Error())
		return
	}
	client.LoadOptionFunc(wechat.WithMchInformation(req.MchId, req.ApiKey), wechat.WithPayHost())
	var p = wechat.TradeWap{}
	p.Body = req.Body
	p.OutTradeNo = req.OutTradeNo
	p.TotalFee = fmt.Sprintf("%f", req.TotalFee*100)
	p.SpbillCreateIp = req.SpbillCreateIp
	p.NotifyUrl = req.NotifyUrl
	payRsp, err := client.TradeWap(p)
	if err != nil {
		err = fmt.Errorf("支付失败：%s", err.Error())
		return
	}
	rpy.AppId = payRsp.AppID
	rpy.MchId = payRsp.MchID
	rpy.NonceStr = payRsp.NonceStr
	rpy.PrepayId = payRsp.PrepayId
	rpy.Sign = payRsp.Sign
	rpy.MWebUrl = payRsp.MWebUrl
	rpy.TradeType = payRsp.TradeType
	return
}

// 扫码支付
func (wp *WxPay) TradeNativePay(ctx context.Context, req *pb.PayRequest) (rpy *pb.NativePayReply, err error) {
	client, err := NewWeChatClient(req.AppId, req.Secret)
	if err != nil {
		err = fmt.Errorf("初始化失败：%s", err.Error())
		return
	}
	client.LoadOptionFunc(wechat.WithMchInformation(req.MchId, req.ApiKey), wechat.WithPayHost())
	var p = wechat.TradeNative{}
	p.Body = req.Body
	p.OutTradeNo = req.OutTradeNo
	p.TotalFee = fmt.Sprintf("%f", req.TotalFee*100)
	p.SpbillCreateIp = req.SpbillCreateIp
	p.NotifyUrl = req.NotifyUrl
	payRsp, err := client.TradeNative(p)
	if err != nil {
		err = fmt.Errorf("支付失败：%s", err.Error())
		return
	}
	rpy.AppId = payRsp.AppID
	rpy.MchId = payRsp.MchID
	rpy.NonceStr = payRsp.NonceStr
	rpy.PrepayId = payRsp.PrepayId
	rpy.Sign = payRsp.Sign
	rpy.TradeType = payRsp.TradeType
	rpy.CodeUrl = payRsp.CodeUrl
	return
}

// 查询订单
func (wp *WxPay) TradeQuery(ctx context.Context, req *pb.PayQueryRequest) (rpy *pb.PayQueryReply, err error) {
	client, err := NewWeChatClient(req.AppId, req.Secret)
	if err != nil {
		err = fmt.Errorf("初始化失败：%s", err.Error())
		return
	}
	client.LoadOptionFunc(wechat.WithMchInformation(req.MchId, req.ApiKey), wechat.WithApiHost("https://api.mch.weixin.qq.com/pay/orderquery"))
	var p = wechat.TradeOrderQuery{}
	if req.GetOutTradeNo() != "" {
		p.OutTradeNo = req.OutTradeNo
	} else if req.GetTransactionId() != "" {
		p.TransactionId = req.TransactionId
	} else {
		err = fmt.Errorf("订单号参数错误")
		return
	}
	queryRsp, err := client.TradeOrderQuery(p)
	if err != nil {
		err = fmt.Errorf("查询失败：%s", err.Error())
		return
	}
	r, err := json.Marshal(queryRsp)
	if err != nil {
		err = fmt.Errorf("序列化失败：%s", err.Error())
		return
	}
	rpy.QueryResult = r
	return
}

// 申请退款
func (wp *WxPay) TradeRefund(ctx context.Context, req *pb.RefundRequest) (rpy *pb.RefundReply, err error) {
	client, err := NewWeChatClient(req.AppId, req.Secret)
	if err != nil {
		err = fmt.Errorf("初始化失败：%s", err.Error())
		return
	}
	client.LoadOptionFunc(wechat.WithMchInformation(req.MchId, req.ApiKey), wechat.WithApiHost("https://api.mch.weixin.qq.com/secapi/pay/refund"))
	var p = wechat.TradeRefund{}
	if req.GetTransactionId() == "" {
		p.TransactionId = req.TransactionId
	} else if req.GetOutTradeNo() == "" {
		p.OutTradeNo = req.OutTradeNo
	} else {
		err = fmt.Errorf("订单号参数错误")
		return
	}
	p.OutRefundNo = req.OutRefundNo
	p.TotalFee = fmt.Sprintf("%f", req.TotalFee*100)
	p.RefundFee = fmt.Sprintf("%f", req.RefundFee*100)
	p.RefundDesc = req.RefundDesc
	p.RefundFeeType = "CNY"
	refundRsp, err := client.TradeRefund(p)
	if err != nil {
		err = fmt.Errorf("退款失败：%s", err.Error())
		return
	}
	r, err := json.Marshal(refundRsp)
	if err != nil {
		err = fmt.Errorf("序列化失败：%s", err.Error())
		return
	}
	rpy.RefundResult = r
	return
}

// 查询退款
func (wp *WxPay) TradeRefundQuery(ctx context.Context, req *pb.RefundQueryRequest) (rpy *pb.RefundQueryReply, err error) {
	client, err := NewWeChatClient(req.AppId, req.Secret)
	if err != nil {
		err = fmt.Errorf("初始化失败：%s", err.Error())
		return
	}
	client.LoadOptionFunc(wechat.WithMchInformation(req.MchId, req.ApiKey), wechat.WithApiHost("https://api.mch.weixin.qq.com/pay/refundquery"))
	var p = wechat.TradeRefundQuery{}
	if req.GetOutRefundNo() != "" {
		p.OutRefundNo = req.OutRefundNo
	} else if req.GetOutTradeNo() != "" {
		p.OutTradeNo = req.OutTradeNo
	} else if req.GetTransactionId() != "" {
		p.TransactionId = req.TransactionId
	} else if req.GetRefundId() != "" {
		p.RefundId = req.RefundId
	} else {
		err = fmt.Errorf("订单号参数错误")
		return
	}
	queryRsp, err := client.TradeRefundQuery(p)
	if err != nil {
		err = fmt.Errorf("查询失败：%s", err.Error())
		return
	}
	r, err := json.Marshal(queryRsp)
	if err != nil {
		err = fmt.Errorf("序列化失败：%s", err.Error())
		return
	}
	rpy.RefundQueryResult = r
	return
}
