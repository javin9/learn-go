package main

import (
	"fmt"
	"math/rand"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func main() {
	fmt.Println("gin demo")
	r := gin.Default()
	// middleware
	logger, _ := zap.NewProduction()
	defer logger.Sync() // flushes buffer, if any
	r.Use(func(ctx *gin.Context) {
		startTime := time.Now()

		logger.Info("request logger", zap.String("url", ctx.Request.URL.Path))
		ctx.Next()
		logger.Info("response logger", zap.Int("StatusOK", ctx.Writer.Status()), zap.Duration("Duration", time.Now().Sub(startTime)))
	}, func(ctx *gin.Context) {
		ctx.Set("RequestID", rand.Int())
		ctx.Next()
	})
	r.GET("/ping", func(ctx *gin.Context) {
		time.Sleep(1000)
		RequestID, exists := ctx.Get("RequestID")
		response := gin.H{"message": "pong"}
		if exists {
			response["request_id"] = RequestID
		}
		ctx.JSON(http.StatusOK, response)
	})
	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
