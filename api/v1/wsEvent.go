package v1

import socketio "github.com/googollee/go-socket.io"

func JoinChat(server *socketio.Server, nameSpace string) {
	server.OnEvent(nameSpace, "joinChat", func(s socketio.Conn) {
		// 获取用户的sendId
		//userInfo =
	})
}
