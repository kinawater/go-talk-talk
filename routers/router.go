package routers

import (
	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	r := gin.New()
	r.Use(gin.Logger())
	// panic 恢复 中间件
	r.Use(gin.Recovery())

}
