package main

// https://pkg.go.dev/fmt
import (
	"fmt"
	"strconv"
)

// Atoi 函数用于将字符串转换为 int 类型，
// Itoa 函数则用于将 int 类型转换为字符串类型
func AtoiAndItoa() {
	fmt.Println("**AtoiAndItoa ")
	str := "123"
	intValue, err := strconv.Atoi(str)
	if err != nil {
		panic("error")
	}
	fmt.Printf("str to int: %d\n", intValue)

	intValue += 1
	str = strconv.Itoa(intValue)
	fmt.Printf("int to str: %s\n", str)
}

func Append() {
	fmt.Println("**Append ")
	// 追加整数到字节数组
	num1 := 123
	byteSlice := []byte("Number: ")
	byteSlice = strconv.AppendInt(byteSlice, int64(num1), 10)
	fmt.Printf("Appended int: %s\n", byteSlice) //Appended int: Number: 123

	// 追加布尔值到字节数组
	boolVal := true
	byteSlice = []byte("Bool: ")
	byteSlice = strconv.AppendBool(byteSlice, boolVal)
	fmt.Printf("Appended bool: %s\n", byteSlice) //Appended bool: Bool: true

	// 追加浮点数到字节数组
	floatVal := 3.14
	byteSlice = []byte("Float: ")
	byteSlice = strconv.AppendFloat(byteSlice, floatVal, 'f', -1, 64)
	fmt.Printf("Appended float: %s\n", byteSlice) //Appended float: Float: 3.14
}

// strconv.Parse 系列函数用于将字符串解析为指定类型。其中常用的函数有 ParseInt、ParseBool 和 ParseFloat。简单使用示例如下
func Parse() {
	fmt.Println("**Parse ")
	// 解析整数
	intStr := "123"
	intValue, _ := strconv.ParseInt(intStr, 10, 64)
	fmt.Printf("Parsed int value: %d\n", intValue)

	// 解析布尔值
	boolStr := "true"
	boolValue, _ := strconv.ParseBool(boolStr)
	fmt.Printf("Parsed bool value: %t\n", boolValue)

	// 解析浮点数
	floatStr := "3.14"
	floatValue, _ := strconv.ParseFloat(floatStr, 64)
	fmt.Printf("Parsed float value: %f\n", floatValue)
}

// strconv.Format 系列函数用于将基本数据类型转换为字符串类型。常用的函数有 FormatInt、FormatBool 和 FormatFloat。简单使用示例如下
func FormatTest() {
	fmt.Println("**FormatTest ")
	// 格式化整数
	intValue := 123
	intStr := strconv.FormatInt(int64(intValue), 10)
	fmt.Printf("Formatted int string: %s\n", intStr)

	// 格式化布尔值
	boolValue := true
	boolStr := strconv.FormatBool(boolValue)
	fmt.Printf("Formatted bool string: %s\n", boolStr)

	// 格式化浮点数
	floatValue := 3.14
	floatStr := strconv.FormatFloat(floatValue, 'f', -1, 64)
	fmt.Printf("Formatted float string: %s\n", floatStr)
}

func Quote() {
	fmt.Println("**Quote ")
	str := `路多辛的, "所思所想"!`

	quoted := strconv.Quote(str)
	fmt.Println("Quoted: ", quoted) //Quoted:  "路多辛的, \"所思所想\"!"

	unquoted, err := strconv.Unquote(quoted)
	if err != nil {
		fmt.Println("Unquote error: ", err)
	} else {
		fmt.Println("Unquoted: ", unquoted) //Unquoted:  路多辛的, "所思所想"!
	}
}

func main() {
	fmt.Println(" strconv function demo")
	Append()
	Quote()
	AtoiAndItoa()
	Parse()
	FormatTest()
}
