// Code generated by hertz generator.

package favorite

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

func _favoriteMw() []app.HandlerFunc {
	return []app.HandlerFunc{
		jwt.JwtMiddleware.MiddlewareFunc(),
	}
}

func _actionMw() []app.HandlerFunc {
	// your code...
	return nil
}

func _articlefavoriteactionMw() []app.HandlerFunc {
	// your code...
	return nil
}

func _commentfavoriteactionMw() []app.HandlerFunc {
	// your code...
	return nil
}