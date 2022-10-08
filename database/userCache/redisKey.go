package userCache

// 一级前缀

const KEY_USER = "u:"
const KEY_NOTICE = "notice:" //通知公告相关

// 二级前缀

const SUB_KEY_USER_EMAIL = "email:" //用户邮件

// 三级前缀

const SUB_KEY_CAPTCHA = "captcha" // 验证码

// UserEmailCaptchaKey 用户验证码key
func UserEmailCaptchaKey(uid string) string {
	return KEY_USER + SUB_KEY_USER_EMAIL + SUB_KEY_CAPTCHA + uid
}
