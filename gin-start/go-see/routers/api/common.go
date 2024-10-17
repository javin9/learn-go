package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Heartbeat(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"message": "success",
	})
}
