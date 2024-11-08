package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	service "github.com/rupid/learn-gin/gin-start/ch16_protobuf/proto/output"
	"google.golang.org/protobuf/proto"
)

func main() {
	router := gin.Default()

	// 1. 序列化过程
	// 实例化结构体
	// 注意这里的包名为我们在proto中定义的！
	// 使用&进行引用
	person := &service.Person{
		Name:  "what the???",
		Age:   10238,
		Email: "no way!",
	}

	// Marshal对结构体执行序列化
	bt, _ := proto.Marshal(person)
	fmt.Printf("%v\n", bt)

	// 2. 反序列化过程
	// 先实例化一个空的结构体
	// Unmarshal反序列化，参数一为序列化的数据，参数二为存储反序列化结果的变量
	unperson := &service.Person{}
	proto.Unmarshal(bt, unperson)

	// 反序列化后得到的结构体可以直接调用对应方法
	// 方法名可以去 user.pb.go 文件内查看
	fmt.Printf("otherArticle.GetAid(): %v\n", unperson.GetName())
	router.GET("/get_person", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, unperson)
	})

	router.Run(":10001")
}
