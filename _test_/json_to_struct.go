package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	var it IT1
	jsonData := `{"company":"太凉","subjects":["hello","world"],"isOk":true,"price":3.2}`
	//转成切片
	//第二个参数 取地址
	err := json.Unmarshal([]byte(jsonData), &it)
	if err != nil {
		fmt.Println("error:", err)
	}
	fmt.Printf("%+v", it)
}

// 必须大写
type IT1 struct {
	Company  string   `json:"company"`
	Subjects []string `json:"subjects"`
	IsOk     bool     `json:"isOk"`
	Price    float32  `json:"price"`
}
