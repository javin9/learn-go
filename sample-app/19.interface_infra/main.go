package main

import (
	"fmt"
	"sample-app/19.interface_infra/tsetse"
)

func getRetrieve() Retriever {
	return tsetse.Retrieve{}
}

// something that can Get
type Retriever interface {
	Get(url string) (string, error)
}

func main() {
	fmt.Println("interface demo")
	// 这里就像封装了一个 类，加了一个方法
	var retrieve Retriever = getRetrieve()
	response, _ := retrieve.Get("https://www.baidu.com")
	fmt.Printf("%s", response)
}
