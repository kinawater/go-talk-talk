package websocketIO

import (
	"fmt"
	"github.com/googollee/go-socket.io"
	"github.com/googollee/go-socket.io/engineio"
	"github.com/googollee/go-socket.io/engineio/transport"
	"github.com/googollee/go-socket.io/engineio/transport/polling"
	"github.com/googollee/go-socket.io/engineio/transport/websocket"
	"go-talk-talk/pkg/logger"
	"log"
	"net/http"
	"time"
)

func WsHandle() *socketio.Server {
	var socketconfig = &engineio.Options{
		PingTimeout:  7 * time.Second,
		PingInterval: 5 * time.Second,
		Transports: []transport.Transport{
			&polling.Transport{
				Client: &http.Client{
					Timeout: time.Minute,
				},
				CheckOrigin: func(r *http.Request) bool {
					return true
				},
			},
			&websocket.Transport{
				CheckOrigin: func(r *http.Request) bool {
					return true
				},
			},
		},
	}

	server := socketio.NewServer(socketconfig)
	// 连接成功
	server.OnConnect("/", func(s socketio.Conn) error {
		s.SetContext("")
		s.Join("bcast")
		log.Println("连接成功：", s.ID())
		return nil
	})

	// 处理event
	server.OnEvent("/", "msg", func(s socketio.Conn, msg string) string {
		s.SetContext(msg)
		fmt.Println("getMsg=====chat====>", msg, "  当前sid", s.ID())
		s.Emit("message", "收到了： "+msg)
		return "recv " + msg
	})
	// 连接错误
	server.OnError("/", func(s socketio.Conn, e error) {
		log.Println("连接错误:", e)
		logger.Info("连接错误:", e)

	})
	// 关闭连接
	server.OnDisconnect("/", func(s socketio.Conn, reason string) {
		log.Println("关闭连接：", reason)
		log.Println("当前时间：", time.Now())
	})
	return server
}
