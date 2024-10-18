package app

import (
	"github/javin9/go-see/pkg/e"

	"github.com/gin-gonic/gin"
)

type Response struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
	Data any    `json:"data"`
}

type Gin struct {
	C *gin.Context
}

func (me *Gin) Response(httpCode int, code int, data interface{}) {
	me.C.JSON(httpCode, Response{
		Code: code,
		Msg:  e.GetMsg(code),
		Data: data,
	})
}
