package app

import (
	"content-sysbem/pkg/e"
	"net/http"

	"github.com/gin-gonic/gin"
)

type ResponseSchema struct {
	Code    int    `json:"code" binding:"required"`
	Message string `json:"message"`
	Data    any    `json:"data"`
}

type AppResponse struct {
	GinContext *gin.Context
}

func (me *AppResponse) ResponseSuccess(data interface{}, message string) {
	me.GinContext.JSON(http.StatusOK, ResponseSchema{
		Code:    e.SUCCESS,
		Message: message,
		Data:    data,
	})
}

func (me *AppResponse) ResponseError(code int, message string) {
	me.GinContext.JSON(http.StatusOK, ResponseSchema{
		Code:    code,
		Message: message,
	})
}

// func (me *AppResponse) ResponseError(code int, message *string) {
// 	if message == nil {
// 		defaultMessage := ""
// 		message = &defaultMessage
// 	}
// 	me.GinContext.JSON(http.StatusOK, ResponseSchema{
// 		Code:    code,
// 		Message: *message,
// 	})
// }
