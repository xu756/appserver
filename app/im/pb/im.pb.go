// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.12
// source: im.proto

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

type ImMeta struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	DetailType string `protobuf:"bytes,1,opt,name=DetailType,proto3" json:"DetailType,omitempty"` //元事件 连接 断开 状态更新 解密错误
	ImId       string `protobuf:"bytes,2,opt,name=ImId,proto3" json:"ImId,omitempty"`             //连接ID
	LoginId    uint64 `protobuf:"varint,3,opt,name=LoginId,proto3" json:"LoginId,omitempty"`      //登录ID
	Ip         string `protobuf:"bytes,4,opt,name=Ip,proto3" json:"Ip,omitempty"`                 //IP
	Issue      string `protobuf:"bytes,5,opt,name=Issue,proto3" json:"Issue,omitempty"`           //设备 谁颁布的jwt
	Version    string `protobuf:"bytes,6,opt,name=Version,proto3" json:"Version,omitempty"`       //版本
	ImServer   string `protobuf:"bytes,7,opt,name=ImServer,proto3" json:"ImServer,omitempty"`     //连接的服务器
	Data       []byte `protobuf:"bytes,8,opt,name=Data,proto3" json:"Data,omitempty"`             //数据
}

func (x *ImMeta) Reset() {
	*x = ImMeta{}
	if protoimpl.UnsafeEnabled {
		mi := &file_im_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ImMeta) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ImMeta) ProtoMessage() {}

func (x *ImMeta) ProtoReflect() protoreflect.Message {
	mi := &file_im_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ImMeta.ProtoReflect.Descriptor instead.
func (*ImMeta) Descriptor() ([]byte, []int) {
	return file_im_proto_rawDescGZIP(), []int{0}
}

func (x *ImMeta) GetDetailType() string {
	if x != nil {
		return x.DetailType
	}
	return ""
}

func (x *ImMeta) GetImId() string {
	if x != nil {
		return x.ImId
	}
	return ""
}

func (x *ImMeta) GetLoginId() uint64 {
	if x != nil {
		return x.LoginId
	}
	return 0
}

func (x *ImMeta) GetIp() string {
	if x != nil {
		return x.Ip
	}
	return ""
}

func (x *ImMeta) GetIssue() string {
	if x != nil {
		return x.Issue
	}
	return ""
}

func (x *ImMeta) GetVersion() string {
	if x != nil {
		return x.Version
	}
	return ""
}

func (x *ImMeta) GetImServer() string {
	if x != nil {
		return x.ImServer
	}
	return ""
}

func (x *ImMeta) GetData() []byte {
	if x != nil {
		return x.Data
	}
	return nil
}

type ImData struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ImId   string `protobuf:"bytes,1,opt,name=ImId,proto3" json:"ImId,omitempty"`
	ImData []byte `protobuf:"bytes,2,opt,name=ImData,proto3" json:"ImData,omitempty"`
}

func (x *ImData) Reset() {
	*x = ImData{}
	if protoimpl.UnsafeEnabled {
		mi := &file_im_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ImData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ImData) ProtoMessage() {}

func (x *ImData) ProtoReflect() protoreflect.Message {
	mi := &file_im_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ImData.ProtoReflect.Descriptor instead.
func (*ImData) Descriptor() ([]byte, []int) {
	return file_im_proto_rawDescGZIP(), []int{1}
}

func (x *ImData) GetImId() string {
	if x != nil {
		return x.ImId
	}
	return ""
}

func (x *ImData) GetImData() []byte {
	if x != nil {
		return x.ImData
	}
	return nil
}

type ImResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Ok bool `protobuf:"varint,1,opt,name=Ok,proto3" json:"Ok,omitempty"`
}

func (x *ImResp) Reset() {
	*x = ImResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_im_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ImResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ImResp) ProtoMessage() {}

func (x *ImResp) ProtoReflect() protoreflect.Message {
	mi := &file_im_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ImResp.ProtoReflect.Descriptor instead.
func (*ImResp) Descriptor() ([]byte, []int) {
	return file_im_proto_rawDescGZIP(), []int{2}
}

func (x *ImResp) GetOk() bool {
	if x != nil {
		return x.Ok
	}
	return false
}

var File_im_proto protoreflect.FileDescriptor

var file_im_proto_rawDesc = []byte{
	0x0a, 0x08, 0x69, 0x6d, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x02, 0x70, 0x62, 0x22, 0xc6,
	0x01, 0x0a, 0x06, 0x49, 0x6d, 0x4d, 0x65, 0x74, 0x61, 0x12, 0x1e, 0x0a, 0x0a, 0x44, 0x65, 0x74,
	0x61, 0x69, 0x6c, 0x54, 0x79, 0x70, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x44,
	0x65, 0x74, 0x61, 0x69, 0x6c, 0x54, 0x79, 0x70, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x49, 0x6d, 0x49,
	0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x49, 0x6d, 0x49, 0x64, 0x12, 0x18, 0x0a,
	0x07, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x49, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x04, 0x52, 0x07,
	0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x49, 0x64, 0x12, 0x0e, 0x0a, 0x02, 0x49, 0x70, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x02, 0x49, 0x70, 0x12, 0x14, 0x0a, 0x05, 0x49, 0x73, 0x73, 0x75, 0x65,
	0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x49, 0x73, 0x73, 0x75, 0x65, 0x12, 0x18, 0x0a,
	0x07, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07,
	0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x1a, 0x0a, 0x08, 0x49, 0x6d, 0x53, 0x65, 0x72,
	0x76, 0x65, 0x72, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x49, 0x6d, 0x53, 0x65, 0x72,
	0x76, 0x65, 0x72, 0x12, 0x12, 0x0a, 0x04, 0x44, 0x61, 0x74, 0x61, 0x18, 0x08, 0x20, 0x01, 0x28,
	0x0c, 0x52, 0x04, 0x44, 0x61, 0x74, 0x61, 0x22, 0x34, 0x0a, 0x06, 0x49, 0x6d, 0x44, 0x61, 0x74,
	0x61, 0x12, 0x12, 0x0a, 0x04, 0x49, 0x6d, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x04, 0x49, 0x6d, 0x49, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x49, 0x6d, 0x44, 0x61, 0x74, 0x61, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x06, 0x49, 0x6d, 0x44, 0x61, 0x74, 0x61, 0x22, 0x18, 0x0a,
	0x06, 0x49, 0x6d, 0x52, 0x65, 0x73, 0x70, 0x12, 0x0e, 0x0a, 0x02, 0x4f, 0x6b, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x08, 0x52, 0x02, 0x4f, 0x6b, 0x32, 0x4e, 0x0a, 0x07, 0x49, 0x6d, 0x4c, 0x6f, 0x67,
	0x69, 0x6e, 0x12, 0x1e, 0x0a, 0x04, 0x4d, 0x65, 0x74, 0x61, 0x12, 0x0a, 0x2e, 0x70, 0x62, 0x2e,
	0x49, 0x6d, 0x4d, 0x65, 0x74, 0x61, 0x1a, 0x0a, 0x2e, 0x70, 0x62, 0x2e, 0x49, 0x6d, 0x52, 0x65,
	0x73, 0x70, 0x12, 0x23, 0x0a, 0x09, 0x47, 0x65, 0x74, 0x49, 0x6d, 0x44, 0x61, 0x74, 0x61, 0x12,
	0x0a, 0x2e, 0x70, 0x62, 0x2e, 0x49, 0x6d, 0x44, 0x61, 0x74, 0x61, 0x1a, 0x0a, 0x2e, 0x70, 0x62,
	0x2e, 0x49, 0x6d, 0x52, 0x65, 0x73, 0x70, 0x42, 0x06, 0x5a, 0x04, 0x2e, 0x2f, 0x70, 0x62, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_im_proto_rawDescOnce sync.Once
	file_im_proto_rawDescData = file_im_proto_rawDesc
)

func file_im_proto_rawDescGZIP() []byte {
	file_im_proto_rawDescOnce.Do(func() {
		file_im_proto_rawDescData = protoimpl.X.CompressGZIP(file_im_proto_rawDescData)
	})
	return file_im_proto_rawDescData
}

var file_im_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_im_proto_goTypes = []interface{}{
	(*ImMeta)(nil), // 0: pb.ImMeta
	(*ImData)(nil), // 1: pb.ImData
	(*ImResp)(nil), // 2: pb.ImResp
}
var file_im_proto_depIdxs = []int32{
	0, // 0: pb.ImLogin.Meta:input_type -> pb.ImMeta
	1, // 1: pb.ImLogin.GetImData:input_type -> pb.ImData
	2, // 2: pb.ImLogin.Meta:output_type -> pb.ImResp
	2, // 3: pb.ImLogin.GetImData:output_type -> pb.ImResp
	2, // [2:4] is the sub-list for method output_type
	0, // [0:2] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_im_proto_init() }
func file_im_proto_init() {
	if File_im_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_im_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ImMeta); i {
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
		file_im_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ImData); i {
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
		file_im_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ImResp); i {
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
			RawDescriptor: file_im_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_im_proto_goTypes,
		DependencyIndexes: file_im_proto_depIdxs,
		MessageInfos:      file_im_proto_msgTypes,
	}.Build()
	File_im_proto = out.File
	file_im_proto_rawDesc = nil
	file_im_proto_goTypes = nil
	file_im_proto_depIdxs = nil
}