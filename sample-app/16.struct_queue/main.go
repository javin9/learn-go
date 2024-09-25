package main

import (
	"fmt"
	"sample-app/16.struct_queue/queue"
)

func main() {
	fmt.Println("----")
	q := queue.Queue{}
	q.Push(1)
	q.Push(2)
	q.Push(3)
	q.Push(4)
	q.Print()

	p := q.Pop()
	fmt.Printf("pop value is %d \n", p)

	q.Print()
}
