package baidu

import "net/url"

// TradePolymerPayment 百度收银台 https://smartprogram.baidu.com/docs/develop/api/open/payment_swan-requestPolymerPayment/#orderInfo
func (c *Client) TradePolymerPayment(param TradePolymerPayment) (result TradePolymerPaymentRsp, err error) {
	result.DealId = c.dealId
	result.AppKey = c.appKey
	// 签名字段
	data := url.Values{}
	data.Set("appKey", result.AppKey)
	data.Set("dealId", result.DealId)
	data.Set("tpOrderId", param.TpOrderId)
	// signFieldsRange 的值：0：原验签字段 appKey+dealId+tpOrderId；1：包含 totalAmount 的验签，验签字段包括appKey+dealId+tpOrderId+totalAmount。固定值为 1
	if param.SignFieldsRange == "" || param.SignFieldsRange == "1" {
		param.SignFieldsRange = "1"
		data.Set("totalAmount", param.TotalAmount)
	} else if param.SignFieldsRange != "0" {
		param.SignFieldsRange = "0"
	}
	result.TradePolymerPayment = param
	// 进行 RSA 加密后的签名，防止订单被伪造
	if result.RsaSign, err = c.sign(data); err != nil {
		return
	}
	return
}
