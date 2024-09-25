package main

import "fmt"

func main() {
	var root Node
	root = Node{Value: 1}
	root.Left = &Node{}
	root.Right = &Node{3, nil, nil}
	root.Right.Left = new(Node)
}

type Node struct {
	Value int
	Left  *Node
	Right *Node
}

//记忆方法：接收者  相当于把 (root treeNode) 变量挪动了前面
func (root Node) ShowValue() {
	fmt.Println(root.Value)
}
