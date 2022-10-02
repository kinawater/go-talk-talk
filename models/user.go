package models

import (
	"go-talk-talk/global"
	"time"
)

// Users undefined
type Users struct {
	ID        int64     `gorm:"id" json:"id"`
	Username  string    `gorm:"username" json:"username"`   // 昵称
	Password  string    `gorm:"password" json:"password"`   // 密码
	AvatarId  string    `gorm:"avatar_id" json:"avatar_id"` // 头像ID
	CreatedAt time.Time `gorm:"created_at" json:"created_at"`
	UpdatedAt time.Time `gorm:"updated_at" json:"updated_at"`
	DeletedAt time.Time `gorm:"deleted_at" json:"deleted_at"`
}

// TableName 表名称
func (*Users) TableName() string {
	return "users"
}

//AddUser 增加一个新用户
func AddUser(data *Users) bool {
	global.MysqlDB.Create(data)
	return true
}

//GetUserById 通过Id查找用户
func GetUserById(id int) Users {
	var user Users
	global.MysqlDB.Where(" id = ?", id).First(&user)
	return user
}

// 通过用户名取得用户
func GetUserByName(name string) Users {
	var user Users
	global.MysqlDB.Where(" username = ?", name).First(&user)
	return user
}

//GetUserByName  通过用户名查用户
func GetUserListByName(name string) []Users {
	var users []Users
	global.MysqlDB.Where(" username like ?", "%"+name+"%").Find(&users)

	return users
}
