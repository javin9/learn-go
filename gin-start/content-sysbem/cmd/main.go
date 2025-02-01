package main

import (
	"content-sysbem/internal/api"
	"content-sysbem/internal/validator_format"
	"fmt"

	"github.com/gin-gonic/gin"
)

func main() {
	// rdb.Setup()
	// model.SetUp()
	validator_format.Setup("zh")
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	api.CmsRouter(r)

	err := r.Run()
	if err != nil {
		fmt.Printf("run error %v", err)
		return
	}
}
