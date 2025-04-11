package logic

import (
	"context"
	"fmt"
	"github.com/Caiqm/bytedance"
	pb "github.com/Caiqm/go_payment_platforms/protoc/douyin_pb"
	"github.com/goccy/go-json"
)

type DyPay struct {
	pb.UnimplementedDyPayServer
}

// 通用交易系统支付
func (dy *DyPay) TradeTransaction(ctx context.Context, req *pb.PayTransactionRequest) (rpy *pb.PayTransactionReply, err error) {
	client, err := NewDouYinClient(req.AppId, req.Secret)
	if err != nil {
		err = fmt.Errorf("初始化失败：%s", err.Error())
		return
	}
	client.LoadOptionFunc(bytedance.WithKeyVersion(req.KeyVersion), bytedance.WithPrivateKey(req.PrivateKey))
	var limitPayWay []int
	// 限制支付方式
	if req.LimitPayWayList != nil && len(req.GetLimitPayWayList()) > 0 {
		for _, v := range req.LimitPayWayList {
			limitPayWay = append(limitPayWay, int(v))
		}
	}
	var skuList []bytedance.SkuList
	// 商品类型
	if req.SkuList != nil && len(req.GetSkuList()) > 0 {
		for _, v := range req.SkuList {
			skuList = append(skuList, bytedance.SkuList{
				SkuId:      v.SkuId,
				Title:      v.Title,
				Price:      int(v.Price * 100),
				Quantity:   int(v.Quantity),
				Type:       int(v.Type),
				TagGroupId: v.TagGroupId,
				SkuAttr:    v.SkuAttr,
				ImageList:  v.ImageList,
				EntrySchema: bytedance.OrderEntrySchema{
					Path:   v.EntrySchema.Path,
					Params: v.EntrySchema.Params,
				},
			})
		}
	}
	var p = bytedance.TradeTransaction{
		OutOrderNo:   req.OutOrderNo,
		TotalAmount:  int(req.TotalAmount * 100),
		PayNotifyUrl: req.PayNotifyUrl,
		OrderEntrySchema: bytedance.OrderEntrySchema{
			Path:   req.OrderEntrySchema.Path,
			Params: req.OrderEntrySchema.Params,
		},
		MerchantUid:      req.MerchantUid,
		PayExpireSeconds: int(req.PayExpireSeconds),
		LimitPayWayList:  limitPayWay,
		SkuList:          skuList,
	}
	payRsp, err := client.TradeTransactionPay(p)
	if err != nil {
		err = fmt.Errorf("支付失败：%s", err.Error())
		return
	}
	rpy = &pb.PayTransactionReply{
		ByteAuthorization: payRsp.ByteAuthorization,
		Data:              payRsp.Data,
	}
	return
}

// 订单查询
func (dy *DyPay) TradeTransactionOrderQuery(ctx context.Context, req *pb.TransactionOrderQueryRequest) (rpy *pb.TransactionOrderQueryReply, err error) {
	client, err := NewDouYinClient(req.AppId, req.Secret)
	if err != nil {
		err = fmt.Errorf("初始化失败：%s", err.Error())
		return
	}
	client.LoadOptionFunc(bytedance.WithApiHost("https://open.douyin.com/api/trade_basic/v1/developer/order_query/"))
	dyLogin := DyMpLogin{AppId: req.AppId, AppSecret: req.Secret}
	accessToken, err := dyLogin.GetAccessToken(ctx)
	if err != nil {
		return
	}
	var p = bytedance.TradeTransactionQuery{}
	if req.GetOrderId() != "" {
		p.OrderId = req.OrderId
	} else if req.GetOutOrderNo() != "" {
		p.OutOrderNo = req.OutOrderNo
	} else {
		err = fmt.Errorf("订单号参数错误")
		return
	}
	p.AccessToken = accessToken
	queryRsp, err := client.TradeTransactionQuery(p)
	if err != nil {
		err = fmt.Errorf("查询失败：%s", err.Error())
		return
	}
	r, err := json.Marshal(queryRsp)
	if err != nil {
		err = fmt.Errorf("序列化失败：%s", err.Error())
		return
	}
	rpy.TransactionOrderQueryResult = r
	return
}

// 发起退款
func (dy *DyPay) TradeRefundTransaction(ctx context.Context, req *pb.RefundTransactionRequest) (rpy *pb.RefundTransactionReply, err error) {
	client, err := NewDouYinClient(req.AppId, req.Secret)
	if err != nil {
		err = fmt.Errorf("初始化失败：%s", err.Error())
		return
	}
	client.LoadOptionFunc(bytedance.WithApiHost("https://open.douyin.com/api/trade_basic/v1/developer/refund_create/"))
	dyLogin := DyMpLogin{AppId: req.AppId, AppSecret: req.Secret}
	accessToken, err := dyLogin.GetAccessToken(ctx)
	if err != nil {
		return
	}
	var refundReason []bytedance.ReasonItem
	// 退款原因
	for _, v := range req.RefundReason {
		refundReason = append(refundReason, bytedance.ReasonItem{
			Text: v.Text,
			Code: v.Code,
		})
	}
	var itemOrderDetail []bytedance.ItemOrderDetail
	// 退款商品详情
	if req.GetItemOrderDetail() != nil && len(req.GetItemOrderDetail()) > 0 {
		for _, detail := range req.ItemOrderDetail {
			itemOrderDetail = append(itemOrderDetail, bytedance.ItemOrderDetail{
				ItemOrderId:  detail.ItemOrderId,
				RefundAmount: int64(detail.RefundAmount * 100),
			})
		}
	}
	var p = bytedance.TradeRefundCreate{}
	p.OrderId = req.OrderId
	p.OutRefundNo = req.OutRefundNo
	p.RefundTotalAmount = int64(req.RefundTotalAmount * 100)
	p.RefundReason = refundReason
	p.RefundAll = req.RefundAll
	p.AccessToken = accessToken
	p.NotifyUrl = req.RefundNotifyUrl
	p.CpExtra = req.CpExtra
	// 部分退款，需商品单信息
	if req.GetRefundAll() != true {
		p.ItemOrderDetail = itemOrderDetail
	}
	return
}

// 退款查询
func (dy *DyPay) TradeRefundTransactionQuery(ctx context.Context, req *pb.RefundTransactionQueryRequest) (rpy *pb.RefundTransactionQueryReply, err error) {
	client, err := NewDouYinClient(req.AppId, req.Secret)
	if err != nil {
		err = fmt.Errorf("初始化失败：%s", err.Error())
		return
	}
	client.LoadOptionFunc(bytedance.WithApiHost("https://open.douyin.com/api/trade_basic/v1/developer/refund_query/"))
	dyLogin := DyMpLogin{AppId: req.AppId, AppSecret: req.Secret}
	accessToken, err := dyLogin.GetAccessToken(ctx)
	if err != nil {
		return
	}
	var p = bytedance.TradeRefundCreateQuery{}
	if req.GetOrderId() != "" {
		p.OrderId = req.OrderId
	} else if req.GetOutRefundNo() != "" {
		p.OutRefundNo = req.OutRefundNo
	} else if req.GetRefundId() != "" {
		p.RefundId = req.RefundId
	} else {
		err = fmt.Errorf("退款单号参数错误")
		return
	}
	p.AccessToken = accessToken
	queryRsp, err := client.TradeRefundCreateQuery(p)
	if err != nil {
		err = fmt.Errorf("查询失败：%s", err.Error())
		return
	}
	r, err := json.Marshal(queryRsp)
	if err != nil {
		err = fmt.Errorf("序列化失败：%s", err.Error())
		return
	}
	rpy.RefundTransactionQueryResult = r
	return
}
