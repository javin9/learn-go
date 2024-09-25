package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	r := gin.Default()
	// 泛 绑定，*action
	//访问 /api/user前缀的，全部打到这个get中
	// curl -X GET http://127.0.0.1:8081/api/user/name
	r.GET("/api/user/*action", func(context *gin.Context) {
		context.JSON(http.StatusOK, gin.H{
			"message": "hello world",
		})
	})
	r.Run(":8081")
}
