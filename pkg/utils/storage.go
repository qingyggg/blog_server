package utils

import (
	"context"
	"fmt"
	"net/url"
	"strings"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/hlog"

	"github.com/qingyggg/blog_server/biz/mw/minio"
)

// NewFileName Splicing user_id and time to make unique filename
func NewFileName(user_id, time int64) string {
	return fmt.Sprintf("%d.%d", user_id, time)
}

// URLconvert Convert the path in the database into a complete url accessible by the front end
func URLconvert(ctx context.Context, c *app.RequestContext, path string) (fullURL string) {
	if len(path) == 0 {
		return ""
	}
	bucket, object, ok := splitBucketObject(path)
	if !ok {
		hlog.CtxInfof(ctx, "MinIO path invalid: %s", path)
		return ""
	}
	exists, err := minio.ObjectExists(ctx, bucket, object)
	if err != nil {
		hlog.CtxInfof(ctx, err.Error())
		return ""
	}
	if !exists {
		hlog.CtxInfof(ctx, "MinIO object missing: %s", path)
		return ""
	}
	return "/src/" + bucket + "/" + object
}

// UrlConvertReverse 从完整URL还原数据库中存储的相对路径
func UrlConvertReverse(ctx context.Context, fullURL string) (path string) {
	fullURL = strings.TrimSpace(fullURL)
	if len(fullURL) == 0 {
		return ""
	}

	// 解析传入的 URL
	u, err := url.Parse(fullURL)
	if err != nil {
		hlog.CtxInfof(ctx, "解析URL失败: %s", err.Error())
		return ""
	}

	urlPath := u.Path
	if strings.HasPrefix(urlPath, "/src/") {
		urlPath = strings.TrimPrefix(urlPath, "/src/")
	} else if strings.HasPrefix(fullURL, "/src/") {
		urlPath = strings.TrimPrefix(fullURL, "/src/")
	} else if _, _, ok := splitBucketObject(fullURL); ok && u.Scheme == "" && u.Host == "" {
		return fullURL
	} else {
		hlog.CtxInfof(ctx, "URL路径无效: %s", fullURL)
		return ""
	}

	bucket, object, ok := splitBucketObject(urlPath)
	if !ok {
		hlog.CtxInfof(ctx, "URL格式无效: %s", fullURL)
		return ""
	}
	return fmt.Sprintf("%s/%s", bucket, object)
}

func splitBucketObject(path string) (bucket, object string, ok bool) {
	path = strings.Trim(path, "/")
	arr := strings.SplitN(path, "/", 2)
	if len(arr) != 2 || arr[0] == "" || arr[1] == "" {
		return "", "", false
	}
	return arr[0], arr[1], true
}
