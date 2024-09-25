package main

import (
	"encoding/json"
	"fmt"
)

type User struct {
	Name    string `json:"name,omitempty"`
	Age     int    `json:"age,omitempty"`
	Address string `json:"address,omitempty"`
}

func main() {
	var user User
	user = User{"tailiang", 18, "北京长安街1号"}
	b, err := json.Marshal(user)
	if err != nil {
		fmt.Println("error:", err)
	}
	fmt.Printf("User结构体转json：%s\n", string(b))
}
