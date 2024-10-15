package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type Product struct {
	gorm.Model
	Name  string
	Code  sql.NullString
	Email string
	Price uint
}

type User struct {
	gorm.Model
	Name     string `gorm:"size:256"`
	Email    string `gorm:"unique"`
	Age      sql.NullInt32
	Birthday sql.NullTime
}

type CreditCard struct {
	gorm.Model
	Number     string
	UserID     uint
	ConsumerID int
}

type Consumer struct {
	gorm.Model
	Name       string
	CreditCard CreditCard
}

func connectDB() (db *gorm.DB) {
	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
		logger.Config{
			SlowThreshold:             time.Second, // Slow SQL threshold
			LogLevel:                  logger.Info, // Log level
			IgnoreRecordNotFoundError: true,        // Ignore ErrRecordNotFound error for logger
			ParameterizedQueries:      true,        // Don't include params in the SQL log
			Colorful:                  true,        // Disable color
		},
	)
	// 参考 https://github.com/go-sql-driver/mysql#dsn-data-source-name 获取详情
	dsn := "root:123456@tcp(127.0.0.1:3306)/gorm_test?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: newLogger,
	})
	if err != nil {
		panic("failed to connect database")
	}
	return db
}

func appendTime(data map[string]any) map[string]any {
	data["CreatedAt"] = time.Now()
	data["UpdatedAt"] = time.Now()
	return data
}

func batchAppendTime(list []map[string]any) []map[string]any {
	var newList []map[string]any
	for _, item := range list {
		newList = append(newList, appendTime(item))
	}
	return newList
}

func main() {
	fmt.Println("--")
	db := connectDB()

	// 迁移 schema
	// db.AutoMigrate(&Consumer{}, &CreditCard{})
	// db.Create(&Consumer{Name: "zhangsan", CreditCard: CreditCard{Number: "132134"}})
	// var user User
	// db.Model(&User{}).Where(User{Name: "lisi"}).Update("age", 18)

	// 注意下面是  Updates，不是update
	// db.Model(&User{}).Where(User{Name: "zhangsan"}).Updates(map[string]any{"age": 19})

	// 注意下面是  Updates，不是update
	// db.Model(&User{}).Where(User{Name: "wangwu"}).Updates(User{Age: sql.NullInt32{Int32: 22, Valid: true}})

	db.Model(User{}).Where(User{Name: "v_zhangsan"}).Updates(User{Age: sql.NullInt32{Int32: 0, Valid: true}})
}
