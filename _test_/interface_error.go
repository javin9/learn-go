package main

import (
	"errors"
	"fmt"
)

func main() {
	const name, id = "bueller", 17
	err := fmt.Errorf("user %q (id %d) not found", name, id)
	fmt.Println(err)
	fmt.Println(err.Error())

	err2 := errors.New("this is erros")
	fmt.Printf("result %v", err2.Error())
	fmt.Printf("result %v", err2)
	//Output:
	//user "bueller" (id 17) not found
}
