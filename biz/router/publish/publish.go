// Code generated by hertz generator. DO NOT EDIT.

package publish

import (
	"github.com/cloudwego/hertz/pkg/app/server"
	publish "github.com/qingyggg/blog_server/biz/handler/publish"
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
			_publish := _blog_server.Group("/publish", _publishMw()...)
			_publish.DELETE("/action", append(_publishdelactionMw(), publish.PublishDelAction)...)
			_publish.PATCH("/action", append(_publishmodifyactionMw(), publish.PublishModifyAction)...)
			_publish.POST("/action", append(_publishactionMw(), publish.PublishAction)...)
			_publish.GET("/detail", append(_publishdetailMw(), publish.PublishDetail)...)
			_publish.GET("/list", append(_publishlistMw(), publish.PublishList)...)
			_publish.POST("/view_add", append(_publishviewcountaddMw(), publish.PublishViewCountAdd)...)
		}
	}
}
