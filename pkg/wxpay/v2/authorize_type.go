package v2

// Code2SessionRsp 小程序登录返回参数 https://developers.weixin.qq.com/miniprogram/dev/OpenApiDoc/user-login/code2Session.html
type Code2SessionRsp struct {
	AppletError
	OpenId     string `json:"openid"`      // 用户唯一标识
	SessionKey string `json:"session_key"` // 会话密钥
	UnionId    string `json:"unionid"`     // 用户在开放平台的唯一标识符，若当前小程序已绑定到微信开放平台账号下会返回
}

// GetAccessTokenRsp 接口调用凭据 https://developers.weixin.qq.com/miniprogram/dev/OpenApiDoc/mp-access-token/getAccessToken.html
type GetAccessTokenRsp struct {
	AppletError
	AccessToken string `json:"access_token"` // 获取到的凭证
	ExpiresIn   int    `json:"expires_in"`   // 凭证有效时间，单位：秒。目前是7200秒之内的值
}

// GetPhoneNumber 获取手机号 https://developers.weixin.qq.com/miniprogram/dev/OpenApiDoc/user-info/phone-number/getPhoneNumber.html
type GetPhoneNumber struct {
	AuxParam
	Code   string `json:"code"`             // 手机号获取凭证
	Openid string `json:"openid,omitempty"` // 用户OPENID
}

func (pn GetPhoneNumber) NeedSign() bool {
	return false
}

func (pn GetPhoneNumber) NeedVerify() bool {
	return false
}

// GetPhoneNumberRsp 获取手机号响应参数
type GetPhoneNumberRsp struct {
	AppletError
	PhoneInfo PhoneNumberInfo `json:"phone_info"` // 用户手机号信息
}

type PhoneNumberInfo struct {
	PhoneNumber     string     `json:"phoneNumber"`     // 用户绑定的手机号（国外手机号会有区号）
	PurePhoneNumber string     `json:"purePhoneNumber"` // 没有区号的手机号
	CountryCode     string     `json:"countryCode"`     // 区号
	Watermark       WatermarkS `json:"watermark"`       // 数据水印
}

type WatermarkS struct {
	Timestamp int    `json:"timestamp"` // 用户获取手机号操作的时间戳
	Appid     string `json:"appid"`     // 小程序appid
}
