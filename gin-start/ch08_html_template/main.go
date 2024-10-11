package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.LoadHTMLGlob("template/*")

	router.GET("/index", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", gin.H{
			"title": "开心每一天",
		})
	})

	router.GET("/qwer", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"title": "开心每一天",
		})
	})
	router.Run(":3333")
}

// go build main.go
//  ./main
