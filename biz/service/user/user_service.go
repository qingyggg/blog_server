package service

import (
	"context"
	"github.com/qingyggg/blog_server/biz/dal/db"
	"github.com/qingyggg/blog_server/biz/model/hertz/basic/user"
	"github.com/qingyggg/blog_server/biz/model/hertz/common"
	"github.com/qingyggg/blog_server/biz/model/orm_gen"
	"github.com/qingyggg/blog_server/pkg/constants"
	"github.com/qingyggg/blog_server/pkg/errno"
	"github.com/qingyggg/blog_server/pkg/utils"
	"sync"

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
func (s *UserService) UserRegister(req *user.UserActionRequest) (user_id int64, err error) {
	isExist, err := db.CheckUserExistByUname(req.Username)
	if err != nil {
		return 0, err
	}
	if isExist {
		return 0, errno.UserAlreadyExistErr
	}
	passWord, err := utils.Crypt(req.Password)
	user_id, err = db.CreateUser(&orm_gen.User{
		UserName:        req.Username,
		Password:        passWord,
		Avatar:          constants.TestAva,
		BackgroundImage: constants.TestBackground,
		Signature:       constants.TestSign,
	})
	return user_id, nil
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
	currentUserId, exists := s.c.Get("current_user_id")
	if !exists {
		currentUserId = 0
	}
	return s.GetUserInfo(queryUserId, currentUserId.(int64))
}

func (s *UserService) UserProfileModify(req *user.UserActionProfileModifyRequest) error {
	uid := req.UserId
	profile := make(map[string]string)
	profile["signature"] = req.User.Signature
	profile["avatar"] = req.User.Avatar
	profile["background_image"] = req.User.BackgroundImage
	err := db.UserProfileModify(uid, profile)
	if err != nil {
		return err
	}
	return nil
}

func (s *UserService) GetUserInfo(queryUserId int64, userId int64) (*common.User, error) {
	u := &common.User{}
	errChan := make(chan error, 7)
	defer close(errChan)
	var wg sync.WaitGroup
	wg.Add(5)
	go func() {
		dbUser, err := db.QueryUserById(queryUserId)
		if err != nil {
			errChan <- err
		} else {
			u.Name = dbUser.UserName
			u.Profile.Avatar = utils.URLconvert(s.ctx, s.c, dbUser.Avatar)
			u.Profile.BackgroundImage = utils.URLconvert(s.ctx, s.c, dbUser.BackgroundImage)
			u.Profile.Signature = dbUser.Signature
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
		if userId != 0 {
			IsFollow, err := db.QueryFollowExist(userId, queryUserId)
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
		return &common.User{}, result
	default:
	}
	u.Id = queryUserId
	return u, nil
}
