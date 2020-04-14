package main

import (
	"encoding/json"
	"fmt"
	"testing"
)

//用户
type User struct {
	UserName string `json:"username"`
	NickName string `json:"nickname"`
	Age      int    `json:"age"`
	Birthday string `json:"birthday"`
	Sex      string `json:"sex"`
	Email    string `json:"email"`
	Phone    string `json:"phone"`
}

//结构体转JSON
func structToJSON() {

	user := User{
		UserName: "itbsl",
		NickName: "jack",
		Age:      18,
		Birthday: "2001-08-15",
		Sex:      "itbsl@gmail.com",
		Phone:    "176XXXX6688",
	}

	data, err := json.Marshal(user)

	if err != nil {
		fmt.Println("json.marshal failed, err:", err)
		return
	}

	fmt.Printf("%s\n", string(data))
}

func TestJsk(t *testing.T) {

	structToJSON()

}
