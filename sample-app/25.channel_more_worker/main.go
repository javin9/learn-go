package main

import "fmt"

func worker(i int, c chan string) {
	for {
		n := <-c
		fmt.Printf("index=%d,channel=%s \n", i, n)
	}
}

func main() {
	fmt.Println("more work")
	var channelList [10]chan string
	for i := 0; i < 10; i++ {
		channelList[i] = make(chan string)
		go worker(i, channelList[i])
	}

	for i := 0; i < 10; i++ {
		channelList[i] <- fmt.Sprintf("hello,%d", i)
	}
}
