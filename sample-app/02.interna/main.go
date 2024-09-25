package main

import (
	"fmt"
	"math"
	"strconv"
)

func main() {
	fmt.Print("hello \n")
	triangle()
	constTest()
}

func triangle() {
	var a = 3
	var b = 4
	c := int(math.Sqrt(float64(a*a + b*b)))
	d := strconv.Itoa(c) //转字符串的方法 strconv

	fmt.Println(c, d)
}

func constTest() {
	const filename = "filename"

	print(filename)
	println(filename)
}

/**
 * book,string
 * int,int8
 * 没有隐式转换
 */
