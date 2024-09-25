package main

import "fmt"

func main() {
	ch := make(chan int)

	go func() {
		for i := 0; i < 5; i++ {
			fmt.Println("--------")
			ch <- i
		}
		close(ch) // 关闭通道
	}()

	for v := range ch {
		fmt.Println("*********")
		fmt.Println(v)
	}
	// 此处会打印 0 到 4
}

/**
// 在这个示例中，当通道被关闭后，for range循环会检测到通道关闭，并会正常退出循环而不会产生错误。

// 总之，close在Go语言中是用于安全地关闭通道并通知接收者，避免再发送数据，确保程序运行正确和高效的一种机制
*/
