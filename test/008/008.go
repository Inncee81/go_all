package main

import (
	"database/sql"
	"fmt"
	"log"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

var db = &sql.DB{}

func init() {
	//	db, _ = sql.Open("mysql", "root:root@/worddb")
	db, _ = sql.Open("mysql", "euser:tjqm4912@tcp(104.238.149.37:3306)/worddb")
}
func update(mp3src, dc string) {
	start := time.Now()
	//方式3 update
	stm, _ := db.Prepare("UPdate enwords set mp3src=? where word=?")
	stm.Exec(mp3src, dc)
	stm.Close()
	//	p(mp3src)
	end := time.Now()
	fmt.Println("update time:", end.Sub(start).Seconds())

}
func q(s string) string {
	Ks := ""
	var html string
	var tf bool
	//方式1 query 查询数据
	rows, _ := db.Query("SELECT * FROM enwords where word=?", s)
	defer rows.Close()
	for rows.Next() {
		var word, translation, mp3src string
		if err := rows.Scan(&word, &translation, &mp3src); err != nil {
			log.Fatal(err)
		}
		my := `{"word":"` + word + `","translation":"` + translation + `","mp3src":"` + mp3src + `"},`
		Ks += my
		tf = true
	}
	//fmt.Println(Ks)
	if tf {
		josna := `{"cmd": "allword","aks": 200,"data": [`
		end := "]}"
		myjosn := josna + Ks[:len(Ks)-1] + end
		html = "myjson(" + myjosn + ")"
	} else {
		html = Ks
	}
	return html
}

func main() {
	fmt.Println(query("a"))
	fmt.Println(q("a"))
}
