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
	deleteById(1)
	//insert()
	//updateDemo()
	//query1 := "select id,name,age from user where id =?"
	//var u user
	//err1 := queryById(query1, 1, &u)
	//if err1 != nil {
	//	fmt.Println(u)
	//}
	//query2 := "select id,name,age from user where id>?"
	//rows, err2 := query(query2, 0)
	//if err2 != nil {
	//	fmt.Println("查询失败")
	//	return
	//}

	//defer rows.Close()
	//for ; rows.Next(); {
	//	var u1 user
	//	err = rows.Scan(&u1.id, &u1.name, &u1.age)
	//	if err != nil {
	//		return
	//	}

	//	fmt.Println(u1)
	//}

}

func updateDemo() {
	dsn := "update user set name=? where id = ?"
	ret, err := db.Exec(dsn, "hha", 1)
	if err != nil {
		fmt.Println("update failed ")
		return
	}
	n, err2 := ret.RowsAffected()
	if err2 != nil {
		fmt.Println("get affected failed")
	}
	fmt.Println(n)
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

func deleteById(id int) {
	strSql := "delete from user where id =?"
	ret, err := db.Exec(strSql, id)
	if err != nil {
		fmt.Println("delete failed", err)
		return
	}
	n, err1 := ret.RowsAffected()
	if err1 != nil {
		fmt.Println("get affected failed", err1)
	}
	fmt.Println(n)
}
func insert() {
	dsn := `insert into user (name,age) values ("ge",29)`
	ret, err := db.Exec(dsn)
	if err != nil {
		fmt.Println("insert failed", err)
		return
	}

	id, err1 := ret.LastInsertId()
	if err1 != nil {
		fmt.Println("get id failed", err1)
	}

	fmt.Println(id)

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
