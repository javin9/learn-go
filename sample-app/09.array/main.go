package main

import "fmt"

func basic() {
	var arr1 [5]int
	fmt.Println(arr1)

	arr2 := [3]int{1, 2, 3}
	fmt.Println(arr2)

	arr3 := [...]int{4, 5, 6, 7, 8, 9} // ... 让编译器 帮我们确定长度
	fmt.Println(arr3)

	var grid1 [4][5]int // 4个长度为5的int 数组
	fmt.Println(grid1)

	for i := 0; i < len(arr3); i++ {
		fmt.Println(arr3[i])
	}
	fmt.Println("通过range 遍历")
	for _, value := range arr3 {
		fmt.Println(value)
	}
	// arr [5]int 是数组
	// arr []int 指的切片
	// ------- 一般不使用数组
}

func main() {
	fmt.Println("array demo")
	basic()
}
