package baidu

// TradePolymerPayment 百度收银台 https://smartprogram.baidu.com/docs/develop/api/open/payment_swan-requestPolymerPayment/#orderInfo
type TradePolymerPayment struct {
	TotalAmount     string `json:"totalAmount"`             // 订单金额（单位：人民币分）。注：小程序测试包测试金额不可超过 1000 分
	TpOrderId       string `json:"tpOrderId"`               // 小程序开发者系统创建的唯一订单 ID ，当支付状态发生变化时，会通过此订单 ID 通知开发者
	NotifyUrl       string `json:"notifyUrl"`               // 通知开发者支付状态的回调地址，必须是合法的 URL ，与开发者平台填写的支付回调地址作用一致，未填写的以平台回调地址为准
	DealTitle       string `json:"dealTitle"`               // 订单的名称
	SignFieldsRange string `json:"signFieldsRange"`         // 用于区分验签字段范围，signFieldsRange 的值：0：原验签字段 appKey+dealId+tpOrderId；1：包含 totalAmount 的验签，验签字段包括appKey+dealId+tpOrderId+totalAmount。固定值为 1 。
	BizInfo         string `json:"bizInfo"`                 // 订单详细信息，需要是一个可解析为 JSON Object 的字符串
	PayResultUrl    string `json:"payResultUrl,omitempty"`  // 当前页面 path。Web 态小程序支付成功后跳回的页面路径，例如：'/pages/payResult/payResult'
	InlinePaySign   string `json:"inlinePaySign,omitempty"` // 内嵌支付组件返回的支付信息加密 key，与 内嵌支付组件配套使用，此值不传或者传空时默认调起支付面板
	PromotionTag    string `json:"promotionTag,omitempty"`  // 平台营销信息，此处传可使用平台券的 spuid，支持通过英文逗号分割传入多个 spuid 值，仅与百度合作平台类目券的开发者需要填写
}

// TradePolymerPaymentRsp 百度收银台响应参数
type TradePolymerPaymentRsp struct {
	DealId  string `json:"dealId"`  // 跳转百度收银台支付必带参数之一，是百度收银台的财务结算凭证，与账号绑定的结算协议一一对应，每笔交易将结算到 dealId 对应的协议主体
	AppKey  string `json:"appKey"`  // 支付能力开通后分配的支付 appKey，用以表示应用身份的唯一 ID ，在应用审核通过后进行分配，一经分配后不会发生更改，来唯一确定一个应用
	RsaSign string `json:"rsaSign"` // 对appKey+dealId+totalAmount+tpOrderId进行 RSA 加密后的签名，防止订单被伪造
	TradePolymerPayment
}

type BizInfo struct {
	TpData TpData `json:"tpData"` // bizInfo 组装键值对集合
}

type TpData struct {
	AppKey      string                 `json:"appKey"`       // 表示应用身份的唯一 ID
	DealId      string                 `json:"dealId"`       // 百度收银台的财务结算凭证
	TpOrderId   string                 `json:"tpOrderId"`    // 业务方唯一订单号
	TotalAmount string                 `json:"totalAmount"`  // 订单总金额（单位：分）
	ReturnData  map[string]interface{} `json:"returnData"`   // 业务方用于透传的业务变量 支付成功后会以 query 形式注入到 payResultUrl 页面中（query 可以在页面的 onLoad 生命周期内获取）
	DisplayData string                 `json:"display_data"` // 收银台定制页面展示属性，非定制业务请置空 用于支付页面展示订单详细信息
}
