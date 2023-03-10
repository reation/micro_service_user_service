// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.19.4
// source: user_info.proto

package protoc

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

type UserDetail struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id       int64  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	UserName string `protobuf:"bytes,2,opt,name=user_name,json=userName,proto3" json:"user_name,omitempty"`
	Nickname string `protobuf:"bytes,3,opt,name=nickname,proto3" json:"nickname,omitempty"`
	Sex      int64  `protobuf:"varint,4,opt,name=sex,proto3" json:"sex,omitempty"`
	Birthday string `protobuf:"bytes,5,opt,name=birthday,proto3" json:"birthday,omitempty"`
	States   int64  `protobuf:"varint,6,opt,name=states,proto3" json:"states,omitempty"`
	IsMember int64  `protobuf:"varint,7,opt,name=is_member,json=isMember,proto3" json:"is_member,omitempty"`
}

func (x *UserDetail) Reset() {
	*x = UserDetail{}
	if protoimpl.UnsafeEnabled {
		mi := &file_user_info_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserDetail) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserDetail) ProtoMessage() {}

func (x *UserDetail) ProtoReflect() protoreflect.Message {
	mi := &file_user_info_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserDetail.ProtoReflect.Descriptor instead.
func (*UserDetail) Descriptor() ([]byte, []int) {
	return file_user_info_proto_rawDescGZIP(), []int{0}
}

func (x *UserDetail) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *UserDetail) GetUserName() string {
	if x != nil {
		return x.UserName
	}
	return ""
}

func (x *UserDetail) GetNickname() string {
	if x != nil {
		return x.Nickname
	}
	return ""
}

func (x *UserDetail) GetSex() int64 {
	if x != nil {
		return x.Sex
	}
	return 0
}

func (x *UserDetail) GetBirthday() string {
	if x != nil {
		return x.Birthday
	}
	return ""
}

func (x *UserDetail) GetStates() int64 {
	if x != nil {
		return x.States
	}
	return 0
}

func (x *UserDetail) GetIsMember() int64 {
	if x != nil {
		return x.IsMember
	}
	return 0
}

type UserInfoRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id int64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *UserInfoRequest) Reset() {
	*x = UserInfoRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_user_info_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserInfoRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserInfoRequest) ProtoMessage() {}

func (x *UserInfoRequest) ProtoReflect() protoreflect.Message {
	mi := &file_user_info_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserInfoRequest.ProtoReflect.Descriptor instead.
func (*UserInfoRequest) Descriptor() ([]byte, []int) {
	return file_user_info_proto_rawDescGZIP(), []int{1}
}

func (x *UserInfoRequest) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

type UserInfoResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	States   int64       `protobuf:"varint,1,opt,name=states,proto3" json:"states,omitempty"`
	UserInfo *UserDetail `protobuf:"bytes,2,opt,name=userInfo,proto3" json:"userInfo,omitempty"`
}

func (x *UserInfoResponse) Reset() {
	*x = UserInfoResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_user_info_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserInfoResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserInfoResponse) ProtoMessage() {}

func (x *UserInfoResponse) ProtoReflect() protoreflect.Message {
	mi := &file_user_info_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserInfoResponse.ProtoReflect.Descriptor instead.
func (*UserInfoResponse) Descriptor() ([]byte, []int) {
	return file_user_info_proto_rawDescGZIP(), []int{2}
}

func (x *UserInfoResponse) GetStates() int64 {
	if x != nil {
		return x.States
	}
	return 0
}

func (x *UserInfoResponse) GetUserInfo() *UserDetail {
	if x != nil {
		return x.UserInfo
	}
	return nil
}

type NormalUserInfoRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id int64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *NormalUserInfoRequest) Reset() {
	*x = NormalUserInfoRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_user_info_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NormalUserInfoRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NormalUserInfoRequest) ProtoMessage() {}

func (x *NormalUserInfoRequest) ProtoReflect() protoreflect.Message {
	mi := &file_user_info_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NormalUserInfoRequest.ProtoReflect.Descriptor instead.
func (*NormalUserInfoRequest) Descriptor() ([]byte, []int) {
	return file_user_info_proto_rawDescGZIP(), []int{3}
}

func (x *NormalUserInfoRequest) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

type NormalUserInfoResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	States   int64       `protobuf:"varint,1,opt,name=states,proto3" json:"states,omitempty"`
	UserInfo *UserDetail `protobuf:"bytes,2,opt,name=userInfo,proto3" json:"userInfo,omitempty"`
}

func (x *NormalUserInfoResponse) Reset() {
	*x = NormalUserInfoResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_user_info_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NormalUserInfoResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NormalUserInfoResponse) ProtoMessage() {}

func (x *NormalUserInfoResponse) ProtoReflect() protoreflect.Message {
	mi := &file_user_info_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NormalUserInfoResponse.ProtoReflect.Descriptor instead.
func (*NormalUserInfoResponse) Descriptor() ([]byte, []int) {
	return file_user_info_proto_rawDescGZIP(), []int{4}
}

func (x *NormalUserInfoResponse) GetStates() int64 {
	if x != nil {
		return x.States
	}
	return 0
}

func (x *NormalUserInfoResponse) GetUserInfo() *UserDetail {
	if x != nil {
		return x.UserInfo
	}
	return nil
}

var File_user_info_proto protoreflect.FileDescriptor

var file_user_info_proto_rawDesc = []byte{
	0x0a, 0x0f, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x6e, 0x66, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x09, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x6e, 0x66, 0x6f, 0x22, 0xb8, 0x01, 0x0a,
	0x0a, 0x55, 0x73, 0x65, 0x72, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x12, 0x0e, 0x0a, 0x02, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x12, 0x1b, 0x0a, 0x09, 0x75,
	0x73, 0x65, 0x72, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08,
	0x75, 0x73, 0x65, 0x72, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x6e, 0x69, 0x63, 0x6b,
	0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6e, 0x69, 0x63, 0x6b,
	0x6e, 0x61, 0x6d, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x73, 0x65, 0x78, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x03, 0x73, 0x65, 0x78, 0x12, 0x1a, 0x0a, 0x08, 0x62, 0x69, 0x72, 0x74, 0x68, 0x64,
	0x61, 0x79, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x62, 0x69, 0x72, 0x74, 0x68, 0x64,
	0x61, 0x79, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x65, 0x73, 0x18, 0x06, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x65, 0x73, 0x12, 0x1b, 0x0a, 0x09, 0x69, 0x73,
	0x5f, 0x6d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x18, 0x07, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x69,
	0x73, 0x4d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x22, 0x21, 0x0a, 0x0f, 0x55, 0x73, 0x65, 0x72, 0x49,
	0x6e, 0x66, 0x6f, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x22, 0x5d, 0x0a, 0x10, 0x55, 0x73,
	0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x16,
	0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x65, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06,
	0x73, 0x74, 0x61, 0x74, 0x65, 0x73, 0x12, 0x31, 0x0a, 0x08, 0x75, 0x73, 0x65, 0x72, 0x49, 0x6e,
	0x66, 0x6f, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x15, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x5f,
	0x69, 0x6e, 0x66, 0x6f, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x52,
	0x08, 0x75, 0x73, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x22, 0x27, 0x0a, 0x15, 0x4e, 0x6f, 0x72,
	0x6d, 0x61, 0x6c, 0x55, 0x73, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02,
	0x69, 0x64, 0x22, 0x63, 0x0a, 0x16, 0x4e, 0x6f, 0x72, 0x6d, 0x61, 0x6c, 0x55, 0x73, 0x65, 0x72,
	0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x16, 0x0a, 0x06,
	0x73, 0x74, 0x61, 0x74, 0x65, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x73, 0x74,
	0x61, 0x74, 0x65, 0x73, 0x12, 0x31, 0x0a, 0x08, 0x75, 0x73, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x15, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x6e,
	0x66, 0x6f, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x52, 0x08, 0x75,
	0x73, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x32, 0xb5, 0x01, 0x0a, 0x09, 0x75, 0x73, 0x65, 0x72,
	0x5f, 0x69, 0x6e, 0x66, 0x6f, 0x12, 0x4a, 0x0a, 0x0f, 0x67, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72,
	0x49, 0x6e, 0x66, 0x6f, 0x42, 0x79, 0x49, 0x44, 0x12, 0x1a, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x5f,
	0x69, 0x6e, 0x66, 0x6f, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x1b, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x6e, 0x66, 0x6f,
	0x2e, 0x55, 0x73, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x5c, 0x0a, 0x15, 0x67, 0x65, 0x74, 0x4e, 0x6f, 0x72, 0x6d, 0x61, 0x6c, 0x55, 0x73,
	0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x42, 0x79, 0x49, 0x44, 0x12, 0x20, 0x2e, 0x75, 0x73, 0x65,
	0x72, 0x5f, 0x69, 0x6e, 0x66, 0x6f, 0x2e, 0x4e, 0x6f, 0x72, 0x6d, 0x61, 0x6c, 0x55, 0x73, 0x65,
	0x72, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x21, 0x2e, 0x75,
	0x73, 0x65, 0x72, 0x5f, 0x69, 0x6e, 0x66, 0x6f, 0x2e, 0x4e, 0x6f, 0x72, 0x6d, 0x61, 0x6c, 0x55,
	0x73, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42,
	0x0a, 0x5a, 0x08, 0x2e, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_user_info_proto_rawDescOnce sync.Once
	file_user_info_proto_rawDescData = file_user_info_proto_rawDesc
)

func file_user_info_proto_rawDescGZIP() []byte {
	file_user_info_proto_rawDescOnce.Do(func() {
		file_user_info_proto_rawDescData = protoimpl.X.CompressGZIP(file_user_info_proto_rawDescData)
	})
	return file_user_info_proto_rawDescData
}

var file_user_info_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_user_info_proto_goTypes = []interface{}{
	(*UserDetail)(nil),             // 0: user_info.UserDetail
	(*UserInfoRequest)(nil),        // 1: user_info.UserInfoRequest
	(*UserInfoResponse)(nil),       // 2: user_info.UserInfoResponse
	(*NormalUserInfoRequest)(nil),  // 3: user_info.NormalUserInfoRequest
	(*NormalUserInfoResponse)(nil), // 4: user_info.NormalUserInfoResponse
}
var file_user_info_proto_depIdxs = []int32{
	0, // 0: user_info.UserInfoResponse.userInfo:type_name -> user_info.UserDetail
	0, // 1: user_info.NormalUserInfoResponse.userInfo:type_name -> user_info.UserDetail
	1, // 2: user_info.user_info.getUserInfoByID:input_type -> user_info.UserInfoRequest
	3, // 3: user_info.user_info.getNormalUserInfoByID:input_type -> user_info.NormalUserInfoRequest
	2, // 4: user_info.user_info.getUserInfoByID:output_type -> user_info.UserInfoResponse
	4, // 5: user_info.user_info.getNormalUserInfoByID:output_type -> user_info.NormalUserInfoResponse
	4, // [4:6] is the sub-list for method output_type
	2, // [2:4] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_user_info_proto_init() }
func file_user_info_proto_init() {
	if File_user_info_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_user_info_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UserDetail); i {
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
		file_user_info_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UserInfoRequest); i {
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
		file_user_info_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UserInfoResponse); i {
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
		file_user_info_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NormalUserInfoRequest); i {
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
		file_user_info_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NormalUserInfoResponse); i {
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
			RawDescriptor: file_user_info_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_user_info_proto_goTypes,
		DependencyIndexes: file_user_info_proto_depIdxs,
		MessageInfos:      file_user_info_proto_msgTypes,
	}.Build()
	File_user_info_proto = out.File
	file_user_info_proto_rawDesc = nil
	file_user_info_proto_goTypes = nil
	file_user_info_proto_depIdxs = nil
}
