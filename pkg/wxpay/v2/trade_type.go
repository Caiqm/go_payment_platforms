package v2

// 请求参数
type Trade struct {
	AuxParam
	NotifyUrl string `xml:"notify_url" json:"notify_url"`
	// 必须业务参数
	Body           string `xml:"body" json:"body"`
	OutTradeNo     string `xml:"out_trade_no" json:"out_trade_no"`
	TotalFee       string `xml:"total_fee" json:"total_fee"`
	SpbillCreateIp string `xml:"spbill_create_ip" json:"spbill_create_ip"`
	TradeType      string `xml:"trade_type" json:"trade_type"`
	// 选填，额外参数
	Attach string `xml:"attach" json:"attach"`
}

/* 小程序支付 */

// TradeApplet 小程序统一下单接口 https://pay.weixin.qq.com/wiki/doc/api/wxa/wxa_api.php?chapter=9_1
type TradeApplet struct {
	Trade
	OpenId string `xml:"openid" json:"openid"`
}

func (t TradeApplet) ReturnType() string {
	return "xml"
}

// TradeAppletRsp 小程序统一下单响应参数
type TradeAppletRsp struct {
	PayError
	AppID     string `xml:"appid"`
	MchID     string `xml:"mch_id"`
	NonceStr  string `xml:"nonce_str"`
	Sign      string `xml:"sign"`
	TradeType string `xml:"trade_type"`
	PrepayId  string `xml:"prepay_id"`
	CodeUrl   string `xml:"code_url,omitempty"`
}

/* APP支付 */

// TradeApp APP统一下单接口 https://pay.weixin.qq.com/wiki/doc/api/app/app.php?chapter=9_1
type TradeApp struct {
	Trade
}

func (t TradeApp) ReturnType() string {
	return "xml"
}

// TradeAppRsp APP统一下单接口响应参数
type TradeAppRsp struct {
	PayError
	AppID     string `xml:"appid"`
	MchID     string `xml:"mch_id"`
	NonceStr  string `xml:"nonce_str"`
	Sign      string `xml:"sign"`
	TradeType string `xml:"trade_type"`
	PrepayId  string `xml:"prepay_id"`
}

/* JSAPI支付·微信内H5支付 */

// TradeJSAPI 微信内H5统一下单接口 https://pay.weixin.qq.com/wiki/doc/api/jsapi.php?chapter=9_1
type TradeJSAPI struct {
	Trade
	OpenId string `xml:"openid" json:"openid"`
}

func (t TradeJSAPI) ReturnType() string {
	return "xml"
}

// TradeJSAPIRsp TradeJSAPI 微信内H5统一下单响应参数
type TradeJSAPIRsp struct {
	PayError
	AppID     string `xml:"appid"`
	MchID     string `xml:"mch_id"`
	NonceStr  string `xml:"nonce_str"`
	Sign      string `xml:"sign"`
	TradeType string `xml:"trade_type"`
	PrepayId  string `xml:"prepay_id"`
	CodeUrl   string `xml:"code_url,omitempty"`
}

/* Native支付·微信扫码支付 */

// TradeNative Native统一下单接口 https://pay.weixin.qq.com/wiki/doc/api/native.php?chapter=9_1
type TradeNative struct {
	Trade
}

func (t TradeNative) ReturnType() string {
	return "xml"
}

// TradeNativeRsp Native统一下单接口响应参数
type TradeNativeRsp struct {
	PayError
	AppID     string `xml:"appid"`
	MchID     string `xml:"mch_id"`
	NonceStr  string `xml:"nonce_str"`
	Sign      string `xml:"sign"`
	TradeType string `xml:"trade_type"`
	PrepayId  string `xml:"prepay_id"`
	CodeUrl   string `xml:"code_url"`
}

/* MWEB支付，微信外H5支付 */

// TradeWap H5支付 https://pay.weixin.qq.com/wiki/doc/api/H5.php?chapter=9_20&index=1
type TradeWap struct {
	Trade
	SceneInfo SceneInfo `xml:"scene_info"`
}

func (t TradeWap) ReturnType() string {
	return "xml"
}

type SceneInfo struct {
	H5Info struct {
		Type        string `json:"type"`
		WapURL      string `json:"wap_url,omitempty"`
		WapName     string `json:"wap_name,omitempty"`
		AppName     string `json:"app_name,omitempty"`
		BundleId    string `json:"bundle_id,omitempty"`
		PackageName string `json:"package_name,omitempty"`
	} `json:"h5_info,omitempty"`
	StoreInfo struct {
		ID       string `json:"id"`
		Name     string `json:"name"`
		AreaCode string `json:"area_code"`
		Address  string `json:"address"`
	} `json:"store_info,omitempty"`
}

// TradeWapRsp Native统一下单接口响应参数
type TradeWapRsp struct {
	PayError
	AppID     string `xml:"appid"`
	MchID     string `xml:"mch_id"`
	NonceStr  string `xml:"nonce_str"`
	Sign      string `xml:"sign"`
	TradeType string `xml:"trade_type"`
	PrepayId  string `xml:"prepay_id"`
	MWebUrl   string `xml:"mweb_url"`
}
