package main

import (
	"fmt"
	"os"
)

func deferFunc() {
	defer fmt.Println("1") //会在函数结束之前打印出来，不受return和panic的影响
	defer fmt.Println("2") //会在函数结束之前打印出来
	fmt.Println("3")
}

func main() {
	fmt.Println("资源管理与出错处理")
	deferFunc() // 3,2,1  //defer里面有个 stack(栈)，是先进后出的
}

func writeFile(fileNme string) {
	files, err := os.OpenFile(fileNme, os.O_EXCL|os.O_CREATE, 0666)

	if err != nil {
		// Type assertion
		if value, ok := err.(*os.PathError); ok {
			fmt.Println(value.Op, value.Path, value.Err)
		} else {
			panic(err)
		}
	}
	defer files.Close()

}
