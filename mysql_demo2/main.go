package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB

type user struct {
	id   int
	age  int
	name string
}

func main() {
	err := initDB()
	if err != nil {
		fmt.Println(err)
	}

	query1 := "select id,name,age from user where id =?"
	var u user
	err1 := queryById(query1, 1, &u)
	if err1 != nil {
		fmt.Println(u)
	}
	query2 := "select id,name,age from user where id>?"
	rows, err2 := query(query2, 0)
	if err2 != nil {
		fmt.Println("查询失败")
		return
	}

	defer rows.Close()
	for ; rows.Next(); {
		var u1 user
		err = rows.Scan(&u1.id, &u1.name, &u1.age)
		if err != nil {
			return
		}

		fmt.Println(u1)
	}

}

func query(query string, a ...interface{}) (rows *sql.Rows, err error) {
	rows, err = db.Query(query, a...)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	return rows, nil
}

func queryById(query string, id int, u *user) (err error) {
	err = db.QueryRow(query, id).Scan(&(*u).id, &(*u).name, &(*u).age)
	if err != nil {
		fmt.Println(err)
		return
	}
	return nil
}

func initDB() (err error) {
	dsn := "root:123456@tcp(127.0.0.1)/demo"
	db, err = sql.Open("mysql", dsn)
	if err != nil {
		return
	}
	err = db.Ping()
	if err != nil {
		return
	}
	return nil
}
