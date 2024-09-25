package main

import "fmt"

func main() {
	var arr = [...]int{1, 2, 3, 4, 5, 6}
	for index, value := range arr {
		fmt.Println(index, value)
	}
}
