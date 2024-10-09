package main

import (
	"fmt"
	"net/http"
	"reflect"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/locales/en"
	"github.com/go-playground/locales/zh"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	en_translations "github.com/go-playground/validator/v10/translations/en"
	zh_translations "github.com/go-playground/validator/v10/translations/zh"
)

type LoginForm struct {
	Name     string `json:"name" binding:"required"`
	Password string `json:"pwd" binding:"required"`
	// `binding:"required,len=10,min=5,oneof=red green,required_with=Field1,required_without=Field2"` // https://pkg.go.dev/github.com/go-playground/validator#hdr-Baked_In_Validators_and_Tags
}

type RegisterForm struct {
	Name            string `json:"name" binding:"required,min=5,max=10"`
	Email           string `json:"email" binding:"required,email"`
	Age             uint8  `json:"age" binding:"lte=130,gte=18"`
	Password        string `json:"password" binding:"required"`
	ConfirmPassword string `json:"confirm_password" binding:"required,eqfield=Password"`
}

var trans ut.Translator
var found bool

func main() {
	if err := initTranslations("zh"); err != nil {
		fmt.Println("初始化翻译器错误")
		return
	}

	router := gin.Default()
	router.POST("/login", login)
	router.POST("/register", register)
	router.Run(":3333")
}

// 通过分割取key
func removeTopStruct(fields map[string]string) map[string]string {
	newMap := make(map[string]string)
	for key, value := range fields {
		newKey := strings.Split(key, ".")[1]
		newMap[newKey] = value
	}
	return newMap
}

// 通过 字符串切片取key
func removeTopStruct2(fields map[string]string) map[string]string {
	newMap := make(map[string]string)
	for key, value := range fields {
		newMap[key[strings.Index(key, ".")+1:]] = value
	}
	return newMap
}

func initTranslations(locale string) (err error) {
	//修改gin框架中的validator引擎属性，实现定制
	if validate, ok := binding.Validator.Engine().(*validator.Validate); ok {
		//  注册一个获取json的tag的自定义方法
		validate.RegisterTagNameFunc(func(fld reflect.StructField) string {
			name := strings.SplitN(fld.Tag.Get("json"), ",", 2)[0]
			// skip if tag key says it should be ignored
			if name == "-" {
				return ""
			}
			return name
		})

		zhL := zh.New()
		enL := en.New()
		// 第一参数是备用的语言环境，后面是应该支持的语言环境
		uni := ut.New(enL, zhL, enL)
		trans, found = uni.GetTranslator(locale)
		if !found {
			return fmt.Errorf("uni.GetTranslator(%s)", locale)
		}

		switch locale {
		case "en":
			en_translations.RegisterDefaultTranslations(validate, trans)
		case "zh":
			zh_translations.RegisterDefaultTranslations(validate, trans)
		default:
			en_translations.RegisterDefaultTranslations(validate, trans)
		}
		return
	}
	return
}

func login(c *gin.Context) {
	var loginForm LoginForm
	err := c.ShouldBind(&loginForm)
	if err != nil {
		fmt.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "成功",
	})
}

func register(c *gin.Context) {
	var registerForm RegisterForm
	if err := c.ShouldBind(&registerForm); err != nil {
		if comma, ok := err.(validator.ValidationErrors); ok {
			c.JSON(http.StatusBadRequest, gin.H{
				"message": removeTopStruct(comma.Translate(trans)),
			})
			return
		}
		fmt.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "成功",
	})
}
