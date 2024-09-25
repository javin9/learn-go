package main

import (
	"fmt"
	"runtime"
)

func main() {
	//子携程还没结束，助携程就退出了
	go func() {
		for i := 0; i < 10; i++ {
			fmt.Printf("sub %v \n", i)
		}
	}()

	for i := 0; i < 3; i++ {
		//让其他协程先执行，其他协程执行完毕后在 执行
		runtime.Gosched()
		fmt.Printf("main %v  \n", i)
	}
}
