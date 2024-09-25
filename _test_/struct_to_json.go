package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	it := IT{Company: "太凉", Subjects: []string{"hello", "world"}, IsOk: true, Price: 3.2}
	fmt.Println(it)
	buf, err := json.Marshal(it)
	if err != nil {
		fmt.Println("error:", err)
	}
	fmt.Println(buf)
	fmt.Println(string(buf))

}

// 必须大写
type IT struct {
	Company  string
	Subjects []string
	IsOk     bool
	Price    float32
}
