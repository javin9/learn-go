package main

import (
	"fmt"
	"math"
	"reflect"
	"runtime"
)

func div(a, b int) (q int, r int) {
	return a / b, a % b
}

// 返回nil
func level(a, b int, op string) (int, error) {
	switch op {
	case "+":
		return a + b, nil
	case "-":
		return a - b, nil
	case "*":
		return a * b, nil
	case "/":
		return a / b, nil
	default:
		return 0, fmt.Errorf("unsupported operation %s", op)
	}
}

// 函数作为参数传递
// reflect可以拿到函数的名字
// runtime
func apply(op func(a, b int) int, a, b int) int {
	p := reflect.ValueOf(op).Pointer()
	opName := runtime.FuncForPC(p).Name()
	fmt.Printf("\n func name is %s ,variable is (%d,%d) \n", opName, a, b)
	return op(a, b)
}

func Pow(a, b int) int {
	return int(math.Pow(float64(a), float64(b)))
}

// range 遍历
func sum(numbers ...int) int {
	result := 0
	for _, value := range numbers {
		// result += numbers[index]
		result += value
	}
	return result
}

func main() {
	fmt.Println("study func")
	a, b := div(4, 2)
	fmt.Printf("a/b=%d,a%%b=%d,\n", a, b)
	result, err := level(3, 4, "x")
	if err != nil {
		fmt.Print("报错了哈")
	} else {
		fmt.Print(result)
	}

	res := apply(Pow, 3, 4)
	// 匿名函数
	res2 := apply(func(a, b int) int {
		return a + b
	}, 3, 4)
	fmt.Println(res, res2)

	// sum
	sum_res := sum(1, 2, 3, 4, 5, 6, 7, 8, 9)
	fmt.Printf("sum_res=%d \n", sum_res)

	// arr := [5]int{10, 20, 30, 40, 50}
	// for index, value := range arr {
	// 	fmt.Printf("Index: %d, Value: %d\n", index, value)
	// }
}
