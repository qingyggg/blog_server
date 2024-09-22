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
	arr := strings.Split(path, "/")
	u, err := minio.GetObjURL(ctx, arr[0], arr[1])
	if err != nil {
		hlog.CtxInfof(ctx, err.Error())
		return ""
	}
	u.Scheme = "https"
	u.Host = string(c.URI().Host())
	u.Path = "/src" + u.Path
	return u.String()
}

// UrlConvertReverse 从完整URL还原数据库中存储的相对路径
func UrlConvertReverse(ctx context.Context, fullURL string) (path string) {
	if len(fullURL) == 0 {
		return ""
	}

	// 解析传入的 URL
	u, err := url.Parse(fullURL)
	if err != nil {
		hlog.CtxInfof(ctx, "解析URL失败: %s", err.Error())
		return ""
	}

	// 假设路径的前缀是 "/src"，需要去掉前缀部分
	urlPath := u.Path
	if strings.HasPrefix(urlPath, "/src/") {
		urlPath = strings.TrimPrefix(urlPath, "/src/")
	} else {
		hlog.CtxInfof(ctx, "URL路径无效: %s", fullURL)
		return ""
	}

	// 将去掉前缀的路径拆分为 bucket 和 object
	arr := strings.Split(urlPath, "/")
	if len(arr) < 2 {
		hlog.CtxInfof(ctx, "URL格式无效: %s", fullURL)
		return ""
	}

	// 拼接数据库存储的相对路径 (bucket/object)
	path = fmt.Sprintf("%s/%s", arr[0], arr[1])
	return path
}
