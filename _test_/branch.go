package main

import (
	"fmt"
	"io/ioutil"
)

func main() {
	if content, err := ioutil.ReadFile("a.txt"); err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("%s \n", content)
	}
}
