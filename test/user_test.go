package test

import (
	"fmt"
	"instant_messaging_system/db"
	"instant_messaging_system/models"
	"instant_messaging_system/utils"
	"testing"
)

func TestInsertUser(t *testing.T) {
	user := &models.User{
		Email:    "2338244917@qq.com",
		Password: utils.GetMd5("123456"),
		Avatar:   "",
		Nickname: "随风",
		Sex:      1,
	}
	err := db.InsertUser(user)
	if err != nil {
		t.Fatal(err)
	}
}

func TestGetUserId(t *testing.T) {
	id, err := db.GetUserIdByEmailPwd("2338244917@qq.com", utils.GetMd5("123456"))
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println("id:", id)
}

func TestGetUserById(t *testing.T) {
	user, err := db.GetUserById(2)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println("user:", user)
}

func TestGetUserByEmail(t *testing.T) {
	user, err := db.GetUserByEmail("3268242396@qq.com")
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println("user:", user)
}
