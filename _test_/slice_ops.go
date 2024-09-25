package main

import "fmt"

func main() {
	var s []int //Zero value for slice is nil .刚分配之后值有0值的

	for n := 0; n < 10; n++ {
		s = append(s, n)
		fmt.Printf("s=%v,len(s)=%v,cap(s)=%v \n", s, len(s), cap(s))
	}
	//fmt.Println(s)
}
