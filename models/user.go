package models

import (
	"go-talk-talk/global"
)

// TalkUser undefined
type TalkUser struct {
	Id                 int64  `json:"Id" gorm:"column:Id"`                                 // Id，自增
	Name               string `json:"Name" gorm:"column:Name"`                             // 用户名
	Nickname           string `json:"NickName" gorm:"column:NickName"`                     // 昵称
	Password           string `json:"Password" gorm:"column:Password"`                     // 密码
	PersonalitySign    string `json:"PersonalitySign" gorm:"column:PersonalitySign"`       // 个性签名
	Mobile             string `json:"Mobile" gorm:"column:Mobile"`                         // 手机号
	OnlineState        int8   `json:"OnlineState" gorm:"column:OnlineState"`               // 在线状态
	Region             string `json:"Region" gorm:"column:Region"`                         // 地区
	Avatar             string `json:"Avatar" gorm:"column:Avatar"`                         // 头像
	RightPadding       int64  `json:"RightPadding" gorm:"column:RightPadding"`             // 右边距，仅需要该字段用于移动端
	Email              string `json:"Email" gorm:"column:Email"`                           // 邮箱
	HistorySessionList string `json:"HistorySessionList" gorm:"column:HistorySessionList"` // 历史会话列表
	OutTradeNo         string `json:"OutTradeNo" gorm:"column:OutTradeNo"`                 // socket.id
	NoCode             string `json:"NoCode" gorm:"column:NoCode"`                         // 时间戳
	CreateAt           int64  `json:"CreateAt" gorm:"column:CreateAt"`                     // 创建时间
	UpdateAt           int64  `json:"UpdateAt" gorm:"column:UpdateAt"`                     // 更新时间
}

// TableName 表名称
func (*TalkUser) TableName() string {
	return "talk_user"
}

//AddUser 增加一个新用户
func AddUser(data *TalkUser) bool {
	global.MysqlDB.Create(data)
	return true
}

//GetUserById 通过Id查找用户
func GetUserById(id int) TalkUser {
	var user TalkUser
	global.MysqlDB.Where(" Id = ?", id).First(&user)
	return user
}

//GetUserByName 通过用户名取得用户
func GetUserByName(name string) TalkUser {
	var user TalkUser
	global.MysqlDB.Where(" Name = ?", name).First(&user)
	return user
}

//GetUserByEmail 通过用户名取得用户
func GetUserByEmail(email string) TalkUser {
	var user TalkUser
	global.MysqlDB.Where(" Email = ?", email).First(&user)
	return user
}

//GetUserListByName  通过用户名查用户
func GetUserListByName(name string) []TalkUser {
	var users []TalkUser
	global.MysqlDB.Where(" Name like ?", "%"+name+"%").Find(&users)

	return users
}

func GetUserListAllByTime(limitNum int, descFlag bool) []TalkUser {
	var users []TalkUser
	var descStr string
	if descFlag {
		descStr = "desc"
	} else {
		descStr = "asc"
	}
	global.MysqlDB.Order("UpdateAt " + descStr).Limit(limitNum).Find(&users)
	return users
}
