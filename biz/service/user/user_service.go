package service

import (
	"context"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/qingyggg/blog_server/biz/dal/db"
	"github.com/qingyggg/blog_server/biz/model/hertz/basic/user"
	"github.com/qingyggg/blog_server/biz/model/hertz/common"
	"github.com/qingyggg/blog_server/biz/model/orm_gen"
	"github.com/qingyggg/blog_server/biz/mw/minio"
	"github.com/qingyggg/blog_server/pkg/constants"
	"github.com/qingyggg/blog_server/pkg/errno"
	"github.com/qingyggg/blog_server/pkg/utils"
	"sync"
	"time"

	"github.com/cloudwego/hertz/pkg/app"
)

type UserService struct {
	ctx context.Context
	c   *app.RequestContext
}

// NewUserService create user service
func NewUserService(ctx context.Context, c *app.RequestContext) *UserService {
	return &UserService{ctx: ctx, c: c}
}

// UserRegister register user return user id.
func (s *UserService) UserRegister(req *user.UserActionRequest) (uHashId string, err error) {
	isExist, err := db.CheckUserExistByUname(req.Username)
	uHashId = utils.GetSHA256String(req.Username + time.Now().String())[:16] //截取哈希值的前16为作为用户的hashId
	if err != nil {
		return "", err
	}
	if isExist {
		return "", errno.UserAlreadyExistErr
	}
	passWord, err := utils.Crypt(req.Password)
	err = db.CreateUser(&orm_gen.User{
		UserName:        req.Username,
		Password:        passWord,
		HashID:          utils.ConvertStringHashToByte(uHashId),
		Avatar:          constants.TestAva,
		BackgroundImage: constants.TestBackground,
		Signature:       constants.TestSign,
	})
	if err != nil {
		return uHashId, err
	}
	err = initUserSpace(s.ctx, uHashId)
	if err != nil {
		return "", err
	}
	return uHashId, nil
}

// initUserSpace 为新建立的用户分配minio存储空间
func initUserSpace(ctx context.Context, uHashId string) error {
	buckName := "user-" + uHashId //username:userid
	err := minio.MakeBucket(ctx, buckName)
	if err != nil {
		hlog.Error("初始化用户桶失败")
		return err
	}
	return nil
}

func (s *UserService) PwdModify(req *user.UserActionPwdModifyRequest) (userId int64, err error) {
	uid, err := db.VerifyUser(req.Username, req.OldPassword)
	if err != nil {
		return 0, err
	}
	crptedPwd, err := utils.Crypt(req.NewPassword)
	if err != nil {
		return 0, err
	}
	err = db.UserPwdModify(uid, crptedPwd)
	if err != nil {
		return 0, err
	}
	return uid, nil
}

// UserInfo the function of user api
func (s *UserService) UserInfo(req *user.UserRequest) (*common.User, error) {
	queryUserId := req.UserId
	//
	exist, err := db.CheckUserExistById(req.UserId)
	if err != nil {
		return nil, err
	}
	if !exist {
		return nil, errno.UserIsNotExistErr
	}
	currentUserId, exists := s.c.Get("current_user_id")
	if !exists {
		currentUserId = 0
	}
	return s.GetUserInfo(queryUserId, currentUserId.(int))
}

func (s *UserService) UserProfileModify(req *user.UserActionProfileModifyRequest) error {
	uid := req.UserId
	profile := map[string]interface{}{
		"signature":        req.User.Signature,
		"avatar":           req.User.Avatar,
		"background_image": req.User.BackgroundImage,
	}

	// 过滤空字符串的字段
	for key, value := range profile {
		if value == "" {
			delete(profile, key)
		}
	}
	err := db.UserProfileModify(uid, profile)
	if err != nil {
		return err
	}
	return nil
}

func (s *UserService) GetUserInfo(queryUserId int64, userId int) (*common.User, error) {
	u := new(common.User)
	errChan := make(chan error, 5)
	defer close(errChan)
	var wg sync.WaitGroup
	wg.Add(5)

	go func() { //
		dbUser, err := db.QueryUserById(queryUserId)
		if err != nil {
			errChan <- err
		} else {
			u.Base = new(common.UserBase)
			u.Base.Name = dbUser.UserName
			u.Base.Profile = new(common.UserProfile)
			u.Base.Profile.Avatar = utils.URLconvert(s.ctx, s.c, dbUser.Avatar)
			u.Base.Profile.BackgroundImage = utils.URLconvert(s.ctx, s.c, dbUser.BackgroundImage)
			u.Base.Profile.Signature = dbUser.Signature
			u.Base.HashId = utils.ConvertByteHashToString(dbUser.HashID)
		}
		wg.Done()
	}()

	go func() {
		WorkCount, err := db.GetWorkCount(queryUserId)
		if err != nil {
			errChan <- err
		} else {
			u.WorkCount = WorkCount
		}
		wg.Done()
	}()

	go func() {
		FollowCount, err := db.GetFollowCount(queryUserId)
		if err != nil {
			errChan <- err
			return
		} else {
			u.FollowCount = FollowCount
		}
		wg.Done()
	}()

	go func() {
		FollowerCount, err := db.GetFollowerCount(queryUserId)
		if err != nil {
			errChan <- err
		} else {
			u.FollowerCount = FollowerCount
		}
		wg.Done()
	}()

	go func() {
		if userId != 0 && int64(userId) != queryUserId {
			IsFollow, err := db.QueryFollowExist(int64(userId), queryUserId)
			if err != nil {
				errChan <- err
			} else {
				u.IsFollow = IsFollow
			}
		} else {
			u.IsFollow = false
		}
		wg.Done()
	}()

	wg.Wait()
	select {
	case result := <-errChan:
		return nil, result
	default:
	}
	u.Base.Id = queryUserId
	return u, nil
}
