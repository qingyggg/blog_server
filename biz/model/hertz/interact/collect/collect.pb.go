// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v5.28.0
// source: collect.proto

package collect

import (
	_ "github.com/qingyggg/blog_server/biz/model/hertz/api"
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

type CollectActionRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AHashId    string `protobuf:"bytes,1,opt,name=AHashId,proto3" form:"AHashId" json:"aHashId,required" query:"AHashId"` //文章的id
	Tag        string `protobuf:"bytes,2,opt,name=Tag,proto3" form:"Tag" json:"tag,required" query:"Tag"`
	ActionType int32  `protobuf:"varint,3,opt,name=ActionType,proto3" form:"ActionType" json:"action_type,required" query:"ActionType" vd:"$==1 || $==2"` // 1-collect, 2-Uncollect
}

func (x *CollectActionRequest) Reset() {
	*x = CollectActionRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_collect_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CollectActionRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CollectActionRequest) ProtoMessage() {}

func (x *CollectActionRequest) ProtoReflect() protoreflect.Message {
	mi := &file_collect_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CollectActionRequest.ProtoReflect.Descriptor instead.
func (*CollectActionRequest) Descriptor() ([]byte, []int) {
	return file_collect_proto_rawDescGZIP(), []int{0}
}

func (x *CollectActionRequest) GetAHashId() string {
	if x != nil {
		return x.AHashId
	}
	return ""
}

func (x *CollectActionRequest) GetTag() string {
	if x != nil {
		return x.Tag
	}
	return ""
}

func (x *CollectActionRequest) GetActionType() int32 {
	if x != nil {
		return x.ActionType
	}
	return 0
}

type CollectActionResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	StatusCode int32  `protobuf:"varint,1,opt,name=status_code,json=statusCode,proto3" form:"status_code" json:"status_code" query:"status_code"` // status code, 0-success, other values-failure
	StatusMsg  string `protobuf:"bytes,2,opt,name=status_msg,json=statusMsg,proto3" form:"status_msg" json:"status_msg" query:"status_msg"`       // status description
}

func (x *CollectActionResponse) Reset() {
	*x = CollectActionResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_collect_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CollectActionResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CollectActionResponse) ProtoMessage() {}

func (x *CollectActionResponse) ProtoReflect() protoreflect.Message {
	mi := &file_collect_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CollectActionResponse.ProtoReflect.Descriptor instead.
func (*CollectActionResponse) Descriptor() ([]byte, []int) {
	return file_collect_proto_rawDescGZIP(), []int{1}
}

func (x *CollectActionResponse) GetStatusCode() int32 {
	if x != nil {
		return x.StatusCode
	}
	return 0
}

func (x *CollectActionResponse) GetStatusMsg() string {
	if x != nil {
		return x.StatusMsg
	}
	return ""
}

var File_collect_proto protoreflect.FileDescriptor

var file_collect_proto_rawDesc = []byte{
	0x0a, 0x0d, 0x63, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x07, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x1a, 0x09, 0x61, 0x70, 0x69, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x22, 0xc9, 0x01, 0x0a, 0x14, 0x43, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x41,
	0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x35, 0x0a, 0x07,
	0x41, 0x48, 0x61, 0x73, 0x68, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x1b, 0xca,
	0xf3, 0x18, 0x17, 0x6a, 0x73, 0x6f, 0x6e, 0x3a, 0x22, 0x61, 0x48, 0x61, 0x73, 0x68, 0x49, 0x64,
	0x2c, 0x72, 0x65, 0x71, 0x75, 0x69, 0x72, 0x65, 0x64, 0x22, 0x52, 0x07, 0x41, 0x48, 0x61, 0x73,
	0x68, 0x49, 0x64, 0x12, 0x29, 0x0a, 0x03, 0x54, 0x61, 0x67, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x42, 0x17, 0xca, 0xf3, 0x18, 0x13, 0x6a, 0x73, 0x6f, 0x6e, 0x3a, 0x22, 0x74, 0x61, 0x67, 0x2c,
	0x72, 0x65, 0x71, 0x75, 0x69, 0x72, 0x65, 0x64, 0x22, 0x52, 0x03, 0x54, 0x61, 0x67, 0x12, 0x4f,
	0x0a, 0x0a, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x05, 0x42, 0x2f, 0xda, 0xbb, 0x18, 0x0c, 0x24, 0x3d, 0x3d, 0x31, 0x20, 0x7c, 0x7c, 0x20,
	0x24, 0x3d, 0x3d, 0x32, 0xca, 0xf3, 0x18, 0x1b, 0x6a, 0x73, 0x6f, 0x6e, 0x3a, 0x22, 0x61, 0x63,
	0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x2c, 0x72, 0x65, 0x71, 0x75, 0x69, 0x72,
	0x65, 0x64, 0x22, 0x52, 0x0a, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x22,
	0x6f, 0x0a, 0x15, 0x43, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x37, 0x0a, 0x0b, 0x73, 0x74, 0x61, 0x74,
	0x75, 0x73, 0x5f, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x42, 0x16, 0xca,
	0xf3, 0x18, 0x12, 0x6a, 0x73, 0x6f, 0x6e, 0x3a, 0x22, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x5f,
	0x63, 0x6f, 0x64, 0x65, 0x22, 0x52, 0x0a, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x43, 0x6f, 0x64,
	0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x5f, 0x6d, 0x73, 0x67, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x4d, 0x73, 0x67,
	0x32, 0x81, 0x01, 0x0a, 0x0e, 0x43, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x48, 0x61, 0x6e, 0x64,
	0x6c, 0x65, 0x72, 0x12, 0x6f, 0x0a, 0x0d, 0x43, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x41, 0x63,
	0x74, 0x69, 0x6f, 0x6e, 0x12, 0x1d, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x43,
	0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x1e, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x43, 0x6f,
	0x6c, 0x6c, 0x65, 0x63, 0x74, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x22, 0x1f, 0xd2, 0xc1, 0x18, 0x1b, 0x2f, 0x62, 0x6c, 0x6f, 0x67, 0x5f, 0x73,
	0x65, 0x72, 0x76, 0x65, 0x72, 0x2f, 0x63, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x2f, 0x61, 0x63,
	0x74, 0x69, 0x6f, 0x6e, 0x42, 0x42, 0x5a, 0x40, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63,
	0x6f, 0x6d, 0x2f, 0x71, 0x69, 0x6e, 0x67, 0x79, 0x67, 0x67, 0x67, 0x2f, 0x62, 0x6c, 0x6f, 0x67,
	0x5f, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2f, 0x62, 0x69, 0x7a, 0x2f, 0x6d, 0x6f, 0x64, 0x65,
	0x6c, 0x2f, 0x68, 0x65, 0x72, 0x74, 0x7a, 0x2f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x61, 0x63, 0x74,
	0x2f, 0x63, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_collect_proto_rawDescOnce sync.Once
	file_collect_proto_rawDescData = file_collect_proto_rawDesc
)

func file_collect_proto_rawDescGZIP() []byte {
	file_collect_proto_rawDescOnce.Do(func() {
		file_collect_proto_rawDescData = protoimpl.X.CompressGZIP(file_collect_proto_rawDescData)
	})
	return file_collect_proto_rawDescData
}

var file_collect_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_collect_proto_goTypes = []interface{}{
	(*CollectActionRequest)(nil),  // 0: comment.CollectActionRequest
	(*CollectActionResponse)(nil), // 1: comment.CollectActionResponse
}
var file_collect_proto_depIdxs = []int32{
	0, // 0: comment.CollectHandler.CollectAction:input_type -> comment.CollectActionRequest
	1, // 1: comment.CollectHandler.CollectAction:output_type -> comment.CollectActionResponse
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_collect_proto_init() }
func file_collect_proto_init() {
	if File_collect_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_collect_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CollectActionRequest); i {
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
		file_collect_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CollectActionResponse); i {
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
			RawDescriptor: file_collect_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_collect_proto_goTypes,
		DependencyIndexes: file_collect_proto_depIdxs,
		MessageInfos:      file_collect_proto_msgTypes,
	}.Build()
	File_collect_proto = out.File
	file_collect_proto_rawDesc = nil
	file_collect_proto_goTypes = nil
	file_collect_proto_depIdxs = nil
}
