package main

import (
	"fmt"
	"github.com/rupid/learn-gin/_test_/learngo/retriever/mock"
	real2 "github.com/rupid/learn-gin/_test_/learngo/retriever/real"
)

type IRetriever interface {
	Get(url string) (htmlContent string)
}

func download(r IRetriever) string {
	return r.Get("http://www.baidu.com")
}

func inspect(r IRetriever) {
	//断言
	switch v := r.(type) {
	case mock.Retriever:
		fmt.Println("mock.Retriever", v)
	case *real2.Retriever:
		fmt.Println("real2.Retriever", v)
	}
}

func inspect2(r IRetriever) {

}

func main() {
	var r IRetriever
	r = mock.Retriever{"fake mock"}
	inspect(r)
	//htmlContent := r.Get("http://www.baidu.com")
	//fmt.Println(htmlContent)

	r = &real2.Retriever{StatusCode: 200}
	inspect(r)
	retriever, ok := r.(*real2.Retriever)
	fmt.Println(ok, retriever.StatusCode)
	//content := r.Get("http://www.baidu.com")
	//fmt.Println(content)
}
