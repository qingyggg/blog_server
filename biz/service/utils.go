package service_utils

import "github.com/cloudwego/hertz/pkg/app"

// get uid from jwt
func GetUid(c *app.RequestContext) int64 {
	uid, exist := c.Get("current_user_id")
	if exist {
		return uid.(int64)
	} else {
		return 0
	}
}
