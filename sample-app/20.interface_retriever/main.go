package main

import (
	"fmt"
	"sample-app/20.interface_retriever/mock"
	realgo "sample-app/20.interface_retriever/real"
)

// Type Assertion
// Type Switch

// 接口变量自带指针
// 1.接口变量是一个指针类型的值和一个指向具体类型的指针
// 1.1在 Go 中，接口类型（interface）是一个抽象类型，它包括两个部分：一个类型信息和一个值信息。类型信息存储的是这个值的具体类型，而值信息则存储这个具体类型的值的地址。例如，如果你有一个接口变量，它可能包含实际类型是 *MyStruct（指向 MyStruct 的指针）的信息，也可能包含实际类型是 MyStruct 本身的值。

// 2.接口变量的赋值和传递行为：

// 2.1当我们将一个值赋值给一个接口变量时，接口会持有这个值的副本，而不是引用。即使我们赋值的是一个指针（例如 &MyStruct），接口变量实际上存储的是该指针的副本，而不是指向原始值的指针。
// 由于接口变量实际上存储的是值信息的地址，我们可以认为接口变量自带了指针，因为它不直接存储值，而是存储值的地址

// 接口变量同样采用值传递，几乎不需要使用接口的指针

type IRetriever interface {
	Get(url string) (string, error) //有点像委托
}

type IRetrieverPost interface {
	Post(url string) string //有点像委托
}

// 接口组合
type ICombine interface {
	IRetriever
	IRetrieverPost
}

func download(r IRetriever, url string) (string, error) {
	return r.Get(url)
}

func main() {
	// retrieve := mock.Retriever{}
	var retrieve IRetriever = &mock.Retriever{}

	fmt.Printf("retrieve type is 【%T】,value=【%s】", retrieve, retrieve)
	download(retrieve, "https://www.baidu.com")
	fmt.Println("interface download demo")

	retrieve = &realgo.Retriever{Header: "我是真正的retrieve"}
	inspect(retrieve)
	//NOTE： Type assertion 断言
	rr := retrieve.(*realgo.Retriever)
	fmt.Println(rr.Header)
	// dd
	var post IRetrieverPost = &mock.Retriever{}
	post.Post("https://wwww.google.com")
}

// j检测
func inspect(r IRetriever) {
	// 实现了string了。打印信息
	fmt.Printf("%s \n", r)
	// NOTE:上面  不能这么写  retrieve := mock.Retriever{}，不然会提示is not Retriever
	switch v := r.(type) {
	case *mock.Retriever:
		fmt.Println("mock.Retriever", v)
	case *realgo.Retriever:
		fmt.Println("mock.Retriever", v.Header)
	default:
		fmt.Println("default ", v)
	}
}
