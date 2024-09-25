package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Printf("%d,%q \n", 133, "hello world \n")
	triangle()
}

func triangle() {
	var a, b = 3, 4
	const aa, bb = 3, 4

	var c int
	c = int(math.Sqrt(float64(a*a + b*b)))

	//const 不需要转换
	c = int(math.Sqrt(aa*aa + bb*bb))
	fmt.Println(c)
}

func enmusTest() {
	const (
		java       = 0
		golang     = 1
		javascript = 2
		python     = 3
	)
	//iota
	const (
		java1 = iota
		_     //跳过1
		golang1
		javascript1
		python1
	)

}
