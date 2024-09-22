package service

import (
	"context"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/qingyggg/blog_server/biz/dal/db"
	"github.com/qingyggg/blog_server/biz/model/hertz/basic/user"
	"github.com/qingyggg/blog_server/biz/model/hertz/common"
	"github.com/qingyggg/blog_server/biz/model/orm_gen"
	"github.com/qingyggg/blog_server/biz/mw/minio"
	service_utils "github.com/qingyggg/blog_server/biz/service"
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
	//检测用户是否存在
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
	//后续开发网盘功能的时候用
	//err = initUserSpace(s.ctx, uHashId)
	//if err != nil {
	//	return "", err
	//}
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

func (s *UserService) PwdModify(req *user.UserActionPwdModifyRequest) (uid int64, uHashId string, err error) {
	uid, uHashId, err = db.VerifyUser(req.Username, req.OldPassword) //验证用户是否存在，旧密码是否正确
	if err != nil {
		return 0, "", err
	}
	crptedPwd, err := utils.Crypt(req.NewPassword)
	if err != nil {
		return 0, "", err
	}
	err = db.UserPwdModify(uid, crptedPwd)
	if err != nil {
		return 0, "", err
	}
	return uid, uHashId, nil
}

// UserInfo the function of user api
func (s *UserService) UserInfo(req *user.UserRequest) (*common.User, error) {
	queryUserId := req.UHashID
	if queryUserId != "current" {
		exist, err := db.CheckUserExistByHashId(queryUserId)
		if err != nil {
			return nil, err
		}
		if !exist {
			return nil, errno.UserIsNotExistErr
		}
	}

	return s.GetUserInfo(queryUserId)
}

func (s *UserService) UserProfileModify(req *user.UserActionProfileModifyRequest) (error, string, int64) {
	uid := service_utils.GetUid(s.c)
	user, err := db.QueryUserById(uid)
	if err != nil {
		return err, "", 0
	}
	profile := map[string]interface{}{
		"signature":        req.User.Signature,
		"avatar":           utils.UrlConvertReverse(s.ctx, req.User.Avatar),
		"background_image": utils.UrlConvertReverse(s.ctx, req.User.BackgroundImage),
	}
	// 过滤空字符串的字段
	for key, value := range profile {
		if value == "" {
			delete(profile, key)
		}
	}
	err = db.UserProfileModify(uid, profile)
	if err != nil {
		return err, "", 0
	}
	return nil, utils.ConvertByteHashToString(user.HashID), uid
}

func (s *UserService) GetUserInfo(queryUHashId string) (*common.User, error) {
	u := new(common.User)
	errChan := make(chan error, 4)
	defer close(errChan)
	var wg sync.WaitGroup
	wg.Add(4)

	uid := service_utils.GetUid(s.c)
	var isLoginedUser bool
	if queryUHashId == "current" {
		isLoginedUser = true
	} else {
		isLoginedUser = false
	}

	//判断请求的用户信息是否为登录的用户-->queryUHashId=='current'
	var dbUser *orm_gen.User
	var err error
	if isLoginedUser {
		dbUser, err = db.QueryUserById(uid)
	} else {
		dbUser, err = db.QueryUserByHashId(queryUHashId)
	}
	if err != nil {
		return nil, err
	} else {
		u.Base = new(common.UserBase)
		u.Base.Name = dbUser.UserName
		u.Base.Profile = new(common.UserProfile)
		u.Base.Profile.Avatar = utils.URLconvert(s.ctx, s.c, dbUser.Avatar)
		u.Base.Profile.BackgroundImage = utils.URLconvert(s.ctx, s.c, dbUser.BackgroundImage)
		u.Base.Profile.Signature = dbUser.Signature
		u.Base.HashId = utils.ConvertByteHashToString(dbUser.HashID)
	}

	go func() {
		WorkCount, err := db.GetWorkCount(queryUHashId)
		if err != nil {
			errChan <- err
			return
		} else {
			u.WorkCount = WorkCount
		}
		defer wg.Done()
	}()

	go func() {
		FollowCount, err := db.GetFollowCount(queryUHashId)
		if err != nil {
			errChan <- err
			return
		} else {
			u.FollowCount = FollowCount
		}
		defer wg.Done()
	}()

	go func() {
		FollowerCount, err := db.GetFollowerCount(queryUHashId)
		if err != nil {
			errChan <- err
			return
		} else {
			u.FollowerCount = FollowerCount
		}
		defer wg.Done()
	}()

	go func() {
		defer wg.Done()
		if isLoginedUser {
			return
		}
		user, err := db.QueryUserById(uid)
		if err != nil {
			errChan <- err
			return
		}
		uHashId := utils.ConvertByteHashToString(user.HashID)
		if uHashId != queryUHashId {
			IsFollow, err := db.QueryFollowExist(queryUHashId, uHashId)
			if err != nil {
				errChan <- err
				return
			} else {
				u.IsFollow = IsFollow
			}
		}
	}()

	wg.Wait()
	select {
	case result := <-errChan:
		return nil, result
	default:
	}
	return u, nil
}
