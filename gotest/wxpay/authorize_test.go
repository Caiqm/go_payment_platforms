package wxpay

import (
	"fmt"
	"github.com/Caiqm/go_payment_platforms/pkg/wxpay/v2"
	"log"
	"testing"
)

var client *v2.Client

func init() {
	var err error
	client, err = v2.New("", "")
	if err != nil {
		log.Fatalln(err)
	}
	client.OnReceivedData(func(method string, data []byte) {
		fmt.Println(method, string(data))
	})
}

func TestClient_Code2Session(t *testing.T) {
	t.Log("========== Code2Session ==========")
	client.LoadOptionFunc(v2.WithJsCodeHost())
	var p v2.Code2Session
	p.JsCode = "" // 前端获取的code值
	r, err := client.Code2Session(p)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(r)
}

func TestClient_GetPhoneNumber(t *testing.T) {
	t.Log("========== GetPhoneNumber ==========")
	client.LoadOptionFunc(v2.WithApiHost("https://api.weixin.qq.com/wxa/business/getuserphonenumber?access_token=ACCESS_TOKEN"))
	var p v2.GetPhoneNumber
	p.Code = "" // 前端获取的code值
	r, err := client.GetPhoneNumber(p)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(r)
}
