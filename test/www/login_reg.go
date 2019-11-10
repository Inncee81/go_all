package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

var db = &sql.DB{}

func init() {
	db, _ = sql.Open("mysql", "root:root@/cms")
}

var p = fmt.Println

func main() {
	p(qq("a", "b"))
}

func qq(username, password string) bool {
	var tf bool
	//方式1 query 查询数据
	rows, _ := db.Query("SELECT username FROM user where username=? and password=?", username, password)
	defer rows.Close()
	for rows.Next() {
		tf = true
	}
	return tf
}
