package service

import (
	"context"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/qingyggg/blog_server/biz/dal/db"
	"github.com/qingyggg/blog_server/biz/model/hertz/basic/publish"
	"github.com/qingyggg/blog_server/biz/model/hertz/basic/user"
	"github.com/qingyggg/blog_server/biz/model/hertz/common"
	"github.com/qingyggg/blog_server/biz/model/orm_gen"
	service "github.com/qingyggg/blog_server/biz/service/user"
	"github.com/qingyggg/blog_server/pkg/utils"
	"time"
)

//create,modify,delete,show single,show multi

type PublishService struct {
	ctx context.Context
	c   *app.RequestContext
}

// NewPublishService create publish service
func NewPublishService(ctx context.Context, c *app.RequestContext) *PublishService {
	return &PublishService{ctx: ctx, c: c}
}

func (s *PublishService) PublishCreate(req *publish.ArticleCreateActionRequest) (err error) {
	_, err = db.CreateArticle(&orm_gen.Article{
		UserID:      req.Uid,
		Title:       req.Payload.Preload.Title,
		Note:        req.Payload.Preload.Note,
		CoverURL:    utils.UrlConvertReverse(s.ctx, req.Payload.Preload.CoverUrl),
		PublishTime: time.Now(),
	}, req.Payload.Content)
	return err
}

func (s *PublishService) PublishModify(req *publish.ArticleModifyActionRequest) (err error) {
	_, err = db.ModifyArticle(&orm_gen.Article{
		UserID: req.Base.Uid,
		HashID: req.Base.AHashID,
		Title:  req.Payload.Preload.Title,
		Note:   req.Payload.Preload.Note,
	}, req.Payload.Content)
	return err
}

func (s *PublishService) PublishDelete(req *publish.ArticleBaseActionRequest) (err error) {
	err = db.DeleteArticle(&orm_gen.Article{UserID: req.Uid, HashID: req.AHashID})
	if err != nil {
		return err
	}
	return err
}

func (s *PublishService) PublishDetail(req *publish.ArticleBaseActionRequest) (aA *common.Article, err error) {
	//获取文章信息
	aInfo, aContent, err := db.TakeArticle(req.AHashID, req.Uid)
	if err != nil {
		return nil, err
	}
	//获取用户信息
	curUser, err := service.NewUserService(s.ctx, s.c).UserInfo(&user.UserRequest{UserId: aInfo.UserID})
	if err != nil {
		return nil, err
	}
	aA = new(common.Article)
	aA.Id = aInfo.ID
	aA.HashId = aInfo.HashID
	aA.Author = curUser //作者信息
	aA.Base = &common.ArticleBase{
		Preload: &common.ArticleBasePreload{
			Title:    aInfo.Title,
			Note:     aInfo.Note,
			CoverUrl: utils.URLconvert(s.ctx, s.c, aInfo.CoverURL),
		},
		Content: aContent,
	}
	aA.Info = &common.ArticleInfo{
		LikeCount:    aInfo.LikeCount,
		CommentCount: aInfo.CommentCount,
		DislikeCount: aInfo.DislikeCount,
		ViewedCount:  aInfo.ViewCount,
	}

	return aA, nil
}

func (s *PublishService) PublishList(req *publish.ArticleCardsRequest) (cards []*common.ArticleCard, err error) {
	aInfos, err := db.GetArticleInfos(req.UserId, int(req.Offset))
	if err != nil {
		return nil, err
	}
	var uids []int64
	for _, aInfo := range aInfos {
		uids = append(uids, aInfo.UserID)
	}
	uMaps, err := db.QueryUserByIds(uids)
	if err != nil {
		return nil, err
	}
	for _, aInfo := range aInfos {
		curUser := (*uMaps)[aInfo.UserID]
		cards = append(cards, &common.ArticleCard{
			Id:     aInfo.ID,
			HashId: aInfo.HashID,
			Author: &common.UserBase{
				Id:   curUser.ID,
				Name: curUser.UserName,
				Profile: &common.UserProfile{
					Avatar:          utils.URLconvert(s.ctx, s.c, curUser.Avatar),
					Signature:       curUser.Signature,
					BackgroundImage: utils.URLconvert(s.ctx, s.c, curUser.BackgroundImage),
				},
			},
			Info: &common.ArticleInfo{
				LikeCount:    aInfo.LikeCount,
				CommentCount: aInfo.CommentCount,
				DislikeCount: aInfo.DislikeCount,
				ViewedCount:  aInfo.ViewCount,
			},
			Pre: &common.ArticleBasePreload{
				Title:    aInfo.Title,
				Note:     aInfo.Note,
				CoverUrl: utils.URLconvert(s.ctx, s.c, aInfo.CoverURL),
			},
		})
	}
	return cards, nil
}

func (s *PublishService) AddViewCount(req *publish.ArticleBaseActionRequest) error {
	err := db.AddViewCount(req.Uid, req.AHashID)
	return err
}
