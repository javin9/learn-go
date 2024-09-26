package main

import (
	"errors"
	"fmt"
)

var ErrorNotFound error = errors.New("Not Found")

func main() {
	// fmt.Println("error 处理")
	// data, err := findData(1)
	// if err != nil {
	// 	if err == ErrorNotFound {
	// 		// 已知错误
	// 		fmt.Println("know error")
	// 	} else {
	// 		// 未知错误
	// 		fmt.Println("unKnow error")
	// 	}
	// 	return
	// }
	// fmt.Println("findData1", data)

	data, err := findData2(2)
	if err != nil {
		fmt.Println("err", err)
		fmt.Println(fmt.Errorf("111 %s \n", err.Error()))
		if err.(*HTTPError).Code == 2001 {
			fmt.Println("参数错误")
		}
		return
	}
	fmt.Println("findData2", data)
}

func findData(id int) (int, error) {
	if id != 1 {
		return 0, ErrorNotFound
	}
	return 3, nil
}

// 通用错误码处理逻辑

type HTTPError struct {
	Code    int
	Message string
}

func (me *HTTPError) Error() string {
	switch me.Code {
	case 2001:
		return "参数错误"
	case 2002:
		return "没有权限"
	default:
		return "服务端错误~"
	}
}

func findData2(id int) (int, error) {
	if id != 1 {
		return 0, &HTTPError{Code: 2001}
	}
	return 3, nil
}
