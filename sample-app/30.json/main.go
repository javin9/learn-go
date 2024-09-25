package main

import (
	"encoding/json"
	"fmt"
)

type OrderStruct struct {
	ID       int
	Name     string
	Quantity int
	Price    int
}

func main() {
	fmt.Println("json practice")
	o := OrderStruct{
		ID:       10,
		Name:     "白菜",
		Quantity: 12,
		Price:    24,
	}

	fmt.Printf("order is  %v \n", o)  //order is  {10 白菜 12 24}
	fmt.Printf("order is  %+v \n", o) //order is  {ID:10 Name:白菜 Quantity:12 Price:24}

	byteList, err := json.Marshal(o)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s \n", byteList) // 虽然是byte类型的切片，里面字符串所以可以直接%s处理
	fmt.Println(string(byteList)) //{"ID":10,"Name":"白菜","Quantity":12,"Price":24}
}
