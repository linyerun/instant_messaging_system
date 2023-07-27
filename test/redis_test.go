package test

import (
	"context"
	"fmt"
	"instant_messaging_system/db"
	"testing"
)

func TestNoKey(t *testing.T) {
	err := db.RedisClient.Get(context.Background(), "哈哈哈").Err()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("没有报错")
}
