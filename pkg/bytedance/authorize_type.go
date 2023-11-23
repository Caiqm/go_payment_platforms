package bytedance

type Applet struct {
	AuxParam
}

func (a Applet) NeedSign() bool {
	return false
}

func (a Applet) NeedSecret() bool {
	return true
}

// GetAccessToken 小程序的全局唯一调用凭据 https://developer.open-douyin.com/docs/resource/zh-CN/mini-app/develop/server/interface-request-credential/non-user-authorization/get-access-token
type GetAccessToken struct {
	Applet
	GrantType string `json:"grant_type"` // 获取 access_token 时值为 client_credential
}

// GetAccessTokenRsp 小程序的全局唯一调用凭据响应参数
type GetAccessTokenRsp struct {
	Error
	Data struct {
		AccessToken string `json:"access_token"` // 获取的 access_token
		ExpiresIn   int64  `json:"expires_in"`   // access_token 有效时间，单位：秒
	} `json:"data"`
}

// Code2Session 小程序登录 https://developer.open-douyin.com/docs/resource/zh-CN/mini-app/develop/server/log-in/code-2-session
type Code2Session struct {
	Applet
	AnonymousCode string `json:"anonymous_code"` // login 接口返回的匿名登录凭证
	Code          string `json:"code"`           // login 接口返回的登录凭证
}

// Code2SessionRsp 小程序登录响应参数
type Code2SessionRsp struct {
	Error
	Data struct {
		SessionKey      string `json:"session_key"`      // 会话密钥，如果请求时有 code 参数才会返回
		Openid          string `json:"openid"`           // 用户在当前小程序的 ID，如果请求时有 code 参数才会返回
		AnonymousOpenid string `json:"anonymous_openid"` // 匿名用户在当前小程序的 ID，如果请求时有 anonymous_code 参数才会返回
		Unionid         string `json:"unionid"`          // 用户在小程序平台的唯一标识符，请求时有 code 参数才会返回。如果开发者拥有多个小程序，可通过 unionid 来区分用户的唯一性。
	} `json:"data"`
}

// GetPhoneNumber 获取手机号 https://developer.open-douyin.com/docs/resource/zh-CN/mini-app/develop/guide/open-capabilities/acquire-phone-number-acquire/
type GetPhoneNumber struct {
	EncryptedData string `json:"encrypted_data"`
	SessionKey    string `json:"session_key"`
	Iv            string `json:"iv"`
}

// GetPhoneNumberRsp 获取手机号响应参数
type GetPhoneNumberRsp struct {
	PhoneNumber     string        `json:"phoneNumber"`     // 用户绑定的手机号（国外手机号会有区号）
	PurePhoneNumber string        `json:"purePhoneNumber"` // 没有区号的手机号
	CountryCode     string        `json:"countryCode"`     // 区号
	Watermark       WatermarkData `json:"watermark"`
}

type WatermarkData struct {
	Appid     string `json:"appid"`
	Timestamp int    `json:"timestamp"`
}
