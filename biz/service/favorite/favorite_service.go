package service

import (
	"context"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/qingyggg/blog_server/biz/dal/db"
	"github.com/qingyggg/blog_server/biz/model/hertz/interact/favorite"
	service_utils "github.com/qingyggg/blog_server/biz/service"
	"github.com/qingyggg/blog_server/pkg/errno"
	"github.com/qingyggg/blog_server/pkg/utils"
)

type FavoriteService struct {
	ctx context.Context
	c   *app.RequestContext
}

// NewCommentService create comment service
func NewFavoriteService(ctx context.Context, c *app.RequestContext) *FavoriteService {
	return &FavoriteService{ctx: ctx, c: c}
}

func (c *FavoriteService) CmtFavoriteAction(req *favorite.FavoriteActionRequest) error {
	exist, err := db.CheckCmtExistById(c.ctx, req.CHashID)
	if err != nil {
		return err
	}
	if !exist {
		return errno.CommentIsNotExistErr
	}
	uid := service_utils.GetUid(c.c)
	user, err := db.QueryUserById(uid)
	if err != nil {
		return err
	}
	uHashId := utils.ConvertByteHashToString(user.HashID)
	err, exist = db.CmtFavorieExist(uHashId, req.CHashID)
	if err != nil {
		return err
	}
	if exist {
		err, status := db.CmtFavoriteStatus(uHashId, req.CHashID)
		if err != nil {
			return err
		}
		if status == req.ActionType {
			return nil
		} else {
			err = db.CmtFavoriteStatusFlush(uHashId, req.CHashID)
			if err != nil {
				return err
			}
		}
	}
	if req.ActionType == 1 || req.ActionType == 2 {
		err = db.CmtFavoriteAction(uHashId, req.CHashID, req.AHashID, req.ActionType)
		if err != nil {
			return err
		}
	}
	return nil
}

func (c *FavoriteService) ArticleFavoriteAction(req *favorite.FavoriteActionRequest) error {
	exist, err := db.CheckArticleExistByHashId(req.AHashID)
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
	err, exSignal := db.AFavoriteExist(req.AHashID, uHashId)
	if err != nil {
		return err
	}
	if exSignal != 3 { //1 or 2
		if exSignal == req.ActionType {
			return nil
		}
		err = db.AFavoriteStatusFlush(req.AHashID, uHashId, exSignal)
		if err != nil {
			return err
		}
	}
	if req.ActionType == 1 || req.ActionType == 2 {
		err = db.AFavorite(req.AHashID, uHashId, req.ActionType)
		if err != nil {
			return err
		}
	}
	return nil
}
