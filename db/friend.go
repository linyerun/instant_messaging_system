package db

import (
	"context"
	"fmt"
	"instant_messaging_system/models"
	"time"
)

func IsFriend(email01, email02 string) (bool, uint) {
	var cnt int64
	var roomId uint
	err := MySQLClient.Table("friend").Select("room_id").Where("((email01=? and email02=?) or (email02=? and email01=?)) and deleted_at is null", email01, email02, email01, email02).Count(&cnt).Scan(&roomId).Error
	if err != nil {
		return false, 0
	}
	return cnt > 0, roomId
}

func AddFriend(f *models.Friend, r *models.Room) (err error) {
	db := MySQLClient.Begin()
	if err = db.Create(r).Error; err != nil {
		db.Rollback()
		return
	}
	f.RoomId = r.ID
	if err = MySQLClient.Create(f).Error; err != nil {
		db.Rollback()
		return
	}
	if err = db.Commit().Error; err != nil {
		db.Rollback()
		return
	}
	return
}

func DeleteFriend(email01, email02 string, roomId uint) bool {
	RedisClient.Del(context.Background(), fmt.Sprintf("%d_%d", roomId, 1))
	db := MySQLClient.Begin()
	if err := db.Table("friend").Where("(email01=? and email02=?) or (email01=? and email02=?)", email01, email02, email02, email01).Update("deleted_at", time.Now()).Error; err != nil {
		db.Rollback()
		return false
	}
	if err := db.Delete(&models.Room{}, roomId).Error; err != nil {
		db.Rollback()
		return false
	}
	if err := db.Commit().Error; err != nil {
		db.Rollback()
		return false
	}
	return true
}
