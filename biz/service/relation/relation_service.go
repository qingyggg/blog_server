package service

import (
	"context"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/qingyggg/blog_server/biz/dal/db"
	"github.com/qingyggg/blog_server/biz/model/hertz/common"
	"github.com/qingyggg/blog_server/biz/model/hertz/social/relation"
	"github.com/qingyggg/blog_server/biz/model/orm_gen"
	service_utils "github.com/qingyggg/blog_server/biz/service"
	service "github.com/qingyggg/blog_server/biz/service/user"
	"github.com/qingyggg/blog_server/pkg/errno"
	"github.com/qingyggg/blog_server/pkg/utils"
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

func (s *RelationService) FollowAction(req *relation.RelationActionRequest) error {
	//1.check user
	exist, err := db.CheckUserExistByHashId(req.UhashID)
	if err != nil {
		return err
	}
	if exist != true {
		return errno.UserIsNotExistErr
	}
	//2.check param user id != to user id,action type==1 or 2
	//获取用户hashid
	user, err := db.QueryUserById(service_utils.GetUid(s.c))
	if err != nil {
		return err
	}
	curUidByte := user.HashID
	curUid := utils.ConvertByteHashToString(curUidByte)
	followedId := req.UhashID
	followedIdByte := utils.ConvertStringHashToByte(req.UhashID)

	if req.UhashID == curUid {
		return errno.NewErrNo(errno.ParamErrCode, "用户不可以自己关注自己")
	}
	//3.follow exist ,db action
	folRelation := orm_gen.Follow{UserID: followedIdByte, FollowerID: curUidByte}
	followExist, _ := db.QueryFollowExist(followedId, curUid)
	if req.ActionType == FOLLOW {
		if followExist {
			return errno.FollowRelationAlreadyExistErr
		}
		err = db.AddNewFollow(&folRelation)
	} else {
		if !followExist {
			return errno.FollowRelationNotExistErr
		}
		err = db.DeleteFollow(&folRelation)
	}
	return err
}

func (s *RelationService) GetRelationList(req *relation.RelationFollowListRequest, reqType int32) ([]*common.User, error) {
	uid := req.UhashID
	var list []*common.User
	var uids []string
	exist, err := db.CheckUserExistByHashId(uid)
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
		user_info, err := service.NewUserService(s.ctx, s.c).GetUserInfo(uid)
		if err != nil {
			return nil, err
		}
		list = append(list, user_info)
	}
	return list, nil
}
