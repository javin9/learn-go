package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	r := gin.Default()
	//curl -X GET http://127.0.0.1:8080/api/get
	r.GET("/api/get", func(context *gin.Context) {
		//context.String(200, "/api/get")
		context.JSON(http.StatusOK, gin.H{"message": "success"})
	})
	//curl -X POST http://127.0.0.1:8080/api/save
	r.POST("/api/save", func(context *gin.Context) {
		context.String(200, "/api/save")
	})
	//curl -X GET http://127.0.0.1:8080/api/generalist
	r.Handle("GET", "/api/generalist", func(context *gin.Context) {
		context.String(200, "/api/generalist")
	})
	//curl -X GET http://127.0.0.1:8080/api/any
	r.Any("/api/any", func(context *gin.Context) {
		context.String(200, "/api/any")
	})

	r.Run()
}
