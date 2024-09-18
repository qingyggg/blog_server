package db

import (
	"github.com/qingyggg/blog_server/biz/model/orm_gen"
	"github.com/qingyggg/blog_server/biz/model/query"
	"github.com/qingyggg/blog_server/biz/mw/redis"
)

// register redis operate strategy
var rdFollows redis.Follows

func AddNewFollow(follow *orm_gen.Follow) (bool, error) {
	var f = query.Follow
	err := f.Create(follow)
	if err != nil {
		return false, err
	}
	// add data to redis
	if rdFollows.CheckFollow(follow.FollowerID) {
		rdFollows.AddFollow(follow.UserID, follow.FollowerID)
	}
	if rdFollows.CheckFollower(follow.UserID) {
		rdFollows.AddFollower(follow.UserID, follow.FollowerID)
	}

	return true, nil
}

// DeleteFollow delete follow relation in db and update redis
func DeleteFollow(follow *orm_gen.Follow) (bool, error) {
	var f = query.Follow
	_, err := f.Where(f.UserID.Eq(follow.UserID), f.FollowerID.Eq(follow.FollowerID)).Delete()
	if err != nil {
		return false, err
	}
	// if redis hit del
	if rdFollows.CheckFollow(follow.FollowerID) {
		rdFollows.DelFollow(follow.UserID, follow.FollowerID)
	}
	if rdFollows.CheckFollower(follow.UserID) {
		rdFollows.DelFollower(follow.UserID, follow.FollowerID)
	}
	return true, nil
}

// QueryFollowExist check the relation of user and follower
func QueryFollowExist(user_id, follower_id int64) (bool, error) {
	var f = query.Follow
	if rdFollows.CheckFollow(follower_id) {
		return rdFollows.ExistFollow(user_id, follower_id), nil
	}
	if rdFollows.CheckFollower(user_id) {
		return rdFollows.ExistFollower(user_id, follower_id), nil
	}

	count, err := f.Where(f.UserID.Eq(user_id), f.FollowerID.Eq(follower_id)).Count()
	if err != nil {
		return false, err
	} else if count != 0 {
		return true, nil
	} else {
		return false, nil
	}
}

// GetFollowCount query the number of users following
func GetFollowCount(follower_id int64) (int64, error) {
	if rdFollows.CheckFollow(follower_id) {
		return rdFollows.CountFollow(follower_id)
	}

	// Not in the cache, go to the database to find and update the cache
	followings, err := getFollowIdList(follower_id)
	if err != nil {
		return 0, err
	}
	// update redis asynchronously
	go addFollowRelationToRedis(follower_id, followings)
	return int64(len(followings)), nil
}

// addFollowRelationToRedis update redis.RdbFollowing
func addFollowRelationToRedis(follower_id int64, followings []int64) {
	for _, following := range followings {
		rdFollows.AddFollow(following, follower_id)
	}
}

// GetFollowerCount query the number of followers of a user
func GetFollowerCount(user_id int64) (int64, error) {
	if rdFollows.CheckFollower(user_id) {
		return rdFollows.CountFollower(user_id)
	}
	// Not in the cache, go to the database to find and update the cache
	followers, err := getFollowerIdList(user_id)
	if err != nil {
		return 0, err
	}
	// update redis asynchronously
	go addFollowerRelationToRedis(user_id, followers)
	return int64(len(followers)), nil
}

// addFollowerRelationToRedis update redis.RdbFollower
func addFollowerRelationToRedis(user_id int64, followers []int64) {
	for _, follower := range followers {
		rdFollows.AddFollower(user_id, follower)
	}
}

// getFollowIdList find user_id follow id list in db
func getFollowIdList(follower_id int64) ([]int64, error) {
	var f = query.Follow
	follows, err := f.Where(f.FollowerID.Eq(follower_id)).Find()
	if err != nil {
		return nil, err
	}
	var result []int64
	for _, v := range follows {
		result = append(result, v.UserID)
	}
	return result, nil
}

// GetFollowIdList find user_id follow id list in db or rdb
func GetFollowIdList(follower_id int64) ([]int64, error) {
	if rdFollows.CheckFollow(follower_id) {
		return rdFollows.GetFollow(follower_id), nil
	}
	return getFollowIdList(follower_id)
}

// getFollowerIdList get follower id list in db
func getFollowerIdList(user_id int64) ([]int64, error) {
	var f = query.Follow
	follows, err := f.Where(f.UserID.Eq(user_id)).Find()
	if err != nil {
		return nil, err
	}
	var result []int64
	for _, v := range follows {
		result = append(result, v.FollowerID)
	}
	return result, nil
}

// GetFollowerIdList get follower id list in db or rdb
func GetFollowerIdList(user_id int64) ([]int64, error) {
	if rdFollows.CheckFollower(user_id) {
		return rdFollows.GetFollower(user_id), nil
	}
	return getFollowerIdList(user_id)
}

func GetFriendIdList(user_id int64) ([]int64, error) {
	if !rdFollows.CheckFollow(user_id) {
		following, err := getFollowIdList(user_id)
		if err != nil {
			return *new([]int64), err
		}
		addFollowRelationToRedis(user_id, following)
	}
	if !rdFollows.CheckFollower(user_id) {
		followers, err := getFollowerIdList(user_id)
		if err != nil {
			return *new([]int64), err
		}
		addFollowerRelationToRedis(user_id, followers)
	}
	return rdFollows.GetFriend(user_id), nil
}
