package tag_service

import (
	"fmt"
	"github/javin9/go-see/model"
	"github/javin9/go-see/pkg/app"
	"github/javin9/go-see/pkg/e"
	validator_translation "github/javin9/go-see/validator"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

// service
type CreateTagDto struct {
	Name  string `json:"name" binding:"required,min=2,max=10"`
	State int    `json:"state"`
}

// type QueryKeyWord struct {
// 	KeyWord string `json:"keyword"`
// 	Page    int    `json:"page"`
// 	Limit   int    `json:"limit"`
// }

// func Paginate(db *gorm.DB, page, pageSize int) ([]Tag, error) {
// 	var tags []Tag
// 	result := db.Offset((page - 1) * pageSize).Limit(pageSize).Find(&tags)
// 	return tags, result.Error
// }

func GetTagList(ctx *gin.Context) {
	appG := app.Gin{C: ctx}
	keyword := ctx.DefaultQuery("keyword", "")
	pageString := ctx.DefaultQuery("page", "1")
	limitString := ctx.DefaultQuery("limit", "20")
	page, err := strconv.Atoi(pageString)
	if err != nil {
		page = 1
	}
	limit, limitErr := strconv.Atoi(limitString)
	if limitErr != nil {
		limit = 20
	}

	fmt.Print(page, limit)
	tagList := model.GetTagList(keyword, page, limit)
	appG.ResponseSuccess(app.Pagination{List: tagList, Page: page, Limit: limit})
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
