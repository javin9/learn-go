package main

import (
	"bytes"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	router := gin.Default()

	router.POST("/post", func(c *gin.Context) {

		id := c.Query("id")
		page := c.DefaultQuery("page", "0")
		bodyContent, err := ioutil.ReadAll(c.Request.Body)
		if err != nil {
			log.Fatal(err)
			c.String(http.StatusBadRequest, err.Error())
			c.Abort()
		}
		//将数据读回request。body里面
		c.Request.Body = ioutil.NopCloser(bytes.NewBuffer(bodyContent))
		//fmt.Printf("id: %s; page: %s; name: %s; message: %s", id, page, name, message)
		c.String(http.StatusOK, "id: %s; page: %s; bodyContent: %s;", id, page, string(bodyContent))

	})
	router.Run(":8080")
}
