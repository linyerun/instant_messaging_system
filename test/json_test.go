package test

import (
	"encoding/json"
	"fmt"
	"testing"
)

type User struct {
	Id   uint `json:"id,omitempty"`
	Name string
}

// 不大写, 反射也获取不到
func TestUserJson(t *testing.T) {
	user := User{Name: "哈哈哈"}
	bytes, err := json.Marshal(user)
	if err != nil {
		t.Fatal(err)
	}
	u := new(User)
	err = json.Unmarshal(bytes, u)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(string(bytes))
	fmt.Println(u)
}
