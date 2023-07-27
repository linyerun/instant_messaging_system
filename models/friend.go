package models

import "time"

type Friend struct {
	Email01   string    `gorm:"column:email01"`
	Email02   string    `gorm:"column:email02"`
	RoomId    uint      `gorm:"column:room_id"`
	CreatedAt time.Time `gorm:"column:created_at"`
	DeletedAt time.Time `gorm:"column:deleted_at;default:NULL"`
}

func (Friend) TableName() string {
	return "friend"
}
