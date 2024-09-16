// Code generated by hertz generator.

package user

import (
	"context"
	"github.com/qingyggg/blog_server/biz/mw/jwt"
	service "github.com/qingyggg/blog_server/biz/service/user"
	"github.com/qingyggg/blog_server/pkg/errno"
	"github.com/qingyggg/blog_server/pkg/utils"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
	user "github.com/qingyggg/blog_server/biz/model/hertz/basic/user"
)

// User .
//
//	@router	/blog_server/user/ [GET]
func User(ctx context.Context, c *app.RequestContext) {
	var err error
	req := new(user.UserRequest)
	err = c.BindAndValidate(req)
	if err != nil {
		utils.ErrResp(c, err)
		return
	}

	payload, err := service.NewUserService(ctx, c).UserInfo(req)
	c.JSON(consts.StatusOK, user.UserResponse{
		StatusCode: errno.SuccessCode,
		StatusMsg:  errno.SuccessMsg,
		User:       payload,
	})
}

// UserRegister 注册用户接口
//
//	@Summary		用户注册
//	@Description	用户通过提供用户名和密码注册账户
//	@Tags			用户相关接口
//	@Accept			json
//	@Produce		json
//	@Param			user	body		user.UserActionRequest	true	"用户注册请求参数"
//	@Success		200		{object}	user.UserActionResponse	"成功返回用户ID及状态信息"
//	@Failure		400		{object}	user.UserActionResponse	"请求参数错误或其他错误信息"
//	@Router			/blog_server/user/register/ [post]
func UserRegister(ctx context.Context, c *app.RequestContext) {
	var err error
	req := new(user.UserActionRequest) //req:username,password
	err = c.BindAndValidate(req)
	if err != nil {
		utils.ErrResp(c, err)
		return
	}

	_, err = service.NewUserService(ctx, c).UserRegister(req)
	if err != nil {
		utils.ErrResp(c, err)
		return
	}
	jwt.JwtMiddleware.LoginHandler(ctx, c)
	v, _ := c.Get("user_id")
	user_id := v.(int64)
	c.JSON(consts.StatusOK, user.UserActionResponse{
		StatusCode: errno.SuccessCode,
		StatusMsg:  errno.SuccessMsg,
		UserId:     user_id,
	})
}

// UserLogin .
//
//	@router	/blog_server/user/login/ [POST]
func UserLogin(ctx context.Context, c *app.RequestContext) {
	var err error
	req := new(user.UserActionRequest)
	err = c.BindAndValidate(req)
	if err != nil {
		utils.ErrResp(c, err)
		return
	}
	jwt.JwtMiddleware.LoginHandler(ctx, c)
	v, _ := c.Get("user_id")
	user_id := v.(int64)
	c.JSON(consts.StatusOK, user.UserActionResponse{
		StatusCode: errno.SuccessCode,
		StatusMsg:  errno.SuccessMsg,
		UserId:     user_id,
	})
}

// UserPwdModify .
//
//	@router	/blog_server/user/pwd_modify/ [POST]
func UserPwdModify(ctx context.Context, c *app.RequestContext) {
	var err error
	var req = new(user.UserActionPwdModifyRequest)
	err = c.BindAndValidate(req)
	if err != nil {
		utils.ErrResp(c, err)
		return
	}
	uid, err := service.NewUserService(ctx, c).PwdModify(req)
	if err != nil {
		utils.ErrResp(c, err)
		return
	}
	c.JSON(consts.StatusOK, user.UserActionResponse{
		StatusCode: errno.SuccessCode,
		StatusMsg:  errno.SuccessMsg,
		UserId:     uid,
	})
}

// UserProfileModify .
//
//	@router	/blog_server/user/profile_modify/ [POST]
func UserProfileModify(ctx context.Context, c *app.RequestContext) {
	var err error
	req := new(user.UserActionProfileModifyRequest)
	err = c.BindAndValidate(req)
	if err != nil {
		utils.ErrResp(c, err)
		return
	}
	err = service.NewUserService(ctx, c).UserProfileModify(req)
	if err != nil {
		utils.ErrResp(c, err)
		return
	}
	c.JSON(consts.StatusOK, user.UserActionResponse{
		StatusCode: errno.SuccessCode,
		StatusMsg:  errno.SuccessMsg,
		UserId:     req.UserId,
	})
}
