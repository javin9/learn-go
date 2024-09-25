package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

type Person struct {
	Name     string    `form:"name"`
	Address  string    `form:"address"`
	Birthday time.Time `form:"birthday" time_format:"2006-01-02"`
}

func main() {
	r := gin.Default()
	r.GET("/api", handleApi)
	r.POST("/api", handleApi)
	r.Run()
}

func handleApi(context *gin.Context) {
	//定义person
	var person Person
	//这里是根据 content-type 来做不同的binding操作
	// & 取值
	err := context.ShouldBind(&person)
	if err == nil {
		context.String(http.StatusOK, "formdata %v", person)
	} else {
		context.String(http.StatusOK, "person  bing error %v", err)
	}

}
