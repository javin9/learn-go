package api

import (
	"content-sysbem/internal/services/cms"
	"content-sysbem/internal/util"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
)

const SESSION_ID = "token"

type SessionAuth struct {
	cmsApp *cms.CmsApp
}

func (me *SessionAuth) Auth(ctx *gin.Context) {
	token := ctx.GetHeader(SESSION_ID)
	if token == "" {
		// ctx.Abort()
		ctx.AbortWithStatusJSON(http.StatusForbidden, gin.H{
			"message": "forbidden",
		})
	}

	token_key := util.GenerateTokenValueKey(token)
	fmt.Printf("token is %s", token_key)
	result, err := me.cmsApp.Get(token_key)
	if err == redis.Nil {
		fmt.Printf("%s does not exist", token_key)
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"message": "登录过期1",
		})
	} else if err != nil {
		fmt.Printf("%s err %v", token_key, err)
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"message": "登录过期3",
		})
	} else {
		fmt.Println("key2", result)
	}

	ctx.Next()
}
