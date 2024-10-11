package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.LoadHTMLGlob("template/**/*")

	// router.GET("/index", func(c *gin.Context) {
	// 	c.HTML(http.StatusOK, "index.html", gin.H{
	// 		"title": "开心每一天",
	// 	})
	// })

	router.GET("/goods/list", func(c *gin.Context) {
		c.HTML(http.StatusOK, "goods/list.html", gin.H{
			"title": "商品列表",
		})
	})

	router.GET("/user/list", func(c *gin.Context) {
		c.HTML(http.StatusOK, "user/list.html", gin.H{
			"title": "用户列表",
		})
	})

	router.GET("/game/list", func(c *gin.Context) {
		c.HTML(http.StatusOK, "game/list.html", gin.H{
			"title": "游戏列表",
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

// docker run --hostname=4621eaf76760 --mac-address=02:42:ac:11:00:02 --env=MYSQL_ROOT_PASSWORD=123456 --env=PATH=/usr/local/sbin:/usr/local/bin:/usr/sbin:/usr/bin:/sbin:/bin --env=GOSU_VERSION=1.12 --env=MYSQL_MAJOR=8.0 --env=MYSQL_VERSION=8.0.27-1debian10 --volume=/Users/liujianwei/Documents/docker_data/my_mysql:/var/lib/mysql --volume=/var/lib/mysql --network=bridge -p 3306:3306 --restart=no --runtime=runc -d mysql
