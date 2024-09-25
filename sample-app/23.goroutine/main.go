package main

import (
	"fmt"
	"runtime"
	"time"
)

// go 有点类似 js中的Promise
// goroutine 是一个协程 Coroutine。是一个轻量级的线程
// 非抢占式多任务处理，由协程交出控制权
// 编译器，解释器，虚拟机层面的多任务
// 多个协程，可能在一个或者多个进程运行
func basic() {
	var a [10]int
	for i := 0; i < 10; i++ {
		go func(i int) {
			fmt.Println(i)
			a[i]++
		}(i)
	}
	time.Sleep(time.Microsecond)
	fmt.Println(a)
}

func basic2() {
	var a [10]int
	for i := 0; i < 10; i++ {
		go func(i int) {
			for {
				a[i]++
				runtime.Gosched() //手动交出控制权
			}
		}(i)
	}
	time.Sleep(time.Microsecond)
	fmt.Println(a)
}

func basic3() {
	for i := 0; i < 1000; i++ {
		go func(i int) {
			for {
				fmt.Println(i)
			}
		}(i)
	}
	time.Sleep(time.Minute)
}

func main() {
	fmt.Println("goroutine")
	// basic()
	// basic2()
	basic3()
}
