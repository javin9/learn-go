package main

//http://c.biancheng.net/view/28.html
import "fmt"

func main() {
	var arr []int
	arr = append(arr, 0)
	arr = append(arr, 1, 2, 3)
	arr = append([]int{-1}, arr...)
	fmt.Println(arr)

	//push 尾部追加
	b := []int{0}
	c := append(b, 1, 2, 3)
	d := append(b, c...)
	fmt.Println(b, c, d)

	// unshift	开头添加元素,其实就是切片拼装
	aa := []int{4, 5}
	value := 3
	bb := append([]int{0, 1, 2, value}, aa...)
	fmt.Println(bb)
}
