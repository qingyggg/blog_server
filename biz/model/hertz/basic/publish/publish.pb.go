// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v5.28.0
// source: publish.proto

package publish

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

type ArticleBaseActionRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Uid     int64  `protobuf:"varint,1,opt,name=uid,proto3" form:"uid" json:"uid,omitempty" query:"uid"`                //作者id
	AHashID string `protobuf:"bytes,2,opt,name=aHashID,proto3" form:"aHashID" json:"aHashID,omitempty" query:"aHashID"` //文章的hashid
}

func (x *ArticleBaseActionRequest) Reset() {
	*x = ArticleBaseActionRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_publish_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ArticleBaseActionRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ArticleBaseActionRequest) ProtoMessage() {}

func (x *ArticleBaseActionRequest) ProtoReflect() protoreflect.Message {
	mi := &file_publish_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ArticleBaseActionRequest.ProtoReflect.Descriptor instead.
func (*ArticleBaseActionRequest) Descriptor() ([]byte, []int) {
	return file_publish_proto_rawDescGZIP(), []int{0}
}

func (x *ArticleBaseActionRequest) GetUid() int64 {
	if x != nil {
		return x.Uid
	}
	return 0
}

func (x *ArticleBaseActionRequest) GetAHashID() string {
	if x != nil {
		return x.AHashID
	}
	return ""
}

// 文章的创建
type ArticleCreateActionRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Uid     int64               `protobuf:"varint,1,opt,name=uid,proto3" form:"uid" json:"uid,omitempty" query:"uid"`
	Payload *common.ArticleBase `protobuf:"bytes,2,opt,name=payload,proto3" form:"payload" json:"payload,omitempty" query:"payload"`
}

func (x *ArticleCreateActionRequest) Reset() {
	*x = ArticleCreateActionRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_publish_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ArticleCreateActionRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ArticleCreateActionRequest) ProtoMessage() {}

func (x *ArticleCreateActionRequest) ProtoReflect() protoreflect.Message {
	mi := &file_publish_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ArticleCreateActionRequest.ProtoReflect.Descriptor instead.
func (*ArticleCreateActionRequest) Descriptor() ([]byte, []int) {
	return file_publish_proto_rawDescGZIP(), []int{1}
}

func (x *ArticleCreateActionRequest) GetUid() int64 {
	if x != nil {
		return x.Uid
	}
	return 0
}

func (x *ArticleCreateActionRequest) GetPayload() *common.ArticleBase {
	if x != nil {
		return x.Payload
	}
	return nil
}

// 文章的修改
type ArticleModifyActionRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Base    *ArticleBaseActionRequest `protobuf:"bytes,1,opt,name=base,proto3" form:"base" json:"base,omitempty" query:"base"`
	Payload *common.ArticleBase       `protobuf:"bytes,2,opt,name=payload,proto3" form:"payload" json:"payload,omitempty" query:"payload"`
}

func (x *ArticleModifyActionRequest) Reset() {
	*x = ArticleModifyActionRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_publish_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ArticleModifyActionRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ArticleModifyActionRequest) ProtoMessage() {}

func (x *ArticleModifyActionRequest) ProtoReflect() protoreflect.Message {
	mi := &file_publish_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ArticleModifyActionRequest.ProtoReflect.Descriptor instead.
func (*ArticleModifyActionRequest) Descriptor() ([]byte, []int) {
	return file_publish_proto_rawDescGZIP(), []int{2}
}

func (x *ArticleModifyActionRequest) GetBase() *ArticleBaseActionRequest {
	if x != nil {
		return x.Base
	}
	return nil
}

func (x *ArticleModifyActionRequest) GetPayload() *common.ArticleBase {
	if x != nil {
		return x.Payload
	}
	return nil
}

type ArticleActionResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	StatusCode int32  `protobuf:"varint,1,opt,name=status_code,json=statusCode,proto3" form:"status_code" json:"status_code" query:"status_code"`
	StatusMsg  string `protobuf:"bytes,2,opt,name=status_msg,json=statusMsg,proto3" form:"status_msg" json:"status_msg,omitempty" query:"status_msg"`
}

func (x *ArticleActionResponse) Reset() {
	*x = ArticleActionResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_publish_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ArticleActionResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ArticleActionResponse) ProtoMessage() {}

func (x *ArticleActionResponse) ProtoReflect() protoreflect.Message {
	mi := &file_publish_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ArticleActionResponse.ProtoReflect.Descriptor instead.
func (*ArticleActionResponse) Descriptor() ([]byte, []int) {
	return file_publish_proto_rawDescGZIP(), []int{3}
}

func (x *ArticleActionResponse) GetStatusCode() int32 {
	if x != nil {
		return x.StatusCode
	}
	return 0
}

func (x *ArticleActionResponse) GetStatusMsg() string {
	if x != nil {
		return x.StatusMsg
	}
	return ""
}

// 文章卡片列表
type ArticleCardsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId int64 `protobuf:"varint,1,opt,name=user_id,json=userId,proto3" form:"user_id" json:"user_id,omitempty" query:"user_id"` //如果id为0，则根据时间的先后去请求文章列表
	Offset int32 `protobuf:"varint,2,opt,name=offset,proto3" form:"offset" json:"offset,omitempty" query:"offset"`
}

func (x *ArticleCardsRequest) Reset() {
	*x = ArticleCardsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_publish_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ArticleCardsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ArticleCardsRequest) ProtoMessage() {}

func (x *ArticleCardsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_publish_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ArticleCardsRequest.ProtoReflect.Descriptor instead.
func (*ArticleCardsRequest) Descriptor() ([]byte, []int) {
	return file_publish_proto_rawDescGZIP(), []int{4}
}

func (x *ArticleCardsRequest) GetUserId() int64 {
	if x != nil {
		return x.UserId
	}
	return 0
}

func (x *ArticleCardsRequest) GetOffset() int32 {
	if x != nil {
		return x.Offset
	}
	return 0
}

type ArticleResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	StatusCode int32           `protobuf:"varint,1,opt,name=status_code,json=statusCode,proto3" form:"status_code" json:"status_code" query:"status_code"`
	StatusMsg  string          `protobuf:"bytes,2,opt,name=status_msg,json=statusMsg,proto3" form:"status_msg" json:"status_msg,omitempty" query:"status_msg"`
	Article    *common.Article `protobuf:"bytes,3,opt,name=article,proto3" form:"article" json:"article,omitempty" query:"article"`
}

func (x *ArticleResponse) Reset() {
	*x = ArticleResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_publish_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ArticleResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ArticleResponse) ProtoMessage() {}

func (x *ArticleResponse) ProtoReflect() protoreflect.Message {
	mi := &file_publish_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ArticleResponse.ProtoReflect.Descriptor instead.
func (*ArticleResponse) Descriptor() ([]byte, []int) {
	return file_publish_proto_rawDescGZIP(), []int{5}
}

func (x *ArticleResponse) GetStatusCode() int32 {
	if x != nil {
		return x.StatusCode
	}
	return 0
}

func (x *ArticleResponse) GetStatusMsg() string {
	if x != nil {
		return x.StatusMsg
	}
	return ""
}

func (x *ArticleResponse) GetArticle() *common.Article {
	if x != nil {
		return x.Article
	}
	return nil
}

type ArticleCardsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	StatusCode int32                 `protobuf:"varint,1,opt,name=status_code,json=statusCode,proto3" form:"status_code" json:"status_code" query:"status_code"`
	StatusMsg  string                `protobuf:"bytes,2,opt,name=status_msg,json=statusMsg,proto3" form:"status_msg" json:"status_msg,omitempty" query:"status_msg"`
	List       []*common.ArticleCard `protobuf:"bytes,3,rep,name=list,proto3" form:"list" json:"list,omitempty" query:"list"`
}

func (x *ArticleCardsResponse) Reset() {
	*x = ArticleCardsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_publish_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ArticleCardsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ArticleCardsResponse) ProtoMessage() {}

func (x *ArticleCardsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_publish_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ArticleCardsResponse.ProtoReflect.Descriptor instead.
func (*ArticleCardsResponse) Descriptor() ([]byte, []int) {
	return file_publish_proto_rawDescGZIP(), []int{6}
}

func (x *ArticleCardsResponse) GetStatusCode() int32 {
	if x != nil {
		return x.StatusCode
	}
	return 0
}

func (x *ArticleCardsResponse) GetStatusMsg() string {
	if x != nil {
		return x.StatusMsg
	}
	return ""
}

func (x *ArticleCardsResponse) GetList() []*common.ArticleCard {
	if x != nil {
		return x.List
	}
	return nil
}

var File_publish_proto protoreflect.FileDescriptor

var file_publish_proto_rawDesc = []byte{
	0x0a, 0x0d, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x73, 0x68, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x07, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x73, 0x68, 0x1a, 0x09, 0x61, 0x70, 0x69, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x0c, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x22, 0x46, 0x0a, 0x18, 0x41, 0x72, 0x74, 0x69, 0x63, 0x6c, 0x65, 0x42, 0x61, 0x73, 0x65,
	0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x10, 0x0a,
	0x03, 0x75, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x03, 0x75, 0x69, 0x64, 0x12,
	0x18, 0x0a, 0x07, 0x61, 0x48, 0x61, 0x73, 0x68, 0x49, 0x44, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x07, 0x61, 0x48, 0x61, 0x73, 0x68, 0x49, 0x44, 0x22, 0x56, 0x0a, 0x1a, 0x41, 0x72, 0x74,
	0x69, 0x63, 0x6c, 0x65, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x10, 0x0a, 0x03, 0x75, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x03, 0x75, 0x69, 0x64, 0x12, 0x26, 0x0a, 0x07, 0x70, 0x61, 0x79,
	0x6c, 0x6f, 0x61, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x41, 0x72, 0x74,
	0x69, 0x63, 0x6c, 0x65, 0x42, 0x61, 0x73, 0x65, 0x52, 0x07, 0x70, 0x61, 0x79, 0x6c, 0x6f, 0x61,
	0x64, 0x22, 0x7b, 0x0a, 0x1a, 0x41, 0x72, 0x74, 0x69, 0x63, 0x6c, 0x65, 0x4d, 0x6f, 0x64, 0x69,
	0x66, 0x79, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x35, 0x0a, 0x04, 0x62, 0x61, 0x73, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x21, 0x2e,
	0x70, 0x75, 0x62, 0x6c, 0x69, 0x73, 0x68, 0x2e, 0x41, 0x72, 0x74, 0x69, 0x63, 0x6c, 0x65, 0x42,
	0x61, 0x73, 0x65, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x52, 0x04, 0x62, 0x61, 0x73, 0x65, 0x12, 0x26, 0x0a, 0x07, 0x70, 0x61, 0x79, 0x6c, 0x6f, 0x61,
	0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x41, 0x72, 0x74, 0x69, 0x63, 0x6c,
	0x65, 0x42, 0x61, 0x73, 0x65, 0x52, 0x07, 0x70, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x22, 0x6f,
	0x0a, 0x15, 0x41, 0x72, 0x74, 0x69, 0x63, 0x6c, 0x65, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x37, 0x0a, 0x0b, 0x73, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x5f, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x42, 0x16, 0xca, 0xf3,
	0x18, 0x12, 0x6a, 0x73, 0x6f, 0x6e, 0x3a, 0x22, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x5f, 0x63,
	0x6f, 0x64, 0x65, 0x22, 0x52, 0x0a, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x43, 0x6f, 0x64, 0x65,
	0x12, 0x1d, 0x0a, 0x0a, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x5f, 0x6d, 0x73, 0x67, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x4d, 0x73, 0x67, 0x22,
	0x46, 0x0a, 0x13, 0x41, 0x72, 0x74, 0x69, 0x63, 0x6c, 0x65, 0x43, 0x61, 0x72, 0x64, 0x73, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12,
	0x16, 0x0a, 0x06, 0x6f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x06, 0x6f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x22, 0x8d, 0x01, 0x0a, 0x0f, 0x41, 0x72, 0x74, 0x69,
	0x63, 0x6c, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x37, 0x0a, 0x0b, 0x73,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x5f, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05,
	0x42, 0x16, 0xca, 0xf3, 0x18, 0x12, 0x6a, 0x73, 0x6f, 0x6e, 0x3a, 0x22, 0x73, 0x74, 0x61, 0x74,
	0x75, 0x73, 0x5f, 0x63, 0x6f, 0x64, 0x65, 0x22, 0x52, 0x0a, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73,
	0x43, 0x6f, 0x64, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x5f, 0x6d,
	0x73, 0x67, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73,
	0x4d, 0x73, 0x67, 0x12, 0x22, 0x0a, 0x07, 0x61, 0x72, 0x74, 0x69, 0x63, 0x6c, 0x65, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x08, 0x2e, 0x41, 0x72, 0x74, 0x69, 0x63, 0x6c, 0x65, 0x52, 0x07,
	0x61, 0x72, 0x74, 0x69, 0x63, 0x6c, 0x65, 0x22, 0x90, 0x01, 0x0a, 0x14, 0x41, 0x72, 0x74, 0x69,
	0x63, 0x6c, 0x65, 0x43, 0x61, 0x72, 0x64, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x37, 0x0a, 0x0b, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x5f, 0x63, 0x6f, 0x64, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x05, 0x42, 0x16, 0xca, 0xf3, 0x18, 0x12, 0x6a, 0x73, 0x6f, 0x6e, 0x3a,
	0x22, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x5f, 0x63, 0x6f, 0x64, 0x65, 0x22, 0x52, 0x0a, 0x73,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x73, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x5f, 0x6d, 0x73, 0x67, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x73,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x4d, 0x73, 0x67, 0x12, 0x20, 0x0a, 0x04, 0x6c, 0x69, 0x73, 0x74,
	0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x41, 0x72, 0x74, 0x69, 0x63, 0x6c, 0x65,
	0x43, 0x61, 0x72, 0x64, 0x52, 0x04, 0x6c, 0x69, 0x73, 0x74, 0x32, 0xd3, 0x05, 0x0a, 0x0e, 0x50,
	0x75, 0x62, 0x6c, 0x69, 0x73, 0x68, 0x48, 0x61, 0x6e, 0x64, 0x6c, 0x65, 0x72, 0x12, 0x75, 0x0a,
	0x0d, 0x50, 0x75, 0x62, 0x6c, 0x69, 0x73, 0x68, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x23,
	0x2e, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x73, 0x68, 0x2e, 0x41, 0x72, 0x74, 0x69, 0x63, 0x6c, 0x65,
	0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x1e, 0x2e, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x73, 0x68, 0x2e, 0x41, 0x72,
	0x74, 0x69, 0x63, 0x6c, 0x65, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x22, 0x1f, 0xd2, 0xc1, 0x18, 0x1b, 0x2f, 0x62, 0x6c, 0x6f, 0x67, 0x5f, 0x73,
	0x65, 0x72, 0x76, 0x65, 0x72, 0x2f, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x73, 0x68, 0x2f, 0x61, 0x63,
	0x74, 0x69, 0x6f, 0x6e, 0x12, 0x7b, 0x0a, 0x13, 0x50, 0x75, 0x62, 0x6c, 0x69, 0x73, 0x68, 0x4d,
	0x6f, 0x64, 0x69, 0x66, 0x79, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x23, 0x2e, 0x70, 0x75,
	0x62, 0x6c, 0x69, 0x73, 0x68, 0x2e, 0x41, 0x72, 0x74, 0x69, 0x63, 0x6c, 0x65, 0x4d, 0x6f, 0x64,
	0x69, 0x66, 0x79, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x1e, 0x2e, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x73, 0x68, 0x2e, 0x41, 0x72, 0x74, 0x69, 0x63,
	0x6c, 0x65, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x22, 0x1f, 0xea, 0xc1, 0x18, 0x1b, 0x2f, 0x62, 0x6c, 0x6f, 0x67, 0x5f, 0x73, 0x65, 0x72, 0x76,
	0x65, 0x72, 0x2f, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x73, 0x68, 0x2f, 0x61, 0x63, 0x74, 0x69, 0x6f,
	0x6e, 0x12, 0x76, 0x0a, 0x10, 0x50, 0x75, 0x62, 0x6c, 0x69, 0x73, 0x68, 0x44, 0x65, 0x6c, 0x41,
	0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x21, 0x2e, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x73, 0x68, 0x2e,
	0x41, 0x72, 0x74, 0x69, 0x63, 0x6c, 0x65, 0x42, 0x61, 0x73, 0x65, 0x41, 0x63, 0x74, 0x69, 0x6f,
	0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1e, 0x2e, 0x70, 0x75, 0x62, 0x6c, 0x69,
	0x73, 0x68, 0x2e, 0x41, 0x72, 0x74, 0x69, 0x63, 0x6c, 0x65, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x1f, 0xe2, 0xc1, 0x18, 0x1b, 0x2f, 0x62,
	0x6c, 0x6f, 0x67, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2f, 0x70, 0x75, 0x62, 0x6c, 0x69,
	0x73, 0x68, 0x2f, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x69, 0x0a, 0x0b, 0x50, 0x75, 0x62,
	0x6c, 0x69, 0x73, 0x68, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x1c, 0x2e, 0x70, 0x75, 0x62, 0x6c, 0x69,
	0x73, 0x68, 0x2e, 0x41, 0x72, 0x74, 0x69, 0x63, 0x6c, 0x65, 0x43, 0x61, 0x72, 0x64, 0x73, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1d, 0x2e, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x73, 0x68,
	0x2e, 0x41, 0x72, 0x74, 0x69, 0x63, 0x6c, 0x65, 0x43, 0x61, 0x72, 0x64, 0x73, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x1d, 0xca, 0xc1, 0x18, 0x19, 0x2f, 0x62, 0x6c, 0x6f, 0x67,
	0x5f, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2f, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x73, 0x68, 0x2f,
	0x6c, 0x69, 0x73, 0x74, 0x12, 0x6d, 0x0a, 0x0d, 0x50, 0x75, 0x62, 0x6c, 0x69, 0x73, 0x68, 0x44,
	0x65, 0x74, 0x61, 0x69, 0x6c, 0x12, 0x21, 0x2e, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x73, 0x68, 0x2e,
	0x41, 0x72, 0x74, 0x69, 0x63, 0x6c, 0x65, 0x42, 0x61, 0x73, 0x65, 0x41, 0x63, 0x74, 0x69, 0x6f,
	0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x18, 0x2e, 0x70, 0x75, 0x62, 0x6c, 0x69,
	0x73, 0x68, 0x2e, 0x41, 0x72, 0x74, 0x69, 0x63, 0x6c, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x22, 0x1f, 0xca, 0xc1, 0x18, 0x1b, 0x2f, 0x62, 0x6c, 0x6f, 0x67, 0x5f, 0x73, 0x65,
	0x72, 0x76, 0x65, 0x72, 0x2f, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x73, 0x68, 0x2f, 0x64, 0x65, 0x74,
	0x61, 0x69, 0x6c, 0x12, 0x7b, 0x0a, 0x13, 0x50, 0x75, 0x62, 0x6c, 0x69, 0x73, 0x68, 0x56, 0x69,
	0x65, 0x77, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x41, 0x64, 0x64, 0x12, 0x21, 0x2e, 0x70, 0x75, 0x62,
	0x6c, 0x69, 0x73, 0x68, 0x2e, 0x41, 0x72, 0x74, 0x69, 0x63, 0x6c, 0x65, 0x42, 0x61, 0x73, 0x65,
	0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1e, 0x2e,
	0x70, 0x75, 0x62, 0x6c, 0x69, 0x73, 0x68, 0x2e, 0x41, 0x72, 0x74, 0x69, 0x63, 0x6c, 0x65, 0x41,
	0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x21, 0xd2,
	0xc1, 0x18, 0x1d, 0x2f, 0x62, 0x6c, 0x6f, 0x67, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2f,
	0x70, 0x75, 0x62, 0x6c, 0x69, 0x73, 0x68, 0x2f, 0x76, 0x69, 0x65, 0x77, 0x5f, 0x61, 0x64, 0x64,
	0x42, 0x3f, 0x5a, 0x3d, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x71,
	0x69, 0x6e, 0x67, 0x79, 0x67, 0x67, 0x67, 0x2f, 0x62, 0x6c, 0x6f, 0x67, 0x5f, 0x73, 0x65, 0x72,
	0x76, 0x65, 0x72, 0x2f, 0x62, 0x69, 0x7a, 0x2f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2f, 0x68, 0x65,
	0x72, 0x74, 0x7a, 0x2f, 0x62, 0x61, 0x73, 0x69, 0x63, 0x2f, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x73,
	0x68, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_publish_proto_rawDescOnce sync.Once
	file_publish_proto_rawDescData = file_publish_proto_rawDesc
)

func file_publish_proto_rawDescGZIP() []byte {
	file_publish_proto_rawDescOnce.Do(func() {
		file_publish_proto_rawDescData = protoimpl.X.CompressGZIP(file_publish_proto_rawDescData)
	})
	return file_publish_proto_rawDescData
}

var file_publish_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_publish_proto_goTypes = []interface{}{
	(*ArticleBaseActionRequest)(nil),   // 0: publish.ArticleBaseActionRequest
	(*ArticleCreateActionRequest)(nil), // 1: publish.ArticleCreateActionRequest
	(*ArticleModifyActionRequest)(nil), // 2: publish.ArticleModifyActionRequest
	(*ArticleActionResponse)(nil),      // 3: publish.ArticleActionResponse
	(*ArticleCardsRequest)(nil),        // 4: publish.ArticleCardsRequest
	(*ArticleResponse)(nil),            // 5: publish.ArticleResponse
	(*ArticleCardsResponse)(nil),       // 6: publish.ArticleCardsResponse
	(*common.ArticleBase)(nil),         // 7: ArticleBase
	(*common.Article)(nil),             // 8: Article
	(*common.ArticleCard)(nil),         // 9: ArticleCard
}
var file_publish_proto_depIdxs = []int32{
	7,  // 0: publish.ArticleCreateActionRequest.payload:type_name -> ArticleBase
	0,  // 1: publish.ArticleModifyActionRequest.base:type_name -> publish.ArticleBaseActionRequest
	7,  // 2: publish.ArticleModifyActionRequest.payload:type_name -> ArticleBase
	8,  // 3: publish.ArticleResponse.article:type_name -> Article
	9,  // 4: publish.ArticleCardsResponse.list:type_name -> ArticleCard
	1,  // 5: publish.PublishHandler.PublishAction:input_type -> publish.ArticleCreateActionRequest
	2,  // 6: publish.PublishHandler.PublishModifyAction:input_type -> publish.ArticleModifyActionRequest
	0,  // 7: publish.PublishHandler.PublishDelAction:input_type -> publish.ArticleBaseActionRequest
	4,  // 8: publish.PublishHandler.PublishList:input_type -> publish.ArticleCardsRequest
	0,  // 9: publish.PublishHandler.PublishDetail:input_type -> publish.ArticleBaseActionRequest
	0,  // 10: publish.PublishHandler.PublishViewCountAdd:input_type -> publish.ArticleBaseActionRequest
	3,  // 11: publish.PublishHandler.PublishAction:output_type -> publish.ArticleActionResponse
	3,  // 12: publish.PublishHandler.PublishModifyAction:output_type -> publish.ArticleActionResponse
	3,  // 13: publish.PublishHandler.PublishDelAction:output_type -> publish.ArticleActionResponse
	6,  // 14: publish.PublishHandler.PublishList:output_type -> publish.ArticleCardsResponse
	5,  // 15: publish.PublishHandler.PublishDetail:output_type -> publish.ArticleResponse
	3,  // 16: publish.PublishHandler.PublishViewCountAdd:output_type -> publish.ArticleActionResponse
	11, // [11:17] is the sub-list for method output_type
	5,  // [5:11] is the sub-list for method input_type
	5,  // [5:5] is the sub-list for extension type_name
	5,  // [5:5] is the sub-list for extension extendee
	0,  // [0:5] is the sub-list for field type_name
}

func init() { file_publish_proto_init() }
func file_publish_proto_init() {
	if File_publish_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_publish_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ArticleBaseActionRequest); i {
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
		file_publish_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ArticleCreateActionRequest); i {
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
		file_publish_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ArticleModifyActionRequest); i {
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
		file_publish_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ArticleActionResponse); i {
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
		file_publish_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ArticleCardsRequest); i {
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
		file_publish_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ArticleResponse); i {
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
		file_publish_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ArticleCardsResponse); i {
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
			RawDescriptor: file_publish_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_publish_proto_goTypes,
		DependencyIndexes: file_publish_proto_depIdxs,
		MessageInfos:      file_publish_proto_msgTypes,
	}.Build()
	File_publish_proto = out.File
	file_publish_proto_rawDesc = nil
	file_publish_proto_goTypes = nil
	file_publish_proto_depIdxs = nil
}