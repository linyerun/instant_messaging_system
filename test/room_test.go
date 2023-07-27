package test

import (
	"fmt"
	"instant_messaging_system/db"
	"instant_messaging_system/models"
	"testing"
)

func TestGetRooms(t *testing.T) {
	rooms, err := db.GetRoomsByCreatorId(1)
	if err != nil {
		t.Fatal(err)
	}
	var room *models.Room
	for _, room = range rooms {
		fmt.Println(room)
	}
}
