package main

import "fmt"

func worker(i int, c chan string) {
	for {
		n := <-c
		fmt.Printf("index=%d,channel=%s \n", i, n)
	}
}

func createWorker(i int) chan string {
	c := make(chan string)
	go func(chan string) {
		for {
			n := <-c
			fmt.Printf("index=%d,channel=%s \n", i, n)
		}
	}(c)
	return c
}

func main() {
	fmt.Println("more work")
	var channelList [10]chan string
	for i := 0; i < 10; i++ {
		channelList[i] = createWorker(i)
	}

	for i := 0; i < 10; i++ {
		channelList[i] <- fmt.Sprintf("hello,%d", i)
	}
	for i := 0; i < 10; i++ {
		channelList[i] <- fmt.Sprintf("Hello,%d", i)
	}
}
