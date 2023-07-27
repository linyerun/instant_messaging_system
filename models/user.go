package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Email    string `gorm:"column:email"`
	Password string `gorm:"column:password"`
	Nickname string `gorm:"column:nickname"`
	Sex      uint8  `gorm:"column:sex"` // 0-未知, 1-男, 2-女
	Avatar   string `gorm:"column:avatar"`
}

func (User) TableName() string {
	return "user"
}
