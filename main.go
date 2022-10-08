package main

import (
	"fmt"
	"go-talk-talk/config"
	_ "go-talk-talk/database"
	"go-talk-talk/routers"
	"net/http"
)

func main() {
	// 初始化路由
	router := routers.InitRouter()
	// 服务器
	server := &http.Server{
		Addr:           fmt.Sprintf(":%d", config.ServerConf.HTTPPort),
		Handler:        router,
		ReadTimeout:    config.ServerConf.ReadTimeout,
		WriteTimeout:   config.ServerConf.WriteTimeout,
		MaxHeaderBytes: 1 << 20,
	}
	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}
