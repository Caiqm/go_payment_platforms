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
	conn, err := grpc.Dial("localhost:8082", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := apb.NewAliMpLoginClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	r, err := c.SystemOauthToken(ctx, &apb.MpLoginRequest{
		GrantType: "authorization_code",
		Code:      "4b203fe6c11548bcabd8da5bb087a83b",
		AppId:     "2021003110607584",
	})
	if err != nil {
		log.Fatalf("find err: %v", err)
	}
	log.Printf("result: %s", r.String())
}
