package main

import "fmt"

func main() {
	//0.创建
	var mp1 map[string]int      //mp1===empty map
	mp2 := make(map[string]int) //nil
	fmt.Println(mp1, mp2)
	obj := make(map[int]string)
	//1.赋值
	//var obj = make(map[int]string)
	obj[1] = "beijing"
	obj[2] = "shanghai"
	obj[3] = "guangzhou"
	obj[4] = "shenzhen"

	for key, value := range obj {
		fmt.Printf("key=%v,value=%v \n", key, value)
	}
	//2.遍历
	if mapValue, ok := obj[1]; ok {
		fmt.Printf("key=1的值%v \n", mapValue)
	} else {
		fmt.Println("key不存在")
	}

	//3.删除 key，value
	delete(obj, 1)
	fmt.Println(obj)
	//4.判断map的key和value是否存在，用ok的方式
	if s, ok := obj[9]; ok {
		fmt.Println(s)
	}
	s, ok := obj[8]
	if ok {
		fmt.Println(s, ok)
	}

	fmt.Println(len(obj))
}
