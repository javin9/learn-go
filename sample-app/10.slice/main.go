package main

import (
	"fmt"
	"reflect"
)

func printSliceInformation(s []int) {
	fmt.Printf("slice=%v,len(s)=%d,cap(s)=%d \n", s, len(s), cap(s))
}

func basic() {
	fmt.Println("slice demo")

	// slice 本身没有数据，是对底层array的一个 view(视图)
	arr := [...]int{0, 1, 2, 3, 4, 5, 6, 7, 8}

	fmt.Println(" arr[2:6]=", arr[2:6]) //  arr[2:6]= [2 3 4 5]
	fmt.Println(" arr[2:]=", arr[2:])   //  arr[2:]= [2 3 4 5 6 7 8]
	fmt.Println(" arr[:6]=", arr[:6])   //  arr[:6]= [0 1 2 3 4 5]
	fmt.Println(" arr[:]=", arr[:])     //  arr[:]= [0 1 2 3 4 5 6 7 8]

	part_slice := arr[:6]
	part_slice = part_slice[2:]
	fmt.Println("part_slice=", part_slice) //part_slice= [2 3 4 5]
	fmt.Printf("part_slice=%v,len(part_slice)=%d,cap(part_slice)=%d \n", part_slice, len(part_slice), cap(part_slice))
	//  可以向扩展，不可以向前扩展
	fmt.Println("part_slice[3:5]=", part_slice[3:5]) //part_slice[3:5]= [5 6]
}

func slice_method() {
	var s []int
	new_s := append(s, 1)
	//
	fmt.Println("append s= ", s)         // append s=  []
	fmt.Println("append new_s =", new_s) // append new_s = [1]

	for i := 0; i < 20; i++ {
		s = append(s, 2*i+1)
	}
	fmt.Printf("s=%v \n", s)
}

func slice_define() {
	fmt.Print("creating slice \n")
	// 方法1
	s1 := []int{2, 4, 6, 8}
	fmt.Println("s1=", s1, reflect.TypeOf(s1))
	printSliceInformation(s1)
	s2 := make([]int, 16)
	printSliceInformation(s2)

	// 长度为4，cap为32
	s3 := make([]int, 4, 32)
	printSliceInformation(s3)
}

func copySlice() {
	fmt.Println("****copy slice data ")
	base_slice := []int{2, 4, 6}
	var dst_slice []int

	printSliceInformation(base_slice)

	res := copy(dst_slice, base_slice)
	printSliceInformation(dst_slice) //WARN:slice=[],len(s)=0,cap(s)=0 没有拷贝进去
	fmt.Printf("res=%d \n", res)

	dst_slice2 := make([]int, 5)
	res2 := copy(dst_slice2, base_slice)
	printSliceInformation(dst_slice2) //slice=[2 4 6 0 0],len(s)=5,cap(s)=5
	fmt.Printf("res2=%d \n", res2)    // 3
}

func deleteSlice() {
	fmt.Println("delete element from slice")
	s := []int{2, 4, 5, 6, 7, 8, 1, 9}
	// delete 5
	new_s := append(s[:2], s[3:]...)
	printSliceInformation(new_s)
}

func main() {
	basic()
	//
	fmt.Println("----")
	slice_method()
	slice_define()
	copySlice()
	deleteSlice()
}
