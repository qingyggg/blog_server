package main

import (
	"context"
	"fmt"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/app/server"
	"github.com/cloudwego/hertz/pkg/app/server/binding"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/hertz-contrib/pprof"
	"github.com/hertz-contrib/reverseproxy"
	"github.com/qingyggg/blog_server/biz/dal"
	"github.com/qingyggg/blog_server/biz/mw/jwt"
	"github.com/qingyggg/blog_server/biz/mw/minio"
	_ "github.com/qingyggg/blog_server/docs"
	"github.com/qingyggg/blog_server/pkg/constants"
	"github.com/qingyggg/blog_server/pkg/utils"
)

//	@title			blog_server tests
//	@version		1.0
//	@description	This is a demo using Hertz.

//	@contact.name	hertz-contrib
//	@contact.url	https://github.com/hertz-contrib

//	@license.name	Apache 2.0
//	@license.url	http://www.apache.org/licenses/LICENSE-2.0.html

// @host		localhost:18005
// @BasePath	/
// @schemes	http
func main() {
	validateConfig := GetCustomValidateConfig()
	h := server.Default(
		server.WithStreamBody(true),
		server.WithHostPorts("0.0.0.0:18005"),
		server.WithValidateConfig(validateConfig),
	)
	// default is "debug/pprof"
	pprof.Register(h, "dev/pprof")

	register(h)
	h.Spin()
}

// Set up /src/*name route forwarding to access minio from external network
func minioReverseProxy(c context.Context, ctx *app.RequestContext) {
	proxyUrl := "http://" + constants.MinioEndPoint
	proxy, _ := reverseproxy.NewSingleHostReverseProxy(proxyUrl)
	ctx.URI().SetPath(ctx.Param("name"))
	hlog.CtxInfof(c, "minio图片访问==>"+string(ctx.Request.URI().Path()))
	proxy.ServeHTTP(c, ctx)
}

func init() {
	utils.EnvInit()
	constants.UrlInit()
	//logger.InitLogger()
	dal.Init() //数据库初始化
	jwt.Init()
	minio.Init() //存储服务初始化
	//mq.Init()     //消息件初始化
	//socket.Init() //初始化socket.io
}
func GetCustomValidateConfig() *binding.ValidateConfig {
	//自定义参数校验
	validateConfig := &binding.ValidateConfig{}
	validateConfig.MustRegValidateFunc("password", func(args ...interface{}) error {
		err := utils.ValidatePassword(fmt.Sprint(args...))
		if err != nil {
			return err
		}
		return nil
	})
	return validateConfig
}
