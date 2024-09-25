package main

import "fmt"

// go 语言只有值传递一种方式

func demo() {
	a := 2
	var pa *int = &a
	*pa = 3
	fmt.Println(*pa, a)
}

func swap(a, b *int) {
	*b, *a = *a, *b
}

func swap_new(a, b int) (int, int) {
	return b, a
}

func main() {
	fmt.Print(" pointer demo \n")
	demo()

	a, b := 1, 2
	swap(&a, &b)
	fmt.Printf("a=%d,b=%d \n", a, b)
	m, n := 6, 8
	m, n = swap_new(m, n)
	fmt.Printf("m=%d,n=%d \n", m, n)
}
