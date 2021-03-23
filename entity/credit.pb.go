// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.15.2
// source: credit.proto

package entity

import (
	_ "github.com/cloudstateio/go-support/cloudstate"
	proto "github.com/golang/protobuf/proto"
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

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

type UserCredit struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Quantity int32 `protobuf:"varint,1,opt,name=quantity,proto3" json:"quantity,omitempty"`
}

func (x *UserCredit) Reset() {
	*x = UserCredit{}
	if protoimpl.UnsafeEnabled {
		mi := &file_credit_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserCredit) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserCredit) ProtoMessage() {}

func (x *UserCredit) ProtoReflect() protoreflect.Message {
	mi := &file_credit_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserCredit.ProtoReflect.Descriptor instead.
func (*UserCredit) Descriptor() ([]byte, []int) {
	return file_credit_proto_rawDescGZIP(), []int{0}
}

func (x *UserCredit) GetQuantity() int32 {
	if x != nil {
		return x.Quantity
	}
	return 0
}

type AddUserCredit struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId   string `protobuf:"bytes,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	Quantity int32  `protobuf:"varint,2,opt,name=quantity,proto3" json:"quantity,omitempty"`
}

func (x *AddUserCredit) Reset() {
	*x = AddUserCredit{}
	if protoimpl.UnsafeEnabled {
		mi := &file_credit_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AddUserCredit) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddUserCredit) ProtoMessage() {}

func (x *AddUserCredit) ProtoReflect() protoreflect.Message {
	mi := &file_credit_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddUserCredit.ProtoReflect.Descriptor instead.
func (*AddUserCredit) Descriptor() ([]byte, []int) {
	return file_credit_proto_rawDescGZIP(), []int{1}
}

func (x *AddUserCredit) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *AddUserCredit) GetQuantity() int32 {
	if x != nil {
		return x.Quantity
	}
	return 0
}

type GetUserCredit struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId string `protobuf:"bytes,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
}

func (x *GetUserCredit) Reset() {
	*x = GetUserCredit{}
	if protoimpl.UnsafeEnabled {
		mi := &file_credit_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetUserCredit) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetUserCredit) ProtoMessage() {}

func (x *GetUserCredit) ProtoReflect() protoreflect.Message {
	mi := &file_credit_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetUserCredit.ProtoReflect.Descriptor instead.
func (*GetUserCredit) Descriptor() ([]byte, []int) {
	return file_credit_proto_rawDescGZIP(), []int{2}
}

func (x *GetUserCredit) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

var File_credit_proto protoreflect.FileDescriptor

var file_credit_proto_rawDesc = []byte{
	0x0a, 0x0c, 0x63, 0x72, 0x65, 0x64, 0x69, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x06,
	0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x1a, 0x1b, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x73, 0x74, 0x61,
	0x74, 0x65, 0x2f, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x5f, 0x6b, 0x65, 0x79, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f,
	0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x22, 0x28, 0x0a, 0x0a, 0x55, 0x73, 0x65, 0x72, 0x43, 0x72, 0x65, 0x64, 0x69, 0x74, 0x12,
	0x1a, 0x0a, 0x08, 0x71, 0x75, 0x61, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x08, 0x71, 0x75, 0x61, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x22, 0x49, 0x0a, 0x0d, 0x41,
	0x64, 0x64, 0x55, 0x73, 0x65, 0x72, 0x43, 0x72, 0x65, 0x64, 0x69, 0x74, 0x12, 0x1c, 0x0a, 0x07,
	0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x03, 0xc0,
	0x43, 0x01, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x71, 0x75,
	0x61, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x71, 0x75,
	0x61, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x22, 0x2d, 0x0a, 0x0d, 0x47, 0x65, 0x74, 0x55, 0x73, 0x65,
	0x72, 0x43, 0x72, 0x65, 0x64, 0x69, 0x74, 0x12, 0x1c, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x03, 0xc0, 0x43, 0x01, 0x52, 0x06, 0x75,
	0x73, 0x65, 0x72, 0x49, 0x64, 0x32, 0xbc, 0x01, 0x0a, 0x0d, 0x43, 0x72, 0x65, 0x64, 0x69, 0x74,
	0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x58, 0x0a, 0x09, 0x41, 0x64, 0x64, 0x43, 0x72,
	0x65, 0x64, 0x69, 0x74, 0x12, 0x15, 0x2e, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x2e, 0x41, 0x64,
	0x64, 0x55, 0x73, 0x65, 0x72, 0x43, 0x72, 0x65, 0x64, 0x69, 0x74, 0x1a, 0x12, 0x2e, 0x65, 0x6e,
	0x74, 0x69, 0x74, 0x79, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x43, 0x72, 0x65, 0x64, 0x69, 0x74, 0x22,
	0x20, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x1a, 0x22, 0x15, 0x2f, 0x63, 0x72, 0x65, 0x64, 0x69, 0x74,
	0x2f, 0x7b, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x7d, 0x2f, 0x61, 0x64, 0x64, 0x3a, 0x01,
	0x2a, 0x12, 0x51, 0x0a, 0x09, 0x47, 0x65, 0x74, 0x43, 0x72, 0x65, 0x64, 0x69, 0x74, 0x12, 0x15,
	0x2e, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x2e, 0x47, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x43,
	0x72, 0x65, 0x64, 0x69, 0x74, 0x1a, 0x12, 0x2e, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x2e, 0x55,
	0x73, 0x65, 0x72, 0x43, 0x72, 0x65, 0x64, 0x69, 0x74, 0x22, 0x19, 0x82, 0xd3, 0xe4, 0x93, 0x02,
	0x13, 0x12, 0x11, 0x2f, 0x63, 0x72, 0x65, 0x64, 0x69, 0x74, 0x2f, 0x7b, 0x75, 0x73, 0x65, 0x72,
	0x5f, 0x69, 0x64, 0x7d, 0x42, 0x53, 0x5a, 0x51, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x77,
	0x64, 0x66, 0x2e, 0x73, 0x61, 0x70, 0x2e, 0x63, 0x6f, 0x72, 0x70, 0x2f, 0x49, 0x43, 0x4e, 0x2d,
	0x4e, 0x61, 0x6e, 0x6a, 0x69, 0x6e, 0x67, 0x2d, 0x50, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x73,
	0x2f, 0x43, 0x6f, 0x6c, 0x6c, 0x61, 0x62, 0x6f, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2d, 0x62,
	0x6f, 0x61, 0x72, 0x64, 0x2f, 0x63, 0x72, 0x65, 0x64, 0x69, 0x74, 0x2f, 0x65, 0x6e, 0x74, 0x69,
	0x74, 0x79, 0x3b, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_credit_proto_rawDescOnce sync.Once
	file_credit_proto_rawDescData = file_credit_proto_rawDesc
)

func file_credit_proto_rawDescGZIP() []byte {
	file_credit_proto_rawDescOnce.Do(func() {
		file_credit_proto_rawDescData = protoimpl.X.CompressGZIP(file_credit_proto_rawDescData)
	})
	return file_credit_proto_rawDescData
}

var file_credit_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_credit_proto_goTypes = []interface{}{
	(*UserCredit)(nil),    // 0: entity.UserCredit
	(*AddUserCredit)(nil), // 1: entity.AddUserCredit
	(*GetUserCredit)(nil), // 2: entity.GetUserCredit
}
var file_credit_proto_depIdxs = []int32{
	1, // 0: entity.CreditService.AddCredit:input_type -> entity.AddUserCredit
	2, // 1: entity.CreditService.GetCredit:input_type -> entity.GetUserCredit
	0, // 2: entity.CreditService.AddCredit:output_type -> entity.UserCredit
	0, // 3: entity.CreditService.GetCredit:output_type -> entity.UserCredit
	2, // [2:4] is the sub-list for method output_type
	0, // [0:2] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_credit_proto_init() }
func file_credit_proto_init() {
	if File_credit_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_credit_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UserCredit); i {
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
		file_credit_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AddUserCredit); i {
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
		file_credit_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetUserCredit); i {
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
			RawDescriptor: file_credit_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_credit_proto_goTypes,
		DependencyIndexes: file_credit_proto_depIdxs,
		MessageInfos:      file_credit_proto_msgTypes,
	}.Build()
	File_credit_proto = out.File
	file_credit_proto_rawDesc = nil
	file_credit_proto_goTypes = nil
	file_credit_proto_depIdxs = nil
}