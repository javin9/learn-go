package main

import (
	"fmt"
	"sample-app/15.struct_module/tree"
	"sample-app/17.struct_embedding/embeddingNode"
)

func main() {
	fmt.Println("embedding ")
	root := tree.CreateNode(12)
	root.Left = &tree.TreeNode{Value: 11, Left: nil, Right: nil}
	root.Left.Right = tree.CreateNode(15)
	root.Right = &tree.TreeNode{Value: 15}

	node := embeddingNode.MyNode{root}
	node.PostData()

	//第二种
	root2 := embeddingNode.MyNode{&tree.TreeNode{Value: 2}}
	root2.Left = &tree.TreeNode{Value: 11, Left: nil, Right: nil}
	root2.Left.Right = tree.CreateNode(15)
	root2.Right = &tree.TreeNode{Value: 15}
	root2.Print()
}
