package routers

import (
	"github.com/gin-gonic/gin"
	v1 "go-talk-talk/api/v1"
	"go-talk-talk/config"
	"go-talk-talk/middleware"
)

func InitRouter() *gin.Engine {
	r := gin.New()
	r.Use(gin.Logger())
	// panic 恢复 中间件
	r.Use(gin.Recovery())
	// 处理因复杂性增加导致的请求变options中间件
	r.Use(middleware.Cors())
	// 设置运行模式
	gin.SetMode(config.RunMode)

	apiV1 := r.Group("/api/v1")

	{
		// 注册
		apiV1.POST("/register", v1.Register)
		// 登录
		apiV1.POST("/login", v1.Login)
		// 发送邮件验证码
		apiV1.POST("/sendVerificationCode", v1.SendVerificationCode)
		// 验证验证码
		apiV1.POST("/checkVerificationCode", v1.CheckVerificationCode)
	}
	{
		// 用户列表
		apiV1.GET("/userList", v1.GetUserList)
	}
	return r
}
