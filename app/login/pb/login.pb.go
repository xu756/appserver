// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.12
// source: login.proto

package pb

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

type LoginReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Mobile  string `protobuf:"bytes,1,opt,name=Mobile,proto3" json:"Mobile,omitempty"`
	SmsCode string `protobuf:"bytes,2,opt,name=SmsCode,proto3" json:"SmsCode,omitempty"`
}

func (x *LoginReq) Reset() {
	*x = LoginReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_login_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LoginReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LoginReq) ProtoMessage() {}

func (x *LoginReq) ProtoReflect() protoreflect.Message {
	mi := &file_login_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LoginReq.ProtoReflect.Descriptor instead.
func (*LoginReq) Descriptor() ([]byte, []int) {
	return file_login_proto_rawDescGZIP(), []int{0}
}

func (x *LoginReq) GetMobile() string {
	if x != nil {
		return x.Mobile
	}
	return ""
}

func (x *LoginReq) GetSmsCode() string {
	if x != nil {
		return x.SmsCode
	}
	return ""
}

type LoginResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AccessToken  string `protobuf:"bytes,1,opt,name=access_token,json=accessToken,proto3" json:"access_token,omitempty"`
	AccessExpire int64  `protobuf:"varint,2,opt,name=access_expire,json=accessExpire,proto3" json:"access_expire,omitempty"`
	RefreshAfter int64  `protobuf:"varint,3,opt,name=refresh_after,json=refreshAfter,proto3" json:"refresh_after,omitempty"`
}

func (x *LoginResp) Reset() {
	*x = LoginResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_login_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LoginResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LoginResp) ProtoMessage() {}

func (x *LoginResp) ProtoReflect() protoreflect.Message {
	mi := &file_login_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LoginResp.ProtoReflect.Descriptor instead.
func (*LoginResp) Descriptor() ([]byte, []int) {
	return file_login_proto_rawDescGZIP(), []int{1}
}

func (x *LoginResp) GetAccessToken() string {
	if x != nil {
		return x.AccessToken
	}
	return ""
}

func (x *LoginResp) GetAccessExpire() int64 {
	if x != nil {
		return x.AccessExpire
	}
	return 0
}

func (x *LoginResp) GetRefreshAfter() int64 {
	if x != nil {
		return x.RefreshAfter
	}
	return 0
}

type MiniAuthReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Code string `protobuf:"bytes,1,opt,name=Code,proto3" json:"Code,omitempty"`
}

func (x *MiniAuthReq) Reset() {
	*x = MiniAuthReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_login_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MiniAuthReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MiniAuthReq) ProtoMessage() {}

func (x *MiniAuthReq) ProtoReflect() protoreflect.Message {
	mi := &file_login_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MiniAuthReq.ProtoReflect.Descriptor instead.
func (*MiniAuthReq) Descriptor() ([]byte, []int) {
	return file_login_proto_rawDescGZIP(), []int{2}
}

func (x *MiniAuthReq) GetCode() string {
	if x != nil {
		return x.Code
	}
	return ""
}

type Empty struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *Empty) Reset() {
	*x = Empty{}
	if protoimpl.UnsafeEnabled {
		mi := &file_login_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Empty) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Empty) ProtoMessage() {}

func (x *Empty) ProtoReflect() protoreflect.Message {
	mi := &file_login_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Empty.ProtoReflect.Descriptor instead.
func (*Empty) Descriptor() ([]byte, []int) {
	return file_login_proto_rawDescGZIP(), []int{3}
}

// 图形验证码
type CaptchaResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Code  int64  `protobuf:"varint,1,opt,name=code,proto3" json:"code,omitempty"`
	Key   string `protobuf:"bytes,2,opt,name=key,proto3" json:"key,omitempty"`
	Image string `protobuf:"bytes,3,opt,name=image,proto3" json:"image,omitempty"`
	Thumb string `protobuf:"bytes,4,opt,name=thumb,proto3" json:"thumb,omitempty"`
}

func (x *CaptchaResp) Reset() {
	*x = CaptchaResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_login_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CaptchaResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CaptchaResp) ProtoMessage() {}

func (x *CaptchaResp) ProtoReflect() protoreflect.Message {
	mi := &file_login_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CaptchaResp.ProtoReflect.Descriptor instead.
func (*CaptchaResp) Descriptor() ([]byte, []int) {
	return file_login_proto_rawDescGZIP(), []int{4}
}

func (x *CaptchaResp) GetCode() int64 {
	if x != nil {
		return x.Code
	}
	return 0
}

func (x *CaptchaResp) GetKey() string {
	if x != nil {
		return x.Key
	}
	return ""
}

func (x *CaptchaResp) GetImage() string {
	if x != nil {
		return x.Image
	}
	return ""
}

func (x *CaptchaResp) GetThumb() string {
	if x != nil {
		return x.Thumb
	}
	return ""
}

// 验证码比对
type CaptchaCheckReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Dots []int64 `protobuf:"varint,1,rep,packed,name=dots,proto3" json:"dots,omitempty"` //  点击的坐标 x,y,x,y 4个点
	Key  string  `protobuf:"bytes,2,opt,name=key,proto3" json:"key,omitempty"`
}

func (x *CaptchaCheckReq) Reset() {
	*x = CaptchaCheckReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_login_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CaptchaCheckReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CaptchaCheckReq) ProtoMessage() {}

func (x *CaptchaCheckReq) ProtoReflect() protoreflect.Message {
	mi := &file_login_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CaptchaCheckReq.ProtoReflect.Descriptor instead.
func (*CaptchaCheckReq) Descriptor() ([]byte, []int) {
	return file_login_proto_rawDescGZIP(), []int{5}
}

func (x *CaptchaCheckReq) GetDots() []int64 {
	if x != nil {
		return x.Dots
	}
	return nil
}

func (x *CaptchaCheckReq) GetKey() string {
	if x != nil {
		return x.Key
	}
	return ""
}

type CaptchaCheckResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Result bool `protobuf:"varint,1,opt,name=result,proto3" json:"result,omitempty"`
}

func (x *CaptchaCheckResp) Reset() {
	*x = CaptchaCheckResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_login_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CaptchaCheckResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CaptchaCheckResp) ProtoMessage() {}

func (x *CaptchaCheckResp) ProtoReflect() protoreflect.Message {
	mi := &file_login_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CaptchaCheckResp.ProtoReflect.Descriptor instead.
func (*CaptchaCheckResp) Descriptor() ([]byte, []int) {
	return file_login_proto_rawDescGZIP(), []int{6}
}

func (x *CaptchaCheckResp) GetResult() bool {
	if x != nil {
		return x.Result
	}
	return false
}

var File_login_proto protoreflect.FileDescriptor

var file_login_proto_rawDesc = []byte{
	0x0a, 0x0b, 0x6c, 0x6f, 0x67, 0x69, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x02, 0x70,
	0x62, 0x22, 0x3c, 0x0a, 0x08, 0x6c, 0x6f, 0x67, 0x69, 0x6e, 0x52, 0x65, 0x71, 0x12, 0x16, 0x0a,
	0x06, 0x4d, 0x6f, 0x62, 0x69, 0x6c, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x4d,
	0x6f, 0x62, 0x69, 0x6c, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x53, 0x6d, 0x73, 0x43, 0x6f, 0x64, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x53, 0x6d, 0x73, 0x43, 0x6f, 0x64, 0x65, 0x22,
	0x78, 0x0a, 0x09, 0x6c, 0x6f, 0x67, 0x69, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x12, 0x21, 0x0a, 0x0c,
	0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x5f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0b, 0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x12,
	0x23, 0x0a, 0x0d, 0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x5f, 0x65, 0x78, 0x70, 0x69, 0x72, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0c, 0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x45, 0x78,
	0x70, 0x69, 0x72, 0x65, 0x12, 0x23, 0x0a, 0x0d, 0x72, 0x65, 0x66, 0x72, 0x65, 0x73, 0x68, 0x5f,
	0x61, 0x66, 0x74, 0x65, 0x72, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0c, 0x72, 0x65, 0x66,
	0x72, 0x65, 0x73, 0x68, 0x41, 0x66, 0x74, 0x65, 0x72, 0x22, 0x21, 0x0a, 0x0b, 0x4d, 0x69, 0x6e,
	0x69, 0x41, 0x75, 0x74, 0x68, 0x52, 0x65, 0x71, 0x12, 0x12, 0x0a, 0x04, 0x43, 0x6f, 0x64, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x43, 0x6f, 0x64, 0x65, 0x22, 0x07, 0x0a, 0x05,
	0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x5f, 0x0a, 0x0b, 0x43, 0x61, 0x70, 0x74, 0x63, 0x68, 0x61,
	0x52, 0x65, 0x73, 0x70, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x69, 0x6d,
	0x61, 0x67, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x69, 0x6d, 0x61, 0x67, 0x65,
	0x12, 0x14, 0x0a, 0x05, 0x74, 0x68, 0x75, 0x6d, 0x62, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x05, 0x74, 0x68, 0x75, 0x6d, 0x62, 0x22, 0x37, 0x0a, 0x0f, 0x43, 0x61, 0x70, 0x74, 0x63, 0x68,
	0x61, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x52, 0x65, 0x71, 0x12, 0x12, 0x0a, 0x04, 0x64, 0x6f, 0x74,
	0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x03, 0x52, 0x04, 0x64, 0x6f, 0x74, 0x73, 0x12, 0x10, 0x0a,
	0x03, 0x6b, 0x65, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x22,
	0x2a, 0x0a, 0x10, 0x43, 0x61, 0x70, 0x74, 0x63, 0x68, 0x61, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x52,
	0x65, 0x73, 0x70, 0x12, 0x16, 0x0a, 0x06, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x08, 0x52, 0x06, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x32, 0xd6, 0x01, 0x0a, 0x08,
	0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x52, 0x70, 0x63, 0x12, 0x30, 0x0a, 0x11, 0x4d, 0x69, 0x6e, 0x69,
	0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x42, 0x79, 0x4d, 0x6f, 0x62, 0x69, 0x6c, 0x65, 0x12, 0x0c, 0x2e,
	0x70, 0x62, 0x2e, 0x6c, 0x6f, 0x67, 0x69, 0x6e, 0x52, 0x65, 0x71, 0x1a, 0x0d, 0x2e, 0x70, 0x62,
	0x2e, 0x6c, 0x6f, 0x67, 0x69, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x12, 0x31, 0x0a, 0x0f, 0x4d, 0x69,
	0x6e, 0x69, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x42, 0x79, 0x41, 0x75, 0x74, 0x68, 0x12, 0x0f, 0x2e,
	0x70, 0x62, 0x2e, 0x4d, 0x69, 0x6e, 0x69, 0x41, 0x75, 0x74, 0x68, 0x52, 0x65, 0x71, 0x1a, 0x0d,
	0x2e, 0x70, 0x62, 0x2e, 0x6c, 0x6f, 0x67, 0x69, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x12, 0x28, 0x0a,
	0x0a, 0x47, 0x65, 0x74, 0x43, 0x61, 0x70, 0x74, 0x63, 0x68, 0x61, 0x12, 0x09, 0x2e, 0x70, 0x62,
	0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x0f, 0x2e, 0x70, 0x62, 0x2e, 0x43, 0x61, 0x70, 0x74,
	0x63, 0x68, 0x61, 0x52, 0x65, 0x73, 0x70, 0x12, 0x3b, 0x0a, 0x0e, 0x43, 0x61, 0x70, 0x74, 0x63,
	0x68, 0x61, 0x43, 0x6f, 0x6d, 0x70, 0x61, 0x72, 0x65, 0x12, 0x13, 0x2e, 0x70, 0x62, 0x2e, 0x43,
	0x61, 0x70, 0x74, 0x63, 0x68, 0x61, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x52, 0x65, 0x71, 0x1a, 0x14,
	0x2e, 0x70, 0x62, 0x2e, 0x43, 0x61, 0x70, 0x74, 0x63, 0x68, 0x61, 0x43, 0x68, 0x65, 0x63, 0x6b,
	0x52, 0x65, 0x73, 0x70, 0x42, 0x06, 0x5a, 0x04, 0x2e, 0x2f, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_login_proto_rawDescOnce sync.Once
	file_login_proto_rawDescData = file_login_proto_rawDesc
)

func file_login_proto_rawDescGZIP() []byte {
	file_login_proto_rawDescOnce.Do(func() {
		file_login_proto_rawDescData = protoimpl.X.CompressGZIP(file_login_proto_rawDescData)
	})
	return file_login_proto_rawDescData
}

var file_login_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_login_proto_goTypes = []interface{}{
	(*LoginReq)(nil),         // 0: pb.loginReq
	(*LoginResp)(nil),        // 1: pb.loginResp
	(*MiniAuthReq)(nil),      // 2: pb.MiniAuthReq
	(*Empty)(nil),            // 3: pb.Empty
	(*CaptchaResp)(nil),      // 4: pb.CaptchaResp
	(*CaptchaCheckReq)(nil),  // 5: pb.CaptchaCheckReq
	(*CaptchaCheckResp)(nil), // 6: pb.CaptchaCheckResp
}
var file_login_proto_depIdxs = []int32{
	0, // 0: pb.LoginRpc.MiniLoginByMobile:input_type -> pb.loginReq
	2, // 1: pb.LoginRpc.MiniLoginByAuth:input_type -> pb.MiniAuthReq
	3, // 2: pb.LoginRpc.GetCaptcha:input_type -> pb.Empty
	5, // 3: pb.LoginRpc.CaptchaCompare:input_type -> pb.CaptchaCheckReq
	1, // 4: pb.LoginRpc.MiniLoginByMobile:output_type -> pb.loginResp
	1, // 5: pb.LoginRpc.MiniLoginByAuth:output_type -> pb.loginResp
	4, // 6: pb.LoginRpc.GetCaptcha:output_type -> pb.CaptchaResp
	6, // 7: pb.LoginRpc.CaptchaCompare:output_type -> pb.CaptchaCheckResp
	4, // [4:8] is the sub-list for method output_type
	0, // [0:4] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_login_proto_init() }
func file_login_proto_init() {
	if File_login_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_login_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LoginReq); i {
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
		file_login_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LoginResp); i {
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
		file_login_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MiniAuthReq); i {
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
		file_login_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Empty); i {
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
		file_login_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CaptchaResp); i {
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
		file_login_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CaptchaCheckReq); i {
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
		file_login_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CaptchaCheckResp); i {
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
			RawDescriptor: file_login_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_login_proto_goTypes,
		DependencyIndexes: file_login_proto_depIdxs,
		MessageInfos:      file_login_proto_msgTypes,
	}.Build()
	File_login_proto = out.File
	file_login_proto_rawDesc = nil
	file_login_proto_goTypes = nil
	file_login_proto_depIdxs = nil
}
