package main

import "fmt"

type Processor interface {
	Process(data string)
}

// 使用函数类型实现接口
type ProcessFunc func(num int)

func (me ProcessFunc) Process(data string) {
	fmt.Println("data", data)
}

func main() {
	fmt.Println("函数实现接口")

	// 定义一个符合 ProcessFunc 签名的函数
	myFunc := ProcessFunc(func(num int) {
		fmt.Println("Processing data:", num)
	})
	myFunc(123)
	// 因为 myFunc 符合 Processor 接口，所以这里可以直接使用
	var p Processor = myFunc
	p.Process("Hello, world!")

}
