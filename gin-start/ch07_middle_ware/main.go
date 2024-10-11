package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func main() {
	fmt.Println("middle_ware")
	router := gin.New()
	router.Use(gin.Logger())
	router.Use(gin.Recovery())
	router.Use(MyLogger())
	router.Use(Author())
	router.GET("/middle_ware/test", MiddleWareTest)
	router.Run(":3333")
}

func Author() gin.HandlerFunc {
	return func(c *gin.Context) {
		var token string
		for key, value := range c.Request.Header {
			// fmt.Printf("%s=%s \n", key, value)
			if key == "X-Token" {
				token = value[0]
			}
		}
		if token != "javin9" {
			c.JSON(http.StatusBadGateway, gin.H{
				"message": "无权限",
			})
			// return //return无效， 必须使用 c.Abort()的方式
			c.Abort()
		}
		c.Next()
	}
}

// https://blog.csdn.net/wohu1104/article/details/126689112

func MyLogger() gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()
		c.Set("example", "12345")

		// before request
		c.Next()
		// after request

		latency := time.Since(start)
		fmt.Printf("latency=%v \n", latency)
		status := c.Writer.Status()
		fmt.Printf("status=%d \n", status)
	}
}

func MiddleWareTest(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "success",
	})
}
