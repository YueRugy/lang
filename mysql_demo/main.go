package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	dsn := "root:123456@tcp(127.0.0.1:3306)/demo"
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		fmt.Println("connection db failed ", dsn)
		return
	}
	err1 := db.Ping()
	if err1 != nil {
		fmt.Println("connection database failed", err1)
		return
	}
	fmt.Println("连接数据库成功")

}
