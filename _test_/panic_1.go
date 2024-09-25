package main

import "fmt"

//panic很致命
//errors 自定义错误
func main() {
	t1()
	t2()
	t3()
}

func t1() {
	fmt.Println("1111")
}

func t2() {
	//显示调用panic，让程序中断
	panic("显示调用了")
	fmt.Println("1111")
}

func t3() {
	fmt.Println("1111")
}
