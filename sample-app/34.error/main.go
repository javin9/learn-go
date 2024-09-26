package main

import (
	"errors"
	"fmt"
)

var ErrorNotFound error = errors.New("Not Found")

func main() {
	fmt.Println("error 处理")
	data, err := findData(1)
	if err != nil {
		if err == ErrorNotFound {
			// 已知错误
			fmt.Println("know error")
		} else {
			// 未知错误
			fmt.Println("unKnow error")
		}
		return
	}
	fmt.Println("data", data)
}

func findData(id int) (int, error) {
	if id != 1 {
		return 0, ErrorNotFound
	}
	return 3, nil
}
