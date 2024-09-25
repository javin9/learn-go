package main

import "fmt"

func main() {
	enums()
}

// 定义新类型
type Color int

func enums() {
	const (
		java Color = iota
		python
		javascript
	)

	language := java

	switch language {
	case java:
		print("java \n")
	case python:
		print("python")
	case javascript:
		print("javascript \n")
	}
	fmt.Println("语言", java)
}
