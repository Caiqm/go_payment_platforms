package bytedance

import (
	"github.com/Caiqm/go_payment_platforms/pkg/bytedance"
	"testing"
)

// 担保支付
func TestClient_TradeEcPay(t *testing.T) {
	t.Log("========== TradeEcPay ==========")
	client.LoadOptionFunc(bytedance.WithApiHost("https://developer.toutiao.com/api/apps/ecpay/v1/create_order"))
	var p = bytedance.TradeEcPay{}
	p.OutOrderNo = "TEST2023112717521212345678"
	p.TotalAmount = 1
	p.Subject = "支付测试"
	p.Body = "支付测试"
	p.ValidTime = 300
	rsp, err := client.TradeEcPay(p)
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("%+v", rsp)
}

// 订单查询
func TestClient_TradeOrderQuery(t *testing.T) {
	t.Log("========== TradeOrderQuery ==========")
	client.LoadOptionFunc(bytedance.WithApiHost("https://developer.toutiao.com/api/apps/ecpay/v1/query_order"))
	var p = bytedance.TradeOrderQuery{}
	p.OutOrderNo = "TEST2023112717521212345678"
	rsp, err := client.TradeOrderQuery(p)
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("%+v", rsp)
}

// 发起退款
func TestClient_TradeRefund(t *testing.T) {
	t.Log("========== TradeRefund ==========")
	client.LoadOptionFunc(bytedance.WithApiHost("https://developer.toutiao.com/api/apps/ecpay/v1/create_refund"))
	var p = bytedance.TradeRefund{}
	p.OutOrderNo = "TEST2023112717521212345678"
	p.OutRefundNo = "TEST2023112717521212345698"
	p.Reason = "测试退款"
	p.RefundAmount = 1
	rsp, err := client.TradeRefund(p)
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("%+v", rsp)
}

// 退款结果查询
func TestClient_TradeRefundQuery(t *testing.T) {
	t.Log("========== TradeRefundQuery ==========")
	client.LoadOptionFunc(bytedance.WithApiHost("https://developer.toutiao.com/api/apps/ecpay/v1/query_refund"))
	var p = bytedance.TradeRefundQuery{}
	p.OutRefundNo = "TEST2023112717521212345678"
	rsp, err := client.TradeRefundQuery(p)
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("%+v", rsp)
}
