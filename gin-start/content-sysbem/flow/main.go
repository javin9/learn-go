package main

import (
	"encoding/json"
	"fmt"
	"math/rand/v2"
	"strconv"
)

func main() {
	fmt.Print("Hello, World! \n")
	fmt.Printf("The number is %d \n", rand.IntN(10))

	// 定义 JSON 字符串
	jsonString := `{"name":"tailiang"}`

	// 定义 map 用于存储解析后的数据
	var result map[string]string

	// 使用 json.Unmarshal 解析 JSON 字符串
	err := json.Unmarshal([]byte(jsonString), &result)
	if err != nil {
		fmt.Println("解析出错:", err)
	}

	// 可以访问 map 中的值
	fmt.Println("Name:", result["name"])
}

func AddOne(data []byte) ([]byte, error) {
	num, _ := strconv.Atoi(string(data))
	output := num + rand.IntN(10) + 1
	fmt.Print(output)
	return []byte(strconv.Itoa(output)), nil
}

func AddTwo(data []byte) ([]byte, error) {
	num, _ := strconv.Atoi(string(data))
	output := num + rand.IntN(101) + 1
	fmt.Print(output)
	return []byte(strconv.Itoa(output)), nil
}
