package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	r := gin.Default()
	r.GET("/query", func(context *gin.Context) {
		//curl -X GET http://127.0.0.1:8080/query\?firstname\=hello\&nickname\="wokrld"                                                                                ⬡ 14.17.5
		//{"firstname":"hello","nickname":"wokrld"}%
		firstname := context.Query("firstname")
		nickname := context.DefaultQuery("nickname", "冰箱太凉")
		context.JSON(http.StatusOK, gin.H{
			"firstname": firstname,
			"nickname":  nickname,
		})
	})
	r.Run()
}
