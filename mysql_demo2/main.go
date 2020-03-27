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
	tran()
	//preparseQuery(3)
	//preparseInsert()
	//deleteById(1)
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
func preparseInsert() {
	dsn := `insert into user (name,age) values(?,?)`

	stmt, err := db.Prepare(dsn)
	if err != nil {
		fmt.Println("preparse failed", err)
		return
	}
	defer stmt.Close()
	m := map[string]int{
		"feng": 23,
		"diao": 25,
		"long": 27,
	}

	for k, v := range m {
		ret, err1 := stmt.Exec(k, v)
		if err1 != nil {
			fmt.Println("insert failed", err1)
			return
		}
		n, err2 := ret.RowsAffected()
		if err2 != nil {
			fmt.Println("get affected failed", err2)
			return
		}

		fmt.Println(n)
	}

}

func preparseQuery(id int) {
	sqlStr := "select id,name,age from user where id >? "

	stmt, err := db.Prepare(sqlStr)
	if err != nil {
		fmt.Println("preparse failed", err)
		return
	}

	defer stmt.Close()
	rows, err1 := stmt.Query(id)
	if err1 != nil {
		fmt.Println("select failed", err1)
		return
	}
	defer rows.Close()
	sli := make([]user, 0, 20)
	for ; rows.Next(); {
		var u user
		err2 := rows.Scan(&u.id, &u.name, &u.age)
		if err2 != nil {
			fmt.Println("select scan failed ", err2)
			return
		}
		sli = append(sli, u)
	}

	fmt.Println(sli)

}

func tran() {
	tx, err := db.Begin()
	if err != nil {
		fmt.Println("open tran failed", err)
		return
	}
	sqlStr1 := "update user set age=8000 where id =?"
	sqlStr2 := "update user set age=80000 where id =?"

	_, err1 := tx.Exec(sqlStr1, 2)
	if err1 != nil {
		fmt.Println("update failed", err1)
		err = tx.Rollback()
		if err != nil {
			fmt.Println("rollback failed", err)
		}
		return
	}
	_, err2 := tx.Exec(sqlStr2, 3)
	if err2 != nil {
		fmt.Println("update failed", err2)
		err = tx.Rollback()
		if err != nil {
			fmt.Println("rollback failed", err)
		}
		return
	}

	err = tx.Commit()
	if err != nil {
		fmt.Println("commit failed ", err)
	}

}
