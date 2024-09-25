package main

import (
	"fmt"
	"sample-app/15.struct_module/tree"
)

// 如何拓充其他人的包
type MyTreeNode struct {
	node *tree.TreeNode
}

func (myNode *MyTreeNode) PostOrderTraversal() {
	if myNode == nil || myNode.node == nil {
		return
	}
	leftNode := &MyTreeNode{node: myNode.node.Left}
	leftNode.PostOrderTraversal()

	(&MyTreeNode{node: myNode.node.Right}).PostOrderTraversal()
	myNode.node.Print()
}

func main() {
	fmt.Println("camelcase 方式 定义访问方式")
	// 首字母大写代表 public
	// 首字母小写代表 private
	// 每个目录一个包
	root := tree.CreateNode(12)
	root.Left = &tree.TreeNode{Value: 11, Left: nil, Right: nil}
	root.Left.Right = tree.CreateNode(15)
	root.Right = &tree.TreeNode{Value: 15}
	root.Print()
	root.Traversal()
	fmt.Println("-----")
	(&MyTreeNode{root}).PostOrderTraversal()
}
