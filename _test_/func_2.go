package main

import "fmt"

func main() {
	a, b := 3, 4
	swapDemo(a, b)
	fmt.Println(a, b)
	swapDemo2(&a, &b)
	fmt.Println(a, b)
}

func swapDemo(a, b int) {
	b, a = a, b
}
func swapDemo2(a, b *int) {
	//第一种：
	//*b, *a = *a, *b

	//第二种：
	temp := *a //去除a的值，给temp int类型
	*a = *b
	*b = temp
}
