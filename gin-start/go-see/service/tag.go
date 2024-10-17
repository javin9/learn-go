package service

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// service

type CreateTagDto struct {
	Name string `json:"name" binding:"required,min=5,max=10"`
}

func CreateTag(ctx *gin.Context) {
	var createTagDto CreateTagDto
	if err := ctx.ShouldBind(&createTagDto); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "success",
	})
}
