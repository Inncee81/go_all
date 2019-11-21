package main

import (
	"database/sql"
	"log"
	_ "github.com/go-sql-driver/mysql"
)

var db = &sql.DB{}

func init() {
	db, _ = sql.Open("mysql", "root:root@/test")
}

func main() {
	insert()
	query()
	update()
	delete()
}

func update() {

	//方式3 update
	stm, _ := db.Prepare("UPdate user set age=? where uid=?")
		stm.Exec("ii", "df")
	stm.Close()

}

func delete() {
	//方式3 delete
	stm, _ := db.Prepare("DELETE FROM user WHERE uid=?")
		stm.Exec("i")
	stm.Close()

}

func query() {

	//方式3 query
	tx, _ := db.Begin()
	defer tx.Commit()
	rows, _ = tx.Query("SELECT uid,username FROM user")
	defer rows.Close()
	for rows.Next() {
		var name string
		var id int
		if err := rows.Scan(&id, &name); err != nil {
			log.Fatal(err)
		}
		//fmt.Printf("name:%s ,id:is %d\n", name, id)
	}
}

func insert() {
	//方式3 insert
	stm, _ := db.Prepare("INSERT INTO user(uid,username,age) values(?,?,?)")
		stm.Exec("i", "user","i", "i")
	stm.Close()
}
