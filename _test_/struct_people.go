package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	xm := &People{Name: "xiaoming", Address: &Address{City: "北京", Street: "上地", Planet: "cunli", Phone: "110"}}
	xmByte, err := json.Marshal(xm)
	if err != nil {
		return
	}
	fmt.Println(string(xmByte))

	tal := &Company{Name: "tal", Address: []*Address{{City: "北京", Street: "上地", Planet: "cunli", Phone: "110"}}}
	talByte, _ := json.Marshal(tal)
	fmt.Println(string(talByte))
	//dog := Dog{}

	//dog.TypeName
	//dog.TypeName
	//dog.Name
}

type People struct {
	Name    string
	Address *Address
}
type Company struct {
	Name    string
	Address []*Address
}
type Address struct {
	Street string
	City   string
	Planet string
	Phone  string
}

type Anminal struct {
	Name     string
	TypeName int
}

type Dog struct {
	*Anminal
}
