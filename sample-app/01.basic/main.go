package main

import "fmt"

var c = 3
var (
	a = 1
	b = 3
)

func main() {
	fmt.Println(a, b, c)
	test()
}

func test() {
	var age = 20
	name := "javin"
	school := "清华啊"

	fmt.Printf("%d,%s,%s\n", age, name, school)
}

// 换句话说，“:=”只能在声明“局部变量”的时候使用，而“var”没有这个限制。
// https://studygolang.com/articles/5294
