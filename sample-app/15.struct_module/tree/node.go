package tree

import "fmt"

// 首字母大写代表 public
// 首字母小写代表 private
// 可以定义很多 东西

type TreeNode struct {
	Value int
	Left  *TreeNode //这里必须是指针 避免拷贝 共享数据
	Right *TreeNode
}

// 给结构体定义方法
func (node TreeNode) Print() {
	fmt.Println(node.Value)
}

func (node *TreeNode) Traversal() {
	if node == nil {
		return
	}
	node.Left.Traversal()
	node.Print()
	node.Right.Traversal()
}

// go中 函数都是传值的，这么写 其实没用
func (node TreeNode) SetValue(value int) {
	node.Value = value
}

func (node *TreeNode) SetValue2(value int) {
	// go 语言中，node虽然是引用类型的结构体，但是可以直接node.Value 。但是 如果是值类型，需要*p=3,可以参考08pointer的示例
	//  在c++中 需要node->Value
	node.Value = value
}

func CreateNode(Value int) *TreeNode {
	return &TreeNode{Value, nil, nil}
}
