package v1

import (
	"github.com/gin-gonic/gin"
	"go-talk-talk/models"
	"go-talk-talk/pkg/errorInfo"
	"net/http"
)

func GetUserList(c *gin.Context) {
	errorCode := errorInfo.SUCCESS
	userList := models.GetUserListAllByTime(100, true)
	c.JSON(http.StatusOK, gin.H{
		"code":    errorCode,
		"data":    userList,
		"message": "验证码尚未过期，60秒后再试",
	})
}

func GetRandomUser(c *gin.Context) {
	errorCode := errorInfo.SUCCESS
	userInfo := models.GetRandomInfo()
	userInfo.Password = "123456"
	c.JSON(http.StatusOK, gin.H{
		"code":    errorCode,
		"data":    userInfo,
		"message": "验证码尚未过期，60秒后再试",
	})
}
