package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	jsonData := `{"age":18,"friends":["zhangsan","lisi","wangwu"],"isUpdate":true,"name":"太凉"}`
	targetMap := make(map[string]interface{})
	err := json.Unmarshal([]byte(jsonData), &targetMap)
	if err != nil {
		fmt.Println(err)
	}
	for key, value := range targetMap {
		fmt.Printf("key=%v,value=%v \n", key, value)
		switch data := value.(type) {
		case string:
			fmt.Printf("map[] ,%v,%v \n", key, data)
		}
	}
}
