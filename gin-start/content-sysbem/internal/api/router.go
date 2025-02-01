package api

import (
	services "content-sysbem/internal/services/cms"

	"github.com/gin-gonic/gin"
)

func CmsRouter(r *gin.Engine) {
	// 创建路由组
	cmsApp := services.NewCmsApp()
	sessionAuth := &SessionAuth{cmsApp: cmsApp}
	apiCmsGroup := r.Group("/api/cms").Use(sessionAuth.Auth)
	{
		apiCmsGroup.GET("/test", cmsApp.Hello)
	}

	apiOutGroup := r.Group("/api/out")
	{
		// /api/out/register
		apiOutGroup.POST("/register", cmsApp.HandleRegister)
		// /api/out/login
		apiOutGroup.POST("/login", cmsApp.LoginController)
	}

	apiContent := r.Group("/api/content").Use(sessionAuth.Auth)
	{
		// 创建文章
		apiContent.POST("/article/create", cmsApp.ContentCreate)
		// 更新文章
		apiContent.POST("/article/update", cmsApp.ContentUpdate)
		// 删除文章
		apiContent.POST("/article/delete", cmsApp.ContentDelete)
		// 查询文章
		apiContent.GET("/article/find", cmsApp.ContentFind)
	}

}
