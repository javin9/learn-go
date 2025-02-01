package cms

import (
	"content-sysbem/internal/app"
	"content-sysbem/internal/dao"
	"content-sysbem/internal/validator_format"
	"content-sysbem/pkg/e"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type ContentDeleteDto struct {
	ID int `json:"id" binding:"required"`
}

type ContentDeleteResponse struct {
	Message string `json:"message"`
}

// ContentDelete 删除内容
func (me *CmsApp) ContentDelete(ctx *gin.Context) {
	// 设置返回状态
	appResponse := app.AppResponse{GinContext: ctx}

	// 校验参数
	var req ContentDeleteDto
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

	// 判断是否存在
	isExist, err := server.ExistByID(req.ID)
	if err != nil {
		appResponse.ResponseError(e.INVALID_PARAMS, err.Error())
		return
	}

	if !isExist {
		appResponse.ResponseError(e.INVALID_PARAMS, "数据不存在")
		return
	}

	// 删除
	err = server.DeleteByID(req.ID)
	if err != nil {
		appResponse.ResponseError(e.ERROR, err.Error())
		return
	}
	appResponse.ResponseSuccess(nil, "删除成功")
}
