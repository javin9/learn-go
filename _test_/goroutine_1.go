package main

import (
	"fmt"
	"time"
)

func main() {
	//新建一个协程，新建一个任务
	go Task1()
	Task2()
	//fmt.Println("start")
}

func Task1() {
	for {
		fmt.Println(time.Now())
		time.Sleep(time.Second)
	}
}

func Task2() {
	for {
		fmt.Println("任务2")
		time.Sleep(time.Second)
	}
}
