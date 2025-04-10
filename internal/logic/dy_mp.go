package logic

import (
	"context"
	"fmt"
	"github.com/Caiqm/bytedance"
	pb "github.com/Caiqm/go_payment_platforms/protoc/douyin_pb"
	"strconv"
)

type DyMpLogin struct {
	AppId     string
	AppSecret string
	pb.UnimplementedDyMpLoginServer
}

// 获取access_token
func (dy *DyMpLogin) GetAccessToken(ctx context.Context) (accessToken string, err error) {
	client, err := NewDouYinClient(dy.AppId, dy.AppSecret)
	if err != nil {
		return "", fmt.Errorf("初始化失败：%s", err.Error())
	}
	client.LoadOptionFunc(bytedance.WithApiHost("https://open.douyin.com/oauth/client_token/"))
	var p bytedance.ClientToken
	p.GrantType = "client_credential"
	tokenRsp, err := client.ClientToken(p)
	if err != nil {
		return "", fmt.Errorf("获取access_token失败：%s", err.Error())
	}
	accessToken = tokenRsp.Data.AccessToken
	return
}

// 获取openid，小程序登录
func (dy *DyMpLogin) Code2Session(ctx context.Context, req *pb.MpLoginRequest) (rsp *pb.MpLoginReply, err error) {
	client, err := NewDouYinClient(req.AppId, req.Secret)
	if err != nil {
		return nil, fmt.Errorf("初始化失败：%s", err.Error())
	}
	client.LoadOptionFunc(bytedance.WithApiHost("https://developer.toutiao.com/api/apps/v2/jscode2session"))
	var p bytedance.Code2Session
	p.Code = req.Code
	if req.GetAnonymousCode() != "" {
		p.AnonymousCode = req.AnonymousCode
	}
	codeRsp, err := client.Code2Session(p)
	if err != nil {
		return nil, fmt.Errorf("获取openid失败：%s", err.Error())
	}
	rsp = &pb.MpLoginReply{
		AnonymousOpenid: codeRsp.Data.AnonymousOpenid,
		OpenId:          codeRsp.Data.Openid,
		SessionKey:      codeRsp.Data.SessionKey,
		UnionId:         codeRsp.Data.Unionid,
	}
	return
}

// 获取手机号
func (dy *DyMpLogin) GetPhoneNumber(ctx context.Context, req *pb.MpPhoneNumberRequest) (rsp *pb.MpPhoneNumberReply, err error) {
	dy.AppId = req.AppId
	dy.AppSecret = req.Secret
	client, err := NewDouYinClient(req.AppId, req.Secret)
	if err != nil {
		return nil, fmt.Errorf("初始化失败：%s", err.Error())
	}
	// 获取access_token
	accessToken, err := dy.GetAccessToken(ctx)
	if err != nil {
		return
	}
	client.LoadOptionFunc(bytedance.WithApiHost("https://open.douyin.com/api/apps/v1/get_phonenumber_info/"))
	// 获取手机号
	var p = bytedance.GetUserPhoneNumber{}
	p.Code = req.Code
	p.AccessToken = accessToken
	phoneRsp, err := client.GetPhoneNumberByCode(p)
	if err != nil {
		return
	}
	// 返回手机号信息
	countryCode, _ := strconv.Atoi(phoneRsp.CountryCode)
	rsp.PhoneNumber = phoneRsp.PhoneNumber
	rsp.PurePhoneNumber = phoneRsp.PurePhoneNumber
	rsp.CountryCode = int32(countryCode)
	rsp.Watermark = &pb.Watermark{
		AppId:     phoneRsp.Watermark.Appid,
		Timestamp: int32(phoneRsp.Watermark.Timestamp),
	}
	return
}
