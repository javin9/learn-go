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
	"golang.org/x/crypto/bcrypt"
)

type RegisterRequest struct {
	Email           string `json:"email" binding:"email,required"`
	Password        string `json:"password" binding:"required"`
	ConfirmPassword string `json:"confirm_password" binding:"required,eqfield=Password"`
	Nickname        string `json:"nickname" binding:"required"`
}

type RegisterResponse struct {
	UserID   string `json:"user_id" binding:"required,min=2,max=20"`
	Password string `json:"password" binding:"required"`
	Nickname string `json:"nickname" binding:"required"`
}

func (me *CmsApp) HandleRegister(ctx *gin.Context) {
	appResponse := app.AppResponse{GinContext: ctx}
	var body RegisterRequest
	if err := ctx.ShouldBindJSON(&body); err != nil {
		if message, ok := err.(validator.ValidationErrors); ok {
			msg := validator_format.FormMessage(message)
			appResponse.ResponseError(e.INVALID_PARAMS, msg)
			return
		}

		appResponse.ResponseError(e.INVALID_PARAMS, err.Error())
		return
	}
	// 判断用户邮箱是否存在
	accountDao := dao.NewAccountDao(me.mysqlDB)
	isExist, err := accountDao.IsExist(body.Email)
	if err != nil {
		appResponse.ResponseError(e.INVALID_PARAMS, err.Error())
		return
	}

	if isExist {
		appResponse.ResponseError(e.ERROR, fmt.Sprintf("%v", "邮箱已经被注册"))
		return
	}

	// TODO: 密码加密
	pwd, err := encryptPwd(body.Password)
	if err != nil {
		appResponse.ResponseError(e.INVALID_PARAMS, "密码加密错误")
		return
	}
	fmt.Printf("crypt result pwd %s", pwd)

	if err := accountDao.Create(&model.Account{Nickname: body.Nickname, Password: pwd, Email: body.Email}); err != nil {
		appResponse.ResponseError(e.ERROR, "保存失败")
		return
	}
	appResponse.ResponseSuccess("", "success")
}

func encryptPwd(pwd string) (string, error) {
	byteList, err := bcrypt.GenerateFromPassword([]byte(pwd), bcrypt.DefaultCost)
	if err != nil {
		fmt.Printf("bcrypt GenerateFromPassword error = %v", err)
		return "", err
	}
	return string(byteList), nil
}
