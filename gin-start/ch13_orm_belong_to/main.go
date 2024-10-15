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
	Name      string `gorm:"size:256"`
	Email     string `gorm:"unique"`
	Age       sql.NullInt32
	Birthday  sql.NullTime
	CompanyID int
	Company   Company
}

type Company struct {
	gorm.Model
	Name string
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
	db.AutoMigrate(&User{}) //
	// db.Create(&Company{Name: "好未来"})
	// db.Create(&[]Company{{Name: "小米"}, {Name: "百度"}, {Name: "腾讯"}, {Name: "滴滴"}})
	var user User
	db.Preload("Company").First(&user)
	fmt.Println(user.Name, user.Company.Name)
	db.Joins("Company").First(&user)
}
