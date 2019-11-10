package main

import (
	"database/sql"
	"fmt"
	"log"
	"time"

	. "../../private"

	_ "github.com/go-sql-driver/mysql"
)

var db = &sql.DB{}

func init() {
	db, _ = sql.Open("mysql", "root:tj4912.huvip@tcp(127.0.0.1:3306)/worddb")
	//	db, _ = sql.Open("mysql", "root:root@/worddb")
	//db, _ = sql.Open("mysql", "hvp:huvip4912@tcp(127.0.0.1:3306)/msg")
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

//方式3 delete total time: 8.3016172
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
func gogo(dc string) {
	if query(dc) {

		//	p("--------", dc)
		update("/download/"+dc+".mp3", dc)
	}
}
func listdir() {
	values, err := AllListDir(`C:\Users\DELL\Desktop\mp3\download`, ".mp3")
	//values, err := AllListDir("Harry Potter And The Chamber Of Secrets", ".txt")
	if err == nil {
		for _, value := range values {
			//p(value) // 转换完成\encn\Harry Potter And The Chamber Of Secrets31.txt
			vh := Ren(value)

			dc := Paths(vh, 4) //Harry Potter And The Chamber Of Secrets31
			gogo(dc)
		}
	}

}

var p = fmt.Println

func main() {
	listdir()
}
