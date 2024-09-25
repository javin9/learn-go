package main

import "fmt"

//http://c.biancheng.net/view/66.html
func main() {
	var version int = 1
	cmd := &Command{} //等价于new(Command)
	cmd.Name = "validate"
	cmd.Var = &version
	cmd.Comment = "validate any form data"
}

type Command struct {
	Name    string
	Var     *int
	Comment string
}

//方法
func (receiver Command) showCommand() {
	fmt.Println(receiver.Var)
}
