package models

import "time"

// Messages undefined
type Messages struct {
	ID        int64     `gorm:"id" json:"id"`
	UserId    int64     `gorm:"user_id" json:"user_id"`       // 用户ID
	RoomId    int64     `gorm:"room_id" json:"room_id"`       // 房间ID
	ToUserId  int64     `gorm:"to_user_id" json:"to_user_id"` // 私聊用户ID
	Content   string    `gorm:"content" json:"content"`       // 聊天内容
	ImageUrl  string    `gorm:"image_url" json:"image_url"`   // 图片URL
	CreatedAt time.Time `gorm:"created_at" json:"created_at"`
	UpdatedAt time.Time `gorm:"updated_at" json:"updated_at"`
	DeletedAt time.Time `gorm:"deleted_at" json:"deleted_at"`
}

// TableName 表名称
func (*Messages) TableName() string {
	return "messages"
}
