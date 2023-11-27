package wxpay

import (
	wxpay "github.com/Caiqm/go_payment_platforms/pkg/wxpay/v2"
	"testing"
)

func TestClient_TradeApplet(t *testing.T) {
	t.Log("========== TradeApplet ==========")
	client.LoadApiHost(wxpay.WithPayHost())
	var p wxpay.TradeApplet
	p.Body = "支付测试"
	p.OutTradeNo = "TEST2023112717521212345678"
	p.TotalFee = 1
	p.SpbillCreateIp = "59.41.161.29"
	p.OpenId = "ob_lN5etmVRNYAKSvV7AZf7ajyyc"
	r, err := client.TradeApplet(p)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(r)
}
