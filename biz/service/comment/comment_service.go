package service

import (
	"context"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/qingyggg/blog_server/biz/dal/db"
	"github.com/qingyggg/blog_server/biz/model/hertz/interact/comment"
	"github.com/qingyggg/blog_server/biz/mw/mongo"
	service_utils "github.com/qingyggg/blog_server/biz/service"
	service "github.com/qingyggg/blog_server/biz/service/user"
	"github.com/qingyggg/blog_server/pkg/errno"
	"github.com/qingyggg/blog_server/pkg/utils"
	"math/rand"
	"strconv"
	"time"
)

type CommentService struct {
	ctx context.Context
	c   *app.RequestContext
}

// NewCommentService create comment service
func NewCommentService(ctx context.Context, c *app.RequestContext) *CommentService {
	return &CommentService{ctx: ctx, c: c}
}

func (c *CommentService) AddNewCmt(req *comment.CommentActionRequest) (error, string) {
	if req.Content == "" {
		return errno.ParamErr.WithMessage("评论内容不可以为空"), ""
	}
	//检查文章是否存在
	aExist, err := db.CheckArticleExistByHashId(req.AHashId)
	if err != nil {
		return err, ""
	}
	if !aExist {
		return errno.ArticleIsNotExistErr, "" //文章不存在
	}
	//检查父评论是否存在
	if req.Degree == 2 {
		exist, err := db.CheckCmtExistById(c.ctx, req.PHashId)
		if err != nil {
			return err, ""
		}
		if !exist {
			return errno.CommentIsNotExistErr.WithMessage("您回复的评论不存在"), ""
		}
	}
	uid := service_utils.GetUid(c.c)
	user, err := db.QueryUserById(uid)
	if err != nil {
		return err, ""
	}
	UHashId := utils.ConvertByteHashToString(user.HashID)
	cHashId := utils.GetSHA256String(req.AHashId + UHashId + strconv.FormatInt(int64(rand.Int()), 10) + time.Now().String())
	err = db.AddNewComment(c.ctx, &mongo.Comment{
		ArticleID: req.AHashId,
		UserID:    UHashId,
		Content:   req.Content,
		ParentID:  req.PHashId,
		Degree:    int8(req.Degree),
		HashID:    cHashId,
	})
	if err != nil {
		return err, ""
	}
	return nil, cHashId
}

func (c *CommentService) DelCmt(req *comment.CommentDelActionRequest) error {
	//
	uid := service_utils.GetUid(c.c)
	user, err := db.QueryUserById(uid)
	if err != nil {
		return err
	}
	exist, err := db.CheckCmtExistById(c.ctx, req.CHashId)
	if err != nil {
		return err
	}
	if !exist {
		return errno.CommentIsNotExistErr
	}
	cmt, err := db.GetCommentByCmtID(c.ctx, req.CHashId)
	if err != nil {
		return err
	}
	if cmt.UserID != utils.ConvertByteHashToString(user.HashID) {
		return errno.ServiceErr.WithMessage("当前用户没有权利删除该评论")
	}
	err = db.DelCommentByHashID(c.ctx, req.CHashId)
	return err
}

func (c *CommentService) GetTopCmtList(req *comment.CommentListRequest) (err error, comments []*comment.Comment) {
	//1.请求
	cmtList, err := db.GetCommentListByArticleID(c.ctx, req.AHashId)
	if err != nil {
		return err, nil
	}
	//2.获取uid
	var uids []string
	for _, cmt := range cmtList {
		uids = append(uids, cmt.UserID)
	}
	uInfoMaps, err := db.QueryUserByHashIds(uids)
	//3.用户信息附加,形成完整的评论表
	for _, cmt := range cmtList {
		curUserPayload := uInfoMaps[cmt.UserID]
		curPayload := comment.Comment{
			CHashId:       cmt.HashID,
			AHashId:       cmt.ArticleID,
			Content:       cmt.Content,
			ChildNum:      cmt.ChildNum,
			User:          service.UserAssign(curUserPayload),
			CreateDate:    utils.ConvertBsonTimeToString(cmt.CreateTime),
			FavoriteCount: 3, //后续完善
		}
		comments = append(comments, &curPayload)
	}
	return nil, comments
}

func (c *CommentService) GetSubCmtList(req *comment.CommentListRequest) (err error, cmts []*comment.Comment) {
	//1.数据库操作，平摊数组，获取uid，对uid进行搜索用户信息
	var uids []string
	list, err := db.GetCommentListByTopCommentID(c.ctx, req.AHashId, req.CHashId)
	if err != nil {
		return err, nil
	}
	for _, v := range list {
		uids = append(uids, v.UserID)
	}
	//请求用户信息
	uMaps, err := db.QueryUserByHashIds(uids)
	if err != nil {
		return err, nil
	}
	for _, v := range list {
		cmts = append(cmts, &comment.Comment{
			CHashId:        v.HashID,
			AHashId:        v.ArticleID,
			User:           service.UserAssign(uMaps[v.UserID]),
			Content:        v.Content,
			CreateDate:     utils.ConvertBsonTimeToString(v.CreateTime),
			RepliedUHashId: v.ParentUID,
			FavoriteCount:  3, //后续完善
		})
	}
	return nil, cmts
}
