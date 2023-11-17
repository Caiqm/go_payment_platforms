package logic

import "github.com/Caiqm/go_payment_platforms/pkg/alipay"

// 初始化
func NewPayClient(appId string, isProduction bool) *alipay.Client {
	return alipay.New(appId, "", isProduction)
}
