package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.GET("/welcome", welcome)
	router.GET("/get_struct_json", getJson)
	router.POST("/form_post", formPost)
	router.Run(":3333")
}

func welcome(ctx *gin.Context) {
	//{{host}}/welcome?name=jack&age=19
	name := ctx.DefaultQuery("name", "javin9")
	// age := ctx.Query("age")
	age := ctx.DefaultQuery("age", "18")
	ctx.JSON(http.StatusOK, gin.H{
		"name": name,
		"age":  age,
	})
}

func formPost(ctx *gin.Context) {
	name := ctx.PostForm("name")
	age := ctx.DefaultPostForm("age", "18")
	ctx.JSON(http.StatusOK, gin.H{
		"name": name,
		"age":  age,
	})
}

// ref: https://swaggo.github.io/swaggo.io/declarative_comments_format/api_operation.html
// @Summary getJson
// @Description get string by ID
// @Tags accounts
// @Accept  json
// @Produce  json
// @Param id path string true "Account ID"
// @Success 200 {object} model.Account
// @Failure 400 {object} model.HTTPError
// @Router /accounts/{id} [get]
func getJson(c *gin.Context) {
	type Person struct {
		Name string `json:"user"`
		Age  int    `json:"age"`
	}

	p := Person{Name: "javin9", Age: 18}
	c.JSON(http.StatusOK, p)
}
