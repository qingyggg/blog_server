package minio

import (
	"bytes"
	"context"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
	"github.com/qingyggg/blog_server/pkg/constants"
	"mime/multipart"
	"net/url"
	"time"
)

var (
	Client *minio.Client
	err    error
)

// MakeBucket create a bucket with a specified name
func MakeBucket(ctx context.Context, bucketName string) error {
	exists, err := Client.BucketExists(ctx, bucketName)
	if err != nil {
		return err
	}
	if !exists {
		err = Client.MakeBucket(ctx, bucketName, minio.MakeBucketOptions{})
		if err != nil {
			return err
		}
		hlog.Infof("Successfully created mybucket %v\n", bucketName)
	}
	return nil
}

// PutToBucket put the file into the bucket by *multipart.FileHeader
func PutToBucket(ctx context.Context, bucketName string, file *multipart.FileHeader) (info minio.UploadInfo, err error) {
	fileObj, _ := file.Open()
	info, err = Client.PutObject(ctx, bucketName, file.Filename, fileObj, file.Size, minio.PutObjectOptions{})
	fileObj.Close()
	return info, err
}

// GetObjURL get the original link of the file in minio
func GetObjURL(ctx context.Context, bucketName, filename string) (u *url.URL, err error) {
	exp := time.Hour * 24
	reqParams := make(url.Values)
	u, err = Client.PresignedGetObject(ctx, bucketName, filename, exp, reqParams)
	return u, err
}

// PutToBucketByBuf put the file into the bucket by *bytes.Buffer
func PutToBucketByBuf(ctx context.Context, bucketName, filename string, buf *bytes.Buffer) (info minio.UploadInfo, err error) {
	info, err = Client.PutObject(ctx, bucketName, filename, buf, int64(buf.Len()), minio.PutObjectOptions{})
	return info, err
}

// PutToBucketByFilePath put the file into the bucket by filepath
func PutToBucketByFilePath(ctx context.Context, bucketName, filename, filepath string) (info minio.UploadInfo, err error) {
	info, err = Client.FPutObject(ctx, bucketName, filename, filepath, minio.PutObjectOptions{})
	return info, err
}

// DelObject delete file in bucket
func DelObject(ctx context.Context, bucketName, fileName string) error {

	// 执行删除操作
	err = Client.RemoveObject(ctx, bucketName, fileName, minio.RemoveObjectOptions{})
	return err
}

func Init() {
	ctx := context.Background()

	// 初始化 MinIO 客户端
	var err error
	Client, err = minio.New(constants.MinioEndPoint, &minio.Options{
		Creds:  credentials.NewStaticV4(constants.MinioAccessKeyID, constants.MinioSecretAccessKey, ""),
		Secure: constants.MinioSSL,
	})
	if err != nil {
		hlog.Fatal("连接 MinIO 失败: ", err)
	}
	hlog.Info("成功连接到 MinIO")

	// 创建 Bucket
	if err = MakeBucket(ctx, constants.MinioImgBucketName); err != nil {
		hlog.Fatal("创建 Bucket 失败: ", err)
	}
	hlog.Info("成功创建 Bucket")

	// 上传文件
	imgNames := []string{"mols.jpg", "marisa.jpg"}
	imgPaths := []string{"static/mols.jpg", "static/marisa.jpg"}
	for idx, imgPath := range imgPaths {
		if _, err = PutToBucketByFilePath(ctx, constants.MinioImgBucketName, imgNames[idx], imgPaths[idx]); err != nil {
			hlog.Fatal("上传文件失败: ", err)
		}
		hlog.Infof("成功上传文件: %s", imgPath)
	}
}
