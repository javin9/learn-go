package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	fmt.Println(converIntBoBin(1))
	fmt.Println(converIntBoBin(13))
	fmt.Println(converIntBoBin(12342134))
	printFileContent("a.txt")
}

func converIntBoBin(n int) string {
	//for ; n > 0; n = n / 2 {
	//	fmt.Println(n)
	//}
	if n == 0 {
		return "0"
	}
	result := ""
	for n > 0 {
		lsb := n % 2
		result = strconv.Itoa(lsb) + result
		n = n / 2
	}
	return result
}

//逐行打印文件内容
func printFileContent(filename string) {
	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
}
