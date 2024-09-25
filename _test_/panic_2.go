package main

import "fmt"

//panic很致命
//errors 自定义错误
func main() {
	t11()
	t22(20)
	t33()
}

func t11() {
	fmt.Println("1111")
}

func t22(x int) {
	var arr [10]int
	arr[x] = 20
	//	panic: runtime error: index out of range [20] with length 10
}

func t33() {
	fmt.Println("1111")
}
