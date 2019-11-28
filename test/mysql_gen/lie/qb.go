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

type Lie struct {
Cname string
Ctype string
}
type Cnt struct {
Nt []Lie	
}
func main() {
//	var nt Cnt
	//"用户名:密码@[连接方式](主机名:端口号)/数据库名"
		rows, _ := db.Query("SELECT COLUMN_NAME,DATA_TYPE FROM INFORMATION_SCHEMA.COLUMNS WHERE TABLE_NAME='users'")
	defer rows.Close()
	for rows.Next() {
		var COLUMN_NAME,DATA_TYPE string
		if err := rows.Scan(&COLUMN_NAME,	&DATA_TYPE ); err != nil {
			p(err)
		}
	p(COLUMN_NAME,DATA_TYPE)
	}
}
