package logic

import (
	"context"
	"fmt"
	"github.com/Caiqm/alipay"
	pb "github.com/Caiqm/go_payment_platforms/protoc/aliyun_pb"
)

type AliMpLogin struct {
	pb.UnimplementedAliMpLoginServer
}

// 换取授权访问令牌
func (mpl *AliMpLogin) SystemOauthToken(ctx context.Context, req *pb.MpLoginRequest) (rpy *pb.MpLoginReply, err error) {
	client, _ := NewPayClient(req.AppId, false)
	client.OnReceivedData(func(method string, data []byte) {
		fmt.Println("【支付宝小程序登录信息】", method, string(data))
	})
	var p = alipay.SystemOauthToken{}
	p.GrantType = req.GrantType
	p.Code = req.Code
	at, err := client.SystemOauthToken(p)
	if err != nil {
		return
	}
	rpy.AccessToken = at.AccessToken
	rpy.UserId = at.UserId
	rpy.OpenId = at.OpenId
	return
}

// 支付宝会员授权信息查询接口
func (mpl *AliMpLogin) UserInfoShare(ctx context.Context, req *pb.MpUserInfoRequest) (rpy *pb.MpUserInfoReply, err error) {
	client, _ := NewPayClient(req.AppId, false)
	client.OnReceivedData(func(method string, data []byte) {
		fmt.Println(method, string(data))
	})
	var p = alipay.UserInfoShare{}
	p.AuthToken = req.AuthToken
	param, err := client.UserInfoShare(p)
	if err != nil {
		return
	}
	rpy.UserId = param.UserId
	rpy.City = param.City
	rpy.Province = param.Province
	rpy.NickName = param.NickName
	rpy.Avatar = param.Avatar
	rpy.Gender = param.Gender
	return
}

// 小程序获取会员手机号
func (mpl *AliMpLogin) DecodePhoneNumber(ctx context.Context, req *pb.MpPhoneNumberRequest) (rpy *pb.MpPhoneNumberReply, err error) {
	client, _ := NewPayClient(req.AppId, false)
	client.OnReceivedData(func(method string, data []byte) {
		fmt.Println(method, string(data))
	})
	err = client.SetEncryptKey(req.AuthToken)
	if err != nil {
		return
	}
	result, err := client.DecodePhoneNumber(req.Response)
	if err != nil {
		return
	}
	rpy.PhoneNumber = result.Mobile
	return
}
