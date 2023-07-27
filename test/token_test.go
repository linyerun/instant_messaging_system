package test

import (
	"fmt"
	"instant_messaging_system/utils"
	"testing"
)

func TestGenerateToken(t *testing.T) {
	token, err := utils.GenerateToken(1, "2338244917@qq.com")
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println("token:", token)
}

func TestParseToken(t *testing.T) {
	id, email, err := utils.ParseToken("f9503fc8780a82d0cc0ea76bd5ce1517")
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println("id:", id)
	fmt.Println("email:", email)
}
