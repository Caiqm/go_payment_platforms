package logic

import "github.com/Caiqm/alipay"

// 初始化
func NewPayClient(appId string, isProduction bool) (*alipay.Client, error) {
	return alipay.New(appId, "", isProduction)
}
