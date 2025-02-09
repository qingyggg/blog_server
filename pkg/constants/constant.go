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
	AmqpDSN              string
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
	AmqpDSN = "amqp://" + os.Getenv("MqUser") + ":" + os.Getenv("MqPwd") + "@" + os.Getenv("MqUrl")
}

// constants in the project
const (
	MinioImgBucketName = "imagebucket"

	DefaultSign       = "该用户没有留下任何签名"
	DefaultAva        = "imagebucket/mols.jpg"
	DefaultBackground = "imagebucket/marisa.jpg"
)
