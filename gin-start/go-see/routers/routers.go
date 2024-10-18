package routers

import (
	"github/javin9/go-see/routers/api"
	"github/javin9/go-see/service/tag_service"

	"github.com/gin-gonic/gin"
)

const (
	rootPathV1 = "/api/v1"
)

func InitRouters() *gin.Engine {
	r := gin.New()
	// 中间件
	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	r.GET("/heartbeat", api.Heartbeat)

	// appV1 := r.Group("/api/v1").Use(session.Auth())
	appV1 := r.Group("/api/v1") //.Use(middleware.SessionAuth())
	{
		appV1.POST("/tag/create", tag_service.CreateTag)
		appV1.POST("/tag/delete", tag_service.DeleteTag)
	}

	return r
}
