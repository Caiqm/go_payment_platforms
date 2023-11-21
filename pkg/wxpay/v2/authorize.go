package v2

import (
	"fmt"
	"net/url"
)

// GetAccessToken 接口调用凭据 https://developers.weixin.qq.com/miniprogram/dev/OpenApiDoc/mp-access-token/getAccessToken.html
// GET https://api.weixin.qq.com/cgi-bin/token
func (c *Client) GetAccessToken() (result *GetAccessTokenRsp, err error) {
	var u url.Values
	u.Add("appid", c.appId)
	u.Add("secret", c.secret)
	u.Add("grant_type", "client_credential")
	c.host = fmt.Sprintf("%s?%s", c.host, u.Encode())
	err = c.doRequest("GET", nil, &result)
	return
}

// Code2Session 小程序登录 https://developers.weixin.qq.com/miniprogram/dev/OpenApiDoc/user-login/code2Session.html
// GET https://api.weixin.qq.com/sns/jscode2session
func (c *Client) Code2Session(jsCode string) (result *Code2SessionRsp, err error) {
	var u url.Values
	u.Add("appid", c.appId)
	u.Add("secret", c.secret)
	u.Add("grant_type", "authorization_code")
	u.Add("js_code", jsCode)
	c.host = fmt.Sprintf("%s?%s", c.host, u.Encode())
	err = c.doRequest("GET", nil, &result)
	return
}

// GetPhoneNumber 获取手机号 https://developers.weixin.qq.com/miniprogram/dev/OpenApiDoc/user-info/phone-number/getPhoneNumber.html
// POST https://api.weixin.qq.com/wxa/business/getuserphonenumber?access_token=ACCESS_TOKEN
func (c *Client) GetPhoneNumber(param GetPhoneNumber) (result *GetPhoneNumberRsp, err error) {
	err = c.doRequest("POST", param, &result)
	return
}
