package service

import (
	"context"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/qingyggg/blog_server/biz/dal/db"
	"github.com/qingyggg/blog_server/biz/model/hertz/basic/publish"
	"github.com/qingyggg/blog_server/biz/model/hertz/basic/user"
	"github.com/qingyggg/blog_server/biz/model/hertz/common"
	"github.com/qingyggg/blog_server/biz/model/orm_gen"
	service_utils "github.com/qingyggg/blog_server/biz/service"
	service "github.com/qingyggg/blog_server/biz/service/user"
	"github.com/qingyggg/blog_server/pkg/constants"
	"github.com/qingyggg/blog_server/pkg/errno"
	"github.com/qingyggg/blog_server/pkg/utils"
	"strconv"
	"sync"
	"time"
)

type PublishService struct {
	ctx context.Context
	c   *app.RequestContext
}

// NewPublishService create publish service
func NewPublishService(ctx context.Context, c *app.RequestContext) *PublishService {
	return &PublishService{ctx: ctx, c: c}
}

func (s *PublishService) PublishCreate(req *publish.CreateActionRequest) (err error, aHashId string) {
	uid := service_utils.GetUid(s.c)
	user, err := db.QueryUserById(uid)
	if err != nil {
		return err, ""
	}
	var coverUrl string
	if req.Payload.Preload.CoverUrl == "" {
		coverUrl = constants.TestBackground
	} else {
		coverUrl = utils.UrlConvertReverse(s.ctx, req.Payload.Preload.CoverUrl)
	}
	aHashId = utils.GetSHA256String(time.Now().String() + strconv.FormatInt(uid, 16))
	//1.数据库创建记录
	err = db.CreateArticle(&orm_gen.Article{
		UserID:      user.HashID,
		Title:       req.Payload.Preload.Title,
		Note:        req.Payload.Preload.Note,
		CoverURL:    coverUrl,
		PublishTime: time.Now(),
		HashID:      utils.ConvertStringHashToByte(aHashId),
	}, req.Payload.Content)
	if err != nil {
		return err, ""
	}
	//2.初始化view count
	err, exist := db.ViewCountExist(aHashId)
	if err != nil {
		return err, ""
	}
	if exist {
		return errno.ServiceErr.WithMessage("当前view count已存在"), ""
	}
	err = db.ViewCountInit(aHashId)
	if err != nil {
		return err, ""
	}
	return nil, aHashId
}

func (s *PublishService) PublishModify(req *publish.ModifyActionRequest) (err error) {
	uid := service_utils.GetUid(s.c)
	exist, err := db.CheckArticleExistByHashId(req.AHashID)
	if err != nil {
		return err
	}
	if !exist {
		return errno.ArticleIsNotExistErr
	}
	user, err := db.QueryUserById(uid)
	if err != nil {
		return err
	}
	_, err = db.ModifyArticle(&orm_gen.Article{
		UserID:   user.HashID,
		HashID:   utils.ConvertStringHashToByte(req.AHashID),
		Title:    req.Payload.Preload.Title,
		Note:     req.Payload.Preload.Note,
		CoverURL: req.Payload.Preload.CoverUrl,
	}, req.Payload.Content)
	return err
}

func (s *PublishService) PublishDelete(req *publish.DelActionRequest) (err error) {
	exist, err := db.CheckArticleExistByHashId(req.AHashID)
	if err != nil {
		return err
	}
	if !exist {
		return errno.ArticleIsNotExistErr
	}
	var wg sync.WaitGroup
	wg.Add(5)
	errChan := make(chan error, 5)
	go func() {
		defer wg.Done()
		//1.删除文章
		uid := service_utils.GetUid(s.c)
		user, err := db.QueryUserById(uid)
		if err != nil {
			errChan <- err
			return
		}
		err = db.DeleteArticle(&orm_gen.Article{UserID: user.HashID, HashID: utils.ConvertStringHashToByte(req.AHashID)})
		if err != nil {
			errChan <- err
			return
		}
	}()
	go func() {
		defer wg.Done()
		//2.删除与文章相关的评论
		err := db.DelCommentByAid(s.ctx, req.AHashID)
		if err != nil {
			errChan <- err
			return
		}
	}()
	go func() {
		defer wg.Done()
		//3.删除article favorite依赖以及article favorite count
		err := db.AFavoriteDeleteByAid(req.AHashID)
		if err != nil {
			errChan <- err
			return
		}
	}()
	go func() {
		defer wg.Done()
		err = db.ViewCountDel(req.AHashID)
		if err != nil {
			errChan <- err
			return
		}
	}()
	//清理掉collect
	go func() {
		defer wg.Done()
		err := db.ACollectDelByAid(req.AHashID)
		if err != nil {
			errChan <- err
			return
		}
	}()
	go func() {
		wg.Wait()
		close(errChan)
	}()

	for err := range errChan {
		if err != nil {
			return err
		}
	}
	return nil
}

func (s *PublishService) PublishDetail(req *publish.DetailRequest) (*common.Article, error) {
	aA := new(common.Article)
	aInfo := new(common.ArticleInfo)
	//1.get article data from db
	if err := s.publishDetail(req.AHashID, aA); err != nil {
		return nil, err
	}
	var wg sync.WaitGroup
	var mu sync.Mutex
	var errChan = make(chan error, 6)
	wg.Add(6)
	//defer close(errChan)
	//1.get author info
	go func() {
		defer wg.Done()
		//获取作者信息
		curUser, err := service.NewUserService(s.ctx, s.c).UserInfo(&user.UserRequest{UHashID: req.UHashID})
		if err != nil {
			errChan <- err
			return
		}
		mu.Lock()
		aA.Author = curUser //作者信息
		mu.Unlock()
	}()
	//2. like count
	go func() {
		defer wg.Done()
		err, ct := db.AFavoriteCtGet(req.AHashID)
		if err != nil {
			errChan <- err
			return
		}
		mu.Lock()
		aInfo.LikeCount = ct
		mu.Unlock()
	}()
	//3.comment count
	go func() {
		defer wg.Done()
		err, ct := db.GetCmtCountByAid(s.ctx, req.AHashID)
		if err != nil {
			errChan <- err
			return
		}
		mu.Lock()
		aInfo.CommentCount = ct
		mu.Unlock()
	}()
	//collect count
	go func() {
		defer wg.Done()
		err, count := db.ACollectCtGet(req.AHashID)
		if err != nil {
			errChan <- err
			return
		}
		mu.Lock()
		aInfo.CollectCount = count
		mu.Unlock()
	}()
	//view count
	go func() {
		defer wg.Done()
		var count int64
		err, exist := db.ViewCountExist(req.AHashID)
		if err != nil {
			errChan <- err
			return
		}
		if exist {
			err, count = db.ViewCountGet(req.AHashID)
			if err != nil {
				errChan <- err
				return
			}
		} else {
			count = 0
		}
		mu.Lock()
		aInfo.ViewedCount = count
		mu.Unlock()
	}()
	//if token exists:
	//is favorite,is collect
	go func() {
		defer wg.Done()
		if uid := service_utils.GetUid(s.c); uid != 0 {

			var err error
			var colEx bool
			var faSig int32
			var faEx bool
			user, err := db.QueryUserById(uid)
			if err != nil {
				errChan <- err
				return
			}
			uHashId := utils.ConvertByteHashToString(user.HashID)
			err, colEx = db.ACollectExist(req.AHashID, uHashId)
			if err != nil {
				errChan <- err
				return
			}
			err, faSig = db.AFavoriteExist(req.AHashID, uHashId)
			if err != nil {
				errChan <- err
				return
			}
			if faSig == 1 {
				faEx = true
			} else {
				faEx = false
			}
			aInfo.IsFavorite = faEx
			aInfo.IsCollect = colEx
		} else {
			aInfo.IsFavorite = false
			aInfo.IsCollect = false
		}
	}()

	go func() {
		wg.Wait()
		close(errChan)
	}()

	for err := range errChan {
		if err != nil {
			return nil, err
		}
	}
	aA.Info = aInfo
	return aA, nil
}

func (s *PublishService) publishDetail(aHashId string, aA *common.Article) error {
	//检查文章是否存在
	exist, err := db.CheckArticleExistByHashId(aHashId)
	if err != nil {
		return err
	}
	if !exist {
		return errno.ArticleIsNotExistErr
	}
	//获取文章信息
	aInfo, aContent, err := db.TakeArticle(aHashId)
	if err != nil {
		return err
	}
	aA.Base = &common.ArticleBase{
		Preload: &common.ArticleBasePreload{
			Title:    aInfo.Title,
			Note:     aInfo.Note,
			CoverUrl: utils.URLconvert(s.ctx, s.c, aInfo.CoverURL),
		},
		Content: aContent,
	}
	aA.Id = aInfo.ID
	aA.HashId = utils.ConvertByteHashToString(aInfo.HashID)
	return nil
}

func (s *PublishService) PublishList(req *publish.CardsRequest) (cards []*common.ArticleCard, err error) {
	var wg sync.WaitGroup
	var aHashIds []string
	var uids []string
	var UMap map[string]*orm_gen.User //用户映射
	var CMap map[string]int64         //评论数映射
	var AfMap map[string]int64        //点赞数映射
	var VMap map[string]int64         //阅读数
	var CoMap map[string]int64        //收藏数量
	var errChan = make(chan error, 5)
	defer close(errChan)
	//1.请求文章信息
	aInfos, err := db.GetArticleInfos(req.UHashID, int(req.Offset))
	if err != nil {
		return nil, err
	}
	for _, v := range aInfos {
		aHashIds = append(aHashIds, utils.ConvertByteHashToString(v.HashID))
		uids = append(uids, utils.ConvertByteHashToString(v.UserID))
	}
	wg.Add(5)
	//2.请求作者信息
	go func() {
		defer wg.Done()
		uMap, err := db.QueryUserByHashIds(uids)
		if err != nil {
			errChan <- err
			return
		}
		UMap = uMap
	}()
	//3.文章被评论的数量
	go func() {
		defer wg.Done()
		err, cMap := db.GetCmtCtByAids(s.ctx, aHashIds)
		if err != nil {
			errChan <- err
			return
		}
		CMap = cMap
	}()
	//4.文章被点赞的数量
	go func() {
		defer wg.Done()
		err, afMap := db.AFavoriteCtGetByAids(aHashIds)
		if err != nil {
			errChan <- err
			return
		}
		AfMap = afMap
	}()
	//5.文章被阅览的数量
	go func() {
		defer wg.Done()
		err, vMap := db.ViewCountGets(aHashIds)
		if err != nil {
			errChan <- err
			return
		}
		VMap = vMap
	}()
	//6.文章的收藏数量
	go func() {
		defer wg.Done()
		err, coMap := db.ACollectCtGetByAids(aHashIds)
		if err != nil {
			errChan <- err
			return
		}
		CoMap = coMap
	}()

	wg.Wait()
	select {
	case err := <-errChan:
		if err != nil {
			return nil, err
		}
	default:
	}
	//处理字段
	var cur *common.ArticleCard
	for idx, info := range aInfos {
		cur = &common.ArticleCard{
			Id:     info.ID,
			HashId: aHashIds[idx],
			Author: service.NewUserService(s.ctx, s.c).UserAssign(UMap[utils.ConvertByteHashToString(info.UserID)]),
			Pre: &common.ArticleBasePreload{
				Title:    info.Title,
				Note:     info.Note,
				CoverUrl: utils.URLconvert(s.ctx, s.c, info.CoverURL),
			},
			Info: &common.ArticleInfo{
				LikeCount:    AfMap[aHashIds[idx]],
				CollectCount: CoMap[aHashIds[idx]],
				ViewedCount:  VMap[aHashIds[idx]],
				CommentCount: CMap[aHashIds[idx]],
			},
		}
		cards = append(cards, cur)
	}
	return cards, nil
}

func (s *PublishService) AddViewCount(req *publish.ActionRequest) error {
	exist, err := db.CheckArticleExistByHashId(req.AHashID)
	if err != nil {
		return err
	}
	if !exist {
		return errno.ArticleIsNotExistErr
	}
	err = db.ViewCountIncr(req.AHashID)
	if err != nil {
		return err
	}
	return nil
}
