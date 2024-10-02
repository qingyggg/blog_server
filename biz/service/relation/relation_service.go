package service

import (
	"context"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/qingyggg/blog_server/biz/dal/db"
	"github.com/qingyggg/blog_server/biz/model/hertz/common"
	"github.com/qingyggg/blog_server/biz/model/hertz/social/relation"
	"github.com/qingyggg/blog_server/biz/model/orm_gen"
	service "github.com/qingyggg/blog_server/biz/service/user"
	"github.com/qingyggg/blog_server/pkg/errno"
)

var (
	FOLLOW    int32 = 1
	UNFOLLOW  int32 = 2
	FOLLOWING int32 = 1
	FOLLOWER  int32 = 2
)

type RelationService struct {
	ctx context.Context
	c   *app.RequestContext
}

// NewPublishService create publish service
func NewRelationService(ctx context.Context, c *app.RequestContext) *RelationService {
	return &RelationService{ctx: ctx, c: c}
}

func (s *RelationService) FollowAction(req *relation.RelationActionRequest) (flag bool, err error) {
	//1.check user
	exist, err := db.CheckUserExistById(req.ToUserId)
	if err != nil {
		return false, err
	}
	if exist != true {
		return false, errno.UserIsNotExistErr
	}
	//2.check param user id != to user id,action type==1 or 2
	curUid, _ := s.c.Get("current_user_id")
	if curUid.(int64) == req.ToUserId {
		return false, errno.NewErrNo(errno.ParamErrCode, "用户不可以自己关注自己")
	}
	//3.follow exist ,db action
	folRelation := orm_gen.Follow{UserID: req.ToUserId, FollowerID: curUid.(int64)}
	followExist, _ := db.QueryFollowExist(folRelation.UserID, folRelation.FollowerID)
	if req.ActionType == FOLLOW {
		if followExist {
			return false, errno.FollowRelationAlreadyExistErr
		}
		flag, err = db.AddNewFollow(&folRelation)
	} else {
		if !followExist {
			return false, errno.FollowRelationNotExistErr
		}
		flag, err = db.DeleteFollow(&folRelation)
	}
	return flag, err
}

func (s *RelationService) GetRelationList(req *relation.RelationFollowListRequest, reqType int32) ([]*common.User, error) {
	uid := req.UserId
	var list []*common.User
	var uids []int64
	curUid, exist := s.c.Get("current_user_id")
	if !exist {
		curUid = int64(0)
	}
	exist, err := db.CheckUserExistById(uid)
	if err != nil {
		return nil, err
	}
	if !exist {
		return nil, errno.UserIsNotExistErr
	}
	if reqType == FOLLOWING {
		uids, err = db.GetFollowIdList(uid)
	} else {
		uids, err = db.GetFollowerIdList(uid)
	}
	if err != nil {
		return nil, err
	}
	for _, uid := range uids {
		user_info, err := service.NewUserService(s.ctx, s.c).GetUserInfo(uid, curUid.(int64))
		if err != nil {
			return nil, err
		}
		list = append(list, user_info)
	}
	return list, nil
}
