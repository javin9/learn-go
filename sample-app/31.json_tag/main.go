package main

import (
	"encoding/json"
	"fmt"
)

type OrderItem struct {
	ID    int     `json:"id"`
	Name  string  `json:"name"`
	Price float32 `json:"price"`
}

// 打个tag标签，json:... 是写个json.Marshal看的
type OrderStruct struct {
	ID         int         `json:"id"`
	Name       string      `json:"name,omitempty"`
	Quantity   int         `json:"quantity"`
	TotalPrice int         `json:"total_price"`
	Item       []OrderItem `json:"item"`
}

func main() {
	fmt.Println("case1:正常使用")
	o := OrderStruct{
		ID:         10,
		Name:       "白菜",
		Quantity:   12,
		TotalPrice: 24,
	}

	fmt.Printf("order is  %v \n", o)  //order is  {10 白菜 12 24}
	fmt.Printf("order is  %+v \n", o) //order is  {ID:10 Name:白菜 Quantity:12 TotalPrice:24}

	byteList, err := json.Marshal(o)
	if err != nil {
		panic(err)
	}

	fmt.Printf("%s \n", byteList) // {"id":10,"name":"白菜","quantity":12,"total_price":24}
	fmt.Println(string(byteList)) // {"id":10,"name":"白菜","quantity":12,"total_price":24}

	// name 没有数值，直接通过 omitempty 排除掉了
	fmt.Println("case2:name 没有数值，直接通过 omitempty 排除掉了")
	o2 := OrderStruct{
		ID:         10,
		Quantity:   12,
		TotalPrice: 24,
	}

	byteList, err = json.Marshal(o2)
	if err != nil {
		panic(err)
	}

	fmt.Printf("%s \n", byteList) //
	fmt.Println(string(byteList)) //

	// 嵌套
	fmt.Println("case3:嵌套")
	o3 := OrderStruct{
		ID:         10,
		Quantity:   12,
		TotalPrice: 24,
		Item:       []OrderItem{{ID: 1, Name: "1111", Price: 64.3}, {ID: 2, Name: "22222", Price: 67.4}},
	}

	byteList, err = json.Marshal(o3)
	if err != nil {
		panic(err)
	}

	fmt.Printf("%s \n", byteList) //
	fmt.Println(string(byteList)) //

	// 还原回去
	orderString := string(byteList)
	var target OrderStruct
	json.Unmarshal([]byte(orderString), &target)
	// fmt.Println(target)
	fmt.Printf("%+v \n", target)
}
