package validator_format

import (
	"fmt"
	"reflect"
	"strings"

	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/locales/en"
	"github.com/go-playground/locales/zh"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	en_translations "github.com/go-playground/validator/v10/translations/en"
	zh_translations "github.com/go-playground/validator/v10/translations/zh"
)

var trans ut.Translator
var found bool

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

func Setup(locale string) (err error) {
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

func FormMessage(message validator.ValidationErrors) string {
	result := removeTopStruct(message.Translate(trans))
	str := ""
	for _, v := range result {
		str = fmt.Sprintf("%s;%s", str, v)
	}
	return str
}
