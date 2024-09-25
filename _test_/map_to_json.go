package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	dic := make(map[string]interface{})
	dic["name"] = "太凉"
	dic["age"] = 18
	dic["isUpdate"] = true
	dic["friends"] = []string{"zhangsan", "lisi", "wangwu"}
	byteValue, errorValue := json.Marshal(dic)
	if errorValue != nil {
		fmt.Println(errorValue)
	}
	//{"age":18,"friends":["zhangsan","lisi","wangwu"],"isUpdate":true,"name":"太凉"}
	fmt.Println(string(byteValue))
}
