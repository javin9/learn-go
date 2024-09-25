package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	r := gin.Default()
	//curl -X GET http://127.0.0.1:8080/123/12323
	r.GET("/:id/:number", func(context *gin.Context) {
		//{"id":"123","number":"12323"}
		fmt.Println(context.Params)
		context.JSON(http.StatusOK, gin.H{
			"id":     context.Param("id"),
			"number": context.Param("number"),
		})
	})
	r.Run(":8080")
}
