package db

import "github.com/go-redis/redis/v8"

var RedisClient = InitRedis()

const RedisNil = "redis: nil"

func InitRedis() *redis.Client {
	return redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
	})
}
