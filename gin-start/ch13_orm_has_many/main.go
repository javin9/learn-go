package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type MallUser struct {
	gorm.Model
	Name        string
	CreditCards []CreditCard `gorm:"foreignKey:MallUserID"`
}

type CreditCard struct {
	gorm.Model
	Number     string
	MallUserID uint
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
	// db.AutoMigrate(MallUser{}, CreditCard{}) //
	db.AutoMigrate(&MallUser{}, &CreditCard{}) //

	// mallUser := MallUser{Name: "javin"}
	// db.Create(&mallUser)
	// db.Create(&[]CreditCard{{Number: "111", MallUserID: mallUser.ID}, {Number: "222", MallUserID: mallUser.ID}, {Number: "333", MallUserID: mallUser.ID}})
	// // db.Create(&MallUser{Name})
	var userList MallUser
	db.Preload("CreditCards").First(&userList) //CreditCards 必须加s
	for _, v := range userList.CreditCards {
		fmt.Println(v.Number)
	}
}
