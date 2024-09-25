package main

import "fmt"

func main() {
	//fmt.Println(evalFunc(1, 3, "+"))
	if i, err := evalFunc(1, 3, "0"); err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(i)
	}
	fmt.Println(getGrade(76))
}

//返回两个值，一个int，一个error
func evalFunc(a int, b int, op string) (int, error) {
	//如果switch后面有值，后面的case必须 是登录
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
		//error返回方式
		return 0, fmt.Errorf("op is not found %v", op)
	}
}

func getGrade(score int) string {
	grade := ""
	//switch可以没有表达式
	switch {
	case score < 0 || score > 100:
		panic("错误")
	case score < 60:
		grade = "D"
	case score >= 60 && score < 80:
		grade = "C"
	case score >= 80 && score < 90:
		grade = "B"
	case score >= 90:
		grade = "A"
	}
	return grade
}
