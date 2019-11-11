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
	tf := qtf("a", "b")
	if tf {
		fmt.Println("ture")
	}

}

func qtf(username, password string) bool {
	var tf bool
	//方式1 query 查询数据
	rows, _ := db.Query("SELECT username FROM user where username=? and password=?", username, password)
	defer rows.Close()
	for rows.Next() {
		tf = true
	}
	return tf
}

func inster(username, password, gender, avatar string) {
	stm, _ := db.Prepare("INSERT INTO user(username,password,gender,avatar) values(?,?,?,?)")
	stm.Exec(username, password, gender, avatar)
	stm.Close()
}
