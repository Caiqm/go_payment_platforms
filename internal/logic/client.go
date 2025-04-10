package logic

import (
	"fmt"
	"github.com/Caiqm/alipay"
	"github.com/Caiqm/bytedance"
	wechat "github.com/Caiqm/wxpay-v2"
)

// 初始化支付宝
func NewPayClient(appId string, isProduction bool) (*alipay.Client, error) {
	return alipay.New(appId, "", isProduction)
}

// 初始化微信
func NewWeChatClient(appId, secret string) (*wechat.Client, error) {
	client, err := wechat.New(appId, secret)
	if err != nil {
		return nil, err
	}
	client.OnReceivedData(func(method string, data []byte) {
		fmt.Println("【微信小程序接收信息】", method, string(data))
	})
	return client, nil
}

// 初始化抖音
func NewDouYinClient(appId, secret string) (*bytedance.Client, error) {
	client, err := bytedance.New(appId, secret)
	if err != nil {
		return nil, err
	}
	client.OnReceivedData(func(method string, data []byte) {
		fmt.Println("【抖音小程序接收信息】", method, string(data))
	})
	return client, nil
}
