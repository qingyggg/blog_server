package socket

import (
	"github.com/qingyggg/blog_server/biz/model/hertz/common"
	"github.com/qingyggg/blog_server/biz/mw/mongo"
)

type Response struct {
	StatusCode int32       `json:"status_code"`    // JSON 中的字段名为 "msg_code"
	StatusMsg  string      `json:"status_msg"`     // JSON 中的字段名为 "msg_msg"
	Data       interface{} `json:"data,omitempty"` // 如果 Data 是空值，序列化时会忽略
}

// System 系统通知
type System struct {
	NID       string `json:"nid"`
	Message   string `json:"message"`
	EventType int32  `json:"event_type"`
}

// AFavorite 文章点赞
type AFavorite struct {
	NID       string           `json:"nid"`
	User      *common.UserBase `json:"user"` // 用户信息
	Aid       string           `json:"aid"`
	ATitle    string           `json:"a_title"` // 文章基础信息
	EventType int32            `json:"event_type"`
}

type ACollect struct {
	NID       string           `json:"nid"`
	User      *common.UserBase `json:"user"` // 用户信息
	Aid       string           `json:"aid"`
	ATitle    string           `json:"a_title"` // 文章基础信息
	EventType int32            `json:"event_type"`
}

// CFavorite 评论点赞
type CFavorite struct {
	NID       string           `json:"nid"`
	User      *common.UserBase `json:"user"`    // 用户信息
	Comment   *mongo.Comment   `json:"comment"` // 评论信息
	EventType int32            `json:"event_type"`
}

// Commented 评论或回复
type Commented struct {
	NID       string           `json:"nid"`
	User      *common.UserBase `json:"user"`    // 用户信息
	Comment   *mongo.Comment   `json:"comment"` // 评论或回复内容
	EventType int32            `json:"event_type"`
}

// Follow 关注
type Follow struct {
	NID       string           `json:"nid"`  // 关注 ID
	User      *common.UserBase `json:"user"` // 用户信息
	EventType int32            `json:"event_type"`
}
