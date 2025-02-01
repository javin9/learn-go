package main

import (
	"flag"
	"fmt"
)

func main() {

	host := flag.String("host", "", "数据库地址")
	dbName := flag.String("db_name", "", "数据库名称")
	user := flag.String("user", "", "数据库用户")
	password := flag.String("password", "", "数据库密码")
	port := flag.Int("port", 3306, "数据库端口")

	flag.Parse()

	fmt.Printf("数据库地址:%s\n", *host)
	fmt.Printf("数据库名称:%s\n", *dbName)
	fmt.Printf("数据库用户:%s\n", *user)
	fmt.Printf("数据库密码:%s\n", *password)
	fmt.Printf("数据库端口:%d\n", *port)
}
