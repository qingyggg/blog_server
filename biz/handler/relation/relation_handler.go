// Code generated by hertz generator.

package relation

import (
	"context"
	service "github.com/qingyggg/blog_server/biz/service/relation"
	"github.com/qingyggg/blog_server/pkg/errno"
	"github.com/qingyggg/blog_server/pkg/utils"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
	relation "github.com/qingyggg/blog_server/biz/model/hertz/social/relation"
)

// RelationAction 处理用户关注或取消关注操作
// @Summary 关注/取消关注
// @Description 用户对其他用户进行关注或取消关注
// @Tags 用户关系
// @Accept application/json
// @Produce application/json
// @Param relation body relation.RelationActionRequest true "RelationAction请求参数"
// @Success 200 {object} relation.RelationActionResponse "成功响应，包含状态码和状态信息"
// @Failure 400 {object} common.BaseResponse "Invalid request"
// @Failure 500 {object} common.BaseResponse "Internal server error"
// @Router /blog_server/relation/action [POST]
func RelationAction(ctx context.Context, c *app.RequestContext) {
	var err error
	req := new(relation.RelationActionRequest)
	err = c.BindAndValidate(req)
	if err != nil {
		utils.ErrResp(c, err)
		return
	}
	err = service.NewRelationService(ctx, c).FollowAction(req)
	if err != nil {
		utils.ErrResp(c, err)
		return
	}

	c.JSON(consts.StatusOK, &relation.RelationActionResponse{
		StatusCode: errno.SuccessCode,
		StatusMsg:  errno.SuccessMsg,
	})
}

// RelationFollowList 获取用户关注列表
// @Summary 获取关注列表
// @Description 获取指定用户的关注列表
// @Tags 用户关系
// @Accept application/json
// @Produce application/json
// @Param  UhashID  query string true "用户ID"
// @Success 200 {object} relation.RelationFollowListResponse "成功响应，包含状态码、状态信息和用户列表"
// @Failure 400 {object} common.BaseResponse "Invalid request"
// @Failure 500 {object} common.BaseResponse "Internal server error"
// @Router /blog_server/relation/follow/list [GET]
func RelationFollowList(ctx context.Context, c *app.RequestContext) {
	var err error
	req := new(relation.RelationFollowListRequest)
	err = c.BindAndValidate(req)
	if err != nil {
		utils.ErrResp(c, err)
		return
	}
	list, err := service.NewRelationService(ctx, c).GetRelationList(req, 1)
	if err != nil {
		utils.ErrResp(c, err)
		return
	}
	c.JSON(consts.StatusOK, &relation.RelationFollowListResponse{
		StatusCode: errno.SuccessCode,
		StatusMsg:  errno.SuccessMsg,
		UserList:   list,
	})
}

// RelationFollowerList 获取用户粉丝列表
// @Summary 获取粉丝列表
// @Description 获取指定用户的粉丝列表
// @Tags 用户关系
// @Accept application/json
// @Produce application/json
// @Param UhashID  query string true "用户ID"
// @Success 200 {object} relation.RelationFollowListResponse "成功响应，包含状态码、状态信息和用户列表"
// @Failure 400 {object} common.BaseResponse "Invalid request"
// @Failure 500 {object} common.BaseResponse "Internal server error"
// @Router /blog_server/relation/follower/list [GET]
func RelationFollowerList(ctx context.Context, c *app.RequestContext) {
	var err error
	req := new(relation.RelationFollowListRequest)
	err = c.BindAndValidate(req)
	if err != nil {
		utils.ErrResp(c, err)
		return
	}
	list, err := service.NewRelationService(ctx, c).GetRelationList(req, 2)
	if err != nil {
		utils.ErrResp(c, err)
		return
	}
	c.JSON(consts.StatusOK, &relation.RelationFollowListResponse{
		StatusCode: errno.SuccessCode,
		StatusMsg:  errno.SuccessMsg,
		UserList:   list,
	})
}