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
	//db, _ = sql.Open("mysql", "hvp:huvip4912@tcp(149.28.27.116:3306)/msg")
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
func query(dc string) bool {

	//方式1 query
	start := time.Now()
	rows, _ := db.Query("SELECT word FROM enwords where word=?", dc)
	// select * from enwords where locate('v.', translation)&& not locate('adv.', translation)
	defer rows.Close()
	for rows.Next() {
		var word string
		if err := rows.Scan(&word); err != nil {
			log.Fatal(err)
		}
		return true
		//	fmt.Printf("word:%s", word)
	}
	end := time.Now()
	fmt.Println("query time:", end.Sub(start).Seconds())
	return false
}
func main() {
	fmt.Println(query("a"))
	fmt.Println(q("a"))
}
