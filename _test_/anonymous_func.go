package main

import "fmt"

func main() {
	list := []int{1, 2, 3, 4, 5}
	iterate(list, handleValue)
}

func iterate(list []int, cb func(value int) int) {
	for _, value := range list {
		result := cb(value)
		fmt.Println(result)
	}
}

func handleValue(value int) int {
	return value * value
}
