package service

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"instant_messaging_system/db"
	"instant_messaging_system/define"
	"instant_messaging_system/dto"
	"instant_messaging_system/models"
	"instant_messaging_system/utils"
	"log"
	"net/http"
	"strconv"
	"strings"
	"sync"
)

var (
	upgrader       = websocket.Upgrader{}
	websocketConns = make(map[string]*websocket.Conn, 100)
	mapMutex       sync.RWMutex
)

// WebsocketMessage 每个用户只允许同时存在一个聊天房间
func WebsocketMessage(c *gin.Context) {
	// 校验参数
	roomIdStr := c.Query("room_id")
	roomTypeStr := c.Query("room_type")
	roomId, err := strconv.Atoi(roomIdStr)
	roomType, err := strconv.Atoi(roomTypeStr)
	if err != nil {
		c.JSON(200, define.NewResponseResult(400, "参数有误", nil))
		return
	}

	// 判断房间是否存在
	if !db.IsRoom(uint(roomId), uint8(roomType)) {
		c.JSON(200, define.NewResponseResult(400, "房间不存在, 发送失败", nil))
		return
	}

	// 获取连接
	conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		c.JSON(http.StatusOK, define.NewResponseResult(400, "系统错误，请稍后再尝试", nil))
		return
	}
	defer conn.Close() // 关闭连接

	// 把连接加入map集合中, 移除之前的连接
	email := c.GetString("user_email")
	mapMutex.Lock()
	key := email + "_" + roomIdStr
	if c, ok := websocketConns[key]; ok {
		c.Close()
		delete(websocketConns, key)
	}
	websocketConns[key] = conn
	mapMutex.Unlock()
	defer func() { // 移除
		mapMutex.Lock()
		delete(websocketConns, email)
		mapMutex.Unlock()
	}()

	msgDto := new(dto.UserMsgDto)
	for {
		msgDto.Token = ""
		err := conn.ReadJSON(msgDto)
		if err != nil {
			log.Panicf("read error: %v\n", err)
			return
		}

		// 判断token是否过期
		uid, _, err := utils.ParseToken(msgDto.Token)
		if err != nil {
			fmt.Printf("token error: %v\n", err)
			return
		}

		// 存储消息
		go func() {
			msg := &models.Message{RoomId: uint(roomId), UserId: uid, Content: msgDto.Content}
			err := db.SaveMessage(msg)
			if err != nil {
				log.Printf("mysql error: %v\n", err)
				db.RedisClient.LPush(context.Background(), "msg_not_save", msg)
			}
		}()

		// 判断哪些用户需要被发送
		mapMutex.RLock()
		for k, conn := range websocketConns {
			if strings.Contains(k, "_"+roomIdStr) {
				conn.WriteJSON(gin.H{"msg_content": msgDto.Content, "msg_user_email": email})
			}
		}
		mapMutex.RUnlock()
	}
}
