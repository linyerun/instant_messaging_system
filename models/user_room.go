package models

import (
	"gorm.io/gorm"
)

type UserRoom struct {
	gorm.Model
	RoomId uint `gorm:"column:room_id"`
	UserId uint `gorm:"column:user_id"`
}

func (UserRoom) TableName() string {
	return "user_room"
}
