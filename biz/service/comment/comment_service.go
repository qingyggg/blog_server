package service

import (
	"context"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/qingyggg/blog_server/biz/dal/db"
	"github.com/qingyggg/blog_server/biz/model/hertz/interact/comment"
	"github.com/qingyggg/blog_server/biz/model/orm_gen"
	"github.com/qingyggg/blog_server/biz/mw/mongo"
	service_utils "github.com/qingyggg/blog_server/biz/service"
	service "github.com/qingyggg/blog_server/biz/service/user"
	"github.com/qingyggg/blog_server/pkg/errno"
	"github.com/qingyggg/blog_server/pkg/utils"
	"math/rand"
	"strconv"
	"sync"
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
	err = db.DelCommentByHashID(c.ctx, req.CHashId, req.AHashId)
	return err
}

func (c *CommentService) GetTopCmtList(req *comment.CommentListRequest) (err error, comments []*comment.Comment) {
	//1.请求
	firstList, err := db.GetCommentListByArticleID(c.ctx, req.AHashId)
	if err != nil {
		return err, nil
	}
	err, cmts := c.getCmtList(firstList)
	if err != nil {
		return err, nil
	}
	return nil, cmts
}

func (c *CommentService) GetSubCmtList(req *comment.CommentListRequest) (err error, cmts []*comment.Comment) {
	firstList, err := db.GetCommentListByTopCommentID(c.ctx, req.AHashId, req.CHashId)
	if err != nil {
		return err, nil
	}
	err, cmts = c.getCmtList(firstList)
	if err != nil {
		return err, nil
	}
	return nil, cmts
}

func (c *CommentService) getCmtList(cmtList []*mongo.CommentItem) (err error, cmts []*comment.Comment) {
	//获取uid,cid
	var uids []string
	var cids []string
	var wg sync.WaitGroup
	cmts = []*comment.Comment{}
	var UInfoMaps map[string]*orm_gen.User
	var CInfoMaps map[string]int32
	var CCtMaps map[string]int64
	var errChan = make(chan error, 3)
	wg.Add(3)
	for _, cmt := range cmtList {
		uids = append(uids, cmt.UserID)
		cids = append(cids, cmt.HashID)
	}
	curUid := service_utils.GetUid(c.c)
	user, err := db.QueryUserById(curUid)
	if err != nil {
		return err, nil
	}

	go func() {
		defer wg.Done()
		uInfoMaps, err := db.QueryUserByHashIds(uids) //请求用户信息
		if err != nil {
			errChan <- err
			return
		}
		UInfoMaps = uInfoMaps
	}()
	go func() {
		defer wg.Done()
		err, cInfoMaps := db.GetCmtFavoriteStatusMap(cids, utils.ConvertByteHashToString(user.HashID))
		if err != nil {
			errChan <- err
			return
		}
		CInfoMaps = cInfoMaps
	}()
	go func() {
		defer wg.Done()
		err, cCtMaps := db.GetCmtFavoriteCtMap(cids)
		if err != nil {
			errChan <- err
			return
		}
		CCtMaps = cCtMaps
	}()
	go func() {
		wg.Wait()
		close(errChan)
	}()
	for err := range errChan {
		if err != nil {
			return err, nil
		}
	}
	//3.用户信息附加,形成完整的评论表
	var isFavorite bool
	var curUserPayload *orm_gen.User
	for _, cmt := range cmtList {
		curUserPayload = UInfoMaps[cmt.UserID]
		if CInfoMaps[cmt.HashID] == 1 {
			isFavorite = true
		} else {
			isFavorite = false
		}
		curPayload := comment.Comment{
			CHashId:       cmt.HashID,
			AHashId:       cmt.ArticleID,
			Content:       cmt.Content,
			ChildNum:      cmt.ChildNum,
			User:          service.UserAssign(curUserPayload),
			CreateDate:    utils.ConvertBsonTimeToString(cmt.CreateTime),
			FavoriteCount: CCtMaps[cmt.HashID], //后续完善
			IsFavorite:    isFavorite,
		}
		cmts = append(cmts, &curPayload)
	}
	return nil, cmts
}
