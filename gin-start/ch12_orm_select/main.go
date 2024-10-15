package main

import (
	"database/sql"
	"encoding/json"
	"errors"
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
	var user User

	result := db.First(&user)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		fmt.Println("未找到")
	} else {
		fmt.Println(user.Name)
	}

	byteList, err := json.Marshal(user)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(byteList))

	db.First(&user, []int{1, 2, 3})

	// ------- 查询列表
	var userList []User
	db.Find(&userList)
	for _, v := range userList {
		fmt.Println(v.ID, v.Name)
	}

	fmt.Println("Get first matched record")

	var currentUser User
	db.Where("Name = ?", "zhangsan").First(&currentUser)
	// SELECT * FROM `users` WHERE Name = ? AND `users`.`deleted_at` IS NULL ORDER BY `users`.`id` LIMIT ?
	db.Where(&User{Name: "zhangsan"}).First(&currentUser)
	// SELECT * FROM `users` WHERE `users`.`name` = ? AND `users`.`deleted_at` IS NULL AND `users`.`id` = ? ORDER BY `users`.`id` LIMIT ?

	db.Select("name", "age").Find(&userList)
	for _, v := range userList {
		fmt.Println(v.Age)
	}
	// 不等于 <>
	db.Where("name <> ?", "javin9").First(&currentUser)
	byteList, err = json.Marshal(currentUser)
	if err != nil {
		return
	}
	fmt.Println(string(byteList))
	//
	//

	db.Where("name <> ?", "javin9").Find(&userList)
	for _, v := range userList {
		fmt.Println(v.ID, v.Name)
	}
}
