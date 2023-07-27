package service

import (
	"github.com/gin-gonic/gin"
	"instant_messaging_system/db"
	"instant_messaging_system/define"
	"strconv"
)

func GetRooms(c *gin.Context) {
	uid, err := strconv.Atoi(c.GetString("user_id"))
	if err != nil {
		c.JSON(200, define.NewResponseResult(500, "系统错误:"+err.Error(), nil))
		return
	}
	creatorId := uint(uid)
	rooms, err := db.GetRoomsByCreatorId(creatorId)
	if err != nil {
		c.JSON(200, define.NewResponseResult(500, "系统错误:"+err.Error(), nil))
		return
	}
	c.JSON(200, define.NewResponseResult(200, "成功", gin.H{"rooms": rooms}))
}
