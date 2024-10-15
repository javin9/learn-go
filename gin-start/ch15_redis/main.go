package main

import (
	"context"
	"fmt"
	"log"

	"github.com/go-redis/redis/v8"
)

func main() {
	// 创建一个上下文，默认的背景上下文
	ctx := context.Background()

	// 配置 Redis 客户端
	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:6399", // Redis 服务器地址
		Password: "123456",         // Redis 连接密码
		DB:       0,                // 使用的数据库
	})

	// 检查连接是否成功
	pong, err := rdb.Ping(ctx).Result()
	if err != nil {
		log.Fatalf("无法连接到 Redis: %v", err)
	}
	fmt.Printf("成功连接到 Redis: %s\n", pong)

	// 执行一个 Redis GET 查询命令
	key := "my-key"
	val, err := rdb.Get(ctx, key).Result()
	if err == redis.Nil {
		fmt.Printf("Key %s does not exist\n", key)
	} else if err != nil {
		log.Fatalf("查询失败: %v", err)
	} else {
		fmt.Printf("Key %s value: %s\n", key, val)
	}
}
