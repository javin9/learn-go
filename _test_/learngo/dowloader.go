package main

import (
	"fmt"
	"github.com/rupid/learn-gin/_test_/learngo/infra"
)

func getRetriever() IRetriever {
	return infra.Retriever{}
}

//接口
//结构体实现接口方法
type IRetriever interface {
	Get(url string) string
}

func main() {
	///写出来retriever的类型
	var retriever IRetriever = getRetriever()
	contnt := retriever.Get("https://www.baidu.com")
	fmt.Println(contnt)
}
