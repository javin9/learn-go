package cms

import (
	"content-sysbem/internal/model"
	"context"
	"fmt"
	"log"
	"time"

	"github.com/redis/go-redis/v9"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// 结构体
type CmsApp struct {
	mysqlDB *gorm.DB
	redisDB *redis.Client
}

func NewCmsApp() *CmsApp {
	mysqlDB := connectMysql()
	redisDB := connectRedisDB()
	app := &CmsApp{mysqlDB: mysqlDB, redisDB: redisDB}

	return app
}

func connectMysql() *gorm.DB {
	// dsn := "root:123456@tcp(127.0.0.1:3306)/gorm_test?charset=utf8mb4&parseTime=True&loc=Local"
	dsn := fmt.Sprintf("root:123456@tcp(127.0.0.1:3306)/cms_content?charset=utf8&parseTime=True&loc=Local")
	mysqlDB, err := gorm.Open(mysql.New(mysql.Config{
		DSN:                       dsn,   // DSN data source name
		DefaultStringSize:         256,   // string 类型字段的默认长度
		DisableDatetimePrecision:  true,  // 禁用 datetime 精度，MySQL 5.6 之前的数据库不支持
		DontSupportRenameIndex:    true,  // 重命名索引时采用删除并新建的方式，MySQL 5.7 之前的数据库和 MariaDB 不支持重命名索引
		DontSupportRenameColumn:   true,  // 用 `change` 重命名列，MySQL 8 之前的数据库和 MariaDB 不支持重命名列
		SkipInitializeWithVersion: false, // 根据当前 MySQL 版本自动配置
	}), &gorm.Config{})

	if err != nil {
		fmt.Printf("mysql connect is error %v", err)
		panic(err)
	}
	sqlDB, err := mysqlDB.DB()
	if err != nil {
		log.Fatalln("failed to  db.DB()")
		panic(err)
	}
	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetMaxOpenConns(100)
	fmt.Println("mysql connect successfully")
	mysqlDB.AutoMigrate(&model.Account{})
	mysqlDB.AutoMigrate(&model.Content{})
	return mysqlDB
}

var ctx = context.Background()

func connectRedisDB() *redis.Client {
	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:6399",
		Password: "123456", // no password set
		DB:       0,        // use default DB
	})
	// 判断链接是否成功
	if _, err := rdb.Ping(context.Background()).Result(); err != nil {
		panic(err)
	}

	err := rdb.Set(ctx, "key", "value", 0).Err()
	if err != nil {
		panic(err)
	}

	val, err := rdb.Get(ctx, "key").Result()
	if err != nil {
		panic(err)
	}
	fmt.Println("key", val)
	return rdb
}

func (me *CmsApp) Set(key string, value interface{}, expiration *time.Duration) error {
	exp := time.Duration(0)
	if expiration != nil {
		exp = *expiration
	}
	err := me.redisDB.Set(ctx, key, value, exp).Err()
	if err != nil {
		fmt.Printf("key is %s,value is %v ,error is %v", key, value, err)
		return err
	}
	return nil
}

func (me *CmsApp) Get(key string) (string, error) {
	return me.redisDB.Get(ctx, key).Result()
}
