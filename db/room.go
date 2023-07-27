package db

import (
	"context"
	"fmt"
	"instant_messaging_system/define"
	"instant_messaging_system/models"
	"time"
)

func AddRoom(r *models.Room) error {
	return MySQLClient.Create(r).Error
}

func GetRoomsByCreatorId(creatorId uint) (rooms []*models.Room, err error) {
	err = MySQLClient.Where(map[string]interface{}{"creator_id": creatorId}).Find(&rooms).Error
	return
}

func IsRoom(roomId uint, roomType uint8) bool {
	key := fmt.Sprintf("%d_%d", roomId, roomType)
	ctx := context.Background()
	if RedisClient.Get(ctx, key).Err().Error() == RedisNil {
		return true
	}
	var cnt int64
	if MySQLClient.Table("room").Where(map[string]interface{}{"id": roomId, "room_type": roomType}).Count(&cnt).Error != nil {
		return false
	}
	if cnt == 0 {
		return false
	}
	RedisClient.Set(ctx, key, "", time.Second*time.Duration(define.CacheExpireTime))
	return true
}
