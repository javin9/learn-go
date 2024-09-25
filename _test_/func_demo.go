package main

import (
	"fmt"
	"reflect"
	"runtime"
)

func main() {
	fmt.Println(div(4, 2))
	q, r := div3(13, 3)
	fmt.Println(q, r)
	result := apply(add, 1, 3)
	fmt.Println(result)
	sum(3, 4)
}

//返回多个值
func div(a, b int) (int, int) {
	return a / b, a % b
}
func div2(a, b int) (q int, r int) {
	//第一种写法：
	return a / b, a % b

	//第二种写法：这种方法，如果函数内容过长，容易混淆
	//q = a / b
	//r = a % b
	//return
}

func div3(a, b int) (q, r int) {
	return a / b, a % b
}

type Person struct {
	Name string `json:"name"`
}

func apply(cb func(int, int) int, a, b int) int {
	//打印cb的名字
	cbPointer := reflect.ValueOf(cb).Pointer()
	cbName := runtime.FuncForPC(cbPointer).Name()
	fmt.Printf("cbName=%s,a=%d,b=%d \n", cbName, a, b)
	return cb(a, b)
}

func add(a, b int) int {
	return a + b
}

func sum(numbers ...int) {
	s := 0
	for _, number := range numbers {
		s += number
	}
	fmt.Println(s)
}
