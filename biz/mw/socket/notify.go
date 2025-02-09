package socket

import (
	"context"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	socketio "github.com/googollee/go-socket.io"
	"github.com/qingyggg/blog_server/biz/dal/db"
	"github.com/qingyggg/blog_server/biz/mw/mongo"
	"github.com/qingyggg/blog_server/biz/mw/redis"
	publish_service "github.com/qingyggg/blog_server/biz/service/publish"
	user_service "github.com/qingyggg/blog_server/biz/service/user"
	"github.com/qingyggg/blog_server/pkg/errno"
)

var NotifyClientEventTypes = struct {
	Read        string //已读某个信息
	ReadAll     string //已读所有信息
	DelOne      string //删除
	DellAllRead string //删除所有已读的信息
	Fetch       string //获取所有通知
}{
	Read:        "notify_read",
	ReadAll:     "notify_read_all",
	DelOne:      "notify_delete_one",
	DellAllRead: "notify_delete_all_read",
	Fetch:       "notify_fetch",
}
var NotifyServerEventTypes = struct {
	FollowEvent      string
	CommentEvent     string
	ACollectEvent    string
	AFavoriteEvent   string
	CmtFavoriteEvent string
}{
	FollowEvent:      "notify_follow",
	CommentEvent:     "notify_comment",
	ACollectEvent:    "notify_article_collect",
	AFavoriteEvent:   "notify_article_favorite",
	CmtFavoriteEvent: "notify_cmt_favorite",
}

const NotifyServerEventName = "notify_server_event"

func initNotify() {
	//获取所有通知
	Server.OnEvent("/", NotifyClientEventTypes.Fetch, func(s socketio.Conn, uid string) string {
		ex, err := db.CheckUserExistByHashId(uid)
		if err != nil {
			return ConvertResToJson(GetErrResponse(err))
		}
		if !ex {
			return ConvertResToJson(GetErrResponse(errno.UserIsNotExistErr))
		}
		//获取所有通知
		err, ns := db.NotifyGetByUid(uid)
		if err != nil {
			return ConvertResToJson(GetErrResponse(err))
		}
		for _, notify := range ns {
			notifyPushToClient(notify)
		}
		return ConvertResToJson(GetBaseResponse())
	})
	//已读通知
	Server.OnEvent("/", NotifyClientEventTypes.Read, func(s socketio.Conn, nid string) string {
		err, ex := db.NotifyExist(nid)
		if err != nil {
			return ConvertResToJson(GetErrResponse(err))
		}
		if !ex {
			return ConvertResToJson(GetErrResponse(errno.NotifyIsNotExistErr))
		}
		err = db.NotifyRead(nid)
		if err != nil {
			return ConvertResToJson(GetErrResponse(err))
		}
		return ConvertResToJson(GetBaseResponse())
	})
	//已读所有
	Server.OnEvent("/", NotifyClientEventTypes.ReadAll, func(s socketio.Conn, uid string) string {
		ex, err := db.CheckUserExistByHashId(uid)
		if err != nil {
			return ConvertResToJson(GetErrResponse(err))
		}
		if !ex {
			return ConvertResToJson(GetErrResponse(errno.UserIsNotExistErr))
		}
		err = db.NotifyReadByUid(uid)
		if err != nil {
			return ConvertResToJson(GetErrResponse(err))
		}
		return ConvertResToJson(GetBaseResponse())
	})
	//删除通知
	Server.OnEvent("/", NotifyClientEventTypes.DelOne, func(s socketio.Conn, nid string) string {
		err, ex := db.NotifyExist(nid)
		if err != nil {
			return ConvertResToJson(GetErrResponse(err))
		}
		if !ex {
			return ConvertResToJson(GetErrResponse(errno.NotifyIsNotExistErr))
		}
		err = db.NotifyDelById(nid)
		if err != nil {
			return ConvertResToJson(GetErrResponse(err))
		}
		return ConvertResToJson(GetBaseResponse())
	})
	//删除所有已读通知
	Server.OnEvent("/", NotifyClientEventTypes.DellAllRead, func(s socketio.Conn, uid string) string {
		ex, err := db.CheckUserExistByHashId(uid)
		if err != nil {
			return ConvertResToJson(GetErrResponse(err))
		}
		if !ex {
			return ConvertResToJson(GetErrResponse(errno.UserIsNotExistErr))
		}
		err = db.NotifyDelByUid(uid)
		if err != nil {
			return ConvertResToJson(GetErrResponse(err))
		}
		return ConvertResToJson(GetBaseResponse())
	})
}

// NotifyPushToClient 推送信息函数
// 新增通知，推送到客户端
func NotifyPushToClient(nid string) {
	//从mongodb获取消息
	err, notify := db.NotifyTake(nid)
	if err != nil {
		hlog.Error("get notify failed!!!", err)
		return
	}
	//检查用户是否在线
	err, ex := userIsOnline(notify.UserID)
	if err != nil {
		hlog.Error("get notify failed!!!", err)
		return
	}
	if ex {
		notifyPushToClient(notify)
	}
}

func notifyPushToClient(notify *mongo.Notification) {
	rdbOnline := new(redis.Online)
	notifyTypes := mongo.NotifyTypes
	sockID := rdbOnline.GetSocketID(notify.UserID)
	var res *Response
	//进行消息处理
	//1.系统消息
	if notify.Type == notifyTypes.SystemNotification {
		res = GetResponse(&System{Message: notify.Data.Message, NID: notify.HashID, EventType: 0})
	} else {
		//2.用户消息
		//2.1.根据actionUid获取用户信息
		aUser, err := user_service.NewUserService(context.TODO(), HertzCtx).QueryUserBase(notify.Data.ActionUid)
		if err != nil {
			broadcastErrToRoom(sockID, err)
			return
		}
		switch notify.Data.EventType {
		case NotifyServerEventTypes.AFavoriteEvent:
			title, err := publish_service.NewPublishService(context.TODO(), HertzCtx).TakeTitle(notify.Data.TargetID)
			if err != nil {
				broadcastErrToRoom(sockID, err)
				return
			}
			res = GetResponse(&AFavorite{NID: notify.HashID, User: aUser, Aid: notify.Data.TargetID, ATitle: title, EventType: 1})
		case NotifyServerEventTypes.CommentEvent:
			cmt, err := db.GetCommentByCmtID(context.TODO(), notify.Data.TargetID)
			if err != nil {
				broadcastErrToRoom(sockID, err)
				return
			}
			res = GetResponse(&Commented{NID: notify.HashID, User: aUser, Comment: cmt, EventType: 4})
		case NotifyServerEventTypes.CmtFavoriteEvent:
			cmt, err := db.GetCommentByCmtID(context.TODO(), notify.Data.TargetID)
			if err != nil {
				broadcastErrToRoom(sockID, err)
				return
			}
			res = GetResponse(&CFavorite{NID: notify.HashID, User: aUser, Comment: cmt, EventType: 3})
		case NotifyServerEventTypes.FollowEvent:
			res = GetResponse(&Follow{NID: notify.HashID, User: aUser, EventType: 5})
		case NotifyServerEventTypes.ACollectEvent:
			title, err := publish_service.NewPublishService(context.TODO(), HertzCtx).TakeTitle(notify.Data.TargetID)
			if err != nil {
				broadcastErrToRoom(sockID, err)
				return
			}
			res = GetResponse(&AFavorite{NID: notify.HashID, User: aUser, Aid: notify.Data.TargetID, ATitle: title, EventType: 2})
		default:
			res = GetErrResponse(errno.ServiceErr.WithMessage("invalid event type"))
		}
	}
	Server.BroadcastToRoom("/", sockID, NotifyServerEventName, ConvertResToJson(res))
}

func broadcastErrToRoom(sockID string, err error) {
	Server.BroadcastToRoom("/", sockID, NotifyServerEventName, ConvertResToJson(GetErrResponse(err)))
}

// UserIsOnline 用以检查用户是否在线
func userIsOnline(uHashId string) (error, bool) {
	onl := new(redis.Online)
	err, ex := onl.OnlExist(uHashId)
	if err != nil {
		return err, false
	}
	return nil, ex
}
