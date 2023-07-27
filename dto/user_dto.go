package dto

import "instant_messaging_system/models"

type UserDto struct {
	ID       uint   `json:"id,omitempty"`
	Email    string `json:"email,omitempty"`
	Nickname string `json:"nickname,omitempty"`
	Sex      string `json:"sex,omitempty"` // 0-未知, 1-男, 2-女
	Avatar   string `json:"avatar,omitempty"`
}

func NewUserDto(user *models.User) *UserDto {
	userDto := new(UserDto)
	userDto.ID = user.ID
	userDto.Email = user.Email
	userDto.Nickname = user.Nickname
	switch user.Sex {
	case 1:
		userDto.Sex = "男"
	case 2:
		userDto.Sex = "女"
	default:
		userDto.Sex = "未知"
	}
	userDto.Avatar = user.Avatar
	return userDto
}
