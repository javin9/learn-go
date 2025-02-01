package model

// var DB *gorm.DB

// func SetUp() {
// 	// dsn := "root:123456@tcp(127.0.0.1:3306)/gorm_test?charset=utf8mb4&parseTime=True&loc=Local"
// 	dsn := fmt.Sprintf("root:123456@tcp(127.0.0.1:3306)/cms_account?charset=utf8&parseTime=True&loc=Local")
// 	var err error
// 	DB, err = gorm.Open(mysql.New(mysql.Config{
// 		DSN:                       dsn,   // DSN data source name
// 		DefaultStringSize:         256,   // string 类型字段的默认长度
// 		DisableDatetimePrecision:  true,  // 禁用 datetime 精度，MySQL 5.6 之前的数据库不支持
// 		DontSupportRenameIndex:    true,  // 重命名索引时采用删除并新建的方式，MySQL 5.7 之前的数据库和 MariaDB 不支持重命名索引
// 		DontSupportRenameColumn:   true,  // 用 `change` 重命名列，MySQL 8 之前的数据库和 MariaDB 不支持重命名列
// 		SkipInitializeWithVersion: false, // 根据当前 MySQL 版本自动配置
// 	}), &gorm.Config{})
// 	if err != nil {
// 		fmt.Printf("mysql connect is error %v", err)
// 		panic(err)
// 	}
// 	sqlDB, err := DB.DB()
// 	if err != nil {
// 		log.Fatalln("failed to  db.DB()")
// 		panic(err)
// 	}
// 	sqlDB.SetMaxIdleConns(10)
// 	sqlDB.SetMaxOpenConns(100)
// 	fmt.Println("mysql connect successfully")
// 	DB.AutoMigrate(&Account{})
// }
