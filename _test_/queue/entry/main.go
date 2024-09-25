package main

import (
	"fmt"
	"github.com/rupid/learn-gin/_test_/queue"
)

func main() {
	q := queue.Queue{}
	q.Push(1)
	q.Push(3)
	q.Push(4)
    q.Push(2)
	fmt.Println(q.Pop())
	fmt.Println(q.Pop())
	fmt.Println(q.Pop())

	//q.Clear()
	fmt.Println(q.IsEmpty())
}
