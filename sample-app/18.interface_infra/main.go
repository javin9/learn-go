package main

import (
	"fmt"
	"sample-app/18.interface/infra"
)

func main() {
	fmt.Println("interface demo")
	// 这里就像封装了一个 类，加了一个方法
	retrieve := infra.Retrieve{}
	response, _ := retrieve.Get("https://www.baidu.com")
	fmt.Printf("%s", response)
}
