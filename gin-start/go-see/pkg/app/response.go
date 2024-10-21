package app

import (
	"github/javin9/go-see/pkg/e"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Pagination struct {
	Page  int `json:"page"`
	Limit int `json:"limit"`
	List  any `json:"list"`
}

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

func (me *Gin) ResponseSuccess(data interface{}) {
	me.C.JSON(http.StatusOK, Response{
		Code: e.SUCCESS,
		Msg:  e.GetMsg(e.SUCCESS),
		Data: data,
	})
}
