// Code generated by hertz generator. DO NOT EDIT.

package comment

import (
	"github.com/cloudwego/hertz/pkg/app/server"
	comment "github.com/qingyggg/blog_server/biz/handler/comment"
)

/*
 This file will register all the routes of the services in the master idl.
 And it will update automatically when you use the "update" command for the idl.
 So don't modify the contents of the file, or your code will be deleted when it is updated.
*/

// Register register routes based on the IDL 'api.${HTTP Method}' annotation.
func Register(r *server.Hertz) {

	root := r.Group("/", rootMw()...)
	{
		_blog_server := root.Group("/blog_server", _blog_serverMw()...)
		{
			_comment := _blog_server.Group("/comment", _commentMw()...)
			_comment.DELETE("/action", append(_commentdelactionMw(), comment.CommentDelAction)...)
			_comment.POST("/action", append(_commentactionMw(), comment.CommentAction)...)
			_comment.GET("/list", append(_commentlistMw(), comment.CommentList)...)
		}
	}
}
