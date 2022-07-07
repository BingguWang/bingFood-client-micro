// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.19.1
// source: v1/bingfood_service.proto

package v1

import (
	_ "google.golang.org/genproto/googleapis/api/annotations"
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

type UserLoginOrRegisterRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserMobile string `protobuf:"bytes,1,opt,name=userMobile,proto3" json:"userMobile,omitempty"` // 手机号
	LoginType  uint32 `protobuf:"varint,2,opt,name=loginType,proto3" json:"loginType,omitempty"`  // 登录方式 0 手机号登录 1 密码登录 2 第三方微信登录
	ValidCode  string `protobuf:"bytes,3,opt,name=validCode,proto3" json:"validCode,omitempty"`   // 账号密码登录不需要验证码
	Password   string `protobuf:"bytes,4,opt,name=password,proto3" json:"password,omitempty"`     // 登录密码
}

func (x *UserLoginOrRegisterRequest) Reset() {
	*x = UserLoginOrRegisterRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v1_bingfood_service_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserLoginOrRegisterRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserLoginOrRegisterRequest) ProtoMessage() {}

func (x *UserLoginOrRegisterRequest) ProtoReflect() protoreflect.Message {
	mi := &file_v1_bingfood_service_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserLoginOrRegisterRequest.ProtoReflect.Descriptor instead.
func (*UserLoginOrRegisterRequest) Descriptor() ([]byte, []int) {
	return file_v1_bingfood_service_proto_rawDescGZIP(), []int{0}
}

func (x *UserLoginOrRegisterRequest) GetUserMobile() string {
	if x != nil {
		return x.UserMobile
	}
	return ""
}

func (x *UserLoginOrRegisterRequest) GetLoginType() uint32 {
	if x != nil {
		return x.LoginType
	}
	return 0
}

func (x *UserLoginOrRegisterRequest) GetValidCode() string {
	if x != nil {
		return x.ValidCode
	}
	return ""
}

func (x *UserLoginOrRegisterRequest) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

type UserLoginOrRegisterReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	RetCode uint32 `protobuf:"varint,1,opt,name=retCode,proto3" json:"retCode,omitempty"`
	RetMsg  string `protobuf:"bytes,2,opt,name=retMsg,proto3" json:"retMsg,omitempty"`
	Token   string `protobuf:"bytes,3,opt,name=token,proto3" json:"token,omitempty"`
}

func (x *UserLoginOrRegisterReply) Reset() {
	*x = UserLoginOrRegisterReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v1_bingfood_service_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserLoginOrRegisterReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserLoginOrRegisterReply) ProtoMessage() {}

func (x *UserLoginOrRegisterReply) ProtoReflect() protoreflect.Message {
	mi := &file_v1_bingfood_service_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserLoginOrRegisterReply.ProtoReflect.Descriptor instead.
func (*UserLoginOrRegisterReply) Descriptor() ([]byte, []int) {
	return file_v1_bingfood_service_proto_rawDescGZIP(), []int{1}
}

func (x *UserLoginOrRegisterReply) GetRetCode() uint32 {
	if x != nil {
		return x.RetCode
	}
	return 0
}

func (x *UserLoginOrRegisterReply) GetRetMsg() string {
	if x != nil {
		return x.RetMsg
	}
	return ""
}

func (x *UserLoginOrRegisterReply) GetToken() string {
	if x != nil {
		return x.Token
	}
	return ""
}

var File_v1_bingfood_service_proto protoreflect.FileDescriptor

var file_v1_bingfood_service_proto_rawDesc = []byte{
	0x0a, 0x19, 0x76, 0x31, 0x2f, 0x62, 0x69, 0x6e, 0x67, 0x66, 0x6f, 0x6f, 0x64, 0x5f, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x13, 0x62, 0x69, 0x6e,
	0x67, 0x66, 0x6f, 0x6f, 0x64, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x76, 0x31,
	0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e,
	0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f,
	0x76, 0x31, 0x2f, 0x62, 0x69, 0x6e, 0x67, 0x66, 0x6f, 0x6f, 0x64, 0x5f, 0x6f, 0x72, 0x64, 0x65,
	0x72, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x1e, 0x76, 0x31, 0x2f, 0x62, 0x69, 0x6e, 0x67, 0x66, 0x6f, 0x6f, 0x64, 0x5f, 0x63, 0x61, 0x72,
	0x74, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22,
	0x94, 0x01, 0x0a, 0x1a, 0x55, 0x73, 0x65, 0x72, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x4f, 0x72, 0x52,
	0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1e,
	0x0a, 0x0a, 0x75, 0x73, 0x65, 0x72, 0x4d, 0x6f, 0x62, 0x69, 0x6c, 0x65, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0a, 0x75, 0x73, 0x65, 0x72, 0x4d, 0x6f, 0x62, 0x69, 0x6c, 0x65, 0x12, 0x1c,
	0x0a, 0x09, 0x6c, 0x6f, 0x67, 0x69, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x0d, 0x52, 0x09, 0x6c, 0x6f, 0x67, 0x69, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x12, 0x1c, 0x0a, 0x09,
	0x76, 0x61, 0x6c, 0x69, 0x64, 0x43, 0x6f, 0x64, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x09, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x61,
	0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x70, 0x61,
	0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x22, 0x62, 0x0a, 0x18, 0x55, 0x73, 0x65, 0x72, 0x4c, 0x6f,
	0x67, 0x69, 0x6e, 0x4f, 0x72, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x52, 0x65, 0x70,
	0x6c, 0x79, 0x12, 0x18, 0x0a, 0x07, 0x72, 0x65, 0x74, 0x43, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0d, 0x52, 0x07, 0x72, 0x65, 0x74, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x16, 0x0a, 0x06,
	0x72, 0x65, 0x74, 0x4d, 0x73, 0x67, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x72, 0x65,
	0x74, 0x4d, 0x73, 0x67, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x32, 0x97, 0x05, 0x0a, 0x0f, 0x42,
	0x69, 0x6e, 0x67, 0x66, 0x6f, 0x6f, 0x64, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x77,
	0x0a, 0x0b, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x53, 0x65, 0x74, 0x74, 0x6c, 0x65, 0x12, 0x27, 0x2e,
	0x62, 0x69, 0x6e, 0x67, 0x66, 0x6f, 0x6f, 0x64, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x2e, 0x76, 0x31, 0x2e, 0x53, 0x65, 0x74, 0x74, 0x6c, 0x65, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x25, 0x2e, 0x62, 0x69, 0x6e, 0x67, 0x66, 0x6f, 0x6f,
	0x64, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x65, 0x74,
	0x74, 0x6c, 0x65, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x18, 0x82,
	0xd3, 0xe4, 0x93, 0x02, 0x12, 0x22, 0x0d, 0x2f, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x2f, 0x73, 0x65,
	0x74, 0x74, 0x6c, 0x65, 0x3a, 0x01, 0x2a, 0x12, 0x77, 0x0a, 0x0b, 0x4f, 0x72, 0x64, 0x65, 0x72,
	0x53, 0x75, 0x62, 0x6d, 0x69, 0x74, 0x12, 0x27, 0x2e, 0x62, 0x69, 0x6e, 0x67, 0x66, 0x6f, 0x6f,
	0x64, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x75, 0x62,
	0x6d, 0x69, 0x74, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x25, 0x2e, 0x62, 0x69, 0x6e, 0x67, 0x66, 0x6f, 0x6f, 0x64, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x75, 0x62, 0x6d, 0x69, 0x74, 0x4f, 0x72, 0x64, 0x65,
	0x72, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x18, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x12, 0x22, 0x0d,
	0x2f, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x2f, 0x73, 0x75, 0x62, 0x6d, 0x69, 0x74, 0x3a, 0x01, 0x2a,
	0x12, 0x77, 0x0a, 0x0b, 0x41, 0x64, 0x64, 0x43, 0x61, 0x72, 0x74, 0x49, 0x74, 0x65, 0x6d, 0x12,
	0x27, 0x2e, 0x62, 0x69, 0x6e, 0x67, 0x66, 0x6f, 0x6f, 0x64, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x41, 0x64, 0x64, 0x43, 0x61, 0x72, 0x74, 0x49, 0x74, 0x65,
	0x6d, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x25, 0x2e, 0x62, 0x69, 0x6e, 0x67, 0x66,
	0x6f, 0x6f, 0x64, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x41,
	0x64, 0x64, 0x43, 0x61, 0x72, 0x74, 0x49, 0x74, 0x65, 0x6d, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22,
	0x18, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x12, 0x22, 0x0d, 0x2f, 0x63, 0x61, 0x72, 0x74, 0x2f, 0x61,
	0x64, 0x64, 0x49, 0x74, 0x65, 0x6d, 0x3a, 0x01, 0x2a, 0x12, 0x7f, 0x0a, 0x0d, 0x47, 0x65, 0x74,
	0x43, 0x61, 0x72, 0x74, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x12, 0x29, 0x2e, 0x62, 0x69, 0x6e,
	0x67, 0x66, 0x6f, 0x6f, 0x64, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x76, 0x31,
	0x2e, 0x47, 0x65, 0x74, 0x43, 0x61, 0x72, 0x74, 0x42, 0x79, 0x43, 0x6f, 0x6e, 0x64, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x27, 0x2e, 0x62, 0x69, 0x6e, 0x67, 0x66, 0x6f, 0x6f, 0x64,
	0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x43,
	0x61, 0x72, 0x74, 0x42, 0x79, 0x43, 0x6f, 0x6e, 0x64, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x1a,
	0x82, 0xd3, 0xe4, 0x93, 0x02, 0x14, 0x22, 0x0f, 0x2f, 0x63, 0x61, 0x72, 0x74, 0x2f, 0x67, 0x65,
	0x74, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x3a, 0x01, 0x2a, 0x12, 0x97, 0x01, 0x0a, 0x13, 0x55,
	0x73, 0x65, 0x72, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x4f, 0x72, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74,
	0x65, 0x72, 0x12, 0x2f, 0x2e, 0x62, 0x69, 0x6e, 0x67, 0x66, 0x6f, 0x6f, 0x64, 0x2e, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x4c, 0x6f, 0x67,
	0x69, 0x6e, 0x4f, 0x72, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x2d, 0x2e, 0x62, 0x69, 0x6e, 0x67, 0x66, 0x6f, 0x6f, 0x64, 0x2e, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x4c, 0x6f,
	0x67, 0x69, 0x6e, 0x4f, 0x72, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x52, 0x65, 0x70,
	0x6c, 0x79, 0x22, 0x20, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x1a, 0x22, 0x15, 0x2f, 0x75, 0x73, 0x65,
	0x72, 0x2f, 0x6c, 0x6f, 0x67, 0x69, 0x6e, 0x4f, 0x72, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65,
	0x72, 0x3a, 0x01, 0x2a, 0x42, 0x1c, 0x5a, 0x1a, 0x61, 0x70, 0x69, 0x2f, 0x62, 0x69, 0x6e, 0x67,
	0x66, 0x6f, 0x6f, 0x64, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2f, 0x76, 0x31, 0x3b,
	0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_v1_bingfood_service_proto_rawDescOnce sync.Once
	file_v1_bingfood_service_proto_rawDescData = file_v1_bingfood_service_proto_rawDesc
)

func file_v1_bingfood_service_proto_rawDescGZIP() []byte {
	file_v1_bingfood_service_proto_rawDescOnce.Do(func() {
		file_v1_bingfood_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_v1_bingfood_service_proto_rawDescData)
	})
	return file_v1_bingfood_service_proto_rawDescData
}

var file_v1_bingfood_service_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_v1_bingfood_service_proto_goTypes = []interface{}{
	(*UserLoginOrRegisterRequest)(nil), // 0: bingfood.service.v1.UserLoginOrRegisterRequest
	(*UserLoginOrRegisterReply)(nil),   // 1: bingfood.service.v1.UserLoginOrRegisterReply
	(*SettleOrderRequest)(nil),         // 2: bingfood.service.v1.SettleOrderRequest
	(*SubmitOrderRequest)(nil),         // 3: bingfood.service.v1.SubmitOrderRequest
	(*AddCartItemRequest)(nil),         // 4: bingfood.service.v1.AddCartItemRequest
	(*GetCartByCondRequest)(nil),       // 5: bingfood.service.v1.GetCartByCondRequest
	(*SettleOrderReply)(nil),           // 6: bingfood.service.v1.SettleOrderReply
	(*SubmitOrderReply)(nil),           // 7: bingfood.service.v1.SubmitOrderReply
	(*AddCartItemReply)(nil),           // 8: bingfood.service.v1.AddCartItemReply
	(*GetCartByCondReply)(nil),         // 9: bingfood.service.v1.GetCartByCondReply
}
var file_v1_bingfood_service_proto_depIdxs = []int32{
	2, // 0: bingfood.service.v1.BingfoodService.OrderSettle:input_type -> bingfood.service.v1.SettleOrderRequest
	3, // 1: bingfood.service.v1.BingfoodService.OrderSubmit:input_type -> bingfood.service.v1.SubmitOrderRequest
	4, // 2: bingfood.service.v1.BingfoodService.AddCartItem:input_type -> bingfood.service.v1.AddCartItemRequest
	5, // 3: bingfood.service.v1.BingfoodService.GetCartDetail:input_type -> bingfood.service.v1.GetCartByCondRequest
	0, // 4: bingfood.service.v1.BingfoodService.UserLoginOrRegister:input_type -> bingfood.service.v1.UserLoginOrRegisterRequest
	6, // 5: bingfood.service.v1.BingfoodService.OrderSettle:output_type -> bingfood.service.v1.SettleOrderReply
	7, // 6: bingfood.service.v1.BingfoodService.OrderSubmit:output_type -> bingfood.service.v1.SubmitOrderReply
	8, // 7: bingfood.service.v1.BingfoodService.AddCartItem:output_type -> bingfood.service.v1.AddCartItemReply
	9, // 8: bingfood.service.v1.BingfoodService.GetCartDetail:output_type -> bingfood.service.v1.GetCartByCondReply
	1, // 9: bingfood.service.v1.BingfoodService.UserLoginOrRegister:output_type -> bingfood.service.v1.UserLoginOrRegisterReply
	5, // [5:10] is the sub-list for method output_type
	0, // [0:5] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_v1_bingfood_service_proto_init() }
func file_v1_bingfood_service_proto_init() {
	if File_v1_bingfood_service_proto != nil {
		return
	}
	file_v1_bingfood_order_service_proto_init()
	file_v1_bingfood_cart_service_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_v1_bingfood_service_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UserLoginOrRegisterRequest); i {
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
		file_v1_bingfood_service_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UserLoginOrRegisterReply); i {
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
			RawDescriptor: file_v1_bingfood_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_v1_bingfood_service_proto_goTypes,
		DependencyIndexes: file_v1_bingfood_service_proto_depIdxs,
		MessageInfos:      file_v1_bingfood_service_proto_msgTypes,
	}.Build()
	File_v1_bingfood_service_proto = out.File
	file_v1_bingfood_service_proto_rawDesc = nil
	file_v1_bingfood_service_proto_goTypes = nil
	file_v1_bingfood_service_proto_depIdxs = nil
}
