package main

type Person struct {
	Name    string `form:"name" binding:"required"`
	Age     int    `form:"age" binding:"required,gt=10"`
	Address string `form:"address"`
}

func main() {
	// instance := gin.Default()
	// instance.POST("/api/createuser", func(context *gin.Context) {
	// 	var person Person
	// 	err := context.ShouldBind(&person)
	// 	if err != nil {
	// 		context.JSON(http.StatusInternalServerError, gin.H{"message": "服务器内部错误"})
	// 		return
	// 	}
	// 	//转换成对象返回
	// 	//buf, jsonerr := json.Marshal(person)
	// 	//if jsonerr != nil {
	// 	//	context.JSON(http.StatusInternalServerError, gin.H{"message": "字符串转换失败"})
	// 	//	return
	// 	//}
	// 	////"{\"Name\":\"太凉\",\"Age\":11,\"Address\":\"地址信息\"}"
	// 	//context.JSON(http.StatusOK, string(buf))
	// 	context.JSON(http.StatusOK, gin.H{
	// 		"name":    person.Name,
	// 		"age":     person.Age,
	// 		"address": person.Address,
	// 	})
	// })
	// instance.Run(":8080")
}
