package baidu

import (
	"github.com/Caiqm/go_payment_platforms/pkg/baidu"
	"log"
	"testing"
)

var client *baidu.Client

func init() {
	var err error
	client, err = baidu.New("", "")
	if err != nil {
		log.Fatalln(err)
	}
	client.OnReceivedData(func(method string, data []byte) {
		log.Println(method, string(data))
	})
}

// 用户登陆凭证
func TestClient_SessionKey(t *testing.T) {
	t.Log("========== SessionKey ==========")
	client.LoadOptionFunc(baidu.WithApiHost("https://openapi.baidu.com/rest/2.0/smartapp/getsessionkey"))
	var p baidu.SessionKey
	p.Code = "123456"
	p.AccessToken = "123456123456123456123456123456123456123456123456123456"
	r, err := client.SessionKey(p)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(r)
}
