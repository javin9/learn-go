package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	//Contains
	fmt.Println(strings.Contains("seafood", "foo"))
	fmt.Println(strings.Contains("seafood", "bar"))
	fmt.Println(strings.Contains("seafood", ""))
	fmt.Println(strings.Contains("", ""))
	//Index
	fmt.Println(strings.Index("chicken", "ken"))
	//Join
	s := []string{"foo", "bar", "baz"}
	fmt.Println(strings.Join(s, ", "))
	//Repeat
	fmt.Println("ba" + strings.Repeat("na", 2))
	//Replace
	fmt.Println(strings.Replace("oink oink oink", "k", "ky", 2))
	fmt.Println(strings.Replace("oink oink oink", "oink", "moo", -1))
	//strings.ReplaceAll()
	fmt.Println(strings.ReplaceAll("oink oink oink", "oink", "moo"))
	//strings.Split()
	fmt.Printf("%q\n", strings.Split("a,b,c", ","))
	fmt.Printf("%q\n", strings.Split("a man a plan a canal panama", "a "))
	fmt.Printf("%q\n", strings.Split(" xyz ", ""))
	fmt.Printf("%q\n", strings.Split("", "Bernardo O'Higgins"))
	//strings.Trim()
	fmt.Print(strings.Trim("¡¡¡Hello, Gophers!!!", "!¡"))
	//strings.Fields()
	fmt.Printf("Fields are: %q", strings.Fields("  foo bar  baz   "))
	//append()
	fmt.Println("\n ===========")
	str := make([]byte, 0, 100)
	str = strconv.AppendInt(str, 10, 10)
	str = strconv.AppendBool(str, false)
	str = strconv.AppendFloat(str, 3.1415926535, 'E', -1, 32)
	fmt.Println(string(str), "\n")
	//	Format
	a := strconv.FormatInt(13423, 10)
	b := strconv.FormatBool(true)
	c := strconv.FormatUint(2341234123, 10)
	fmt.Println(a, b, c)
	//整形转字符串 常用
	var age = strconv.Itoa(122)
	fmt.Printf("\n age=%v \n", age)
	//	字符串转成其他类型
	var d string
	d = "true"
	fmt.Println(strconv.ParseBool(d))
}
