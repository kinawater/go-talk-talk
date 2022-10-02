package errorInfo

const (
	SUCCESS        = 200
	ERROR          = 500
	INVALID_PARAMS = 400

	ERROR_EXIST_TAG         = 10001
	ERROR_NOT_EXIST_TAG     = 10002
	ERROR_NOT_EXIST_ARTICLE = 10003
	ERROR_NOT_EXIST_USER    = 10004
	ERROR_EXIST_USER        = 10005

	ERROR_AUTH_NO_USERNAME_OR_PASSWORD = 20001
	ERROR_AUTH_CHECK_TOKEN_FAIL        = 20002
	ERROR_AUTH_CHECK_TOKEN_TIMEOUT     = 20003
	ERROR_AUTH_TOKEN                   = 20004
	ERROR_AUTH                         = 20005
)

const (
	SUCCESS_CN                            = "ok"
	ERROR_CN                              = "fail"
	INVALID_PARAMS_CN                     = "请求参数错误"
	ERROR_EXIST_TAG_CN                    = "已存在该标签名称"
	ERROR_NOT_EXIST_TAG_CN                = "该标签不存在"
	ERROR_NOT_EXIST_ARTICLE_CN            = "该文章不存在"
	ERROR_NOT_EXIST_USER_CN               = "该用户不存在"
	ERROR_EXIST_USER_CN                   = "该用户名已经存在"
	ERROR_AUTH_NO_USERNAME_OR_PASSWORD_CN = "请输入正确的用户名或者密码"
	ERROR_AUTH_CHECK_TOKEN_FAIL_CN        = "Token鉴权失败"
	ERROR_AUTH_CHECK_TOKEN_TIMEOUT_CN     = "Token已超时"
	ERROR_AUTH_TOKEN_CN                   = "Token生成失败"
	ERROR_AUTH_CN                         = "Token错误"
)

var ErrorMsgCN = map[int]string{
	SUCCESS:                            SUCCESS_CN,
	ERROR:                              ERROR_CN,
	INVALID_PARAMS:                     INVALID_PARAMS_CN,
	ERROR_EXIST_TAG:                    ERROR_EXIST_TAG_CN,
	ERROR_NOT_EXIST_TAG:                ERROR_NOT_EXIST_TAG_CN,
	ERROR_NOT_EXIST_ARTICLE:            ERROR_NOT_EXIST_ARTICLE_CN,
	ERROR_NOT_EXIST_USER:               ERROR_NOT_EXIST_USER_CN,
	ERROR_AUTH_CHECK_TOKEN_FAIL:        ERROR_AUTH_CHECK_TOKEN_FAIL_CN,
	ERROR_AUTH_CHECK_TOKEN_TIMEOUT:     ERROR_AUTH_CHECK_TOKEN_TIMEOUT_CN,
	ERROR_AUTH_TOKEN:                   ERROR_AUTH_TOKEN_CN,
	ERROR_AUTH:                         ERROR_AUTH_CN,
	ERROR_AUTH_NO_USERNAME_OR_PASSWORD: ERROR_AUTH_NO_USERNAME_OR_PASSWORD_CN,
	ERROR_EXIST_USER:                   ERROR_EXIST_USER_CN,
}
