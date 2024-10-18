package tag_service

import (
	"github/javin9/go-see/model"
	"github/javin9/go-see/pkg/app"
	"github/javin9/go-see/pkg/e"
	validator_translation "github/javin9/go-see/validator"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

// service
type CreateTagDto struct {
	Name  string `json:"name" binding:"required,min=2,max=10"`
	State int    `json:"state"`
}

func CreateTag(ctx *gin.Context) {
	appG := app.Gin{C: ctx}
	var createTagDto CreateTagDto
	if err := ctx.ShouldBind(&createTagDto); err != nil {
		if message, ok := err.(validator.ValidationErrors); ok {
			appG.Response(http.StatusBadRequest, e.INVALID_PARAMS, validator_translation.FormMessage(message))
			return
		}
		appG.Response(http.StatusBadRequest, e.INVALID_PARAMS, err.Error())
		return
	}
	id, err := model.AddTag(createTagDto.Name, createTagDto.State)
	if err != nil {
		appG.Response(http.StatusBadRequest, e.ERROR, err.Error())
		return
	}
	appG.Response(http.StatusOK, e.SUCCESS, map[string]uint{"id": id})
}

// service
type DeleteTagDto struct {
	ID uint `json:"id" binding:"required"`
}

func DeleteTag(ctx *gin.Context) {
	appG := app.Gin{C: ctx}
	var deleteTagDto DeleteTagDto
	if err := ctx.ShouldBind(&deleteTagDto); err != nil {
		if message, ok := err.(validator.ValidationErrors); ok {
			appG.Response(http.StatusBadRequest, e.INVALID_PARAMS, validator_translation.FormMessage(message))
			return
		}
		appG.Response(http.StatusBadRequest, e.INVALID_PARAMS, err.Error())
		return
	}

	if err := model.DeleteTag(deleteTagDto.ID); err != nil {
		appG.Response(http.StatusOK, e.ERROR_DELETE_ARTICLE_FAIL, err.Error())
		return
	}
	appG.Response(http.StatusOK, e.SUCCESS, "")
}