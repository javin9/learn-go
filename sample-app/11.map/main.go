package main

import (
	"fmt"
)

// 1.每个后面都需要 ,
func createMap() {
	javin1 := map[string]any{
		"name":   "javin",
		"avatar": "http://www.google.com",
		"age":    18,
	}
	// 可以指定长度
	javin2 := make(map[string]int, 30) //empty map

	// 通过var定义
	var javin3 map[string]int //javin3 == nil true

	fmt.Println(javin1, javin2, javin3, javin3 == nil)
}

// 遍历,use range
// map的key是无序的
func traversingMap() {
	fmt.Println("** traversing map")
	javin1 := map[string]any{
		"name":   "javin",
		"avatar": "http://www.google.com",
		"age":    18,
	}
	for key, v := range javin1 {
		fmt.Printf("key=%s,value=%v \n", key, v)
	}

	// 如果值没有就是nil
	fmt.Printf("name=%s,name=%s \n", javin1["name"], javin1["fname"]) //name=javin,name=%!s(<nil>)
	fmt.Println(javin1["fname"] == nil)                               // true

	// 如何判断key是否存在
	name, ok := javin1["errorname"]
	fmt.Println(name, ok)
}

// 删除
func deleteKey() {
	javin1 := map[string]any{
		"name":   "javin",
		"avatar": "http://www.google.com",
		"age":    18,
	}

	delete(javin1, "age")
	fmt.Println(javin1)
}

// 添加
func appendData() {
	javin1 := map[string]any{
		"name":   "javin",
		"avatar": "http://www.google.com",
		"age":    18,
	}

	javin1["nickname"] = "tailaing"
	fmt.Println(javin1)
}

func main() {
	fmt.Println("map new")
	createMap()
	traversingMap()
	deleteKey()
	appendData()
}
