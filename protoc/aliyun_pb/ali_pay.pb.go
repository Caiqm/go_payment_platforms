// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v4.25.0
// source: ali_pay.proto

package aliyun_pb

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// 支付请求参数
type PayRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Subject       string `protobuf:"bytes,1,opt,name=Subject,proto3" json:"Subject,omitempty"`
	OutTradeNo    string `protobuf:"bytes,2,opt,name=OutTradeNo,proto3" json:"OutTradeNo,omitempty"`
	TotalAmount   string `protobuf:"bytes,3,opt,name=TotalAmount,proto3" json:"TotalAmount,omitempty"`
	ProductCode   string `protobuf:"bytes,4,opt,name=ProductCode,proto3" json:"ProductCode,omitempty"`
	NotifyURL     string `protobuf:"bytes,5,opt,name=NotifyURL,proto3" json:"NotifyURL,omitempty"`
	ReturnURL     string `protobuf:"bytes,6,opt,name=ReturnURL,proto3" json:"ReturnURL,omitempty"`
	QuitURL       string `protobuf:"bytes,7,opt,name=QuitURL,proto3" json:"QuitURL,omitempty"`
	BuyerId       string `protobuf:"bytes,8,opt,name=BuyerId,proto3" json:"BuyerId,omitempty"`
	BuyerOpenId   string `protobuf:"bytes,9,opt,name=BuyerOpenId,proto3" json:"BuyerOpenId,omitempty"`
	OpBuyerOpenId string `protobuf:"bytes,10,opt,name=OpBuyerOpenId,proto3" json:"OpBuyerOpenId,omitempty"`
	OpAppId       string `protobuf:"bytes,11,opt,name=OpAppId,proto3" json:"OpAppId,omitempty"`
	AuthToken     string `protobuf:"bytes,12,opt,name=AuthToken,proto3" json:"AuthToken,omitempty"`
	AppAuthToken  string `protobuf:"bytes,13,opt,name=AppAuthToken,proto3" json:"AppAuthToken,omitempty"`
	AppId         string `protobuf:"bytes,14,opt,name=AppId,proto3" json:"AppId,omitempty"`
}

func (x *PayRequest) Reset() {
	*x = PayRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ali_pay_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PayRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PayRequest) ProtoMessage() {}

func (x *PayRequest) ProtoReflect() protoreflect.Message {
	mi := &file_ali_pay_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PayRequest.ProtoReflect.Descriptor instead.
func (*PayRequest) Descriptor() ([]byte, []int) {
	return file_ali_pay_proto_rawDescGZIP(), []int{0}
}

func (x *PayRequest) GetSubject() string {
	if x != nil {
		return x.Subject
	}
	return ""
}

func (x *PayRequest) GetOutTradeNo() string {
	if x != nil {
		return x.OutTradeNo
	}
	return ""
}

func (x *PayRequest) GetTotalAmount() string {
	if x != nil {
		return x.TotalAmount
	}
	return ""
}

func (x *PayRequest) GetProductCode() string {
	if x != nil {
		return x.ProductCode
	}
	return ""
}

func (x *PayRequest) GetNotifyURL() string {
	if x != nil {
		return x.NotifyURL
	}
	return ""
}

func (x *PayRequest) GetReturnURL() string {
	if x != nil {
		return x.ReturnURL
	}
	return ""
}

func (x *PayRequest) GetQuitURL() string {
	if x != nil {
		return x.QuitURL
	}
	return ""
}

func (x *PayRequest) GetBuyerId() string {
	if x != nil {
		return x.BuyerId
	}
	return ""
}

func (x *PayRequest) GetBuyerOpenId() string {
	if x != nil {
		return x.BuyerOpenId
	}
	return ""
}

func (x *PayRequest) GetOpBuyerOpenId() string {
	if x != nil {
		return x.OpBuyerOpenId
	}
	return ""
}

func (x *PayRequest) GetOpAppId() string {
	if x != nil {
		return x.OpAppId
	}
	return ""
}

func (x *PayRequest) GetAuthToken() string {
	if x != nil {
		return x.AuthToken
	}
	return ""
}

func (x *PayRequest) GetAppAuthToken() string {
	if x != nil {
		return x.AppAuthToken
	}
	return ""
}

func (x *PayRequest) GetAppId() string {
	if x != nil {
		return x.AppId
	}
	return ""
}

// 支付返回参数，只有APP与小程序
type PayReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PayURL     string `protobuf:"bytes,1,opt,name=PayURL,proto3" json:"PayURL,omitempty"`
	TradeNo    string `protobuf:"bytes,2,opt,name=TradeNo,proto3" json:"TradeNo,omitempty"`
	OutTradeNo string `protobuf:"bytes,3,opt,name=OutTradeNo,proto3" json:"OutTradeNo,omitempty"`
}

func (x *PayReply) Reset() {
	*x = PayReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ali_pay_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PayReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PayReply) ProtoMessage() {}

func (x *PayReply) ProtoReflect() protoreflect.Message {
	mi := &file_ali_pay_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PayReply.ProtoReflect.Descriptor instead.
func (*PayReply) Descriptor() ([]byte, []int) {
	return file_ali_pay_proto_rawDescGZIP(), []int{1}
}

func (x *PayReply) GetPayURL() string {
	if x != nil {
		return x.PayURL
	}
	return ""
}

func (x *PayReply) GetTradeNo() string {
	if x != nil {
		return x.TradeNo
	}
	return ""
}

func (x *PayReply) GetOutTradeNo() string {
	if x != nil {
		return x.OutTradeNo
	}
	return ""
}

// 支付查询
type PayQueryRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	OutTradeNo   string `protobuf:"bytes,1,opt,name=OutTradeNo,proto3" json:"OutTradeNo,omitempty"`
	TradeNo      string `protobuf:"bytes,2,opt,name=TradeNo,proto3" json:"TradeNo,omitempty"`
	AppAuthToken string `protobuf:"bytes,3,opt,name=AppAuthToken,proto3" json:"AppAuthToken,omitempty"`
	AppId        string `protobuf:"bytes,4,opt,name=AppId,proto3" json:"AppId,omitempty"`
}

func (x *PayQueryRequest) Reset() {
	*x = PayQueryRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ali_pay_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PayQueryRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PayQueryRequest) ProtoMessage() {}

func (x *PayQueryRequest) ProtoReflect() protoreflect.Message {
	mi := &file_ali_pay_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PayQueryRequest.ProtoReflect.Descriptor instead.
func (*PayQueryRequest) Descriptor() ([]byte, []int) {
	return file_ali_pay_proto_rawDescGZIP(), []int{2}
}

func (x *PayQueryRequest) GetOutTradeNo() string {
	if x != nil {
		return x.OutTradeNo
	}
	return ""
}

func (x *PayQueryRequest) GetTradeNo() string {
	if x != nil {
		return x.TradeNo
	}
	return ""
}

func (x *PayQueryRequest) GetAppAuthToken() string {
	if x != nil {
		return x.AppAuthToken
	}
	return ""
}

func (x *PayQueryRequest) GetAppId() string {
	if x != nil {
		return x.AppId
	}
	return ""
}

// 支付查询返回
type PayQueryReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	QueryResult []byte `protobuf:"bytes,1,opt,name=QueryResult,proto3" json:"QueryResult,omitempty"`
}

func (x *PayQueryReply) Reset() {
	*x = PayQueryReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ali_pay_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PayQueryReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PayQueryReply) ProtoMessage() {}

func (x *PayQueryReply) ProtoReflect() protoreflect.Message {
	mi := &file_ali_pay_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PayQueryReply.ProtoReflect.Descriptor instead.
func (*PayQueryReply) Descriptor() ([]byte, []int) {
	return file_ali_pay_proto_rawDescGZIP(), []int{3}
}

func (x *PayQueryReply) GetQueryResult() []byte {
	if x != nil {
		return x.QueryResult
	}
	return nil
}

// 退款请求
type RefundRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	OutTradeNo   string `protobuf:"bytes,1,opt,name=OutTradeNo,proto3" json:"OutTradeNo,omitempty"`
	TradeNo      string `protobuf:"bytes,2,opt,name=TradeNo,proto3" json:"TradeNo,omitempty"`
	RefundAmount string `protobuf:"bytes,3,opt,name=RefundAmount,proto3" json:"RefundAmount,omitempty"`
	RefundReason string `protobuf:"bytes,4,opt,name=RefundReason,proto3" json:"RefundReason,omitempty"`
	OutRequestNo string `protobuf:"bytes,5,opt,name=OutRequestNo,proto3" json:"OutRequestNo,omitempty"`
	AppId        string `protobuf:"bytes,6,opt,name=AppId,proto3" json:"AppId,omitempty"`
	AppAuthToken string `protobuf:"bytes,7,opt,name=AppAuthToken,proto3" json:"AppAuthToken,omitempty"`
}

func (x *RefundRequest) Reset() {
	*x = RefundRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ali_pay_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RefundRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RefundRequest) ProtoMessage() {}

func (x *RefundRequest) ProtoReflect() protoreflect.Message {
	mi := &file_ali_pay_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RefundRequest.ProtoReflect.Descriptor instead.
func (*RefundRequest) Descriptor() ([]byte, []int) {
	return file_ali_pay_proto_rawDescGZIP(), []int{4}
}

func (x *RefundRequest) GetOutTradeNo() string {
	if x != nil {
		return x.OutTradeNo
	}
	return ""
}

func (x *RefundRequest) GetTradeNo() string {
	if x != nil {
		return x.TradeNo
	}
	return ""
}

func (x *RefundRequest) GetRefundAmount() string {
	if x != nil {
		return x.RefundAmount
	}
	return ""
}

func (x *RefundRequest) GetRefundReason() string {
	if x != nil {
		return x.RefundReason
	}
	return ""
}

func (x *RefundRequest) GetOutRequestNo() string {
	if x != nil {
		return x.OutRequestNo
	}
	return ""
}

func (x *RefundRequest) GetAppId() string {
	if x != nil {
		return x.AppId
	}
	return ""
}

func (x *RefundRequest) GetAppAuthToken() string {
	if x != nil {
		return x.AppAuthToken
	}
	return ""
}

// 退款结果
type RefundReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	RefundResult []byte `protobuf:"bytes,1,opt,name=RefundResult,proto3" json:"RefundResult,omitempty"`
}

func (x *RefundReply) Reset() {
	*x = RefundReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ali_pay_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RefundReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RefundReply) ProtoMessage() {}

func (x *RefundReply) ProtoReflect() protoreflect.Message {
	mi := &file_ali_pay_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RefundReply.ProtoReflect.Descriptor instead.
func (*RefundReply) Descriptor() ([]byte, []int) {
	return file_ali_pay_proto_rawDescGZIP(), []int{5}
}

func (x *RefundReply) GetRefundResult() []byte {
	if x != nil {
		return x.RefundResult
	}
	return nil
}

// 退款查询请求
type RefundQueryRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	OutTradeNo   string `protobuf:"bytes,1,opt,name=OutTradeNo,proto3" json:"OutTradeNo,omitempty"`
	TradeNo      string `protobuf:"bytes,2,opt,name=TradeNo,proto3" json:"TradeNo,omitempty"`
	OutRequestNo string `protobuf:"bytes,3,opt,name=OutRequestNo,proto3" json:"OutRequestNo,omitempty"`
	AppId        string `protobuf:"bytes,4,opt,name=AppId,proto3" json:"AppId,omitempty"`
	AppAuthToken string `protobuf:"bytes,5,opt,name=AppAuthToken,proto3" json:"AppAuthToken,omitempty"`
}

func (x *RefundQueryRequest) Reset() {
	*x = RefundQueryRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ali_pay_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RefundQueryRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RefundQueryRequest) ProtoMessage() {}

func (x *RefundQueryRequest) ProtoReflect() protoreflect.Message {
	mi := &file_ali_pay_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RefundQueryRequest.ProtoReflect.Descriptor instead.
func (*RefundQueryRequest) Descriptor() ([]byte, []int) {
	return file_ali_pay_proto_rawDescGZIP(), []int{6}
}

func (x *RefundQueryRequest) GetOutTradeNo() string {
	if x != nil {
		return x.OutTradeNo
	}
	return ""
}

func (x *RefundQueryRequest) GetTradeNo() string {
	if x != nil {
		return x.TradeNo
	}
	return ""
}

func (x *RefundQueryRequest) GetOutRequestNo() string {
	if x != nil {
		return x.OutRequestNo
	}
	return ""
}

func (x *RefundQueryRequest) GetAppId() string {
	if x != nil {
		return x.AppId
	}
	return ""
}

func (x *RefundQueryRequest) GetAppAuthToken() string {
	if x != nil {
		return x.AppAuthToken
	}
	return ""
}

// 退款查询结果
type RefundQueryReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	RefundQueryResult []byte `protobuf:"bytes,1,opt,name=RefundQueryResult,proto3" json:"RefundQueryResult,omitempty"`
}

func (x *RefundQueryReply) Reset() {
	*x = RefundQueryReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ali_pay_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RefundQueryReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RefundQueryReply) ProtoMessage() {}

func (x *RefundQueryReply) ProtoReflect() protoreflect.Message {
	mi := &file_ali_pay_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RefundQueryReply.ProtoReflect.Descriptor instead.
func (*RefundQueryReply) Descriptor() ([]byte, []int) {
	return file_ali_pay_proto_rawDescGZIP(), []int{7}
}

func (x *RefundQueryReply) GetRefundQueryResult() []byte {
	if x != nil {
		return x.RefundQueryResult
	}
	return nil
}

var File_ali_pay_proto protoreflect.FileDescriptor

var file_ali_pay_proto_rawDesc = []byte{
	0x0a, 0x0d, 0x61, 0x6c, 0x69, 0x5f, 0x70, 0x61, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x09, 0x61, 0x6c, 0x69, 0x79, 0x75, 0x6e, 0x5f, 0x70, 0x62, 0x22, 0xb4, 0x03, 0x0a, 0x0a, 0x50,
	0x61, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x53, 0x75, 0x62,
	0x6a, 0x65, 0x63, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x53, 0x75, 0x62, 0x6a,
	0x65, 0x63, 0x74, 0x12, 0x1e, 0x0a, 0x0a, 0x4f, 0x75, 0x74, 0x54, 0x72, 0x61, 0x64, 0x65, 0x4e,
	0x6f, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x4f, 0x75, 0x74, 0x54, 0x72, 0x61, 0x64,
	0x65, 0x4e, 0x6f, 0x12, 0x20, 0x0a, 0x0b, 0x54, 0x6f, 0x74, 0x61, 0x6c, 0x41, 0x6d, 0x6f, 0x75,
	0x6e, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x54, 0x6f, 0x74, 0x61, 0x6c, 0x41,
	0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x20, 0x0a, 0x0b, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74,
	0x43, 0x6f, 0x64, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x50, 0x72, 0x6f, 0x64,
	0x75, 0x63, 0x74, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x4e, 0x6f, 0x74, 0x69, 0x66,
	0x79, 0x55, 0x52, 0x4c, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x4e, 0x6f, 0x74, 0x69,
	0x66, 0x79, 0x55, 0x52, 0x4c, 0x12, 0x1c, 0x0a, 0x09, 0x52, 0x65, 0x74, 0x75, 0x72, 0x6e, 0x55,
	0x52, 0x4c, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x52, 0x65, 0x74, 0x75, 0x72, 0x6e,
	0x55, 0x52, 0x4c, 0x12, 0x18, 0x0a, 0x07, 0x51, 0x75, 0x69, 0x74, 0x55, 0x52, 0x4c, 0x18, 0x07,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x51, 0x75, 0x69, 0x74, 0x55, 0x52, 0x4c, 0x12, 0x18, 0x0a,
	0x07, 0x42, 0x75, 0x79, 0x65, 0x72, 0x49, 0x64, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07,
	0x42, 0x75, 0x79, 0x65, 0x72, 0x49, 0x64, 0x12, 0x20, 0x0a, 0x0b, 0x42, 0x75, 0x79, 0x65, 0x72,
	0x4f, 0x70, 0x65, 0x6e, 0x49, 0x64, 0x18, 0x09, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x42, 0x75,
	0x79, 0x65, 0x72, 0x4f, 0x70, 0x65, 0x6e, 0x49, 0x64, 0x12, 0x24, 0x0a, 0x0d, 0x4f, 0x70, 0x42,
	0x75, 0x79, 0x65, 0x72, 0x4f, 0x70, 0x65, 0x6e, 0x49, 0x64, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0d, 0x4f, 0x70, 0x42, 0x75, 0x79, 0x65, 0x72, 0x4f, 0x70, 0x65, 0x6e, 0x49, 0x64, 0x12,
	0x18, 0x0a, 0x07, 0x4f, 0x70, 0x41, 0x70, 0x70, 0x49, 0x64, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x07, 0x4f, 0x70, 0x41, 0x70, 0x70, 0x49, 0x64, 0x12, 0x1c, 0x0a, 0x09, 0x41, 0x75, 0x74,
	0x68, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x41, 0x75,
	0x74, 0x68, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x22, 0x0a, 0x0c, 0x41, 0x70, 0x70, 0x41, 0x75,
	0x74, 0x68, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x41,
	0x70, 0x70, 0x41, 0x75, 0x74, 0x68, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x14, 0x0a, 0x05, 0x41,
	0x70, 0x70, 0x49, 0x64, 0x18, 0x0e, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x41, 0x70, 0x70, 0x49,
	0x64, 0x22, 0x5c, 0x0a, 0x08, 0x50, 0x61, 0x79, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x16, 0x0a,
	0x06, 0x50, 0x61, 0x79, 0x55, 0x52, 0x4c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x50,
	0x61, 0x79, 0x55, 0x52, 0x4c, 0x12, 0x18, 0x0a, 0x07, 0x54, 0x72, 0x61, 0x64, 0x65, 0x4e, 0x6f,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x54, 0x72, 0x61, 0x64, 0x65, 0x4e, 0x6f, 0x12,
	0x1e, 0x0a, 0x0a, 0x4f, 0x75, 0x74, 0x54, 0x72, 0x61, 0x64, 0x65, 0x4e, 0x6f, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0a, 0x4f, 0x75, 0x74, 0x54, 0x72, 0x61, 0x64, 0x65, 0x4e, 0x6f, 0x22,
	0x85, 0x01, 0x0a, 0x0f, 0x50, 0x61, 0x79, 0x51, 0x75, 0x65, 0x72, 0x79, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x1e, 0x0a, 0x0a, 0x4f, 0x75, 0x74, 0x54, 0x72, 0x61, 0x64, 0x65, 0x4e,
	0x6f, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x4f, 0x75, 0x74, 0x54, 0x72, 0x61, 0x64,
	0x65, 0x4e, 0x6f, 0x12, 0x18, 0x0a, 0x07, 0x54, 0x72, 0x61, 0x64, 0x65, 0x4e, 0x6f, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x54, 0x72, 0x61, 0x64, 0x65, 0x4e, 0x6f, 0x12, 0x22, 0x0a,
	0x0c, 0x41, 0x70, 0x70, 0x41, 0x75, 0x74, 0x68, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0c, 0x41, 0x70, 0x70, 0x41, 0x75, 0x74, 0x68, 0x54, 0x6f, 0x6b, 0x65,
	0x6e, 0x12, 0x14, 0x0a, 0x05, 0x41, 0x70, 0x70, 0x49, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x05, 0x41, 0x70, 0x70, 0x49, 0x64, 0x22, 0x31, 0x0a, 0x0d, 0x50, 0x61, 0x79, 0x51, 0x75,
	0x65, 0x72, 0x79, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x20, 0x0a, 0x0b, 0x51, 0x75, 0x65, 0x72,
	0x79, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x0b, 0x51,
	0x75, 0x65, 0x72, 0x79, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x22, 0xef, 0x01, 0x0a, 0x0d, 0x52,
	0x65, 0x66, 0x75, 0x6e, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1e, 0x0a, 0x0a,
	0x4f, 0x75, 0x74, 0x54, 0x72, 0x61, 0x64, 0x65, 0x4e, 0x6f, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0a, 0x4f, 0x75, 0x74, 0x54, 0x72, 0x61, 0x64, 0x65, 0x4e, 0x6f, 0x12, 0x18, 0x0a, 0x07,
	0x54, 0x72, 0x61, 0x64, 0x65, 0x4e, 0x6f, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x54,
	0x72, 0x61, 0x64, 0x65, 0x4e, 0x6f, 0x12, 0x22, 0x0a, 0x0c, 0x52, 0x65, 0x66, 0x75, 0x6e, 0x64,
	0x41, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x52, 0x65,
	0x66, 0x75, 0x6e, 0x64, 0x41, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x22, 0x0a, 0x0c, 0x52, 0x65,
	0x66, 0x75, 0x6e, 0x64, 0x52, 0x65, 0x61, 0x73, 0x6f, 0x6e, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0c, 0x52, 0x65, 0x66, 0x75, 0x6e, 0x64, 0x52, 0x65, 0x61, 0x73, 0x6f, 0x6e, 0x12, 0x22,
	0x0a, 0x0c, 0x4f, 0x75, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x4e, 0x6f, 0x18, 0x05,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x4f, 0x75, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x4e, 0x6f, 0x12, 0x14, 0x0a, 0x05, 0x41, 0x70, 0x70, 0x49, 0x64, 0x18, 0x06, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x05, 0x41, 0x70, 0x70, 0x49, 0x64, 0x12, 0x22, 0x0a, 0x0c, 0x41, 0x70, 0x70, 0x41,
	0x75, 0x74, 0x68, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c,
	0x41, 0x70, 0x70, 0x41, 0x75, 0x74, 0x68, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x22, 0x31, 0x0a, 0x0b,
	0x52, 0x65, 0x66, 0x75, 0x6e, 0x64, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x22, 0x0a, 0x0c, 0x52,
	0x65, 0x66, 0x75, 0x6e, 0x64, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0c, 0x52, 0x0c, 0x52, 0x65, 0x66, 0x75, 0x6e, 0x64, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x22,
	0xac, 0x01, 0x0a, 0x12, 0x52, 0x65, 0x66, 0x75, 0x6e, 0x64, 0x51, 0x75, 0x65, 0x72, 0x79, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1e, 0x0a, 0x0a, 0x4f, 0x75, 0x74, 0x54, 0x72, 0x61,
	0x64, 0x65, 0x4e, 0x6f, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x4f, 0x75, 0x74, 0x54,
	0x72, 0x61, 0x64, 0x65, 0x4e, 0x6f, 0x12, 0x18, 0x0a, 0x07, 0x54, 0x72, 0x61, 0x64, 0x65, 0x4e,
	0x6f, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x54, 0x72, 0x61, 0x64, 0x65, 0x4e, 0x6f,
	0x12, 0x22, 0x0a, 0x0c, 0x4f, 0x75, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x4e, 0x6f,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x4f, 0x75, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x4e, 0x6f, 0x12, 0x14, 0x0a, 0x05, 0x41, 0x70, 0x70, 0x49, 0x64, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x05, 0x41, 0x70, 0x70, 0x49, 0x64, 0x12, 0x22, 0x0a, 0x0c, 0x41, 0x70,
	0x70, 0x41, 0x75, 0x74, 0x68, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0c, 0x41, 0x70, 0x70, 0x41, 0x75, 0x74, 0x68, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x22, 0x40,
	0x0a, 0x10, 0x52, 0x65, 0x66, 0x75, 0x6e, 0x64, 0x51, 0x75, 0x65, 0x72, 0x79, 0x52, 0x65, 0x70,
	0x6c, 0x79, 0x12, 0x2c, 0x0a, 0x11, 0x52, 0x65, 0x66, 0x75, 0x6e, 0x64, 0x51, 0x75, 0x65, 0x72,
	0x79, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x11, 0x52,
	0x65, 0x66, 0x75, 0x6e, 0x64, 0x51, 0x75, 0x65, 0x72, 0x79, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74,
	0x32, 0xdf, 0x03, 0x0a, 0x06, 0x41, 0x6c, 0x69, 0x50, 0x61, 0x79, 0x12, 0x3b, 0x0a, 0x0b, 0x54,
	0x72, 0x61, 0x64, 0x65, 0x41, 0x70, 0x70, 0x50, 0x61, 0x79, 0x12, 0x15, 0x2e, 0x61, 0x6c, 0x69,
	0x79, 0x75, 0x6e, 0x5f, 0x70, 0x62, 0x2e, 0x50, 0x61, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x13, 0x2e, 0x61, 0x6c, 0x69, 0x79, 0x75, 0x6e, 0x5f, 0x70, 0x62, 0x2e, 0x50, 0x61,
	0x79, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x00, 0x12, 0x3c, 0x0a, 0x0c, 0x54, 0x72, 0x61, 0x64,
	0x65, 0x50, 0x61, 0x67, 0x65, 0x50, 0x61, 0x79, 0x12, 0x15, 0x2e, 0x61, 0x6c, 0x69, 0x79, 0x75,
	0x6e, 0x5f, 0x70, 0x62, 0x2e, 0x50, 0x61, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x13, 0x2e, 0x61, 0x6c, 0x69, 0x79, 0x75, 0x6e, 0x5f, 0x70, 0x62, 0x2e, 0x50, 0x61, 0x79, 0x52,
	0x65, 0x70, 0x6c, 0x79, 0x22, 0x00, 0x12, 0x3b, 0x0a, 0x0b, 0x54, 0x72, 0x61, 0x64, 0x65, 0x43,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x12, 0x15, 0x2e, 0x61, 0x6c, 0x69, 0x79, 0x75, 0x6e, 0x5f, 0x70,
	0x62, 0x2e, 0x50, 0x61, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x13, 0x2e, 0x61,
	0x6c, 0x69, 0x79, 0x75, 0x6e, 0x5f, 0x70, 0x62, 0x2e, 0x50, 0x61, 0x79, 0x52, 0x65, 0x70, 0x6c,
	0x79, 0x22, 0x00, 0x12, 0x3b, 0x0a, 0x0b, 0x54, 0x72, 0x61, 0x64, 0x65, 0x57, 0x61, 0x70, 0x50,
	0x61, 0x79, 0x12, 0x15, 0x2e, 0x61, 0x6c, 0x69, 0x79, 0x75, 0x6e, 0x5f, 0x70, 0x62, 0x2e, 0x50,
	0x61, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x13, 0x2e, 0x61, 0x6c, 0x69, 0x79,
	0x75, 0x6e, 0x5f, 0x70, 0x62, 0x2e, 0x50, 0x61, 0x79, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x00,
	0x12, 0x44, 0x0a, 0x0a, 0x54, 0x72, 0x61, 0x64, 0x65, 0x51, 0x75, 0x65, 0x72, 0x79, 0x12, 0x1a,
	0x2e, 0x61, 0x6c, 0x69, 0x79, 0x75, 0x6e, 0x5f, 0x70, 0x62, 0x2e, 0x50, 0x61, 0x79, 0x51, 0x75,
	0x65, 0x72, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x18, 0x2e, 0x61, 0x6c, 0x69,
	0x79, 0x75, 0x6e, 0x5f, 0x70, 0x62, 0x2e, 0x50, 0x61, 0x79, 0x51, 0x75, 0x65, 0x72, 0x79, 0x52,
	0x65, 0x70, 0x6c, 0x79, 0x22, 0x00, 0x12, 0x41, 0x0a, 0x0b, 0x54, 0x72, 0x61, 0x64, 0x65, 0x52,
	0x65, 0x66, 0x75, 0x6e, 0x64, 0x12, 0x18, 0x2e, 0x61, 0x6c, 0x69, 0x79, 0x75, 0x6e, 0x5f, 0x70,
	0x62, 0x2e, 0x52, 0x65, 0x66, 0x75, 0x6e, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x16, 0x2e, 0x61, 0x6c, 0x69, 0x79, 0x75, 0x6e, 0x5f, 0x70, 0x62, 0x2e, 0x52, 0x65, 0x66, 0x75,
	0x6e, 0x64, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x00, 0x12, 0x57, 0x0a, 0x17, 0x54, 0x72, 0x61,
	0x64, 0x65, 0x46, 0x61, 0x73, 0x74, 0x50, 0x61, 0x79, 0x52, 0x65, 0x66, 0x75, 0x6e, 0x64, 0x51,
	0x75, 0x65, 0x72, 0x79, 0x12, 0x1d, 0x2e, 0x61, 0x6c, 0x69, 0x79, 0x75, 0x6e, 0x5f, 0x70, 0x62,
	0x2e, 0x52, 0x65, 0x66, 0x75, 0x6e, 0x64, 0x51, 0x75, 0x65, 0x72, 0x79, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x1b, 0x2e, 0x61, 0x6c, 0x69, 0x79, 0x75, 0x6e, 0x5f, 0x70, 0x62, 0x2e,
	0x52, 0x65, 0x66, 0x75, 0x6e, 0x64, 0x51, 0x75, 0x65, 0x72, 0x79, 0x52, 0x65, 0x70, 0x6c, 0x79,
	0x22, 0x00, 0x42, 0x18, 0x5a, 0x16, 0x2e, 0x2e, 0x2f, 0x61, 0x6c, 0x69, 0x79, 0x75, 0x6e, 0x5f,
	0x70, 0x62, 0x3b, 0x61, 0x6c, 0x69, 0x79, 0x75, 0x6e, 0x5f, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_ali_pay_proto_rawDescOnce sync.Once
	file_ali_pay_proto_rawDescData = file_ali_pay_proto_rawDesc
)

func file_ali_pay_proto_rawDescGZIP() []byte {
	file_ali_pay_proto_rawDescOnce.Do(func() {
		file_ali_pay_proto_rawDescData = protoimpl.X.CompressGZIP(file_ali_pay_proto_rawDescData)
	})
	return file_ali_pay_proto_rawDescData
}

var file_ali_pay_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_ali_pay_proto_goTypes = []interface{}{
	(*PayRequest)(nil),         // 0: aliyun_pb.PayRequest
	(*PayReply)(nil),           // 1: aliyun_pb.PayReply
	(*PayQueryRequest)(nil),    // 2: aliyun_pb.PayQueryRequest
	(*PayQueryReply)(nil),      // 3: aliyun_pb.PayQueryReply
	(*RefundRequest)(nil),      // 4: aliyun_pb.RefundRequest
	(*RefundReply)(nil),        // 5: aliyun_pb.RefundReply
	(*RefundQueryRequest)(nil), // 6: aliyun_pb.RefundQueryRequest
	(*RefundQueryReply)(nil),   // 7: aliyun_pb.RefundQueryReply
}
var file_ali_pay_proto_depIdxs = []int32{
	0, // 0: aliyun_pb.AliPay.TradeAppPay:input_type -> aliyun_pb.PayRequest
	0, // 1: aliyun_pb.AliPay.TradePagePay:input_type -> aliyun_pb.PayRequest
	0, // 2: aliyun_pb.AliPay.TradeCreate:input_type -> aliyun_pb.PayRequest
	0, // 3: aliyun_pb.AliPay.TradeWapPay:input_type -> aliyun_pb.PayRequest
	2, // 4: aliyun_pb.AliPay.TradeQuery:input_type -> aliyun_pb.PayQueryRequest
	4, // 5: aliyun_pb.AliPay.TradeRefund:input_type -> aliyun_pb.RefundRequest
	6, // 6: aliyun_pb.AliPay.TradeFastPayRefundQuery:input_type -> aliyun_pb.RefundQueryRequest
	1, // 7: aliyun_pb.AliPay.TradeAppPay:output_type -> aliyun_pb.PayReply
	1, // 8: aliyun_pb.AliPay.TradePagePay:output_type -> aliyun_pb.PayReply
	1, // 9: aliyun_pb.AliPay.TradeCreate:output_type -> aliyun_pb.PayReply
	1, // 10: aliyun_pb.AliPay.TradeWapPay:output_type -> aliyun_pb.PayReply
	3, // 11: aliyun_pb.AliPay.TradeQuery:output_type -> aliyun_pb.PayQueryReply
	5, // 12: aliyun_pb.AliPay.TradeRefund:output_type -> aliyun_pb.RefundReply
	7, // 13: aliyun_pb.AliPay.TradeFastPayRefundQuery:output_type -> aliyun_pb.RefundQueryReply
	7, // [7:14] is the sub-list for method output_type
	0, // [0:7] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_ali_pay_proto_init() }
func file_ali_pay_proto_init() {
	if File_ali_pay_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_ali_pay_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PayRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_ali_pay_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PayReply); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_ali_pay_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PayQueryRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_ali_pay_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PayQueryReply); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_ali_pay_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RefundRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_ali_pay_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RefundReply); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_ali_pay_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RefundQueryRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_ali_pay_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RefundQueryReply); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_ali_pay_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_ali_pay_proto_goTypes,
		DependencyIndexes: file_ali_pay_proto_depIdxs,
		MessageInfos:      file_ali_pay_proto_msgTypes,
	}.Build()
	File_ali_pay_proto = out.File
	file_ali_pay_proto_rawDesc = nil
	file_ali_pay_proto_goTypes = nil
	file_ali_pay_proto_depIdxs = nil
}
