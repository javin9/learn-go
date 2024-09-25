package main

import "fmt"

//copy( destSlice, srcSlice []T) int
func main() {
	slice1 := []int{1, 2, 3, 4, 5, 6}
	slice2 := []int{7, 8, 9}
	copy(slice2, slice1)
	//[1 2 3 4 5 6] [1 2 3]
	fmt.Println(slice1, slice2)

	slice3 := []int{1, 2, 3, 4, 5, 6}
	slice4 := []int{7, 8, 9}
	copy(slice3, slice4)
	//[7 8 9 4 5 6] [7 8 9]
	fmt.Println(slice3, slice4)

	//以下代码展示了切片复制是非引用
	const elementCount = 10000

	srcData := make([]int, elementCount)

	for i := 0; i < elementCount; i++ {
		srcData[i] = i
	}
	//	引用切片数据
	refData := srcData
	copyData := make([]int, elementCount)
	//非引用类型
	copy(copyData, srcData)
	srcData[0] = -1
	// 打印引用切片的第一个元素
	fmt.Println(refData[0])
	//
	fmt.Println(copyData[0], copyData[elementCount-1])
}
