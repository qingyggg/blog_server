package main

import (
	"context"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/app/server"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/hertz-contrib/reverseproxy"
	"github.com/hertz-contrib/swagger"
	"github.com/qingyggg/blog_server/biz/dal"
	"github.com/qingyggg/blog_server/biz/mw/jwt"
	"github.com/qingyggg/blog_server/biz/mw/minio"
	_ "github.com/qingyggg/blog_server/docs"
	swaggerFiles "github.com/swaggo/files"
)

// Set up /src/*name route forwarding to access minio from external network
func minioReverseProxy(c context.Context, ctx *app.RequestContext) {
	proxy, _ := reverseproxy.NewSingleHostReverseProxy("http://localhost:18001")
	ctx.URI().SetPath(ctx.Param("name"))
	hlog.CtxInfof(c, string(ctx.Request.URI().Path()))
	proxy.ServeHTTP(c, ctx)
}

func Init() {
	dal.Init()
	jwt.Init()
	minio.Init()
}

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
	Init()
	h := server.Default(
		server.WithStreamBody(true),
		server.WithHostPorts("0.0.0.0:18005"),
	)
	//oss
	h.GET("/src/*name", minioReverseProxy)
	//swagger
	url := swagger.URL("http://localhost:18005/swagger/doc.json") // The url pointing to API definition
	h.GET("/swagger/*any", swagger.WrapHandler(swaggerFiles.Handler, url))

	register(h)
	h.Spin()
}
