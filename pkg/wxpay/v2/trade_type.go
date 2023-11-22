package v2

// 请求参数
type Trade struct {
	AuxParam
	NotifyUrl string `xml:"notify_url" json:"notify_url"`
	// 必须业务参数
	Body           string `xml:"body" json:"body"`
	OutTradeNo     string `xml:"out_trade_no" json:"out_trade_no"`
	TotalFee       int    `xml:"total_fee" json:"total_fee"`
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

/* 查询订单 */

// TradeOrderQuery 查询订单 https://pay.weixin.qq.com/wiki/doc/api/wxa/wxa_api.php?chapter=9_2
type TradeOrderQuery struct {
	AuxParam
	OutTradeNo    string `xml:"out_trade_no,omitempty"`
	TransactionId string `xml:"transaction_id,omitempty"`
}

func (t TradeOrderQuery) ReturnType() string {
	return "xml"
}

// TradeOrderQueryRsp 查询订单响应参数
type TradeOrderQueryRsp struct {
	PayError
	AppID          string `xml:"appid"`
	MchID          string `xml:"mch_id"`
	NonceStr       string `xml:"nonce_str"`
	Sign           string `xml:"sign"`
	DeviceInfo     string `xml:"device_info"`
	OpenId         string `xml:"openid"`
	IsSubscribe    string `xml:"is_subscribe"`
	TradeType      string `xml:"trade_type"`
	TradeState     string `xml:"trade_state"`
	BankType       string `xml:"bank_type"`
	TotalFee       int    `xml:"total_fee"`
	CashFee        int    `xml:"cash_fee"`
	TransactionId  string `xml:"transaction_id"`
	OutTradeNo     string `xml:"out_trade_no"`
	Attach         string `xml:"attach"`
	TimeEnd        string `xml:"time_end"`
	TradeStateDesc string `xml:"trade_state_desc"`
}

/* 关闭订单 */

// TradeCloseOrder 关闭订单 https://pay.weixin.qq.com/wiki/doc/api/wxa/wxa_api.php?chapter=9_3
type TradeCloseOrder struct {
	AuxParam
	OutTradeNo string `xml:"out_trade_no,omitempty"`
}

func (t TradeCloseOrder) ReturnType() string {
	return "xml"
}

// TradeCloseOrderRsp 关闭订单响应参数
type TradeCloseOrderRsp struct {
	PayError
	AppID     string `xml:"appid"`
	MchID     string `xml:"mch_id"`
	NonceStr  string `xml:"nonce_str"`
	Sign      string `xml:"sign"`
	ResultMsg string `xml:"result_msg"`
}

/* 申请退款 */

// TradeRefund 申请退款 https://pay.weixin.qq.com/wiki/doc/api/wxa/wxa_api.php?chapter=9_4
type TradeRefund struct {
	AuxParam
	OutTradeNo    string `xml:"out_trade_no,omitempty"`
	TransactionId string `xml:"transaction_id,omitempty"`
	OutRefundNo   string `xml:"out_refund_no"`
	TotalFee      int    `xml:"total_fee"`
	RefundFee     int    `xml:"refund_fee"`
	RefundDesc    string `xml:"refund_desc"`
	NotifyUrl     string `xml:"notify_url"`
}

func (t TradeRefund) NeedTlsCert() bool {
	return true
}

func (t TradeRefund) ReturnType() string {
	return "xml"
}

// TradeRefundRsp TradeRefund 申请退款响应参数
type TradeRefundRsp struct {
	PayError
	AppID         string `xml:"appid"`
	MchID         string `xml:"mch_id"`
	NonceStr      string `xml:"nonce_str"`
	Sign          string `xml:"sign"`
	TransactionId string `xml:"transaction_id"`
	OutTradeNo    string `xml:"out_trade_no"`
	OutRefundNo   string `xml:"out_refund_no"`
	RefundId      string `xml:"refund_id"`
	RefundFee     int    `xml:"refund_fee"`
	TotalFee      int    `xml:"total_fee"`
	CashFee       int    `xml:"cash_fee"`
}

/* 查询退款 */

// TradeRefundQuery 查询退款 https://pay.weixin.qq.com/wiki/doc/api/wxa/wxa_api.php?chapter=9_5
type TradeRefundQuery struct {
	AuxParam
	TransactionId string `xml:"transaction_id"`
	OutTradeNo    string `xml:"out_trade_no"`
	OutRefundNo   string `xml:"out_refund_no"`
	RefundId      string `xml:"refund_id"`
	Offset        int    `xml:"offset,omitempty"`
}

func (t TradeRefundQuery) ReturnType() string {
	return "xml"
}

// TradeRefundQueryRsp 查询退款响应参数
type TradeRefundQueryRsp struct {
	PayError
	AppID                string `xml:"appid"`
	MchID                string `xml:"mch_id"`
	NonceStr             string `xml:"nonce_str"`
	Sign                 string `xml:"sign"`
	TotalRefundCount     int    `xml:"total_refund_count"`
	TransactionId        string `xml:"transaction_id"`
	OutTradeNo           string `xml:"out_trade_no"`
	TotalFee             int    `xml:"total_fee"`
	SettlementTotalFee   int    `xml:"settlement_total_fee"`
	FeeType              string `xml:"fee_type"`
	CashFee              int    `xml:"cash_fee"`
	RefundCount          int    `xml:"refund_count"`
	OutRefundNo1         string `xml:"out_refund_no_1"`
	RefundId1            string `xml:"refund_id_1"`
	RefundChannel1       string `xml:"refund_channel_1"`
	RefundFee1           int    `xml:"refund_fee_1"`
	RefundFee            int    `xml:"refund_fee"`
	CouponRefundFee      int    `xml:"coupon_refund_fee"`
	SettlementRefundFee1 int    `xml:"settlement_refund_fee_1"`
	RefundStatus1        string `xml:"refund_status_1"`
	RefundAccount1       string `xml:"refund_account_1"`
	RefundRecvAccout1    string `xml:"refund_recv_accout_1"`
	RefundSuccessTime1   string `xml:"refund_success_time_1"`
	CashRefundFee        int    `xml:"cash_refund_fee"`
}
