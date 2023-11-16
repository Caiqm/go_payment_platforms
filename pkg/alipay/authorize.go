package alipay

// SystemOauthToken 换取授权访问令牌接口 https://docs.open.alipay.com/api_9/alipay.system.oauth.token
func (c *Client) SystemOauthToken(param SystemOauthToken) (result *SystemOauthTokenRsp, err error) {
	err = c.doRequest("POST", param, &result)
	return result, err
}

// UserInfoShare 支付宝会员授权信息查询接口 https://docs.open.alipay.com/api_2/alipay.user.info.share
func (c *Client) UserInfoShare(param UserInfoShare) (result *UserInfoShareRsp, err error) {
	err = c.doRequest("POST", param, &result)
	return result, err
}
