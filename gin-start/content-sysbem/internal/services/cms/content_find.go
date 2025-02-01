package cms

import (
	"content-sysbem/internal/app"
	"content-sysbem/internal/dao"
	"content-sysbem/internal/model"
	"content-sysbem/internal/validator_format"
	"content-sysbem/pkg/e"
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type ContentFindDto struct {
	ID       int    `json:"id,omitempty"`
	Author   string `json:"author,omitempty"`
	Title    string `json:"title,omitempty"`
	Page     int    `json:"page,omitempty"`
	PageSize int    `json:"page_size,omitempty"`
}

type ContentListResponse struct {
	List  []model.Content `json:"list"`
	Total int             `json:"total"`
}

// ContentFind 查询内容
func (me *CmsApp) ContentFind(ctx *gin.Context) {
	// 设置返回状态
	appResponse := app.AppResponse{GinContext: ctx}

	// 校验参数
	var req ContentFindDto
	if err := ctx.ShouldBindJSON(&req); err != nil {
		if message, ok := err.(validator.ValidationErrors); ok {
			msg := validator_format.FormMessage(message)
			appResponse.ResponseError(e.INVALID_PARAMS, msg)
			return
		}

		appResponse.ResponseError(e.INVALID_PARAMS, err.Error())
		return
	}

	server := dao.NewContentDao(me.mysqlDB)

	// 查询
	contentList, total, err := server.Find(&dao.FindParams{ID: req.ID, Page: req.Page, PageSize: req.PageSize, Title: req.Title, Author: req.Author})
	if err != nil {
		appResponse.ResponseError(e.ERROR, err.Error())
		return
	}
	fmt.Print(contentList)

	list := make([]model.Content, 0, len(contentList))
	// 转换
	for _, v := range contentList {
		list = append(list, model.Content{
			ID:          v.ID,
			Title:       v.Title,
			Description: v.Description,
			Author:      v.Author,
		})
	}
	appResponse.ResponseSuccess(ContentListResponse{List: list, Total: total}, "查询成功")
}
