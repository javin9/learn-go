package main

import (
	"fmt"
)

func main() {
	ch := make(chan int, 3)

	go func() {
		for i := 0; i < 3; i++ {
			ch <- i
			fmt.Printf("子协程[%d],len(ch)=%d, cap(ch)=%d \n", i, len(ch), cap(ch))
		}
	}()

	//time.Sleep(3 * time.Second)

	for j := 6; j < 8; j++ {
		value := <-ch
		fmt.Printf("value=%d,main=%d \n", value, j)
	}
}
