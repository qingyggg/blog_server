package main

import (
	"context"
	"fmt"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/app/server"
	"github.com/cloudwego/hertz/pkg/app/server/binding"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/hertz-contrib/cors"
	"github.com/hertz-contrib/gzip"
	"github.com/hertz-contrib/pprof"
	"github.com/hertz-contrib/reverseproxy"
	"github.com/hertz-contrib/swagger"
	"github.com/qingyggg/blog_server/biz/dal"
	"github.com/qingyggg/blog_server/biz/mw/jwt"
	"github.com/qingyggg/blog_server/biz/mw/logger"
	"github.com/qingyggg/blog_server/biz/mw/minio"
	_ "github.com/qingyggg/blog_server/docs"
	"github.com/qingyggg/blog_server/pkg/constants"
	"github.com/qingyggg/blog_server/pkg/utils"
	swaggerFiles "github.com/swaggo/files"
	"time"
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
	h.Use(gzip.Gzip(gzip.DefaultCompression)) //gzip压缩
	// default is "debug/pprof"
	pprof.Register(h, "dev/pprof")
	//cors config
	h.Use(cors.New(cors.Config{
		//AllowWildcard: 	  true,
		//AllowOrigins:     []string{"*"},
		AllowAllOrigins:  true,
		AllowMethods:     []string{"PUT", "PATCH", "GET", "POST", "DELETE"},
		AllowHeaders:     []string{"Origin", "Content-Type,AccessToken,X-CSRF-Token, Authorization, Token, x-token"},
		ExposeHeaders:    []string{"Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers, Content-Type"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))
	//oss
	h.GET("/src/*name", minioReverseProxy)
	//swagger
	url := swagger.URL("http://localhost:18005/swagger/doc.json") // The url pointing to API definition
	h.GET("/swagger/*any", swagger.WrapHandler(swaggerFiles.Handler, url))

	register(h)
	h.Spin()
}

// Set up /src/*name route forwarding to access minio from external network
func minioReverseProxy(c context.Context, ctx *app.RequestContext) {
	proxy, _ := reverseproxy.NewSingleHostReverseProxy("http://localhost:18001")
	ctx.URI().SetPath(ctx.Param("name"))
	hlog.CtxInfof(c, string(ctx.Request.URI().Path()))
	proxy.ServeHTTP(c, ctx)
}

func init() {
	utils.EnvInit()
	constants.UrlInit()
	logger.InitLogger()
	dal.Init()
	jwt.Init()
	minio.Init()
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
