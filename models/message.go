package models

import "gorm.io/gorm"

type Message struct {
	gorm.Model
	RoomId  uint   `gorm:"column:room_id"`
	UserId  uint   `gorm:"column:user_id"`
	Content string `gorm:"column:Content"`
}

func (Message) TableName() string {
	return "message"
}
