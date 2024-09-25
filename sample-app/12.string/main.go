package main

import (
	"fmt"
	"strings"
	"unicode/utf8"
)

func basic() {
	s := "Yes我爱慕课网!"
	fmt.Println(len(s))            //19
	fmt.Printf("%s \n", []byte(s)) // Yes我爱慕课网!
	fmt.Printf("%x \n", []byte(s)) //596573e68891e788b1e68595e8afbee7bd9121
	// length = len([]byte(s))

	for index, v := range []byte(s) {
		if len([]byte(s))-1 == index {
			fmt.Printf("%X \n", v)
		} else {
			fmt.Printf("%X ", v)
		}
	}

	for _, v := range s {
		fmt.Printf("%X \n", v) //v 是 rune类型
	}
	//
	bytes := []byte(s)
	for len(bytes) > 0 {
		r, size := utf8.DecodeRune(bytes)
		fmt.Printf("%c ", r)
		bytes = bytes[size:]
	}

	// for _, ch := range []rune(s) {
	// 	fmt.Printf("%c ", ch)
	// }
}

func printNewLine(s any) {
	fmt.Printf("%v \n", s)
}

func stringsFunction() {
	s := "要安装 Gin 软件包，需要先安装 Go 并设置 Go 工作Go区。"
	book := "《红楼梦》|《西游记》|《水浒传》|《三国演义》"
	trimString := " hello boy ! "
	trimString2 := "@@@hello boy !@@"
	printNewLine(strings.Contains(s, "Gin"))
	printNewLine(strings.Clone(s))
	printNewLine(strings.Count(s, "Go"))
	printNewLine(strings.Fields(s))
	printNewLine(strings.Index(s, "Go"))
	printNewLine(strings.Join([]string{"hello", "world", "!"}, " "))
	printNewLine(strings.Repeat(s, 3))
	printNewLine(strings.Replace(s, "Go", "Golang", -1)) // 要安装 Gin 软件包，需要先安装 Golang 并设置 Golang 工作Golang区。
	printNewLine(strings.Replace(s, "Go", "Golang", 2))  // 要安装 Gin 软件包，需要先安装 Golang 并设置 Golang 工作Go区。
	printNewLine(strings.ReplaceAll(s, "Go", "Java"))    // 要安装 Gin 软件包，需要先安装 Java 并设置 Java 工作Java区。
	printNewLine(strings.Split(book, "|"))               //[《红楼梦》 《西游记》 《水浒传》 《三国演义》]
	printNewLine(strings.SplitAfter(book, "|"))          //[《红楼梦》| 《西游记》| 《水浒传》| 《三国演义》]
	printNewLine(strings.ToLower(s))                     //要安装 gin 软件包，需要先安装 go 并设置 go 工作go区。
	printNewLine(strings.ToUpper(s))
	printNewLine(strings.Trim(trimString, " "))       //hello boy !
	printNewLine(strings.TrimSpace(trimString))       //hello boy !
	printNewLine(strings.Trim(trimString2, "@"))      //hello boy !
	printNewLine(strings.TrimLeft(trimString2, "@"))  //hello boy !@@
	printNewLine(strings.TrimRight(trimString2, "@")) //@@@hello boy !

	// 定义一个函数，用于判断字符是否为空白字符
	// func isSpace(r rune) bool {
	//    return r == ' ' || r == '\t' || r == '\n'
	// }

	str := "\t\n  Hello, World!  \n\t"
	fmt.Println("Original:", fmt.Sprintf("%q", str))

	// 使用 strings.TrimFunc 移除字符串开头和结尾的空白字符
	trimmedStr := strings.TrimFunc(str, func(r rune) bool {
		return r == ' ' || r == '\t' || r == '\n'
	})
	fmt.Println("Trimmed:", fmt.Sprintf("%q", trimmedStr)) //Trimmed: "Hello, World!"
}

func main() {
	basic()
	stringsFunction()
}
