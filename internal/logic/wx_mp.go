package logic

import (
	"context"
	"fmt"
	pb "github.com/Caiqm/go_payment_platforms/protoc/weixin_pb"
	wechat "github.com/Caiqm/wxpay-v2"
	"strconv"
)

type WxMpLogin struct {
	pb.UnimplementedWxMpLoginServer
	AppId     string
	AppSecret string
}

// 获取access_token
func (wx *WxMpLogin) GetAccessToken(ctx context.Context) (accessToken string, err error) {
	client, err := NewWeChatClient(wx.AppId, wx.AppSecret)
	if err != nil {
		return "", fmt.Errorf("初始化失败：%s", err.Error())
	}
	client.LoadOptionFunc(wechat.WithApiHost("https://api.weixin.qq.com/cgi-bin/token"))
	var p wechat.GetAccessToken
	p.GrantType = "client_credential"
	tokenRsp, err := client.GetAccessToken(p)
	if err != nil {
		return "", fmt.Errorf("获取access_token失败：%s", err.Error())
	}
	accessToken = tokenRsp.AccessToken
	return
}

// 小程序登录
func (wx *WxMpLogin) Code2Session(ctx context.Context, req *pb.MpLoginRequest) (rsp *pb.MpLoginReply, err error) {
	client, err := NewWeChatClient(req.AppId, req.Secret)
	if err != nil {
		return nil, fmt.Errorf("初始化失败：%s", err.Error())
	}
	client.LoadOptionFunc(wechat.WithJsCodeHost())
	var p = wechat.Code2Session{}
	p.JsCode = req.JsCode
	if req.GetGrantType() == "" {
		p.GrantType = "authorization_code"
	} else {
		p.GrantType = req.GrantType
	}
	loginRsp, err := client.Code2Session(p)
	if err != nil {
		return
	}
	rsp.SessionKey = loginRsp.SessionKey
	rsp.OpenId = loginRsp.OpenId
	rsp.UnionId = loginRsp.UnionId
	return
}

// 小程序获取手机号
func (wx *WxMpLogin) GetPhoneNumber(ctx context.Context, req *pb.MpPhoneNumberRequest) (rsp *pb.MpPhoneNumberReply, err error) {
	wx.AppId = req.AppId
	wx.AppSecret = req.Secret
	client, err := NewWeChatClient(req.AppId, req.Secret)
	if err != nil {
		return nil, fmt.Errorf("初始化失败：%s", err.Error())
	}
	// 获取access_token
	accessToken, err := wx.GetAccessToken(ctx)
	if err != nil {
		return
	}
	client.LoadOptionFunc(wechat.WithApiHost(fmt.Sprintf("https://api.weixin.qq.com/wxa/business/getuserphonenumber?access_token=%s", accessToken)))
	// 获取手机号
	var p = wechat.GetPhoneNumber{}
	p.Code = req.Code
	if req.GetOpenId() != "" {
		p.Openid = req.OpenId
	}
	phoneRsp, err := client.GetPhoneNumber(p)
	if err != nil {
		return
	}
	// 返回手机号信息
	countryCode, _ := strconv.Atoi(phoneRsp.PhoneInfo.CountryCode)
	rsp.PhoneNumber = phoneRsp.PhoneInfo.PhoneNumber
	rsp.PurePhoneNumber = phoneRsp.PhoneInfo.PurePhoneNumber
	rsp.CountryCode = int32(countryCode)
	rsp.Watermark = &pb.Watermark{
		AppId:     phoneRsp.PhoneInfo.Watermark.Appid,
		Timestamp: int32(phoneRsp.PhoneInfo.Watermark.Timestamp),
	}
	return
}
