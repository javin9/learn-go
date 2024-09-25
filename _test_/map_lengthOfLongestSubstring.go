package main

import (
	"fmt"
	"strings"
)

// 数据的操作都在strings上面
func main() {
	fmt.Println(strings.Fields("你好啊"))
}

func lengthOfLongestSubstring(s string) int {

	if sLen := len(s); sLen == 0 || sLen == 1 {
		return sLen
	}
	var sliceString []byte = []byte(s)

	var left, max int = 0, 0
	mapInstance := make(map[byte]int)
	for right, ch := range sliceString {
		v, ok := mapInstance[ch]
		if ok && v >= left {
			left = v + 1
		}
		temp := right - left + 1
		if max < temp {
			max = temp
		}
		mapInstance[ch] = right
		right++
	}
	return max
}

func lengthOfLongestSubstring2(s string) {

}
