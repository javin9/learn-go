package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	v1_component := r.Group("/v1/component")

	list := []int{1, 2, 3, 4, 5}
	// 可以放到一个中括号里面，也可以不妨到里面
	{
		v1_component.GET("/list", func(ctx *gin.Context) {
			ctx.JSON(http.StatusOK, gin.H{
				"data": list,
			})
		})

		v1_component.GET("/detail/:name", func(ctx *gin.Context) {
			name := ctx.Param("name")
			ctx.JSON(http.StatusOK, gin.H{
				"name": name,
			})
		})
	}

	//
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})
	r.Run(":3333") // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
