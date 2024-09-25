package embeddingNode

import "sample-app/15.struct_module/tree"

// 内嵌的方式，省略了前面的名字

type MyNode struct {
	*tree.TreeNode
}

// 鼠标hover上去看一下，解释。TreeNode的属性和方法，都是 tansfer给MyNode

func (me *MyNode) PostData() {
	if me == nil {
		return
	}

	meLeftNode := MyNode{me.Left}
	meLeftNode.PostData()

	meRightNode := MyNode{me.Right}
	meRightNode.PostData()
	me.Print()
}
