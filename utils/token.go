package utils

import (
	"context"
	"fmt"
	"instant_messaging_system/db"
	"instant_messaging_system/define"
	"strconv"
	"strings"
	"time"
)

func GenerateToken(id uint, email string) (token string, err error) {
	val := fmt.Sprintf("%d_%s", id, email)
	token = GetMd5(val)
	err = db.RedisClient.Set(context.Background(), define.TokenPrefix+token, val, time.Duration(define.TokenExpireTime)*time.Second).Err()
	if err != nil {
		return "", err
	}
	return
}

func ParseToken(token string) (id uint, email string, err error) {
	res := db.RedisClient.Get(context.Background(), define.TokenPrefix+token)
	if err = res.Err(); err != nil {
		return 0, "", err
	}
	defer db.RedisClient.Expire(context.Background(), define.TokenPrefix+token, time.Duration(define.TokenExpireTime)*time.Second)
	// 解析结果
	var val = res.Val()
	a := strings.Split(val, "_")
	email = a[1]
	var uid int
	uid, err = strconv.Atoi(a[0])
	if err != nil {
		return 0, "", err
	}
	id = uint(uid)
	return
}
