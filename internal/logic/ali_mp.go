package logic

import (
	"context"
	"fmt"
	"github.com/Caiqm/go_payment_platforms/pkg/alipay"
	pb "github.com/Caiqm/go_payment_platforms/protoc/aliyun_pb"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type AliMpLogin struct {
	pb.UnimplementedAliMpLoginServer
}

const (
	appId  = "2021003110607584"
	priKey = "MIIEvgIBADANBgkqhkiG9w0BAQEFAASCBKgwggSkAgEAAoIBAQDQgsQwb2UQ67KCcrEL8KUKrz1J0MaE2Dv1vm230SEz1rjGww+s/JHUrvmw0vb9MhpBXO0COhjtROmqbaol5bOchgEB/gd4l34WFJwsrlCcIUkMmmxvl6JF1hGgBLy5UUnrL61H2CkiMtsChWICPtZV/xHV45BgpIOnAGYse8x6B2n9MV6pPvmCPHjIOlPWzkWLc0Lo4jK40kY031E/aZYNi+W7QI7L9ExhNms7qYFGH3W3nn81sgmiZxIw3slyuX6h1KQO/O1wSnMrpEFbrcJG0qtULrPH7eoL6yHhXXxJdKFwSFN75ErZqvAS+ZxAgZV2OKEmcfHYOe6OBOT9hWkFAgMBAAECggEAePd0ZHkr1B45iO6LAldRGJrD3sAte8KLPq2Er8C3F4D53+4oeP5qiE4e89PgkNHxY3xK0CDudzCks0g+DxBtHGCt0v8STIbvElZQqKdUgs/YtmQchr0c2pEv1nsJFiYv0S7uw6CHdJy2bIb/bjgQVwHM0h7ckwS+kDNuE16bKSt5zlNUnw5FKziADdbosKBwRGdl7ZX+LfULw3c8jBGWfsBY3U0ei7re1gbN3HIo08KKtDUqfxC//Fm4Bag2z6c24iCu9MMmowtELFhvkbKKKidZfmb4+C/kf/qFhhXm47tNAib2Fs6WNUQv9jTKs34DU/q47FxonlTDnIyttq7SYQKBgQDskWJqVVBBawx7lKGgToYk5pZpk91nZyEsVIjY3IngGMSFJqxLhggf5yrAJUNqp/nZcMT++p60DwqdVVCC+r+FucprwgMQxt1j8a37BI91NLdx5vowfa3Y8Bn3HtB+Wm/vsVhYDRKVm6us3jG0dNabB+mnKxTkFqCGvSMqwM+inQKBgQDho2OwfTNVt2QN8BO2PEcEa3jJKn4sfDCTDIHN9k7gpgryXI9IA6cB3shAFULAR4y9EZkl0R1ffirS+jeAnCxtO3PnBa2j7V7fD2534bgLy5ETlo6Lr2cibxCyh33el+yT+cWKqwraFH+MEcHAgzwh0B7GVODLS5y39wQlOxH/iQKBgQCSzWFyNQD5p5F2evSVg0URTxqV59FkIoZeeRL26FyS+SuyqR4QO9t82LoZxMYCQLz2J1nIFQQ5HoBqxA+TSEFjnbi0iA/Y8F9gbxCVDe5sQEt6XmhxgPxJ/C1QwgmF7185MidtdKdMue6d9sCHiF3IlO908nIjXLs/Ac358O6orQKBgCpsb/fk8lvxaSx9sRcYajnYgxM/nA19JBzhwx4Ya2gtj+1VndLqbOIbIwd5d91zFnEN9/92O9GShTa5uOnDc07uWLJdDPK/VJEX86syk0oUcih+rDpzNi7xNvB6LR5G7kI8OfoCgr3SENEjHYy2n/2zhbXAi6ttgf/Mv8jIPHARAoGBAIGisKsE10myQwrkbE7PPleXlBQidC4DCO9sWJ17qBWvcxAbNjh64o2YD6nEhIVWYzJN1+P7BfVt+xSs5Y8Qhk6gjgWCh5SENAqsZlIVlmsDzi2PYvG07X3mDB0+tDjuMc2vXyOoZHtDvbfzG8YSGSoHUArA3K31ehicnv+0Ba2Y"
)

// 换取授权访问令牌
func (mpl *AliMpLogin) SystemOauthToken(ctx context.Context, req *pb.MpLoginRequest) (rpy *pb.MpLoginReply, err error) {
	client := alipay.New(req.AppId, priKey, true)
	client.OnReceivedData(func(method string, data []byte) {
		fmt.Println(method, string(data))
	})
	var p = alipay.SystemOauthToken{}
	p.GrantType = req.GrantType
	p.Code = req.Code
	at, err := client.SystemOauthToken(p)
	if err != nil {
		err = status.Errorf(codes.Unavailable, err.Error())
		return
	}
	fmt.Println(at)
	rpy.AccessToken = at.AccessToken
	rpy.UserId = at.UserId
	rpy.OpenId = at.OpenId
	return
}

// 支付宝会员授权信息查询接口
func (mpl *AliMpLogin) UserInfoShare(ctx context.Context, req *pb.MpUserInfoRequest) (rpy *pb.MpUserInfoReply, err error) {
	return nil, status.Errorf(codes.Unimplemented, "method UserInfoShare not implemented")
}
