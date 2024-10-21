package service

import (
	"context"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/qingyggg/blog_server/biz/dal/db"
	"github.com/qingyggg/blog_server/biz/model/hertz/interact/collect"
	service_utils "github.com/qingyggg/blog_server/biz/service"
	"github.com/qingyggg/blog_server/pkg/errno"
	"github.com/qingyggg/blog_server/pkg/utils"
)

type CollectService struct {
	ctx context.Context
	c   *app.RequestContext
}

// NewCommentService create comment service
func NewCollectService(ctx context.Context, c *app.RequestContext) *CollectService {
	return &CollectService{ctx: ctx, c: c}
}

func (c *CollectService) ACollectAction(req *collect.CollectActionRequest) error {
	exist, err := db.CheckArticleExistByHashId(req.AHashId)
	if err != nil {
		return err
	}
	if !exist {
		return errno.ArticleIsNotExistErr
	}
	uid := service_utils.GetUid(c.c)
	user, err := db.QueryUserById(uid)
	if err != nil {
		return err
	}
	uHashId := utils.ConvertByteHashToString(user.HashID)
	err, exist = db.ACollectExist(req.AHashId, uHashId)
	if err != nil {
		return err
	}
	if exist {
		if req.ActionType == 1 {
			return errno.CollectAlreadyExistErr
		} else {
			return db.ACollectDel(req.AHashId, uHashId)
		}

	} else {
		if req.ActionType == 2 {
			return errno.CollectIsNotExistErr
		} else {
			return db.ACollectAdd(req.AHashId, uHashId, req.Tag)
		}
	}
}
