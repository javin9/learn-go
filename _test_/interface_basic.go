package main

import "fmt"

type Humaner interface {
	SayHi() string
}

type Student struct {
	Id   int
	Name string
}

func (student *Student) SayHi() string {
	fmt.Printf(`id=%v,name=%v `, student.Id, student.Name)
	return "is student"
}

type Teacher struct {
	Id   int
	Name string
}

func (teacher *Teacher) SayHi() string {
	fmt.Printf(`id=%v,teachername=%v \n`, teacher.Id, teacher.Name)
	return "isstenc"
}

//多态
func whoSayHi(who Humaner) {
	who.SayHi()
}

func main() {
	s1 := Student{Id: 1, Name: "tailiang"}
	s1.SayHi()

	var t1 Teacher
	t1.Id = 5
	t1.Name = "wangwu"

	t1.SayHi()
	//多态
	whoSayHi(&s1)
	whoSayHi(&t1)
}
