package main

import (
	"context"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	r := gin.Default()
	r.GET("/api/get", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "success"})
	})

	srv := &http.Server{
		Addr:    "8085",
		Handler: r,
	}
	//开启一个协程
	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen:%s \n", err)
		}
	}()
	//类型 os.Signal
	quit := make(chan os.Signal)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	//添加阻塞
	<-quit
	log.Println("shutdown server")
	ctx, cancelFunc := context.WithTimeout(context.Background(), 5*time.Second)
	//取消
	defer cancelFunc()
	//关闭
	err := srv.Shutdown(ctx)
	if err != nil {
		log.Fatalln("server shutdown", err)
	}
	log.Println("server exit")
}
