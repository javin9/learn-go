package main

import (
	"fmt"
	"time"
)

func main() {
	timer := time.NewTimer(3 * time.Second)
	fmt.Println(time.Now())
	afterTime := <-timer.C
	fmt.Println(afterTime)
}
