// Code generated by hertz generator.

package user

import (
	"context"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
	user "github.com/qingyggg/blog_server/biz/model/hertz/basic/user"
	"github.com/qingyggg/blog_server/biz/mw/jwt"
	service_utils "github.com/qingyggg/blog_server/biz/service"
	service "github.com/qingyggg/blog_server/biz/service/user"
	"github.com/qingyggg/blog_server/pkg/errno"
	"github.com/qingyggg/blog_server/pkg/utils"
)

// User .
//
// @Summary 获取用户信息
// @Description 根据用户请求获取对应的用户信息
// @Tags 用户相关接口
// @Accept json
// @Produce json
// @Param user_id query int true "用户ID"
// @Success 200 {object} user.UserResponse "成功获取用户信息"
// @Failure 400 {object} user.UserActionResponse "请求参数错误"
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
	if err != nil {
		utils.ErrResp(c, err)
		return
	}
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
	uHashId, err := service.NewUserService(ctx, c).UserRegister(req)
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
		UHashId:    uHashId,
	})
}

// UserLogin 登录接口
//
//	@Summary		用户登录
//	@Description	用户通过提供用户名和密码登录账户
//	@Tags			用户相关接口
//	@Accept			json
//	@Produce		json
//	@Param			user	body		user.UserActionRequest	true	"用户登录请求参数"
//	@Success		200		{object}	user.UserActionResponse	"成功返回用户ID及状态信息"
//	@Failure		400		{object}	user.UserActionResponse	"请求参数错误或其他错误信息"
//	@router	/blog_server/user/login/ [POST]
func UserLogin(ctx context.Context, c *app.RequestContext) {
	var err error
	req := new(user.UserActionRequest)
	err = c.BindAndValidate(req)
	if err != nil {
		utils.ErrResp(c, err)
		return
	}
	//在这里使用jwt login handler进行,如果登录成功，这个中间件会设置cookie,但若是失败了，或者用户他不存在，那么这个插件会自动进行错误响应，因此，该loginhandler不需要再c.json....
	jwt.JwtMiddleware.LoginHandler(ctx, c)
	hasErr, _ := c.Get("hasErr")
	if hasErr.(bool) {
		return
	}
	v, _ := c.Get("user_id")
	user_id := v.(int64)
	c.JSON(consts.StatusOK, user.UserActionResponse{
		StatusCode: errno.SuccessCode,
		StatusMsg:  errno.SuccessMsg,
		UserId:     user_id,
	})
}

// UserPwdModify 密码修改接口
//
//	@Summary		用户修改密码
//	@Description	用户通过提供用户名，旧密码，新密码进行修改账户密码
//	@Tags			用户相关接口
//	@Accept			json
//	@Produce		json
//	@Param			user	body		user.UserActionPwdModifyRequest	true	"用户修改密码请求参数"
//	@Success		200		{object}	user.UserActionResponse	"成功返回用户ID及状态信息"
//	@Failure		400		{object}	user.UserActionResponse	"请求参数错误或其他错误信息"
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
//	@Summary		用户修改资料
//	@Description	用户通过上传新的头像，背景图片，或者更改用户签名来进行修改资料
//	@Tags			用户相关接口
//	@Accept			json
//	@Produce		json
//	@Param			user	body		user.UserActionProfileModifyRequest	true	"用户修改资料请求参数"
//	@Success		200		{object}	user.UserActionResponse	"成功返回用户ID及状态信息"
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
		UserId:     service_utils.GetUid(c),
	})
}
