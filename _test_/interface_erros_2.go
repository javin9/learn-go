package main

import (
	"errors"
	"fmt"
)

func main() {
	result, err := MyDiv(1, 1)
	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println(result)
	}
}

func MyDiv(a int, b int) (result int, err error) {
	err = nil
	if b == 0 {
		err = errors.New("除数不能为空")
		return
	}
	result = a / b
	return
}
