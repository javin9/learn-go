package main

import "fmt"

func main() {
	type Student struct {
		name string
		age  int
	}
	var s *Student
	s = new(Student)  //分配空间
	s.name = "dequan" //(*s).name="dequan"
	(*s).age = 18     //s.age=18
	//取值
	fmt.Println(*s)
}
