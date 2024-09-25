package main

import "fmt"

// 闭包
func add(initialValue int) func(value int) int {
	start := initialValue
	// 闭包的写法
	return func(value int) int {
		start += value
		return start
	}
}

// fibonacci
func fibonacci() func() int {
	a, b := 0, 1
	return func() int {
		a, b = b, a+b
		return a
	}
}

// type 定义一种类型

// 函数是一等公民，表示 参数、变量、返回值 都可以是 函数
// 高阶函数:把函数当做参数传递
func main() {
	fmt.Println("closure typeString")
	fmt.Println("函数式变成和函数指针")

	adder := add(2)
	fmt.Println(adder(1))
	fmt.Println(adder(2))
	fmt.Println(adder(3))
}
