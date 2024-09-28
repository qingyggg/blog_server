package db

import (
	"github.com/qingyggg/blog_server/biz/model/orm_gen"
	"github.com/qingyggg/blog_server/biz/model/query"
	"github.com/qingyggg/blog_server/pkg/errno"
	"github.com/qingyggg/blog_server/pkg/utils"
	"golang.org/x/exp/constraints"
)

//type User

// CreateUser create user info
func CreateUser(user *orm_gen.User) error {
	var u = query.User
	err := u.Create(user)
	return err
}

// QueryUser query User by user_name
func QueryUser(userName string) (*orm_gen.User, error) {
	var u = query.User
	user, err := u.Where(u.UserName.Eq(userName)).Take()
	if err != nil {
		return nil, err
	}
	if user.ID == 0 {
		err := errno.UserIsNotExistErr
		return nil, err
	}
	return user, nil
}

// QueryUserById get user in the database by user id
func QueryUserById(userId int64) (*orm_gen.User, error) {
	var u = query.User
	user, err := u.Where(u.ID.Eq(userId)).Take()
	if err != nil {
		return nil, err
	}
	if user.ID == 0 {
		err := errno.UserIsNotExistErr
		return nil, err
	}
	return user, nil
}

func QueryUserByHashId(uHashId string) (*orm_gen.User, error) {
	var u = query.User
	user, err := u.Where(u.HashID.Eq(utils.ConvertStringHashToByte(uHashId))).Take()
	if err != nil {
		return nil, err
	}
	if user.ID == 0 {
		err := errno.UserIsNotExistErr
		return nil, err
	}
	return user, nil
}

// VerifyUser verify username and password in the db
func VerifyUser(userName string, password string) (int64, error) {
	user, err := QueryUser(userName)
	if err != nil {
		if err.Error() == "record not found" {
			return 0, errno.UserIsNotExistErr
		}
		return 0, err
	}
	if ok := utils.VerifyPassword(password, user.Password); !ok {
		err = errno.PasswordIsNotVerified
		return 0, err
	} else {
		return user.ID, nil
	}
}

// CheckUserExistById find if user exists
func CheckUserExistById(userId int64) (bool, error) {
	var u = query.User
	count, err := u.Where(u.ID.Eq(userId)).Count()
	if err != nil {
		return false, err
	} else {
		if count == 1 {
			return true, nil
		} else {
			return false, nil
		}
	}
}

func CheckUserExistByHashId(hashId string) (bool, error) {
	var u = query.User
	count, err := u.Where(u.HashID.Eq(utils.ConvertStringHashToByte(hashId))).Count()
	if err != nil {
		return false, err
	} else {
		if count == 1 {
			return true, nil
		} else {
			return false, nil
		}
	}
}

func CheckUserExistByUname(username string) (bool, error) {
	var u = query.User
	count, err := u.Where(u.UserName.Eq(username)).Count()
	if err != nil {
		return false, err
	} else {
		if count == 1 {
			return true, nil
		} else {
			return false, nil
		}
	}
}

func UserPwdModify(uid int64, new_pwd string) error {
	var u = query.User
	_, err := u.Where(u.ID.Eq(uid)).Update(u.Password, new_pwd)
	if err != nil {
		return err
	}
	return nil
}

func UserProfileModify(user_id int64, payload map[string]interface{}) error {
	var u = query.User
	_, err := u.Where(u.ID.Eq(user_id)).Updates(payload)
	return err //err =err or err=nil
}

// map[user id] : user payload
func QueryUserByIds(uids []int64) (*map[int64]*orm_gen.User, error) {
	var u = query.User
	//对uids进行去重操作
	uMaps := make(map[int64]*orm_gen.User)
	var uniqueIDs []int64
	for _, uid := range uids {
		uMaps[uid] = new(orm_gen.User)
	}
	for k := range uMaps {
		uniqueIDs = append(uniqueIDs, k)
	}
	users, err := u.Where(u.ID.In(uniqueIDs...)).Find()
	if err != nil {
		return nil, err
	}
	for _, cu := range users {
		uMaps[cu.ID] = cu
	}

	return &uMaps, nil
}

func QueryUserByHashIds(uids []string) (uMaps map[string]*orm_gen.User, err error) {
	var u = query.User
	var uidsByte [][]byte
	uMaps = make(map[string]*orm_gen.User)
	//对uids进行去重操作
	uniqueIDs := trimIds(uids)
	for _, v := range uniqueIDs {
		uidsByte = append(uidsByte, utils.ConvertStringHashToByte(v))
	}
	users, err := u.Where(u.HashID.In(uidsByte...)).Find()
	if err != nil {
		return nil, err
	}
	for _, cu := range users {
		uMaps[utils.ConvertByteHashToString(cu.HashID)] = cu
	}

	return uMaps, nil
}

func trimIds[T constraints.Ordered](uids []T) (uniqueIDs []T) {
	//对uids进行去重操作
	uMaps := make(map[T]interface{})
	for _, uid := range uids {
		uMaps[uid] = new(orm_gen.User)
	}
	for k := range uMaps {
		uniqueIDs = append(uniqueIDs, k)
	}
	return
}
