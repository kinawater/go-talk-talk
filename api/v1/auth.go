package v1

import (
	"errors"
	"github.com/astaxie/beego/validation"
	"github.com/gin-gonic/gin"
	"go-talk-talk/models"
	"go-talk-talk/pkg/errorInfo"
	"golang.org/x/crypto/bcrypt"
	"log"
	"net/http"
)

type auth struct {
	Username string `valid:"Required;MaxSize(50)"`
	Password string `valid:"Required;MaxSize(50)"`
}

func Register(c *gin.Context) {
	var errorCode = errorInfo.SUCCESS
	username, password, err := checkUserNameAndPassWordFromForm(c)
	avatarId, avatarOk := c.GetPostForm("AvatarId")
	if err != nil {
		return
	}
	// 查询用户是否已经存在
	userInfo := models.GetUserByName(username)
	if userInfo.ID > 0 {
		errorCode = errorInfo.ERROR_EXIST_USER
		c.JSON(http.StatusOK, gin.H{
			"code": errorCode,
			"msg":  errorInfo.GetMsg(errorCode),
		})
		return
	}
	screatPassword, _ := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	var user models.Users
	user.Username = username
	user.Password = string(screatPassword)
	if avatarOk {
		user.AvatarId = avatarId
	} else {
		// 默认头像是1
		user.AvatarId = "1"
	}
}

func Login(c *gin.Context) {
	var errorCode = errorInfo.SUCCESS
	username, password, err := checkUserNameAndPassWordFromForm(c)
	if err != nil {
		return
	}

	userInfo := models.GetUserByName(username)

	if userInfo.ID > 0 {
		// 判断密码是否正确
		err := bcrypt.CompareHashAndPassword([]byte(userInfo.Password), []byte(password))
		if err != nil {
			// 密码错误
			errorCode = errorInfo.ERROR_AUTH_NO_USERNAME_OR_PASSWORD
			c.JSON(http.StatusOK, gin.H{
				"code": errorCode,
				"msg":  errorInfo.GetMsg(errorCode),
			})
			return
		} else {
			// 生成jwt，存入cookie
			// TODO
			c.JSON(http.StatusOK, gin.H{
				"code": errorCode,
				"msg":  errorInfo.GetMsg(errorCode),
			})
		}

	} else {
		errorCode = errorInfo.ERROR_NOT_EXIST_USER
		c.JSON(http.StatusOK, gin.H{
			"code": errorCode,
			"msg":  errorInfo.GetMsg(errorCode),
		})
		return
	}
	return
}

func checkUserNameAndPassWordFromForm(c *gin.Context) (string, string, error) {
	var errorCode int
	username, ok := c.GetPostForm("username")
	password, ok := c.GetPostForm("password")

	if !ok {
		errorCode = errorInfo.ERROR_AUTH_NO_USERNAME_OR_PASSWORD
		c.JSON(http.StatusOK, gin.H{
			"code": errorCode,
			"msg":  errorInfo.GetMsg(errorCode),
		})
		return "", "", errors.New(errorInfo.GetMsg(errorCode))
	}
	// 验证用户名和密码
	valid := validation.Validation{}
	authInfo := auth{Username: username, Password: password}
	ok, _ = valid.Valid(&authInfo)
	if !ok {
		for _, err := range valid.Errors {
			log.Println(err.Key, err.Message)
		}
		errorCode = errorInfo.INVALID_PARAMS
		c.JSON(http.StatusOK, gin.H{
			"code": errorCode,
			"msg":  errorInfo.GetMsg(errorCode),
		})
		return "", "", errors.New(errorInfo.GetMsg(errorCode))
	}
	return username, password, nil
}
