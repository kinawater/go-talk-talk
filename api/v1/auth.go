package v1

import (
	"errors"
	"github.com/astaxie/beego/validation"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/gomodule/redigo/redis"
	"go-talk-talk/database/userCache"
	"go-talk-talk/global"
	"go-talk-talk/models"
	"go-talk-talk/pkg/SendEmail"
	"go-talk-talk/pkg/errorInfo"
	"go-talk-talk/pkg/logger"
	"go-talk-talk/pkg/randCode"
	"golang.org/x/crypto/bcrypt"
	"log"
	"net/http"
)

type auth struct {
	Name     string `valid:"Required;MaxSize(50)"`
	Password string `valid:"Required;MaxSize(50)"`
}
type loginByEmail struct {
	Email    string `valid:"Required;MaxSize(50)"`
	Password string `valid:"Required;MaxSize(50)"`
}
type sendVerificationCodeBody struct {
	Email     string `valid:"Required" json:"email"`
	Timestamp int64  `json:"timestamp"`
}
type checkVerificationCodeBody struct {
	Email            string `valid:"Required" json:"email"`
	VerificationCode string `valid:"Required" json:"verificationCode"`
	Timestamp        int64  `json:"timestamp"`
}
type registerBody struct {
	Name         string `json:"name"`
	Email        string `json:"email"`
	Password     string `json:"password"`
	SurePassword string `json:"surePassword"`
	Avatar       string `json:"avatar"`
}

func SendVerificationCode(c *gin.Context) {
	var errorCode = errorInfo.SUCCESS
	redisConn := global.RedisPool.Get()
	defer redisConn.Close()

	getBody := sendVerificationCodeBody{}
	err := c.ShouldBindBodyWith(&getBody, binding.JSON)
	if err != nil {
		logger.Error("解析参数失败")
		errorCode = errorInfo.ERROR
		c.JSON(http.StatusOK, gin.H{
			"code":    errorCode,
			"data":    nil,
			"message": "验证码尚未过期，60秒后再试",
		})
		return
	}
	emailValue := getBody.Email
	var redisKey = userCache.UserEmailCaptchaKey(emailValue)

	// 验证码是否过期
	isKeyExists, err := redis.Bool(redisConn.Do("EXISTS", redisKey))
	if err != nil {
		logger.Error(redisKey, " set error,error info", err)
	}
	if isKeyExists {
		logger.Info(redisKey, " not exists ,验证码已经过期了")
		errorCode = errorInfo.ERROR_EMAIL_SEND_FAIL
		c.JSON(http.StatusOK, gin.H{
			"code":    errorCode,
			"data":    nil,
			"message": "验证码尚未过期",
		})
		return
	}
	// 生成验证码
	verificationCode := string(randCode.GetRandCode(emailValue, 6))

	// 通过邮箱发送验证码
	err = SendEmail.SendEmail([]string{emailValue}, "你好，这是你的验证码请查收", []byte("你的验证码是："+string(verificationCode)+" 祝您生活愉快"))
	if err != nil {
		errorCode = errorInfo.ERROR_EMAIL_SEND_FAIL
		c.JSON(http.StatusOK, gin.H{
			"code":    errorCode,
			"data":    nil,
			"message": errorInfo.GetMsg(errorCode),
		})
		return
	}
	// 存入redis
	outTime := 301 //过期时间多1秒是怕网络延迟
	_, err = redisConn.Do("SETEX", redisKey, outTime, verificationCode)
	if err != nil {
		errorCode = errorInfo.ERROR_EMAIL_SEND_FAIL
		c.JSON(http.StatusOK, gin.H{
			"code":    errorCode,
			"data":    nil,
			"message": errorInfo.GetMsg(errorCode),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code":    200,
		"data":    nil,
		"message": "发送验证码成功",
	})
	return
}

func CheckVerificationCode(c *gin.Context) {
	var errorCode = errorInfo.SUCCESS
	redisConn := global.RedisPool.Get()
	defer redisConn.Close()
	getBody := checkVerificationCodeBody{}
	err := c.ShouldBindBodyWith(&getBody, binding.JSON)
	if err != nil {
		logger.Error("解析参数失败")
		errorCode = errorInfo.ERROR
		c.JSON(http.StatusOK, gin.H{
			"code":    errorCode,
			"data":    nil,
			"message": "验证码尚未过期，60秒后再试",
		})
		return
	}
	var redisKey = userCache.UserEmailCaptchaKey(getBody.Email)
	// 验证码验证
	VerificationCode, err := redis.String(redisConn.Do("GET", redisKey))
	if err != nil {
		logger.Error(redisKey, " set error,error info", err)
	}
	logger.Info(VerificationCode, getBody.VerificationCode, VerificationCode == getBody.VerificationCode)
	if VerificationCode == getBody.VerificationCode {
		c.JSON(http.StatusOK, gin.H{
			"code":    errorCode,
			"data":    nil,
			"message": "验证成功",
		})
	} else {
		errorCode = errorInfo.ERROR_VERIFI_CODE_CHECK_FAIL
		c.JSON(http.StatusOK, gin.H{
			"code":    errorCode,
			"data":    nil,
			"message": errorInfo.GetMsg(errorCode),
		})
	}
	return

}
func Register(c *gin.Context) {
	var errorCode = errorInfo.SUCCESS
	userRegisterBody := registerBody{}
	err := c.ShouldBindBodyWith(&userRegisterBody, binding.JSON)
	err = checkUserNameAndPassWordFromForm(c, userRegisterBody.Name, userRegisterBody.Password, userRegisterBody.SurePassword)
	if err != nil {
		return
	}
	// 查询用户是否已经存在
	userInfo := models.GetUserByName(userRegisterBody.Name)
	if userInfo.Id > 0 {
		errorCode = errorInfo.ERROR_EXIST_USER
		c.JSON(http.StatusOK, gin.H{
			"code": errorCode,
			"msg":  errorInfo.GetMsg(errorCode),
		})
		return
	}
	secretPassword, _ := bcrypt.GenerateFromPassword([]byte(userRegisterBody.Password), bcrypt.DefaultCost)
	var user models.TalkUser
	user.Name = userRegisterBody.Name
	user.Password = string(secretPassword)
	user.Avatar = userRegisterBody.Avatar
	user.Email = userRegisterBody.Email

	addUserFlag := models.AddUser(&user)
	if addUserFlag {
		c.JSON(http.StatusOK, gin.H{
			"code": errorCode,
			"msg":  errorInfo.GetMsg(errorCode),
		})
	} else {
		errorCode = errorInfo.ERROR
		c.JSON(http.StatusOK, gin.H{
			"code": errorCode,
			"msg":  "注册失败",
		})
	}
	return
}

func Login(c *gin.Context) {
	var errorCode = errorInfo.SUCCESS
	userRegisterBody := registerBody{}
	err := c.ShouldBindBodyWith(&userRegisterBody, binding.JSON)
	err = checkEmailAndPassWordFromForm(c, userRegisterBody.Email, userRegisterBody.Password)
	if err != nil {
		return
	}

	userInfo := models.GetUserByEmail(userRegisterBody.Email)

	if userInfo.Id > 0 {
		// 判断密码是否正确
		err := bcrypt.CompareHashAndPassword([]byte(userInfo.Password), []byte(userRegisterBody.Password))
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

func checkUserNameAndPassWordFromForm(c *gin.Context, username, password, surePassword string) error {
	var errorCode int
	// 两次密码不同
	if password != surePassword {
		errorCode = errorInfo.ERROR_AUTH_PASSWORD_NOT_EQUAL_SUREPASSWORD
		c.JSON(http.StatusOK, gin.H{
			"code": errorCode,
			"msg":  errorInfo.GetMsg(errorCode),
		})
		return errors.New(errorInfo.GetMsg(errorCode))
	}

	// 验证用户名和密码
	valid := validation.Validation{}
	authInfo := auth{Name: username, Password: password}
	ok, _ := valid.Valid(&authInfo)
	if !ok {
		for _, err := range valid.Errors {
			log.Println(err.Key, err.Message)
		}
		errorCode = errorInfo.INVALID_PARAMS
		c.JSON(http.StatusOK, gin.H{
			"code": errorCode,
			"msg":  errorInfo.GetMsg(errorCode),
		})
		return errors.New(errorInfo.GetMsg(errorCode))
	}
	return nil
}
func checkEmailAndPassWordFromForm(c *gin.Context, email, password string) error {
	var errorCode int
	valid := validation.Validation{}
	authInfo := loginByEmail{Email: email, Password: password}
	ok, _ := valid.Valid(&authInfo)
	if !ok {
		for _, err := range valid.Errors {
			log.Println(err.Key, err.Message)
		}
		errorCode = errorInfo.INVALID_PARAMS
		c.JSON(http.StatusOK, gin.H{
			"code": errorCode,
			"msg":  errorInfo.GetMsg(errorCode),
		})
		return errors.New(errorInfo.GetMsg(errorCode))
	}
	return nil
}
