package main

import "fmt"

//panic很致命
//errors 自定义错误
//recover 恢复
func main() {
	t111()
	t222(20)
	t222(0)
	t333()
}

func t111() {
	fmt.Println("1111")
}

func t222(x int) {
	//设置 recover
	defer func() {
		//recover()
		//fmt.Println(recover())

		if panicresult := recover(); panicresult != nil {
			fmt.Println(panicresult)
		}
	}()
	var arr [10]int
	arr[x] = 20
	//	panic: runtime error: index out of range [20] with length 10
}

func t333() {
	fmt.Println("333")
}
