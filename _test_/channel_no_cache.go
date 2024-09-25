package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan int, 0)

	go func() {
		for i := 0; i < 4; i++ {
			fmt.Printf("sub=%d \n", i)
			ch <- i
		}
	}()

	time.Sleep(3 * time.Second)

	for j := 6; j < 10; j++ {
		value := <-ch
		fmt.Printf("value=%d,main=%d \n", value, j)
	}
}
