package cms

import (
	"content-sysbem/internal/app"

	"github.com/gin-gonic/gin"
)

func (me *CmsApp) Hello(ctx *gin.Context) {
	appResponse := app.AppResponse{GinContext: ctx}
	appResponse.ResponseSuccess("hello world", "success")
}
