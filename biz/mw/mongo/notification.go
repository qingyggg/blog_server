package mongo

import (
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"log"
)

var (
	NotificationCol *mongo.Collection
)

var NotifyTypes = struct {
	SystemNotification int32
	UserNotification   int32
}{
	1, 2,
}

func InitNotification() {
	NotificationCol = DB.Collection("notification") // 存储用户通知的表

	if err := CreateIndex(NotificationCol, bson.M{"hash_id": 1}, true); err != nil {
		log.Fatal("Failed to create hash_id index for NotificationCol")
	}
	if err := CreateIndex(NotificationCol, bson.M{"user_id": 1}, false); err != nil {
		log.Fatal("Failed to create user_id index for NotificationCol")
	}
}

type Notification struct {
	UserID     string         `bson:"user_id" json:"user_id"`
	HashID     string         `bson:"hash_id" json:"hash_id"`
	Type       int32          `bson:"type" json:"type"`
	Data       *NotifyPayload `bson:"content" json:"content"`
	IsRead     bool           `bson:"is_read" json:"is_read"`
	CreateTime bson.DateTime  `bson:"create_time" json:"create_time"`
}

type NotifyPayload struct {
	EventType string                 `bson:"event_type" json:"event_type"`
	TargetID  string                 `bson:"target_id" json:"target_id"`   //aid(跳转到文章) cid(跳转到评论) uid(跳转到该用户的follower列表)
	ActionUid string                 `bson:"action_uid" json:"action_uid"` //关注你的用户，评论你的用户，点赞你的用户
	Message   string                 `bson:"message" json:"message"`
	ExtraData map[string]interface{} `bson:"extra_data,omitempty" json:"extra_data,omitempty"`
}
