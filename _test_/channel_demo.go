package main

import (
	"fmt"
	"time"
)

var isFinish = make(chan int)

func main() {
	go p1()
	go p2()
	for {

	}
}

func p1() {
	PrintString("bing")
	//给管道添加数据
	isFinish <- 666
}

func p2() {
	//管道取数据，如果管道没有数据就会阻塞
	<-isFinish
	PrintString("xiang")
}

func PrintString(str string) {
	for _, charVar := range str {
		fmt.Printf("%c", charVar)
		time.Sleep(time.Second * 2)
	}
	fmt.Println("")
}
