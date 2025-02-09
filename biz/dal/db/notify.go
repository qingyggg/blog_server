package db

import (
	"context"
	mymongo "github.com/qingyggg/blog_server/biz/mw/mongo"
	"go.mongodb.org/mongo-driver/v2/bson"
	"time"
)

func NotifySave(notification *mymongo.Notification) error {
	notification.CreateTime = bson.NewDateTimeFromTime(time.Now())
	_, err := mymongo.NotificationCol.InsertOne(context.TODO(), notification)
	if err != nil {
		return err
	}
	return nil
}

func NotifyExist(nid string) (error, bool) {
	count, err := mymongo.NotificationCol.CountDocuments(context.TODO(), bson.M{
		"hash_id": nid,
	})
	if err != nil {
		return err, false
	}
	if count != 0 {
		return nil, true
	} else {
		return nil, false
	}
}
func NotifyRead(nid string) error {
	filter := bson.M{"hash_id": nid}
	update := bson.M{
		"$set": bson.M{
			"is_read": true,
		},
	}
	_, err := mymongo.NotificationCol.UpdateOne(context.TODO(), filter, update)
	if err != nil {
		return err
	}
	return nil
}

func NotifyReadByUid(uid string) error {
	filter := bson.M{"user_id": uid}
	update := bson.M{
		"$set": bson.M{
			"is_read": true,
		},
	}
	_, err := mymongo.NotificationCol.UpdateMany(context.TODO(), filter, update)
	if err != nil {
		return err
	}
	return nil
}

func NotifyDelById(nid string) error {
	_, err := mymongo.NotificationCol.DeleteOne(context.TODO(), bson.M{"hash_id": nid})
	if err != nil {
		return err
	}
	return nil
}

func NotifyDelByUid(uid string) error {
	_, err := mymongo.NotificationCol.DeleteMany(context.TODO(), bson.M{"user_id": uid})
	if err != nil {
		return err
	}
	return nil
}

func NotifyGetByUid(uid string) (error, []*mymongo.Notification) {
	cursor, err := mymongo.NotificationCol.Find(context.TODO(), bson.M{"user_id": uid})
	if err != nil {
		return err, nil
	}
	defer cursor.Close(context.TODO())
	var notifications []*mymongo.Notification
	for cursor.Next(context.TODO()) {
		notify := new(mymongo.Notification)
		if err := cursor.Decode(notify); err != nil {
			return err, nil
		}
		notifications = append(notifications, notify)
	}
	return nil, notifications
}
func NotifyTake(nid string) (error, *mymongo.Notification) {
	notify := new(mymongo.Notification)
	err := mymongo.NotificationCol.FindOne(context.TODO(), bson.M{"hash_id": nid}).Decode(notify)
	if err != nil {
		return err, nil
	}
	return nil, notify
}
