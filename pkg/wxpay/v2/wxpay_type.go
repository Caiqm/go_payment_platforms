package v2

const (
	kContentType = "application/x-www-form-urlencoded;charset=utf-8"
	kTimeFormat  = "2006-01-02 15:04:05"
)

const (
	kFieldAppId = "appid"
	kFieldMchId = "mch_id"
	kFieldSign  = "sign"
)

type Param interface {
	// APIName 用于提供访问的 method，即接口名称
	APIName() string

	// Params 公共请求参数
	Params() map[string]string

	// NeedSign 是否需要签名，有的接口不需要签名，比如：小程序登录与获取手机号接口
	NeedSign() bool
}
