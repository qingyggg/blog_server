// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v5.28.0
// source: relation.proto

package relation

import (
	_ "github.com/qingyggg/blog_server/biz/model/hertz/api"
	common "github.com/qingyggg/blog_server/biz/model/hertz/common"
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

type RelationActionRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ToUserId   int64 `protobuf:"varint,1,opt,name=ToUserId,proto3" form:"ToUserId" json:"to_user_id,required" query:"ToUserId"`
	ActionType int32 `protobuf:"varint,2,opt,name=ActionType,proto3" form:"ActionType" json:"action_type,required" query:"ActionType" vd:"$==1 || $==2"` // 1-Follow, 2-Unfollow
}

func (x *RelationActionRequest) Reset() {
	*x = RelationActionRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_relation_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RelationActionRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RelationActionRequest) ProtoMessage() {}

func (x *RelationActionRequest) ProtoReflect() protoreflect.Message {
	mi := &file_relation_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RelationActionRequest.ProtoReflect.Descriptor instead.
func (*RelationActionRequest) Descriptor() ([]byte, []int) {
	return file_relation_proto_rawDescGZIP(), []int{0}
}

func (x *RelationActionRequest) GetToUserId() int64 {
	if x != nil {
		return x.ToUserId
	}
	return 0
}

func (x *RelationActionRequest) GetActionType() int32 {
	if x != nil {
		return x.ActionType
	}
	return 0
}

type RelationActionResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	StatusCode int32  `protobuf:"varint,1,opt,name=status_code,json=statusCode,proto3" form:"status_code" json:"status_code" query:"status_code"`
	StatusMsg  string `protobuf:"bytes,2,opt,name=status_msg,json=statusMsg,proto3" form:"status_msg" json:"status_msg,omitempty" query:"status_msg"`
}

func (x *RelationActionResponse) Reset() {
	*x = RelationActionResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_relation_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RelationActionResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RelationActionResponse) ProtoMessage() {}

func (x *RelationActionResponse) ProtoReflect() protoreflect.Message {
	mi := &file_relation_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RelationActionResponse.ProtoReflect.Descriptor instead.
func (*RelationActionResponse) Descriptor() ([]byte, []int) {
	return file_relation_proto_rawDescGZIP(), []int{1}
}

func (x *RelationActionResponse) GetStatusCode() int32 {
	if x != nil {
		return x.StatusCode
	}
	return 0
}

func (x *RelationActionResponse) GetStatusMsg() string {
	if x != nil {
		return x.StatusMsg
	}
	return ""
}

type RelationFollowListRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId int64 `protobuf:"varint,1,opt,name=UserId,proto3" form:"UserId" json:"UserId,omitempty" query:"user_id,required"`
}

func (x *RelationFollowListRequest) Reset() {
	*x = RelationFollowListRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_relation_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RelationFollowListRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RelationFollowListRequest) ProtoMessage() {}

func (x *RelationFollowListRequest) ProtoReflect() protoreflect.Message {
	mi := &file_relation_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RelationFollowListRequest.ProtoReflect.Descriptor instead.
func (*RelationFollowListRequest) Descriptor() ([]byte, []int) {
	return file_relation_proto_rawDescGZIP(), []int{2}
}

func (x *RelationFollowListRequest) GetUserId() int64 {
	if x != nil {
		return x.UserId
	}
	return 0
}

type RelationFollowListResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	StatusCode int32          `protobuf:"varint,1,opt,name=status_code,json=statusCode,proto3" form:"status_code" json:"status_code" query:"status_code"`
	StatusMsg  string         `protobuf:"bytes,2,opt,name=status_msg,json=statusMsg,proto3" form:"status_msg" json:"status_msg,omitempty" query:"status_msg"`
	UserList   []*common.User `protobuf:"bytes,3,rep,name=user_list,json=userList,proto3" form:"user_list" json:"user_list,omitempty" query:"user_list"` // User information list
}

func (x *RelationFollowListResponse) Reset() {
	*x = RelationFollowListResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_relation_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RelationFollowListResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RelationFollowListResponse) ProtoMessage() {}

func (x *RelationFollowListResponse) ProtoReflect() protoreflect.Message {
	mi := &file_relation_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RelationFollowListResponse.ProtoReflect.Descriptor instead.
func (*RelationFollowListResponse) Descriptor() ([]byte, []int) {
	return file_relation_proto_rawDescGZIP(), []int{3}
}

func (x *RelationFollowListResponse) GetStatusCode() int32 {
	if x != nil {
		return x.StatusCode
	}
	return 0
}

func (x *RelationFollowListResponse) GetStatusMsg() string {
	if x != nil {
		return x.StatusMsg
	}
	return ""
}

func (x *RelationFollowListResponse) GetUserList() []*common.User {
	if x != nil {
		return x.UserList
	}
	return nil
}

var File_relation_proto protoreflect.FileDescriptor

var file_relation_proto_rawDesc = []byte{
	0x0a, 0x0e, 0x72, 0x65, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x08, 0x72, 0x65, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x1a, 0x09, 0x61, 0x70, 0x69, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x0c, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x22, 0xa4, 0x01, 0x0a, 0x15, 0x52, 0x65, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x3a, 0x0a,
	0x08, 0x54, 0x6f, 0x55, 0x73, 0x65, 0x72, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x42,
	0x1e, 0xca, 0xf3, 0x18, 0x1a, 0x6a, 0x73, 0x6f, 0x6e, 0x3a, 0x22, 0x74, 0x6f, 0x5f, 0x75, 0x73,
	0x65, 0x72, 0x5f, 0x69, 0x64, 0x2c, 0x72, 0x65, 0x71, 0x75, 0x69, 0x72, 0x65, 0x64, 0x22, 0x52,
	0x08, 0x54, 0x6f, 0x55, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x4f, 0x0a, 0x0a, 0x41, 0x63, 0x74,
	0x69, 0x6f, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x42, 0x2f, 0xda,
	0xbb, 0x18, 0x0c, 0x24, 0x3d, 0x3d, 0x31, 0x20, 0x7c, 0x7c, 0x20, 0x24, 0x3d, 0x3d, 0x32, 0xca,
	0xf3, 0x18, 0x1b, 0x6a, 0x73, 0x6f, 0x6e, 0x3a, 0x22, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x5f,
	0x74, 0x79, 0x70, 0x65, 0x2c, 0x72, 0x65, 0x71, 0x75, 0x69, 0x72, 0x65, 0x64, 0x22, 0x52, 0x0a,
	0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x22, 0x70, 0x0a, 0x16, 0x52, 0x65,
	0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x37, 0x0a, 0x0b, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x5f, 0x63,
	0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x42, 0x16, 0xca, 0xf3, 0x18, 0x12, 0x6a,
	0x73, 0x6f, 0x6e, 0x3a, 0x22, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x5f, 0x63, 0x6f, 0x64, 0x65,
	0x22, 0x52, 0x0a, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x1d, 0x0a,
	0x0a, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x5f, 0x6d, 0x73, 0x67, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x09, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x4d, 0x73, 0x67, 0x22, 0x51, 0x0a, 0x19,
	0x52, 0x65, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x46, 0x6f, 0x6c, 0x6c, 0x6f, 0x77, 0x4c, 0x69,
	0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x34, 0x0a, 0x06, 0x55, 0x73, 0x65,
	0x72, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x42, 0x1c, 0xca, 0xf3, 0x18, 0x18, 0x71,
	0x75, 0x65, 0x72, 0x79, 0x3a, 0x22, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x2c, 0x72, 0x65,
	0x71, 0x75, 0x69, 0x72, 0x65, 0x64, 0x22, 0x52, 0x06, 0x55, 0x73, 0x65, 0x72, 0x49, 0x64, 0x22,
	0x98, 0x01, 0x0a, 0x1a, 0x52, 0x65, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x46, 0x6f, 0x6c, 0x6c,
	0x6f, 0x77, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x37,
	0x0a, 0x0b, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x5f, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x05, 0x42, 0x16, 0xca, 0xf3, 0x18, 0x12, 0x6a, 0x73, 0x6f, 0x6e, 0x3a, 0x22, 0x73,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x5f, 0x63, 0x6f, 0x64, 0x65, 0x22, 0x52, 0x0a, 0x73, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x73, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x5f, 0x6d, 0x73, 0x67, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x73, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x4d, 0x73, 0x67, 0x12, 0x22, 0x0a, 0x09, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x6c,
	0x69, 0x73, 0x74, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x05, 0x2e, 0x55, 0x73, 0x65, 0x72,
	0x52, 0x08, 0x75, 0x73, 0x65, 0x72, 0x4c, 0x69, 0x73, 0x74, 0x32, 0x9e, 0x03, 0x0a, 0x0f, 0x52,
	0x65, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x48, 0x61, 0x6e, 0x64, 0x6c, 0x65, 0x72, 0x12, 0x75,
	0x0a, 0x0e, 0x52, 0x65, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e,
	0x12, 0x1f, 0x2e, 0x72, 0x65, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x52, 0x65, 0x6c, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x20, 0x2e, 0x72, 0x65, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x52, 0x65, 0x6c,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x22, 0x20, 0xd2, 0xc1, 0x18, 0x1c, 0x2f, 0x62, 0x6c, 0x6f, 0x67, 0x5f, 0x73,
	0x65, 0x72, 0x76, 0x65, 0x72, 0x2f, 0x72, 0x65, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x61,
	0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x86, 0x01, 0x0a, 0x12, 0x52, 0x65, 0x6c, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x46, 0x6f, 0x6c, 0x6c, 0x6f, 0x77, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x23, 0x2e, 0x72,
	0x65, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x52, 0x65, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x46, 0x6f, 0x6c, 0x6c, 0x6f, 0x77, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x24, 0x2e, 0x72, 0x65, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x52, 0x65, 0x6c,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x46, 0x6f, 0x6c, 0x6c, 0x6f, 0x77, 0x4c, 0x69, 0x73, 0x74, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x25, 0xca, 0xc1, 0x18, 0x21, 0x2f, 0x62, 0x6c,
	0x6f, 0x67, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2f, 0x72, 0x65, 0x6c, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x2f, 0x66, 0x6f, 0x6c, 0x6c, 0x6f, 0x77, 0x2f, 0x6c, 0x69, 0x73, 0x74, 0x12, 0x8a,
	0x01, 0x0a, 0x14, 0x52, 0x65, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x46, 0x6f, 0x6c, 0x6c, 0x6f,
	0x77, 0x65, 0x72, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x23, 0x2e, 0x72, 0x65, 0x6c, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x2e, 0x52, 0x65, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x46, 0x6f, 0x6c, 0x6c, 0x6f,
	0x77, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x24, 0x2e, 0x72,
	0x65, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x52, 0x65, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x46, 0x6f, 0x6c, 0x6c, 0x6f, 0x77, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x22, 0x27, 0xca, 0xc1, 0x18, 0x23, 0x2f, 0x62, 0x6c, 0x6f, 0x67, 0x5f, 0x73, 0x65,
	0x72, 0x76, 0x65, 0x72, 0x2f, 0x72, 0x65, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x66, 0x6f,
	0x6c, 0x6c, 0x6f, 0x77, 0x65, 0x72, 0x2f, 0x6c, 0x69, 0x73, 0x74, 0x42, 0x41, 0x5a, 0x3f, 0x67,
	0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x71, 0x69, 0x6e, 0x67, 0x79, 0x67,
	0x67, 0x67, 0x2f, 0x62, 0x6c, 0x6f, 0x67, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2f, 0x62,
	0x69, 0x7a, 0x2f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2f, 0x68, 0x65, 0x72, 0x74, 0x7a, 0x2f, 0x73,
	0x6f, 0x63, 0x69, 0x61, 0x6c, 0x2f, 0x72, 0x65, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_relation_proto_rawDescOnce sync.Once
	file_relation_proto_rawDescData = file_relation_proto_rawDesc
)

func file_relation_proto_rawDescGZIP() []byte {
	file_relation_proto_rawDescOnce.Do(func() {
		file_relation_proto_rawDescData = protoimpl.X.CompressGZIP(file_relation_proto_rawDescData)
	})
	return file_relation_proto_rawDescData
}

var file_relation_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_relation_proto_goTypes = []interface{}{
	(*RelationActionRequest)(nil),      // 0: relation.RelationActionRequest
	(*RelationActionResponse)(nil),     // 1: relation.RelationActionResponse
	(*RelationFollowListRequest)(nil),  // 2: relation.RelationFollowListRequest
	(*RelationFollowListResponse)(nil), // 3: relation.RelationFollowListResponse
	(*common.User)(nil),                // 4: User
}
var file_relation_proto_depIdxs = []int32{
	4, // 0: relation.RelationFollowListResponse.user_list:type_name -> User
	0, // 1: relation.RelationHandler.RelationAction:input_type -> relation.RelationActionRequest
	2, // 2: relation.RelationHandler.RelationFollowList:input_type -> relation.RelationFollowListRequest
	2, // 3: relation.RelationHandler.RelationFollowerList:input_type -> relation.RelationFollowListRequest
	1, // 4: relation.RelationHandler.RelationAction:output_type -> relation.RelationActionResponse
	3, // 5: relation.RelationHandler.RelationFollowList:output_type -> relation.RelationFollowListResponse
	3, // 6: relation.RelationHandler.RelationFollowerList:output_type -> relation.RelationFollowListResponse
	4, // [4:7] is the sub-list for method output_type
	1, // [1:4] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_relation_proto_init() }
func file_relation_proto_init() {
	if File_relation_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_relation_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RelationActionRequest); i {
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
		file_relation_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RelationActionResponse); i {
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
		file_relation_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RelationFollowListRequest); i {
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
		file_relation_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RelationFollowListResponse); i {
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
			RawDescriptor: file_relation_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_relation_proto_goTypes,
		DependencyIndexes: file_relation_proto_depIdxs,
		MessageInfos:      file_relation_proto_msgTypes,
	}.Build()
	File_relation_proto = out.File
	file_relation_proto_rawDesc = nil
	file_relation_proto_goTypes = nil
	file_relation_proto_depIdxs = nil
}
