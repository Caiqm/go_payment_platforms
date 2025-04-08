package main

import (
	"github.com/Caiqm/go_payment_platforms/internal/logic"
	apb "github.com/Caiqm/go_payment_platforms/protoc/aliyun_pb"
	wpb "github.com/Caiqm/go_payment_platforms/protoc/weixin_pb"
	"google.golang.org/grpc"
	"log"
	"net"
)

func main() {
	// 监听端口
	lis, err := net.Listen("tcp", ":8092")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	// 注册服务
	// 支付宝服务
	apb.RegisterAliMpLoginServer(s, &logic.AliMpLogin{})
	apb.RegisterAliPayServer(s, &logic.AliPay{})
	// 微信服务
	wpb.RegisterWxMpLoginServer(s, &logic.WxMpLogin{})
	// 打印监听
	log.Printf("server listening at %v", lis.Addr())
	if err = s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
