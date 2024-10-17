package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

const SESSION_KEY = "SESSION_ID"

func SessionAuth() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		sessionID := ctx.GetHeader(SESSION_KEY)
		if sessionID == "" {
			ctx.AbortWithStatusJSON(http.StatusForbidden, gin.H{
				"message": "session is null",
			})
		}
		ctx.Next()
	}
}
