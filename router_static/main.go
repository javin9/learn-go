package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	r := gin.Default()
	// 必须这种方式运行：go build -o router_s && ./router_s
	// 设置路由和对应的资源
	r.Static("/assets", "./assets")
	r.StaticFS("/static", http.Dir("static"))
	r.StaticFile("/favico.icon", "./facname.ico")
	//curl -X GET http://127.0.0.1:8080/
	r.StaticFile("/", "./index.html")
	r.Run()
}
