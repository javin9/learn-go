package main

import "fmt"

func basic() {
	fmt.Println("chan is demo")

	d := make(chan int)
	// 接手 必须在前面，发送在后面
	go func() {
		n := <-d
		fmt.Println(n)
	}()

	d <- 1 //往channel发送数据
	d <- 2 // 这个发送会报错，应为已经发出去之后，没人接收了
}

// 版本2
func dee() {
	d := make(chan int)
	// 改进 写了一个for循环，不停的接收
	go func() {
		for {
			n := <-d
			fmt.Println(n)
		}
	}()
	//
	// d <- 1
	// d <- 2
	for i := 0; i < 100; i++ {
		d <- i
	}
}

// 版本 3
func work(c chan int) {
	for {
		n := <-c
		fmt.Println(n)
	}
}
func dee2() {
	d := make(chan int)
	//
	go work(d)
	//
	// d <- 1
	// d <- 2
	for i := 0; i < 100; i++ {
		d <- i
	}
}
func main() {
	basic() //这种情况下会报错，因为goroutine 需要接收
	// dee()
	// dee2()
}
