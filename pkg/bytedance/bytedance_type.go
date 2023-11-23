package bytedance

import "fmt"

const (
	kContentType = "application/json;charset=utf-8"
	kTimeFormat  = "2006-01-02 15:04:05"
)

const (
	kFieldAppId        = "app_id"
	kFieldSecret       = "secret"
	kFieldSign         = "sign"
	kFieldMsgSignature = "msg_signature"
	kFieldTimestamp    = "timestamp"
	kFieldNonce        = "nonce"
	kFieldMsg          = "msg"
	kFieldErrNo        = "err_no"
)

type Param interface {
	// NeedSign 是否需要签名，有的接口不需要签名，比如：小程序登录与获取手机号接口
	NeedSign() bool

	// NeedSecret 是否需要密钥
	NeedSecret() bool
}

type AuxParam struct {
}

func (aux AuxParam) NeedSign() bool {
	return true
}

func (aux AuxParam) NeedSecret() bool {
	return false
}

// Error 支付错误类
type Error struct {
	ErrNo   int    `json:"err_no"`
	ErrTips string `json:"err_tips"`
}

func (e Error) Error() string {
	return fmt.Sprintf("%d - %s", e.ErrNo, e.ErrTips)
}
