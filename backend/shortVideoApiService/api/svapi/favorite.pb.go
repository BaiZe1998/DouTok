// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.12.4
// source: svapi/favorite.proto

package svapi

import (
	_ "buf.build/gen/go/bufbuild/protovalidate/protocolbuffers/go/buf/validate"
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

type FavoriteTarget int32

const (
	FavoriteTarget_VIDEO   FavoriteTarget = 0
	FavoriteTarget_COMMENT FavoriteTarget = 1
)

// Enum value maps for FavoriteTarget.
var (
	FavoriteTarget_name = map[int32]string{
		0: "VIDEO",
		1: "COMMENT",
	}
	FavoriteTarget_value = map[string]int32{
		"VIDEO":   0,
		"COMMENT": 1,
	}
)

func (x FavoriteTarget) Enum() *FavoriteTarget {
	p := new(FavoriteTarget)
	*p = x
	return p
}

func (x FavoriteTarget) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (FavoriteTarget) Descriptor() protoreflect.EnumDescriptor {
	return file_svapi_favorite_proto_enumTypes[0].Descriptor()
}

func (FavoriteTarget) Type() protoreflect.EnumType {
	return &file_svapi_favorite_proto_enumTypes[0]
}

func (x FavoriteTarget) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use FavoriteTarget.Descriptor instead.
func (FavoriteTarget) EnumDescriptor() ([]byte, []int) {
	return file_svapi_favorite_proto_rawDescGZIP(), []int{0}
}

type FavoriteType int32

const (
	FavoriteType_FAVORITE FavoriteType = 0 // 点赞
	FavoriteType_UNLIKE   FavoriteType = 1 // 点踩
)

// Enum value maps for FavoriteType.
var (
	FavoriteType_name = map[int32]string{
		0: "FAVORITE",
		1: "UNLIKE",
	}
	FavoriteType_value = map[string]int32{
		"FAVORITE": 0,
		"UNLIKE":   1,
	}
)

func (x FavoriteType) Enum() *FavoriteType {
	p := new(FavoriteType)
	*p = x
	return p
}

func (x FavoriteType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (FavoriteType) Descriptor() protoreflect.EnumDescriptor {
	return file_svapi_favorite_proto_enumTypes[1].Descriptor()
}

func (FavoriteType) Type() protoreflect.EnumType {
	return &file_svapi_favorite_proto_enumTypes[1]
}

func (x FavoriteType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use FavoriteType.Descriptor instead.
func (FavoriteType) EnumDescriptor() ([]byte, []int) {
	return file_svapi_favorite_proto_rawDescGZIP(), []int{1}
}

type AddFavoriteRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Target FavoriteTarget `protobuf:"varint,1,opt,name=target,proto3,enum=svapi.FavoriteTarget" json:"target,omitempty"`
	Type   FavoriteType   `protobuf:"varint,2,opt,name=type,proto3,enum=svapi.FavoriteType" json:"type,omitempty"`
	// @gotags: json:"id,omitempty,string"
	Id int64 `protobuf:"varint,3,opt,name=id,proto3" json:"id,omitempty,string"`
}

func (x *AddFavoriteRequest) Reset() {
	*x = AddFavoriteRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_svapi_favorite_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AddFavoriteRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddFavoriteRequest) ProtoMessage() {}

func (x *AddFavoriteRequest) ProtoReflect() protoreflect.Message {
	mi := &file_svapi_favorite_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddFavoriteRequest.ProtoReflect.Descriptor instead.
func (*AddFavoriteRequest) Descriptor() ([]byte, []int) {
	return file_svapi_favorite_proto_rawDescGZIP(), []int{0}
}

func (x *AddFavoriteRequest) GetTarget() FavoriteTarget {
	if x != nil {
		return x.Target
	}
	return FavoriteTarget_VIDEO
}

func (x *AddFavoriteRequest) GetType() FavoriteType {
	if x != nil {
		return x.Type
	}
	return FavoriteType_FAVORITE
}

func (x *AddFavoriteRequest) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

type AddFavoriteResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *AddFavoriteResponse) Reset() {
	*x = AddFavoriteResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_svapi_favorite_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AddFavoriteResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddFavoriteResponse) ProtoMessage() {}

func (x *AddFavoriteResponse) ProtoReflect() protoreflect.Message {
	mi := &file_svapi_favorite_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddFavoriteResponse.ProtoReflect.Descriptor instead.
func (*AddFavoriteResponse) Descriptor() ([]byte, []int) {
	return file_svapi_favorite_proto_rawDescGZIP(), []int{1}
}

type RemoveFavoriteRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Target FavoriteTarget `protobuf:"varint,1,opt,name=target,proto3,enum=svapi.FavoriteTarget" json:"target,omitempty"`
	Type   FavoriteType   `protobuf:"varint,2,opt,name=type,proto3,enum=svapi.FavoriteType" json:"type,omitempty"`
	// @gotags: json:"id,omitempty,string"
	Id int64 `protobuf:"varint,3,opt,name=id,proto3" json:"id,omitempty,string"`
}

func (x *RemoveFavoriteRequest) Reset() {
	*x = RemoveFavoriteRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_svapi_favorite_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RemoveFavoriteRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RemoveFavoriteRequest) ProtoMessage() {}

func (x *RemoveFavoriteRequest) ProtoReflect() protoreflect.Message {
	mi := &file_svapi_favorite_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RemoveFavoriteRequest.ProtoReflect.Descriptor instead.
func (*RemoveFavoriteRequest) Descriptor() ([]byte, []int) {
	return file_svapi_favorite_proto_rawDescGZIP(), []int{2}
}

func (x *RemoveFavoriteRequest) GetTarget() FavoriteTarget {
	if x != nil {
		return x.Target
	}
	return FavoriteTarget_VIDEO
}

func (x *RemoveFavoriteRequest) GetType() FavoriteType {
	if x != nil {
		return x.Type
	}
	return FavoriteType_FAVORITE
}

func (x *RemoveFavoriteRequest) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

type RemoveFavoriteResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *RemoveFavoriteResponse) Reset() {
	*x = RemoveFavoriteResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_svapi_favorite_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RemoveFavoriteResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RemoveFavoriteResponse) ProtoMessage() {}

func (x *RemoveFavoriteResponse) ProtoReflect() protoreflect.Message {
	mi := &file_svapi_favorite_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RemoveFavoriteResponse.ProtoReflect.Descriptor instead.
func (*RemoveFavoriteResponse) Descriptor() ([]byte, []int) {
	return file_svapi_favorite_proto_rawDescGZIP(), []int{3}
}

var File_svapi_favorite_proto protoreflect.FileDescriptor

var file_svapi_favorite_proto_rawDesc = []byte{
	0x0a, 0x14, 0x73, 0x76, 0x61, 0x70, 0x69, 0x2f, 0x66, 0x61, 0x76, 0x6f, 0x72, 0x69, 0x74, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x73, 0x76, 0x61, 0x70, 0x69, 0x1a, 0x1c, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1b, 0x62, 0x75, 0x66,
	0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61,
	0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x10, 0x73, 0x76, 0x61, 0x70, 0x69, 0x2f,
	0x62, 0x61, 0x73, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x7c, 0x0a, 0x12, 0x41, 0x64,
	0x64, 0x46, 0x61, 0x76, 0x6f, 0x72, 0x69, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x2d, 0x0a, 0x06, 0x74, 0x61, 0x72, 0x67, 0x65, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e,
	0x32, 0x15, 0x2e, 0x73, 0x76, 0x61, 0x70, 0x69, 0x2e, 0x46, 0x61, 0x76, 0x6f, 0x72, 0x69, 0x74,
	0x65, 0x54, 0x61, 0x72, 0x67, 0x65, 0x74, 0x52, 0x06, 0x74, 0x61, 0x72, 0x67, 0x65, 0x74, 0x12,
	0x27, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x13, 0x2e,
	0x73, 0x76, 0x61, 0x70, 0x69, 0x2e, 0x46, 0x61, 0x76, 0x6f, 0x72, 0x69, 0x74, 0x65, 0x54, 0x79,
	0x70, 0x65, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x22, 0x15, 0x0a, 0x13, 0x41, 0x64, 0x64, 0x46,
	0x61, 0x76, 0x6f, 0x72, 0x69, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22,
	0x7f, 0x0a, 0x15, 0x52, 0x65, 0x6d, 0x6f, 0x76, 0x65, 0x46, 0x61, 0x76, 0x6f, 0x72, 0x69, 0x74,
	0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x2d, 0x0a, 0x06, 0x74, 0x61, 0x72, 0x67,
	0x65, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x15, 0x2e, 0x73, 0x76, 0x61, 0x70, 0x69,
	0x2e, 0x46, 0x61, 0x76, 0x6f, 0x72, 0x69, 0x74, 0x65, 0x54, 0x61, 0x72, 0x67, 0x65, 0x74, 0x52,
	0x06, 0x74, 0x61, 0x72, 0x67, 0x65, 0x74, 0x12, 0x27, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x13, 0x2e, 0x73, 0x76, 0x61, 0x70, 0x69, 0x2e, 0x46, 0x61,
	0x76, 0x6f, 0x72, 0x69, 0x74, 0x65, 0x54, 0x79, 0x70, 0x65, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65,
	0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64,
	0x22, 0x18, 0x0a, 0x16, 0x52, 0x65, 0x6d, 0x6f, 0x76, 0x65, 0x46, 0x61, 0x76, 0x6f, 0x72, 0x69,
	0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x2a, 0x28, 0x0a, 0x0e, 0x46, 0x61,
	0x76, 0x6f, 0x72, 0x69, 0x74, 0x65, 0x54, 0x61, 0x72, 0x67, 0x65, 0x74, 0x12, 0x09, 0x0a, 0x05,
	0x56, 0x49, 0x44, 0x45, 0x4f, 0x10, 0x00, 0x12, 0x0b, 0x0a, 0x07, 0x43, 0x4f, 0x4d, 0x4d, 0x45,
	0x4e, 0x54, 0x10, 0x01, 0x2a, 0x28, 0x0a, 0x0c, 0x46, 0x61, 0x76, 0x6f, 0x72, 0x69, 0x74, 0x65,
	0x54, 0x79, 0x70, 0x65, 0x12, 0x0c, 0x0a, 0x08, 0x46, 0x41, 0x56, 0x4f, 0x52, 0x49, 0x54, 0x45,
	0x10, 0x00, 0x12, 0x0a, 0x0a, 0x06, 0x55, 0x4e, 0x4c, 0x49, 0x4b, 0x45, 0x10, 0x01, 0x32, 0xd2,
	0x01, 0x0a, 0x0f, 0x46, 0x61, 0x76, 0x6f, 0x72, 0x69, 0x74, 0x65, 0x53, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x12, 0x5a, 0x0a, 0x0b, 0x41, 0x64, 0x64, 0x46, 0x61, 0x76, 0x6f, 0x72, 0x69, 0x74,
	0x65, 0x12, 0x19, 0x2e, 0x73, 0x76, 0x61, 0x70, 0x69, 0x2e, 0x41, 0x64, 0x64, 0x46, 0x61, 0x76,
	0x6f, 0x72, 0x69, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1a, 0x2e, 0x73,
	0x76, 0x61, 0x70, 0x69, 0x2e, 0x41, 0x64, 0x64, 0x46, 0x61, 0x76, 0x6f, 0x72, 0x69, 0x74, 0x65,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x14, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x0e,
	0x22, 0x09, 0x2f, 0x66, 0x61, 0x76, 0x6f, 0x72, 0x69, 0x74, 0x65, 0x3a, 0x01, 0x2a, 0x12, 0x63,
	0x0a, 0x0e, 0x52, 0x65, 0x6d, 0x6f, 0x76, 0x65, 0x46, 0x61, 0x76, 0x6f, 0x72, 0x69, 0x74, 0x65,
	0x12, 0x1c, 0x2e, 0x73, 0x76, 0x61, 0x70, 0x69, 0x2e, 0x52, 0x65, 0x6d, 0x6f, 0x76, 0x65, 0x46,
	0x61, 0x76, 0x6f, 0x72, 0x69, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1d,
	0x2e, 0x73, 0x76, 0x61, 0x70, 0x69, 0x2e, 0x52, 0x65, 0x6d, 0x6f, 0x76, 0x65, 0x46, 0x61, 0x76,
	0x6f, 0x72, 0x69, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x14, 0x82,
	0xd3, 0xe4, 0x93, 0x02, 0x0e, 0x2a, 0x09, 0x2f, 0x66, 0x61, 0x76, 0x6f, 0x72, 0x69, 0x74, 0x65,
	0x3a, 0x01, 0x2a, 0x42, 0x29, 0x5a, 0x27, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f,
	0x6d, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x7a, 0x65, 0x6e, 0x69, 0x74, 0x68, 0x2f, 0x44, 0x6f,
	0x75, 0x54, 0x6f, 0x6b, 0x2f, 0x2e, 0x2e, 0x2e, 0x3b, 0x73, 0x76, 0x61, 0x70, 0x69, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_svapi_favorite_proto_rawDescOnce sync.Once
	file_svapi_favorite_proto_rawDescData = file_svapi_favorite_proto_rawDesc
)

func file_svapi_favorite_proto_rawDescGZIP() []byte {
	file_svapi_favorite_proto_rawDescOnce.Do(func() {
		file_svapi_favorite_proto_rawDescData = protoimpl.X.CompressGZIP(file_svapi_favorite_proto_rawDescData)
	})
	return file_svapi_favorite_proto_rawDescData
}

var file_svapi_favorite_proto_enumTypes = make([]protoimpl.EnumInfo, 2)
var file_svapi_favorite_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_svapi_favorite_proto_goTypes = []interface{}{
	(FavoriteTarget)(0),            // 0: svapi.FavoriteTarget
	(FavoriteType)(0),              // 1: svapi.FavoriteType
	(*AddFavoriteRequest)(nil),     // 2: svapi.AddFavoriteRequest
	(*AddFavoriteResponse)(nil),    // 3: svapi.AddFavoriteResponse
	(*RemoveFavoriteRequest)(nil),  // 4: svapi.RemoveFavoriteRequest
	(*RemoveFavoriteResponse)(nil), // 5: svapi.RemoveFavoriteResponse
}
var file_svapi_favorite_proto_depIdxs = []int32{
	0, // 0: svapi.AddFavoriteRequest.target:type_name -> svapi.FavoriteTarget
	1, // 1: svapi.AddFavoriteRequest.type:type_name -> svapi.FavoriteType
	0, // 2: svapi.RemoveFavoriteRequest.target:type_name -> svapi.FavoriteTarget
	1, // 3: svapi.RemoveFavoriteRequest.type:type_name -> svapi.FavoriteType
	2, // 4: svapi.FavoriteService.AddFavorite:input_type -> svapi.AddFavoriteRequest
	4, // 5: svapi.FavoriteService.RemoveFavorite:input_type -> svapi.RemoveFavoriteRequest
	3, // 6: svapi.FavoriteService.AddFavorite:output_type -> svapi.AddFavoriteResponse
	5, // 7: svapi.FavoriteService.RemoveFavorite:output_type -> svapi.RemoveFavoriteResponse
	6, // [6:8] is the sub-list for method output_type
	4, // [4:6] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_svapi_favorite_proto_init() }
func file_svapi_favorite_proto_init() {
	if File_svapi_favorite_proto != nil {
		return
	}
	file_svapi_base_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_svapi_favorite_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AddFavoriteRequest); i {
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
		file_svapi_favorite_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AddFavoriteResponse); i {
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
		file_svapi_favorite_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RemoveFavoriteRequest); i {
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
		file_svapi_favorite_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RemoveFavoriteResponse); i {
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
			RawDescriptor: file_svapi_favorite_proto_rawDesc,
			NumEnums:      2,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_svapi_favorite_proto_goTypes,
		DependencyIndexes: file_svapi_favorite_proto_depIdxs,
		EnumInfos:         file_svapi_favorite_proto_enumTypes,
		MessageInfos:      file_svapi_favorite_proto_msgTypes,
	}.Build()
	File_svapi_favorite_proto = out.File
	file_svapi_favorite_proto_rawDesc = nil
	file_svapi_favorite_proto_goTypes = nil
	file_svapi_favorite_proto_depIdxs = nil
}
