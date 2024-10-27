package constants

import (
	"os"
)

// connection information
var (
	MySQLDefaultDSN      string
	MinioEndPoint        string
	MinioAccessKeyID     string
	MinioSecretAccessKey string
	MinioSSL             bool
	RedisAddr            string
	RedisPassword        string
	MongoDefaultDSN      string
)

// UrlInit 调用该函数前，先要加载env
func UrlInit() {
	MySQLDefaultDSN = os.Getenv("MariaDBUser") + ":" + os.Getenv("MariaDBPwd") + "@tcp(" + os.Getenv("MariaDBUrl") + ")/storybook?charset=utf8mb4&parseTime=True&loc=Local"
	MinioEndPoint = os.Getenv("MinioEndPoint")
	MinioAccessKeyID = os.Getenv("MinioAccessKeyID")
	MinioSecretAccessKey = os.Getenv("MinioSecretAccessKey")
	MinioSSL = false
	RedisAddr = os.Getenv("RedisUrl")
	RedisPassword = os.Getenv("RedisPassword")
	MongoDefaultDSN = "mongodb://" + os.Getenv("MongoUser") + ":" + os.Getenv("MongoPwd") + "@" + os.Getenv("MongoUrl") + "/?connect=direct"
	println("mongo dsn", MongoDefaultDSN)
}

// constants in the project
const (
	MinioVideoBucketName = "videobucket"
	MinioImgBucketName   = "imagebucket"

	TestSign       = "测试账号！ offer"
	TestAva        = "avatar/test1.jpg"
	TestBackground = "background/test1.png"
)
