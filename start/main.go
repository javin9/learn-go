package main

import "github.com/gin-gonic/gin"

func main() {
	r := gin.Default()
	r.GET("/api/hello", func(context *gin.Context) {
		context.JSON(200, gin.H{
			"name": "tailiang",
		})
	})
	r.Run()
}
