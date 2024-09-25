package main

import (
	"fmt"
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
	r.Use(func(context *gin.Context) {
		startTime := time.Now()
		logger.Info("request logger", zap.String("url", context.Request.URL.Path))
		context.Next()
		logger.Info("response logger", zap.Int("StatusOK", context.Writer.Status()), zap.Duration("Duration", time.Now().Sub(startTime)))
	})
	r.GET("/ping", func(context *gin.Context) {
		time.Sleep(1000)
		context.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})
	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
