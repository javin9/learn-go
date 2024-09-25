package main

import (
	"bufio"
	"fmt"
	"os"
	"path"
	"strconv"
)

func sum(number int) int {
	sum := 0
	for i := 0; i < number; i++ {
		sum += i
	}
	return sum
}

// 在 Go 语言中，并没有专门的 while 关键字来实现循环
func convertToBinary(n int) string {
	result := ""
	if n == 0 {
		return "0"
	}
	for n > 0 {
		// result = fmt.Sprintf("%d", n%2) + result
		result = strconv.Itoa(n%2) + result
		n = n / 2
	}
	return result
}

func printFile() {
	cwd, err := os.Getwd()
	if err != nil {
		panic(err)
	}

	filename := path.Join(cwd, "sample-app/passwd.txt")
	fmt.Println(filename)

	// 打开一个文件
	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	// 挨个扫描
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
}

func main() {
	fmt.Println("loop statement")
	value := 10
	result := sum(10)
	fmt.Printf("value %d sum result is %d \n", value, result)

	res := convertToBinary(134)
	fmt.Printf("%s \n", res)

	printFile()
}
