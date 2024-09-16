package jwt

import (
	"context"
	"github.com/cloudwego/hertz/pkg/protocol"
	"github.com/qingyggg/blog_server/biz/dal/db"
	"github.com/qingyggg/blog_server/biz/model/hertz/basic/user"
	"github.com/qingyggg/blog_server/pkg/errno"
	"github.com/qingyggg/blog_server/pkg/utils"
	"time"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
	"github.com/hertz-contrib/jwt"
)

var (
	JwtMiddleware *jwt.HertzJWTMiddleware
	identity      = "user_id"
)

func Init() {
	JwtMiddleware, _ = jwt.New(&jwt.HertzJWTMiddleware{
		Key:         []byte("tiktok secret key"),
		TokenLookup: "cookie:token",
		Timeout:     24 * time.Hour,
		MaxRefresh:  time.Hour * 6,
		IdentityKey: identity,
		// Verify password at login
		Authenticator: func(ctx context.Context, c *app.RequestContext) (interface{}, error) {
			//var loginRequest
			user := new(user.UserActionRequest)
			if err := c.BindAndValidate(user); err != nil {
				return nil, err
			}
			uid, err := db.VerifyUser(user.Username, user.Password)
			if err != nil {
				return nil, err
			}
			c.Set("user_id", uid)
			return uid, nil
		},
		// Set the payload in the token
		PayloadFunc: func(data interface{}) jwt.MapClaims {
			if v, ok := data.(int64); ok {
				return jwt.MapClaims{
					identity: v,
				}
			}
			return jwt.MapClaims{}
		},
		// build login response if verify password successfully
		LoginResponse: func(ctx context.Context, c *app.RequestContext, code int, token string, expire time.Time) {
			hlog.CtxInfof(ctx, "Login success ，token is issued clientIP: "+c.ClientIP())
			c.SetCookie("token", token, int(24*time.Hour), "/", "localhost", protocol.CookieSameSiteLaxMode, true, true)
		},
		// Verify token and get the id of logged-in user
		Authorizator: func(data interface{}, ctx context.Context, c *app.RequestContext) bool {
			if v, ok := data.(float64); ok {
				current_user_id := int64(v)
				c.Set("current_user_id", current_user_id)
				hlog.CtxInfof(ctx, "Token is verified clientIP: "+c.ClientIP())
				return true
			}
			return false
		},
		// Validation failed, build the message
		Unauthorized: func(ctx context.Context, c *app.RequestContext, code int, message string) {
			c.JSON(consts.StatusOK, user.UserActionResponse{
				StatusCode: errno.AuthorizationFailedErrCode,
				StatusMsg:  message,
			})
		},
		HTTPStatusMessageFunc: func(e error, ctx context.Context, c *app.RequestContext) string {
			resp := utils.BuildBaseResp(e)
			return resp.StatusMsg
		},
	})
}
