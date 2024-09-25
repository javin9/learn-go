package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	group := ColorGroup{ID: 1, Name: "Reds", IsOk: true, IsOk2: true, Colors: []string{"Crimson", "Red", "Ruby", "Maroon"}}
	b, err := json.Marshal(group)
	if err != nil {
		fmt.Println("error:", err)
	}
	fmt.Println(string(b))
	//{"id":1,"name":"Reds","isOk":"true","isOk2":true}

}

type ColorGroup struct {
	ID     int      `json:"id"`
	Name   string   `json:"name"`
	IsOk   bool     `json:"isOk,string"` //字符串方式输出
	IsOk2  bool     `json:"isOk2"`       //boolean 类型
	Colors []string `json:"-"`           //代表不会输出
}
