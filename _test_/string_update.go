package main

import "fmt"

func main() {
	//字符串不能直接修改，只能转成byte 修改，在转成字符串
	str := "hello"
	byteStr := []byte(str)
	byteStr[2] = 'g'

	str = string(byteStr)
	fmt.Println(str)
}
