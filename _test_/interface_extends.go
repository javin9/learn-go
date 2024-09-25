package main

import (
	"errors"
	"fmt"
)

//接口名 是er结尾
type Animaler interface {
	sayHi()
}
type Tigerer interface {
	Animaler //继承
	sing()
}

type Monkey struct {
	name string
}

func (t *Monkey) sayHi() {
	errors.New("name")
	fmt.Println("sayHi")
}

func (t *Monkey) sing() {
	fmt.Println("sing")
}

func main() {
	m := Monkey{name: "ddd"}
	m.sing()
	m.sayHi()

	var abn Animaler
	abn = &m
	abn.sayHi()
}
