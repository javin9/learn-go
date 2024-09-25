package main

import (
	"fmt"
	"os"
)

func main() {
	WriteContentByPath("./demo.txt")
	ReadFileByPath("./demo.txt")
}
func WriteContentByPath(path string) {
	f, err := os.Create(path)
	if err != nil {
		fmt.Println("文件创建失败", err)
		return
	}
	//延迟关闭
	defer f.Close()
	for i := 0; i < 10; i++ {
		content := fmt.Sprintf("我是第%v \n", i)
		_, err := f.WriteString(content)
		if err != nil {
			fmt.Println(err)
		}
	}
}

func ReadFileByPath(path string) {
	data, err := os.ReadFile(path)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(data))
	os.Stdout.Write(data)
}
