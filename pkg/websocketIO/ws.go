package websocketIO

import (
	"fmt"
	"github.com/googollee/go-socket.io"
)

func wsHandle() *socketio.Server {
	server := socketio.NewServer(nil)
	// 连接成功
	server.OnConnect("/", func(s socketio.Conn) error {
		s.SetContext("")
		// 申请一个房间
		s.Join("bcast")
		fmt.Println("连接成功：", s.ID())
		return nil
	})
}
