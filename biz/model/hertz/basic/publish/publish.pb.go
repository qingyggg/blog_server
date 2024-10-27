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

type DetailRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AHashID string `protobuf:"bytes,1,opt,name=AHashID,proto3" form:"AHashID" json:"AHashID" query:"aHashID,required"` //文章的hashid
	UHashID string `protobuf:"bytes,2,opt,name=UHashID,proto3" form:"UHashID" json:"UHashID" query:"uHashID,required"`
}

func (x *DetailRequest) Reset() {
	*x = DetailRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_publish_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DetailRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DetailRequest) ProtoMessage() {}

func (x *DetailRequest) ProtoReflect() protoreflect.Message {
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

// Deprecated: Use DetailRequest.ProtoReflect.Descriptor instead.
func (*DetailRequest) Descriptor() ([]byte, []int) {
	return file_publish_proto_rawDescGZIP(), []int{0}
}

func (x *DetailRequest) GetAHashID() string {
	if x != nil {
		return x.AHashID
	}
	return ""
}

func (x *DetailRequest) GetUHashID() string {
	if x != nil {
		return x.UHashID
	}
	return ""
}

type ActionRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AHashID string `protobuf:"bytes,1,opt,name=AHashID,proto3" form:"AHashID" json:"aHashID,required" query:"AHashID"`
}

func (x *ActionRequest) Reset() {
	*x = ActionRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_publish_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ActionRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ActionRequest) ProtoMessage() {}

func (x *ActionRequest) ProtoReflect() protoreflect.Message {
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

// Deprecated: Use ActionRequest.ProtoReflect.Descriptor instead.
func (*ActionRequest) Descriptor() ([]byte, []int) {
	return file_publish_proto_rawDescGZIP(), []int{1}
}

func (x *ActionRequest) GetAHashID() string {
	if x != nil {
		return x.AHashID
	}
	return ""
}

type DelActionRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AHashID string `protobuf:"bytes,1,opt,name=AHashID,proto3" form:"AHashID" json:"aHashID,required" query:"AHashID"` //文章的hashid
}

func (x *DelActionRequest) Reset() {
	*x = DelActionRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_publish_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DelActionRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DelActionRequest) ProtoMessage() {}

func (x *DelActionRequest) ProtoReflect() protoreflect.Message {
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

// Deprecated: Use DelActionRequest.ProtoReflect.Descriptor instead.
func (*DelActionRequest) Descriptor() ([]byte, []int) {
	return file_publish_proto_rawDescGZIP(), []int{2}
}

func (x *DelActionRequest) GetAHashID() string {
	if x != nil {
		return x.AHashID
	}
	return ""
}

// 文章的创建
type CreateActionRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Payload *common.ArticleBase `protobuf:"bytes,1,opt,name=Payload,proto3" form:"Payload" json:"payload,required" query:"Payload"`
}

func (x *CreateActionRequest) Reset() {
	*x = CreateActionRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_publish_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateActionRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateActionRequest) ProtoMessage() {}

func (x *CreateActionRequest) ProtoReflect() protoreflect.Message {
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

// Deprecated: Use CreateActionRequest.ProtoReflect.Descriptor instead.
func (*CreateActionRequest) Descriptor() ([]byte, []int) {
	return file_publish_proto_rawDescGZIP(), []int{3}
}

func (x *CreateActionRequest) GetPayload() *common.ArticleBase {
	if x != nil {
		return x.Payload
	}
	return nil
}

// 文章的修改
type ModifyActionRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AHashID string              `protobuf:"bytes,1,opt,name=AHashID,proto3" form:"AHashID" json:"aHashID,required" query:"AHashID"` //文章的hashid
	Payload *common.ArticleBase `protobuf:"bytes,2,opt,name=Payload,proto3" form:"Payload" json:"payload,required" query:"Payload"`
}

func (x *ModifyActionRequest) Reset() {
	*x = ModifyActionRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_publish_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ModifyActionRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ModifyActionRequest) ProtoMessage() {}

func (x *ModifyActionRequest) ProtoReflect() protoreflect.Message {
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

// Deprecated: Use ModifyActionRequest.ProtoReflect.Descriptor instead.
func (*ModifyActionRequest) Descriptor() ([]byte, []int) {
	return file_publish_proto_rawDescGZIP(), []int{4}
}

func (x *ModifyActionRequest) GetAHashID() string {
	if x != nil {
		return x.AHashID
	}
	return ""
}

func (x *ModifyActionRequest) GetPayload() *common.ArticleBase {
	if x != nil {
		return x.Payload
	}
	return nil
}

type ActionResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	StatusCode int32  `protobuf:"varint,1,opt,name=status_code,json=statusCode,proto3" form:"status_code" json:"status_code" query:"status_code"`
	StatusMsg  string `protobuf:"bytes,2,opt,name=status_msg,json=statusMsg,proto3" form:"status_msg" json:"status_msg" query:"status_msg"`
}

func (x *ActionResponse) Reset() {
	*x = ActionResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_publish_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ActionResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ActionResponse) ProtoMessage() {}

func (x *ActionResponse) ProtoReflect() protoreflect.Message {
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

// Deprecated: Use ActionResponse.ProtoReflect.Descriptor instead.
func (*ActionResponse) Descriptor() ([]byte, []int) {
	return file_publish_proto_rawDescGZIP(), []int{5}
}

func (x *ActionResponse) GetStatusCode() int32 {
	if x != nil {
		return x.StatusCode
	}
	return 0
}

func (x *ActionResponse) GetStatusMsg() string {
	if x != nil {
		return x.StatusMsg
	}
	return ""
}

type CreateActionResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	StatusCode int32  `protobuf:"varint,1,opt,name=status_code,json=statusCode,proto3" form:"status_code" json:"status_code" query:"status_code"`
	StatusMsg  string `protobuf:"bytes,2,opt,name=status_msg,json=statusMsg,proto3" form:"status_msg" json:"status_msg" query:"status_msg"`
	AHashId    string `protobuf:"bytes,3,opt,name=AHashId,proto3" form:"AHashId" json:"AHashId" query:"AHashId"`
}

func (x *CreateActionResponse) Reset() {
	*x = CreateActionResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_publish_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateActionResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateActionResponse) ProtoMessage() {}

func (x *CreateActionResponse) ProtoReflect() protoreflect.Message {
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

// Deprecated: Use CreateActionResponse.ProtoReflect.Descriptor instead.
func (*CreateActionResponse) Descriptor() ([]byte, []int) {
	return file_publish_proto_rawDescGZIP(), []int{6}
}

func (x *CreateActionResponse) GetStatusCode() int32 {
	if x != nil {
		return x.StatusCode
	}
	return 0
}

func (x *CreateActionResponse) GetStatusMsg() string {
	if x != nil {
		return x.StatusMsg
	}
	return ""
}

func (x *CreateActionResponse) GetAHashId() string {
	if x != nil {
		return x.AHashId
	}
	return ""
}

// 文章卡片列表
type CardsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UHashID string `protobuf:"bytes,1,opt,name=UHashID,proto3" form:"UHashID" json:"UHashID" query:"uHashId,required"` //如果id为0，则根据时间的先后去请求文章列表
	Offset  int32  `protobuf:"varint,2,opt,name=Offset,proto3" form:"Offset" json:"Offset" query:"offset,required"`
}

func (x *CardsRequest) Reset() {
	*x = CardsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_publish_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CardsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CardsRequest) ProtoMessage() {}

func (x *CardsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_publish_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CardsRequest.ProtoReflect.Descriptor instead.
func (*CardsRequest) Descriptor() ([]byte, []int) {
	return file_publish_proto_rawDescGZIP(), []int{7}
}

func (x *CardsRequest) GetUHashID() string {
	if x != nil {
		return x.UHashID
	}
	return ""
}

func (x *CardsRequest) GetOffset() int32 {
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
	StatusMsg  string          `protobuf:"bytes,2,opt,name=status_msg,json=statusMsg,proto3" form:"status_msg" json:"status_msg" query:"status_msg"`
	Article    *common.Article `protobuf:"bytes,3,opt,name=article,proto3" form:"article" json:"article" query:"article"`
}

func (x *ArticleResponse) Reset() {
	*x = ArticleResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_publish_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ArticleResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ArticleResponse) ProtoMessage() {}

func (x *ArticleResponse) ProtoReflect() protoreflect.Message {
	mi := &file_publish_proto_msgTypes[8]
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
	return file_publish_proto_rawDescGZIP(), []int{8}
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

type CardsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	StatusCode int32                 `protobuf:"varint,1,opt,name=status_code,json=statusCode,proto3" form:"status_code" json:"status_code" query:"status_code"`
	StatusMsg  string                `protobuf:"bytes,2,opt,name=status_msg,json=statusMsg,proto3" form:"status_msg" json:"status_msg" query:"status_msg"`
	List       []*common.ArticleCard `protobuf:"bytes,3,rep,name=List,proto3" form:"List" json:"List" query:"List"`
}

func (x *CardsResponse) Reset() {
	*x = CardsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_publish_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CardsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CardsResponse) ProtoMessage() {}

func (x *CardsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_publish_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CardsResponse.ProtoReflect.Descriptor instead.
func (*CardsResponse) Descriptor() ([]byte, []int) {
	return file_publish_proto_rawDescGZIP(), []int{9}
}

func (x *CardsResponse) GetStatusCode() int32 {
	if x != nil {
		return x.StatusCode
	}
	return 0
}

func (x *CardsResponse) GetStatusMsg() string {
	if x != nil {
		return x.StatusMsg
	}
	return ""
}

func (x *CardsResponse) GetList() []*common.ArticleCard {
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
	0x6f, 0x22, 0x7f, 0x0a, 0x0d, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x36, 0x0a, 0x07, 0x41, 0x48, 0x61, 0x73, 0x68, 0x49, 0x44, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x42, 0x1c, 0xca, 0xf3, 0x18, 0x18, 0x71, 0x75, 0x65, 0x72, 0x79, 0x3a, 0x22,
	0x61, 0x48, 0x61, 0x73, 0x68, 0x49, 0x44, 0x2c, 0x72, 0x65, 0x71, 0x75, 0x69, 0x72, 0x65, 0x64,
	0x22, 0x52, 0x07, 0x41, 0x48, 0x61, 0x73, 0x68, 0x49, 0x44, 0x12, 0x36, 0x0a, 0x07, 0x55, 0x48,
	0x61, 0x73, 0x68, 0x49, 0x44, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x42, 0x1c, 0xca, 0xf3, 0x18,
	0x18, 0x71, 0x75, 0x65, 0x72, 0x79, 0x3a, 0x22, 0x75, 0x48, 0x61, 0x73, 0x68, 0x49, 0x44, 0x2c,
	0x72, 0x65, 0x71, 0x75, 0x69, 0x72, 0x65, 0x64, 0x22, 0x52, 0x07, 0x55, 0x48, 0x61, 0x73, 0x68,
	0x49, 0x44, 0x22, 0x46, 0x0a, 0x0d, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x35, 0x0a, 0x07, 0x41, 0x48, 0x61, 0x73, 0x68, 0x49, 0x44, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x42, 0x1b, 0xca, 0xf3, 0x18, 0x17, 0x6a, 0x73, 0x6f, 0x6e, 0x3a, 0x22,
	0x61, 0x48, 0x61, 0x73, 0x68, 0x49, 0x44, 0x2c, 0x72, 0x65, 0x71, 0x75, 0x69, 0x72, 0x65, 0x64,
	0x22, 0x52, 0x07, 0x41, 0x48, 0x61, 0x73, 0x68, 0x49, 0x44, 0x22, 0x49, 0x0a, 0x10, 0x44, 0x65,
	0x6c, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x35,
	0x0a, 0x07, 0x41, 0x48, 0x61, 0x73, 0x68, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42,
	0x1b, 0xca, 0xf3, 0x18, 0x17, 0x6a, 0x73, 0x6f, 0x6e, 0x3a, 0x22, 0x61, 0x48, 0x61, 0x73, 0x68,
	0x49, 0x44, 0x2c, 0x72, 0x65, 0x71, 0x75, 0x69, 0x72, 0x65, 0x64, 0x22, 0x52, 0x07, 0x41, 0x48,
	0x61, 0x73, 0x68, 0x49, 0x44, 0x22, 0x5a, 0x0a, 0x13, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x41,
	0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x43, 0x0a, 0x07,
	0x50, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0c, 0x2e,
	0x41, 0x72, 0x74, 0x69, 0x63, 0x6c, 0x65, 0x42, 0x61, 0x73, 0x65, 0x42, 0x1b, 0xca, 0xf3, 0x18,
	0x17, 0x6a, 0x73, 0x6f, 0x6e, 0x3a, 0x22, 0x70, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x2c, 0x72,
	0x65, 0x71, 0x75, 0x69, 0x72, 0x65, 0x64, 0x22, 0x52, 0x07, 0x50, 0x61, 0x79, 0x6c, 0x6f, 0x61,
	0x64, 0x22, 0x91, 0x01, 0x0a, 0x13, 0x4d, 0x6f, 0x64, 0x69, 0x66, 0x79, 0x41, 0x63, 0x74, 0x69,
	0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x35, 0x0a, 0x07, 0x41, 0x48, 0x61,
	0x73, 0x68, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x1b, 0xca, 0xf3, 0x18, 0x17,
	0x6a, 0x73, 0x6f, 0x6e, 0x3a, 0x22, 0x61, 0x48, 0x61, 0x73, 0x68, 0x49, 0x44, 0x2c, 0x72, 0x65,
	0x71, 0x75, 0x69, 0x72, 0x65, 0x64, 0x22, 0x52, 0x07, 0x41, 0x48, 0x61, 0x73, 0x68, 0x49, 0x44,
	0x12, 0x43, 0x0a, 0x07, 0x50, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x0c, 0x2e, 0x41, 0x72, 0x74, 0x69, 0x63, 0x6c, 0x65, 0x42, 0x61, 0x73, 0x65, 0x42,
	0x1b, 0xca, 0xf3, 0x18, 0x17, 0x6a, 0x73, 0x6f, 0x6e, 0x3a, 0x22, 0x70, 0x61, 0x79, 0x6c, 0x6f,
	0x61, 0x64, 0x2c, 0x72, 0x65, 0x71, 0x75, 0x69, 0x72, 0x65, 0x64, 0x22, 0x52, 0x07, 0x50, 0x61,
	0x79, 0x6c, 0x6f, 0x61, 0x64, 0x22, 0x68, 0x0a, 0x0e, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x37, 0x0a, 0x0b, 0x73, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x5f, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x42, 0x16, 0xca, 0xf3,
	0x18, 0x12, 0x6a, 0x73, 0x6f, 0x6e, 0x3a, 0x22, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x5f, 0x63,
	0x6f, 0x64, 0x65, 0x22, 0x52, 0x0a, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x43, 0x6f, 0x64, 0x65,
	0x12, 0x1d, 0x0a, 0x0a, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x5f, 0x6d, 0x73, 0x67, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x4d, 0x73, 0x67, 0x22,
	0x88, 0x01, 0x0a, 0x14, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x37, 0x0a, 0x0b, 0x73, 0x74, 0x61, 0x74,
	0x75, 0x73, 0x5f, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x42, 0x16, 0xca,
	0xf3, 0x18, 0x12, 0x6a, 0x73, 0x6f, 0x6e, 0x3a, 0x22, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x5f,
	0x63, 0x6f, 0x64, 0x65, 0x22, 0x52, 0x0a, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x43, 0x6f, 0x64,
	0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x5f, 0x6d, 0x73, 0x67, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x4d, 0x73, 0x67,
	0x12, 0x18, 0x0a, 0x07, 0x41, 0x48, 0x61, 0x73, 0x68, 0x49, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x07, 0x41, 0x48, 0x61, 0x73, 0x68, 0x49, 0x64, 0x22, 0x7b, 0x0a, 0x0c, 0x43, 0x61,
	0x72, 0x64, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x36, 0x0a, 0x07, 0x55, 0x48,
	0x61, 0x73, 0x68, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x1c, 0xca, 0xf3, 0x18,
	0x18, 0x71, 0x75, 0x65, 0x72, 0x79, 0x3a, 0x22, 0x75, 0x48, 0x61, 0x73, 0x68, 0x49, 0x64, 0x2c,
	0x72, 0x65, 0x71, 0x75, 0x69, 0x72, 0x65, 0x64, 0x22, 0x52, 0x07, 0x55, 0x48, 0x61, 0x73, 0x68,
	0x49, 0x44, 0x12, 0x33, 0x0a, 0x06, 0x4f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x05, 0x42, 0x1b, 0xca, 0xf3, 0x18, 0x17, 0x71, 0x75, 0x65, 0x72, 0x79, 0x3a, 0x22, 0x6f,
	0x66, 0x66, 0x73, 0x65, 0x74, 0x2c, 0x72, 0x65, 0x71, 0x75, 0x69, 0x72, 0x65, 0x64, 0x22, 0x52,
	0x06, 0x4f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x22, 0x8d, 0x01, 0x0a, 0x0f, 0x41, 0x72, 0x74, 0x69,
	0x63, 0x6c, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x37, 0x0a, 0x0b, 0x73,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x5f, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05,
	0x42, 0x16, 0xca, 0xf3, 0x18, 0x12, 0x6a, 0x73, 0x6f, 0x6e, 0x3a, 0x22, 0x73, 0x74, 0x61, 0x74,
	0x75, 0x73, 0x5f, 0x63, 0x6f, 0x64, 0x65, 0x22, 0x52, 0x0a, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73,
	0x43, 0x6f, 0x64, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x5f, 0x6d,
	0x73, 0x67, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73,
	0x4d, 0x73, 0x67, 0x12, 0x22, 0x0a, 0x07, 0x61, 0x72, 0x74, 0x69, 0x63, 0x6c, 0x65, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x08, 0x2e, 0x41, 0x72, 0x74, 0x69, 0x63, 0x6c, 0x65, 0x52, 0x07,
	0x61, 0x72, 0x74, 0x69, 0x63, 0x6c, 0x65, 0x22, 0x89, 0x01, 0x0a, 0x0d, 0x43, 0x61, 0x72, 0x64,
	0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x37, 0x0a, 0x0b, 0x73, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x5f, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x42, 0x16,
	0xca, 0xf3, 0x18, 0x12, 0x6a, 0x73, 0x6f, 0x6e, 0x3a, 0x22, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73,
	0x5f, 0x63, 0x6f, 0x64, 0x65, 0x22, 0x52, 0x0a, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x43, 0x6f,
	0x64, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x5f, 0x6d, 0x73, 0x67,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x4d, 0x73,
	0x67, 0x12, 0x20, 0x0a, 0x04, 0x4c, 0x69, 0x73, 0x74, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x0c, 0x2e, 0x41, 0x72, 0x74, 0x69, 0x63, 0x6c, 0x65, 0x43, 0x61, 0x72, 0x64, 0x52, 0x04, 0x4c,
	0x69, 0x73, 0x74, 0x32, 0x83, 0x05, 0x0a, 0x0e, 0x50, 0x75, 0x62, 0x6c, 0x69, 0x73, 0x68, 0x48,
	0x61, 0x6e, 0x64, 0x6c, 0x65, 0x72, 0x12, 0x6d, 0x0a, 0x0d, 0x50, 0x75, 0x62, 0x6c, 0x69, 0x73,
	0x68, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x1c, 0x2e, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x73,
	0x68, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1d, 0x2e, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x73, 0x68, 0x2e,
	0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x22, 0x1f, 0xd2, 0xc1, 0x18, 0x1b, 0x2f, 0x62, 0x6c, 0x6f, 0x67, 0x5f,
	0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2f, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x73, 0x68, 0x2f, 0x61,
	0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x6d, 0x0a, 0x13, 0x50, 0x75, 0x62, 0x6c, 0x69, 0x73, 0x68,
	0x4d, 0x6f, 0x64, 0x69, 0x66, 0x79, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x1c, 0x2e, 0x70,
	0x75, 0x62, 0x6c, 0x69, 0x73, 0x68, 0x2e, 0x4d, 0x6f, 0x64, 0x69, 0x66, 0x79, 0x41, 0x63, 0x74,
	0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x17, 0x2e, 0x70, 0x75, 0x62,
	0x6c, 0x69, 0x73, 0x68, 0x2e, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x22, 0x1f, 0xea, 0xc1, 0x18, 0x1b, 0x2f, 0x62, 0x6c, 0x6f, 0x67, 0x5f, 0x73,
	0x65, 0x72, 0x76, 0x65, 0x72, 0x2f, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x73, 0x68, 0x2f, 0x61, 0x63,
	0x74, 0x69, 0x6f, 0x6e, 0x12, 0x67, 0x0a, 0x10, 0x50, 0x75, 0x62, 0x6c, 0x69, 0x73, 0x68, 0x44,
	0x65, 0x6c, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x19, 0x2e, 0x70, 0x75, 0x62, 0x6c, 0x69,
	0x73, 0x68, 0x2e, 0x44, 0x65, 0x6c, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x17, 0x2e, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x73, 0x68, 0x2e, 0x41, 0x63,
	0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x1f, 0xe2, 0xc1,
	0x18, 0x1b, 0x2f, 0x62, 0x6c, 0x6f, 0x67, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2f, 0x70,
	0x75, 0x62, 0x6c, 0x69, 0x73, 0x68, 0x2f, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x5b, 0x0a,
	0x0b, 0x50, 0x75, 0x62, 0x6c, 0x69, 0x73, 0x68, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x15, 0x2e, 0x70,
	0x75, 0x62, 0x6c, 0x69, 0x73, 0x68, 0x2e, 0x43, 0x61, 0x72, 0x64, 0x73, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x73, 0x68, 0x2e, 0x43, 0x61,
	0x72, 0x64, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x1d, 0xca, 0xc1, 0x18,
	0x19, 0x2f, 0x62, 0x6c, 0x6f, 0x67, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2f, 0x70, 0x75,
	0x62, 0x6c, 0x69, 0x73, 0x68, 0x2f, 0x6c, 0x69, 0x73, 0x74, 0x12, 0x62, 0x0a, 0x0d, 0x50, 0x75,
	0x62, 0x6c, 0x69, 0x73, 0x68, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x12, 0x16, 0x2e, 0x70, 0x75,
	0x62, 0x6c, 0x69, 0x73, 0x68, 0x2e, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x18, 0x2e, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x73, 0x68, 0x2e, 0x41, 0x72,
	0x74, 0x69, 0x63, 0x6c, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x1f, 0xca,
	0xc1, 0x18, 0x1b, 0x2f, 0x62, 0x6c, 0x6f, 0x67, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2f,
	0x70, 0x75, 0x62, 0x6c, 0x69, 0x73, 0x68, 0x2f, 0x64, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x12, 0x69,
	0x0a, 0x13, 0x50, 0x75, 0x62, 0x6c, 0x69, 0x73, 0x68, 0x56, 0x69, 0x65, 0x77, 0x43, 0x6f, 0x75,
	0x6e, 0x74, 0x41, 0x64, 0x64, 0x12, 0x16, 0x2e, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x73, 0x68, 0x2e,
	0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x17, 0x2e,
	0x70, 0x75, 0x62, 0x6c, 0x69, 0x73, 0x68, 0x2e, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x21, 0xd2, 0xc1, 0x18, 0x1d, 0x2f, 0x62, 0x6c, 0x6f,
	0x67, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2f, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x73, 0x68,
	0x2f, 0x76, 0x69, 0x65, 0x77, 0x5f, 0x61, 0x64, 0x64, 0x42, 0x3f, 0x5a, 0x3d, 0x67, 0x69, 0x74,
	0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x71, 0x69, 0x6e, 0x67, 0x79, 0x67, 0x67, 0x67,
	0x2f, 0x62, 0x6c, 0x6f, 0x67, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2f, 0x62, 0x69, 0x7a,
	0x2f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2f, 0x68, 0x65, 0x72, 0x74, 0x7a, 0x2f, 0x62, 0x61, 0x73,
	0x69, 0x63, 0x2f, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x73, 0x68, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
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

var file_publish_proto_msgTypes = make([]protoimpl.MessageInfo, 10)
var file_publish_proto_goTypes = []interface{}{
	(*DetailRequest)(nil),        // 0: publish.DetailRequest
	(*ActionRequest)(nil),        // 1: publish.ActionRequest
	(*DelActionRequest)(nil),     // 2: publish.DelActionRequest
	(*CreateActionRequest)(nil),  // 3: publish.CreateActionRequest
	(*ModifyActionRequest)(nil),  // 4: publish.ModifyActionRequest
	(*ActionResponse)(nil),       // 5: publish.ActionResponse
	(*CreateActionResponse)(nil), // 6: publish.CreateActionResponse
	(*CardsRequest)(nil),         // 7: publish.CardsRequest
	(*ArticleResponse)(nil),      // 8: publish.ArticleResponse
	(*CardsResponse)(nil),        // 9: publish.CardsResponse
	(*common.ArticleBase)(nil),   // 10: ArticleBase
	(*common.Article)(nil),       // 11: Article
	(*common.ArticleCard)(nil),   // 12: ArticleCard
}
var file_publish_proto_depIdxs = []int32{
	10, // 0: publish.CreateActionRequest.Payload:type_name -> ArticleBase
	10, // 1: publish.ModifyActionRequest.Payload:type_name -> ArticleBase
	11, // 2: publish.ArticleResponse.article:type_name -> Article
	12, // 3: publish.CardsResponse.List:type_name -> ArticleCard
	3,  // 4: publish.PublishHandler.PublishAction:input_type -> publish.CreateActionRequest
	4,  // 5: publish.PublishHandler.PublishModifyAction:input_type -> publish.ModifyActionRequest
	2,  // 6: publish.PublishHandler.PublishDelAction:input_type -> publish.DelActionRequest
	7,  // 7: publish.PublishHandler.PublishList:input_type -> publish.CardsRequest
	0,  // 8: publish.PublishHandler.PublishDetail:input_type -> publish.DetailRequest
	1,  // 9: publish.PublishHandler.PublishViewCountAdd:input_type -> publish.ActionRequest
	6,  // 10: publish.PublishHandler.PublishAction:output_type -> publish.CreateActionResponse
	5,  // 11: publish.PublishHandler.PublishModifyAction:output_type -> publish.ActionResponse
	5,  // 12: publish.PublishHandler.PublishDelAction:output_type -> publish.ActionResponse
	9,  // 13: publish.PublishHandler.PublishList:output_type -> publish.CardsResponse
	8,  // 14: publish.PublishHandler.PublishDetail:output_type -> publish.ArticleResponse
	5,  // 15: publish.PublishHandler.PublishViewCountAdd:output_type -> publish.ActionResponse
	10, // [10:16] is the sub-list for method output_type
	4,  // [4:10] is the sub-list for method input_type
	4,  // [4:4] is the sub-list for extension type_name
	4,  // [4:4] is the sub-list for extension extendee
	0,  // [0:4] is the sub-list for field type_name
}

func init() { file_publish_proto_init() }
func file_publish_proto_init() {
	if File_publish_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_publish_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DetailRequest); i {
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
			switch v := v.(*ActionRequest); i {
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
			switch v := v.(*DelActionRequest); i {
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
			switch v := v.(*CreateActionRequest); i {
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
			switch v := v.(*ModifyActionRequest); i {
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
			switch v := v.(*ActionResponse); i {
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
			switch v := v.(*CreateActionResponse); i {
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
		file_publish_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CardsRequest); i {
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
		file_publish_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
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
		file_publish_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CardsResponse); i {
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
			NumMessages:   10,
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
