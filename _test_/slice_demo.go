package main

import "fmt"

func main() {
	//extendElement()
	appendElement()
}

func extendElement() {
	//slice是对数组的view
	var arr = [...]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	s1 := arr[2:6] //[2,3,4,5]
	s2 := s1[3:5]  //[5,6] 注意这个结果
	fmt.Println(arr)
	fmt.Printf("s1=%v, len(s1)=%v, cap(s1)=%v \n", s1, len(s1), cap(s1))
	fmt.Printf("s2=%v, len(s2)=%v, cap(s2)=%v \n", s2, len(s2), cap(s2))
	// 容量 一直到数组最后，试着修改数组的arr的值，可以看到容易在发生发变化
	//	slice可以向后拓展，不是向前扩展
	//  s[i] 不可逾越len(s),向后拓展不可以超越底层数组cap(s)
}

func appendElement() {
	var arr = [...]int{0, 1, 2, 3, 4, 5, 6, 7}
	s1 := arr[2:6]
	s2 := s1[3:5]
	s3 := append(s2, 10)
	s4 := append(s3, 11)
	s5 := append(s4, 12)
	s6 := append(s5, 13)
	s7 := append(s6, 14)
	s8 := append(s7, 15)
	s9 := append(s8, 16)
	fmt.Printf("s1=%v, len(s1)=%v, cap(s1)=%v \n", s1, len(s1), cap(s1))
	fmt.Printf("s2=%v, len(s2)=%v, cap(s2)=%v \n", s2, len(s2), cap(s2))
	fmt.Printf("s3=%v,len(s3)=%v,cap(s3)=%v \n", s3, len(s3), cap(s3))
	// s4和s5不在是对arr的view，而且重新创建了一个数组，背书是容量的2倍
	fmt.Printf("s4=%v,len(s4)=%v,cap(s4)=%v \n", s4, len(s4), cap(s4))
	fmt.Printf("s5=%v,len(s5)=%v,cap(s5)=%v \n", s5, len(s5), cap(s5))
	fmt.Printf("s6=%v,len(s6)=%v,cap(s6)=%v \n", s6, len(s6), cap(s6))
	fmt.Printf("s7=%v,len(s7)=%v,cap(s7)=%v \n", s7, len(s7), cap(s7))
	fmt.Printf("s8=%v,len(s8)=%v,cap(s8)=%v \n", s8, len(s8), cap(s8))
	fmt.Printf("s9=%v,len(s9)=%v,cap(s9)=%v \n", s9, len(s9), cap(s9))
	//- 添加元素时超宇cap，系统会重新分配更大的底层数据，2倍关系，原来的数组会默认被垃圾回收
	//	append 必须接受返回值
}
