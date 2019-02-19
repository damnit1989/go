package main

import (
	"fmt"
	"encoding/json"
)

type User struct{
	UserName string `json:"username"`
	NickName string `json:"nickname"`
	Age int
	Birthday string
	Sex string
	Email string
	Phone string
}

//结构体转json
func structToJson() {
	user1 := &User{
		UserName: "user1",
		NickName: "Murphy",
		Age:      18,
		Birthday: "2008/8/8",
		Sex:      "男",
		Email:    "123456@qq.com",
		Phone:    "15600000000",
	}
	
	data, err := json.Marshal(user1)
	if err != nil {
		fmt.Printf("json.marshal failed,err:", err)
		return
	}

	fmt.Printf("json:%s; type:%T\n", string(data), string(data))
}


func main() {
	structToJson()
	fmt.Printf("---------")
}