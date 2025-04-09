package main

import (
	"context"
	apb "github.com/Caiqm/go_payment_platforms/protoc/aliyun_pb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
	"time"
)

func main() {
	conn, err := grpc.NewClient("localhost:8092", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer func(conn *grpc.ClientConn) {
		err1 := conn.Close()
		if err1 != nil {
			log.Fatalf("conn close err: %v", err1)
		}
	}(conn)
	aliPayMpLogin(conn)
}

// 支付宝小程序登录
func aliPayMpLogin(conn *grpc.ClientConn) {
	c := apb.NewAliMpLoginClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	r, err := c.SystemOauthToken(ctx, &apb.MpLoginRequest{
		GrantType: "authorization_code",
		Code:      "4b203fe6c11548bcabd8da5bb087a83b",
		AppId:     "",
	})
	if err != nil {
		log.Fatalf("find err: %v", err)
	}
	log.Printf("result: %s", r.String())
}

// 支付宝小程序支付
func aliPayTradeCreate(conn *grpc.ClientConn) {
	c := apb.NewAliPayClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	r, err := c.TradeCreate(ctx, &apb.PayRequest{
		Subject:     "测试订单",
		TotalAmount: "0.01",
		OutTradeNo:  "202201011111",
		BuyerOpenId: "",
		NotifyURL:   "",
		AppId:       "",
	})
	if err != nil {
		log.Fatalf("trade create err: %v", err)
	}
	log.Printf("result: %s", r.String())
}
