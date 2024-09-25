package main

import (
	"fmt"
	"os"
	"path"
)

func main() {
	fmt.Println("conditional statement")
	conditionalStatement(18)
	getPassWd()
	switchTest(2, 3, "+")
	level := getLevel(80)
	fmt.Printf("level %s \n", level)
}

// 条件语句 条件不需要括号
func conditionalStatement(age int) {
	println(age)
	if age > 18 {
		fmt.Print("is over 18")
	} else {
		fmt.Println("is child")
	}
}

func getPassWd() {
	cwd, err := os.Getwd()

	if err != nil {
		fmt.Println(err)
	}
	// 写法1
	content, err := os.ReadFile(path.Join(cwd, "sample-app/passwd.txt"))
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(string(content))
	}

	if content, err := os.ReadFile(path.Join(cwd, "sample-app/passwd.txt")); err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(string(content))
	}
}

// switch会自动break，除非使用fallthrough
// 函数的返回值
func switchTest(a, b int, op string) int {
	var result int
	switch op {
	case "+":
		result = a + b
	case "-":
		result = a + b
	default:
		panic("valid date")
	}
	return result
}

func getLevel(score int) string {
	result := ""
	switch {
	case score > 90:
		result = "A"
	case score >= 80:
		result = "B"
	case score >= 60:
		result = "C"
	default:
		panic(fmt.Sprintf("wrong score %d \n", score)) // fmt.Sprintf可以拼接字符串返回
	}
	return result
}
