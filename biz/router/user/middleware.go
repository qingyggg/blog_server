// Code generated by hertz generator.

package user

import (
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/qingyggg/blog_server/biz/mw/jwt"
)

func rootMw() []app.HandlerFunc {
	// your code...
	return nil
}

func _blog_serverMw() []app.HandlerFunc {
	// your code...
	return nil
}

func _userMw() []app.HandlerFunc {
	// your code...
	return nil
}

func _user0Mw() []app.HandlerFunc {
	// your code...
	return nil
}

func _loginMw() []app.HandlerFunc {
	// your code...
	return nil
}

func _userloginMw() []app.HandlerFunc {
	// your code...
	return nil
}

func _profile_modifyMw() []app.HandlerFunc {
	return []app.HandlerFunc{
		jwt.JwtMiddleware.MiddlewareFunc(),
	}
}

func _userprofilemodifyMw() []app.HandlerFunc {
	// your code...
	return nil
}

func _pwd_modifyMw() []app.HandlerFunc {
	return []app.HandlerFunc{
		jwt.JwtMiddleware.MiddlewareFunc(),
	}
}

func _userpwdmodifyMw() []app.HandlerFunc {
	// your code...
	return nil
}

func _registerMw() []app.HandlerFunc {
	// your code...
	return nil
}

func _userregisterMw() []app.HandlerFunc {
	// your code...
	return nil
}
