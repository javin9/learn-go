package cms

import (
	"content-sysbem/internal/app"
	"content-sysbem/internal/dao"
	"content-sysbem/internal/util"
	"content-sysbem/internal/validator_format"
	"content-sysbem/pkg/e"
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

type LoginRequest struct {
	Email    string `json:"email" binding:"email,required"`
	Password string `json:"password" binding:"required"`
}

type LoginResponse struct {
	UserID   int    `json:"user_id" binding:"required,min=2,max=20"`
	Nickname string `json:"nickname" binding:"required"`
	Token    string `json:"token"`
	Email    string `json:"email"`
}

func (me *CmsApp) LoginController(ctx *gin.Context) {
	// 设置返回状态
	appResponse := app.AppResponse{GinContext: ctx}

	// 校验参数
	var req LoginRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		if message, ok := err.(validator.ValidationErrors); ok {
			msg := validator_format.FormMessage(message)
			appResponse.ResponseError(e.INVALID_PARAMS, msg)
			return
		}

		appResponse.ResponseError(e.INVALID_PARAMS, err.Error())
		return
	}
	// server 层
	accountDao := dao.NewAccountDao(me.mysqlDB)
	account, err := accountDao.LoginByEmail(req.Email)
	if err != nil {
		appResponse.ResponseError(e.ERROR, "用户不存在")
		return
	}

	// 判断密码是否正确
	if err := bcrypt.CompareHashAndPassword([]byte(account.Password), []byte(req.Password)); err != nil {
		appResponse.ResponseError(e.ERROR, "密码错误")
		return
	}

	token, err := me.generateToken(account.ID)
	if err != nil {
		appResponse.ResponseError(e.ERROR, "生成token失败")
		return
	}

	appResponse.ResponseSuccess(LoginResponse{
		UserID:   account.ID,
		Nickname: account.Nickname,
		Email:    account.Email,
		Token:    token}, "success")
}

// 生成token
func (me *CmsApp) generateToken(user_id int) (string, error) {
	token := uuid.New().String()
	exp := time.Hour * 8
	key := util.GenerateTokenKey(user_id)
	if err := me.Set(key, token, &exp); err != nil {
		fmt.Printf("generateToken user_id %d err %v", user_id, err)
		return "", err
	}
	token_value_key := util.GenerateTokenValueKey(token)
	fmt.Printf("token_value_key is  %s", token_value_key)
	if err := me.Set(token_value_key, time.Now().Unix(), &exp); err != nil {
		return "", err
	}
	return token, nil
}
