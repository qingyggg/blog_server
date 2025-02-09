package socket

import (
	"fmt"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	socketio "github.com/googollee/go-socket.io"
	"github.com/qingyggg/blog_server/biz/mw/redis"
)

var Server *socketio.Server
var rdbOnline = new(redis.Online)
var HertzCtx *app.RequestContext
var HertzCtxIsAssigned bool

func Init() {
	Server = socketio.NewServer(nil)

	Server.OnConnect("/", func(s socketio.Conn) error {
		hlog.Info("socket.io==>connected,ID:", s.ID())
		return nil
	})
	//设置用户在线状态，客户端连接后，优先调用该请求,否则服务端推送不了信息
	Server.OnEvent("/", "set_online", func(s socketio.Conn, uid string) string {
		err := rdbOnline.SetOnline(uid, s.ID())
		fmt.Println(111)
		if err != nil {
			s.Close() //状态设置失败，主动断开连接
			hlog.Error("socket.io==>无法设置用户在线状态,ID:", s.ID(), "Reason:", err.Error())
			return ""
		}
		s.SetContext(uid)
		s.Join(s.ID()) // 将客户端加入以其 ID 命名的房间
		hlog.Info("socket.io==>用户已在线,ID:", s.ID(), "uid:", uid)
		return ConvertResToJson(GetBaseResponse())
	})

	Server.OnError("/", func(s socketio.Conn, e error) {
		hlog.Error("socket.io==>客户端发生了错误,ID:", s.ID(), "Reason:", e.Error())
	})

	Server.OnDisconnect("/", func(s socketio.Conn, reason string) {
		if s.Context() != nil {
			err := rdbOnline.RemOnline(s.Context().(string)) //uHashId
			if err != nil {
				hlog.Error("socket.io==>移除用户在线状态失败,ID:", s.ID(), "Reason:", err.Error())
				return
			}
		}
		s.LeaveAll()
		hlog.Info("socket.io==>disconnected,ID:", s.ID(), ",Reason:", reason)
	})

	//处理子socket请求
	initNotify()

	// 启动协程处理 Socket.IO
	go func() {
		if err := Server.Serve(); err != nil {
			hlog.Fatalf("Socket.IO listen error: %v", err)
			return
		}
		hlog.Info("socket io初始化成功")
	}()
}
