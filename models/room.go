package models

import "gorm.io/gorm"

type Room struct {
	gorm.Model
	Name      string `gorm:"column:name" json:"name"`
	Info      string `gorm:"column:info" json:"info"`           // 房间介绍
	RoomType  uint8  `gorm:"column:room_type" json:"room_type"` // 0: 多人, 1:单人
	CreatorId uint   `gorm:"column:creator_id" json:"creator_id"`
}

func (Room) TableName() string {
	return "room"
}
