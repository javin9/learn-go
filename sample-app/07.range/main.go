package main

import "fmt"

func sum(number ...int) int {
	result := 0
	for _, value := range number {
		result += value
	}
	return result
}

func main() {
	fmt.Print("range demo \n")
	fmt.Printf("sum result is %d \n", sum(2, 3, 4, 5, 6, 7, 4))
}
