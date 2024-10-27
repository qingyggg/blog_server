package utils

import (
	"go.mongodb.org/mongo-driver/v2/bson"
	"time"
)

// MillTimeStampToTime convert ms timestamp to time.Time
func MillTimeStampToTime(timestamp int64) time.Time {
	second := timestamp / 1000
	nano := timestamp % 1000 * 1000000
	return time.Unix(second, nano)
}

// SecondTimeStampToTime convert s timestamp to time.Time
func SecondTimeStampToTime(timestamp int64) time.Time {
	return time.Unix(timestamp, 0)
}

func ConvertBsonTimeToString(dateTime bson.DateTime) string {
	// 假设你有一个 bson.DateTime 对象

	// 将 bson.DateTime 转换为 Go 的 time.Time
	t := dateTime.Time()

	// 将 time.Time 转换为字符串，使用你需要的格式
	timeStr := t.Format("2006-01-02 15:04:05")

	return timeStr
}
