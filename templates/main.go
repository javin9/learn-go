package main

import (
	"fmt"
	"html/template"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func formatAsDate(t time.Time) string {
	year, month, day := t.Date()
	return fmt.Sprintf("%d%02d/%02d", year, month, day)
}

//https://gin-gonic.com/docs/examples/bind-single-binary-with-template/
func main() {
	router := gin.Default()
	router.Delims("{{", "}}")
	router.SetFuncMap(template.FuncMap{
		"formatAsDate": formatAsDate,
	})
	router.LoadHTMLFiles("./template/index.tmpl")

	router.GET("/index", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.tmpl", gin.H{
			"now":  time.Date(2017, 07, 01, 0, 0, 0, 0, time.UTC),
			"name": "太凉",
		})
	})

	router.Run(":8080")
}
