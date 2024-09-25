package main

import (
	"fmt"
	"runtime"
)

func main() {
	go func() {
		fmt.Println("aaaaaaaaa")
		test()
		fmt.Println("bbbbbbbbbb")
	}()
	for {

	}
}

func test() {
	defer fmt.Println("cccccccc")
	runtime.Goexit()
	fmt.Println("ddddddd")

}
