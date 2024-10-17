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

func connectDB() (db *gorm.DB) {
	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
		logger.Config{
			SlowThreshold:             time.Second, // Slow SQL threshold
			LogLevel:                  logger.Info, // Log level
			IgnoreRecordNotFoundError: true,        // Ignore ErrRecordNotFound error for logger
			ParameterizedQueries:      true,        // Don't include params in the SQL log
			Colorful:                  false,       // Disable color
		},
	)
	// 参考 https://gorm.io/zh_CN/docs/connecting_to_the_database.html
	dsn := "root:123456@tcp(127.0.0.1:3306)/gorm_test?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: newLogger,
	})
	if err != nil {
		panic("failed to connect database")
	}
	return db
}

func main() {
	fmt.Println("--")
	db := connectDB()
	// 设置全局的logger，打印没条sql语句
	// 迁移 schema
	db.AutoMigrate(&Product{})

	// Create
	db.Create(&Product{Code: sql.NullString{String: "R67", Valid: true}, Price: 100, Name: "橘子"})

	// Read
	var product Product
	db.First(&product, 1)                 // 根据整型主键查找
	db.First(&product, "code = ?", "D42") // 查找 code 字段值为 D42 的记录
	var product2 Product
	db.Where("code = ? AND name = ?", "D43", "苹果").First(&product2)
	fmt.Printf("product2=%v \n", product2.Code)
	fmt.Printf("product=%v \n", product)
	fmt.Println(product.Code)
	// fmt.Println()

	// // Update - 将 product 的 price 更新为 200

	// // Update - 更新多个字段
	db.Model(&product).Updates(Product{Price: 200, Code: sql.NullString{String: "Name", Valid: true}, Email: "13123@qq.com"}) // 仅更新非零值字段
	db.Model(&product).Update("Email", "")
	// db.Model(&product).Updates(map[string]interface{}{"Price": 200, "Code": "F42"})

	// // Delete - 删除 product
	// db.Delete(&product, 1)

}
