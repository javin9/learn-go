package queue

type Queue []int

//使用*Queue是因为 使用指针，当改变之后，外界可以感知到
func (receiver *Queue) Push(v int) {
	*receiver = append(*receiver, v)
}

func (receiver *Queue) Pop() int {
	head := (*receiver)[0]
	*receiver = (*receiver)[1:]
	return head
}

func (receiver *Queue) Clear() bool {
   *receiver=[]int{}
   return true
}

func (receiver *Queue) IsEmpty() bool{
	return len(*receiver)==0
}
