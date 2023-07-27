package router

import (
	"github.com/gin-gonic/gin"
	"instant_messaging_system/middlewares"
	"instant_messaging_system/service"
)

func Router() (r *gin.Engine) {
	r = gin.Default()

	// 获取验证码
	r.POST("/send_code", service.SendCode)
	// 注册
	r.POST("/register", service.Register)
	// 登录
	r.POST("/login", service.Login)

	auth := r.Group("/auth", middlewares.AuthCheck())
	{

		// 用户模块
		user := auth.Group("/user")
		{
			// 查看个人信息
			user.GET("/detail", service.PersonalMsg)
			// 查询其他用户信息
			user.GET("/query", service.UserQuery)
			// 添加用户
			user.POST("/add", service.AddFriend)
			// 删除好友
			user.DELETE("/delete", service.DeleteFriend)
		}

		// 通信模块
		chat := auth.Group("/chat")
		{
			// 发送、接受消息
			chat.GET("/websocket/message", service.WebsocketMessage)
			// 聊天记录列表
			chat.GET("/chat/list", service.ChatList)
		}

		// 获取文件
		staticFile := auth.Group("/static")
		{
			staticFile.Static("/img", "./static/images")
		}

		// 房间模块
		room := auth.Group("/room")
		{
			room.GET("/get", service.GetRooms)
		}
	}
	return
}
