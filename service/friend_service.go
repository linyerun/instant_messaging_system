package service

import (
	"github.com/gin-gonic/gin"
	"instant_messaging_system/db"
	"instant_messaging_system/define"
	"instant_messaging_system/models"
	"instant_messaging_system/utils"
	"net/http"
	"time"
)

func AddFriend(c *gin.Context) {
	// 插入操作
	email01 := c.GetString("user_email")
	email02 := c.PostForm("email")
	if !utils.VerifyEmail(email02) {
		c.JSON(http.StatusOK, define.NewResponseResult(401, "参数有误", nil))
		return
	}
	if ok, _ := db.IsFriend(email01, email02); ok {
		c.JSON(http.StatusOK, define.NewResponseResult(401, "添加好友失败, 你们已是好友", nil))
		return
	}
	// 新增操作
	r := &models.Room{RoomType: 1, CreatorId: c.GetUint("user_id")}
	f := &models.Friend{Email01: email01, Email02: email02, CreatedAt: time.Now()}
	if err := db.AddFriend(f, r); err != nil {
		c.JSON(http.StatusOK, define.NewResponseResult(500, "系统错误", nil))
		return
	}
	c.JSON(http.StatusOK, define.NewResponseResult(200, "添加好友成功", nil))
}

func DeleteFriend(c *gin.Context) {
	email01 := c.GetString("user_email")
	email02 := c.Query("email")
	if !utils.VerifyEmail(email02) {
		c.JSON(http.StatusOK, define.NewResponseResult(401, "参数有误", nil))
		return
	}
	ok, roomId := db.IsFriend(email01, email02)
	if !ok {
		c.JSON(http.StatusOK, define.NewResponseResult(401, "删除好友失败, 你们不是好友", nil))
		return
	}
	// 执行删除操作
	if ok = db.DeleteFriend(email01, email02, roomId); !ok {
		c.JSON(http.StatusOK, define.NewResponseResult(500, "系统错误，删除失败。", nil))
		return
	}
	c.JSON(http.StatusOK, define.NewResponseResult(401, "删除成功", nil))
}
