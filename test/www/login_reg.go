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
		inster("myma","passwordl","itime19","src//png",12)
	}

}

func qtf(username, password string) bool {
	var tf bool
	//方式1 query 查询数据
	rows, _ := db.Query("SELECT username FROM users where username=? and password=?", username, password)
	defer rows.Close()
	for rows.Next() {
		tf = true
	}
	return tf
}


表名(列名.值, )

func inster(username, password, createtime, avatar_src string,gender int) {
	stm, _ := db.Prepare("INSERT INTO users(username, password, createtime, avatar_src,gender) values(?,?,?,?,?)")
	stm.Exec(username, password, createtime, avatar_src,gender)
	stm.Close()
}
