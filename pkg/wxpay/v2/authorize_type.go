package v2

// Code2SessionRsp 小程序登录返回参数 https://developers.weixin.qq.com/miniprogram/dev/OpenApiDoc/user-login/code2Session.html
type Code2SessionRsp struct {
	AppletError
	OpenId     string `json:"openid"`
	SessionKey string `json:"session_key"`
	UnionId    string `json:"unionid"`
}

// GetAccessTokenRsp 接口调用凭据 https://developers.weixin.qq.com/miniprogram/dev/OpenApiDoc/mp-access-token/getAccessToken.html
type GetAccessTokenRsp struct {
	AppletError
	AccessToken string `json:"access_token"`
	ExpiresIn   int32  `json:"expires_in"`
}

// GetPhoneNumber 获取手机号 https://developers.weixin.qq.com/miniprogram/dev/OpenApiDoc/user-info/phone-number/getPhoneNumber.html
type GetPhoneNumber struct {
	AuxParam
	Code string `json:"code"`
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
	PhoneInfo PhoneNumberInfo `json:"phone_info"`
}

type PhoneNumberInfo struct {
	PhoneNumber     string     `json:"phoneNumber"`
	PurePhoneNumber string     `json:"purePhoneNumber"`
	CountryCode     string     `json:"countryCode"`
	Watermark       WatermarkS `json:"watermark"`
}

type WatermarkS struct {
	Timestamp int32  `json:"timestamp"`
	Appid     string `json:"appid"`
}
