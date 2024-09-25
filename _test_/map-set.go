package main

import (
	"fmt"
	mapset "github.com/deckarep/golang-set"
)

func main() {
	required := mapset.NewSet[string]()
	required.Add("cooking")
	fmt.Println(required)
	sliceData := []interface{}{"hao", "weilai", "laiaaa"}
	sliceSet := mapset.NewSetFromSlice(sliceData)
	fmt.Println(sliceSet.Contains("hao"))

}
