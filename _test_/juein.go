package main

//参考资料：https://juejin.cn/post/6954670529357807647
import (
	"fmt"
	"reflect"
)

type UserModel struct {
	Name string `data:"test" test:"hah"`
	Age  int    `json:"age"`
}

func main() {
	user := UserModel{"太凉", 18}
	userValue := reflect.ValueOf(user)

	userValueField0 := userValue.Field(0)
	userValueFieldName := userValue.FieldByName("Name")
	fmt.Println(userValue, userValueField0, userValueFieldName)

	userType := reflect.TypeOf(user)
	tagGet := userType.Field(0).Tag.Get("data")
	fmt.Println(userType, tagGet)
}
