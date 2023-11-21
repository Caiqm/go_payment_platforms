package v2

// TradeApplet 小程序统一下单接口 https://pay.weixin.qq.com/wiki/doc/api/wxa/wxa_api.php?chapter=9_1
// POST https://api.mch.weixin.qq.com/pay/unifiedorder
func (c *Client) TradeApplet(param TradeApplet) (result *TradeAppletRsp, err error) {
	if param.TradeType == "" {
		param.TradeType = "JSAPI"
	}
	err = c.doRequest("POST", param, &result)
	return
}

// TradeApp APP统一下单接口 https://pay.weixin.qq.com/wiki/doc/api/app/app.php?chapter=9_1
// POST https://api.mch.weixin.qq.com/pay/unifiedorder
func (c *Client) TradeApp(param TradeApp) (result *TradeAppRsp, err error) {
	if param.TradeType == "" {
		param.TradeType = "APP"
	}
	err = c.doRequest("POST", param, &result)
	return
}

// TradeJSAPI 微信内H5统一下单 https://pay.weixin.qq.com/wiki/doc/api/jsapi.php?chapter=9_1
// POST https://api.mch.weixin.qq.com/pay/unifiedorder
func (c *Client) TradeJSAPI(param TradeJSAPI) (result *TradeJSAPIRsp, err error) {
	if param.TradeType == "" {
		param.TradeType = "JSAPI"
	}
	err = c.doRequest("POST", param, &result)
	return
}

// TradeNative Native统一下单接口 https://pay.weixin.qq.com/wiki/doc/api/native.php?chapter=9_1
// POST https://api.mch.weixin.qq.com/pay/unifiedorder
func (c *Client) TradeNative(param TradeNative) (result *TradeNativeRsp, err error) {
	if param.TradeType == "" {
		param.TradeType = "NATIVE"
	}
	err = c.doRequest("POST", param, &result)
	return
}

// TradeWap H5支付 https://pay.weixin.qq.com/wiki/doc/api/H5.php?chapter=9_20&index=1
// POST https://api.mch.weixin.qq.com/pay/unifiedorder
func (c *Client) TradeWap(param TradeWap) (result *TradeWapRsp, err error) {
	if param.TradeType == "" {
		param.TradeType = "MWEB"
	}
	err = c.doRequest("POST", param, &result)
	return
}
