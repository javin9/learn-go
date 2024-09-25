package queue

import "fmt"

type Queue []int

func (me *Queue) Push(value int) {
	*me = append(*me, value)
}

func (me *Queue) Pop() int {
	meValue := *me
	head := meValue[0]
	*me = meValue[1:]
	return head
}

func (me *Queue) Print() {
	for _, v := range *me {
		fmt.Println(v)
	}
}

func (me *Queue) Size() int {
	return len(*me)
}

func (me *Queue) IsEmpty() bool {
	return len(*me) == 0
}
