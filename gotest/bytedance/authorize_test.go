package bytedance

import (
	"github.com/Caiqm/go_payment_platforms/pkg/bytedance"
	"log"
	"testing"
)

var client *bytedance.Client

func init() {
	var err error
	client, err = bytedance.New("", "", bytedance.WithSalt(""))
	if err != nil {
		log.Fatalln(err)
	}
	client.OnReceivedData(func(method string, data []byte) {
		log.Println(method, string(data))
	})
}

// 小程序登录
func TestClient_Code2Session(t *testing.T) {
	t.Log("========== Code2Session ==========")
	client.LoadOptionFunc(bytedance.WithApiHost("https://open-sandbox.douyin.com/api/apps/v2/jscode2session"))
	var p = bytedance.Code2Session{}
	p.Code = "647f16afe0b44c49a8eb1cb3c02aXX31"
	rsp, err := client.Code2Session(p)
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("%v", rsp)
}

// 手机号解密
func TestClient_GetPhoneNumber(t *testing.T) {
	t.Log("========== GetPhoneNumber ==========")
	var p = bytedance.GetPhoneNumber{}
	p.Iv = ""
	p.EncryptedData = ""
	p.SessionKey = ""
	rsp, err := client.GetPhoneNumber(p)
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("%v", rsp)
}
