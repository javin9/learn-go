package main

import "fmt"

// go 语言仅支持封装，不支持继承和多肽。继承和多肽是通过接口来完成的
//	go语言 没有class ，只有struct
// 为结构体定义方法

type TreeNode struct {
	value int
	left  *TreeNode //这里必须是指针 避免拷贝 共享数据
	right *TreeNode
}

// 给结构体定义方法
func (node TreeNode) print() {
	fmt.Println(node.value)
}

func (node *TreeNode) traversal() {
	if node == nil {
		return
	}
	node.left.traversal()
	node.print()
	node.right.traversal()
}

// go中 函数都是传值的，这么写 其实没用
func (node TreeNode) setValue(value int) {
	node.value = value
}

func (node *TreeNode) setValue2(value int) {
	// go 语言中，node虽然是引用类型的结构体，但是可以直接node.value 。但是 如果是值类型，需要*p=3,可以参考08pointer的示例
	//  在c++中 需要node->value
	node.value = value
}

func basic() {

	root := TreeNode{value: 1, left: nil, right: nil}
	root.left = &TreeNode{value: 2} //left是个指针，在go中，root是可以直接root.left访问
	fmt.Println(root.value, root.left.value, root.right)

	// 函数内容创建函数
	createNode := func(value int) *TreeNode {
		return &TreeNode{value: value, left: nil, right: nil}
	}
	root.right = createNode(3)
	root.left.left = createNode(4)
	root.right.left = createNode(5)

	// nil 代表empty value
	nodes := []TreeNode{
		{value: 2, left: nil, right: nil},
		{3, nil, nil},
	}
	fmt.Println(nodes)

	res := traversal(&root, []int{})
	fmt.Println(res)
	//
	root.print() //1
	root.setValue(13)
	root.print() //1
	root.setValue2(16)
	root.print() //16
	//  go 会做一些人性化的操作 无论node是值传递，直接复制一份。如果是指针 直接调用
	pRoot := &root
	pRoot.print() //
	pRoot.setValue2(22)
}

func traversal(root *TreeNode, arr []int) []int {
	// root.print()
	arr = append(arr, root.value)
	if root.left != nil {
		arr = traversal(root.left, arr)
	}
	if root.right != nil {
		arr = traversal(root.right, arr)
	}
	return arr
}

func main() {
	fmt.Println("struct practice")
	basic()
}
