package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	fmt.Println("go wait")
	numberGroup := 4
	var wg sync.WaitGroup
	wg.Add(numberGroup)
	for i := 0; i < numberGroup; i++ {
		go func(index int) {
			doSomething(index, &wg)
		}(i)
	}
	wg.Wait() // 等待所有goroutine完成
	fmt.Println("All goroutines are done")
}

func doSomething(index int, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Printf("index=%d is running \n", index)
	time.Sleep(time.Microsecond * 3)
	fmt.Printf("index=%d is done \n", index)
}
