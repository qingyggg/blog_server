package handler

import (
	"context"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/protocol"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
	"github.com/qingyggg/blog_server/biz/model/hertz/common"
	"github.com/qingyggg/blog_server/pkg/errno"
)

func Logout(ctx context.Context, c *app.RequestContext) {
	c.SetCookie("token", "", -1, "/", "", protocol.CookieSameSiteDefaultMode, false, true)
	c.JSON(consts.StatusOK, common.BaseResponse{
		StatusCode: errno.SuccessCode,
		StatusMsg:  errno.SuccessMsg,
	})
}
