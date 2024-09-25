package main

import (
	"fmt"
	"os"
)

func main() {
	filepath := "./demo.txt"
	if fileInfo, err := os.Stat(filepath); err != nil {
		if os.IsNotExist(err) {
			// 文件不存在，则创建文件
			file, err := os.Create(filepath)
			if err != nil {
				fmt.Println("创建文件失败，错误信息：", err)
				return
			}
			// 关闭文件操作
			file.Close()
		} else {
			// 其它错误
			panic(err)
		}
	} else {
		if fileInfo.IsDir() {
			// 件是目录，则创建文件
			file, err := os.Create(filepath)
			if err != nil {
				fmt.Println("创建文件失败，错误信息：", err)
				return
			}
			// 关闭文件操作
			file.Close()
		}
	}
}
