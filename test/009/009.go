<<<<<<< HEAD
//
package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)
=======

>>>>>>> c4889c8c845dff14412cc2b43cfc1da5f856db94


<<<<<<< HEAD
func Get(s string) string {

	resp, err := http.Get(s)
	if err != nil {
		// handle error
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		// handle error
	}

	fmt.Println(string(body))
	return string(body)
}

func main() {
	Get("http://dict.youdao.com/dictvoice?audio=a")
=======





package main

import (
	"database/sql"
	"fmt"
	"log"
	_"time"
 ."../../private"
	_ "github.com/go-sql-driver/mysql"
)

var db = &sql.DB{}

func init() {
	//	db, _ = sql.Open("mysql", "root:root@/worddb")
	db, _ = sql.Open("mysql", "euser:tjqm4912@tcp(104.238.149.37:3306)/worddb")
	//db, _ = sql.Open("mysql", "hvp:huvip4912@tcp(149.28.27.116:3306)/msg")
}
func update(mp3src, dc string) {
	//方式3 update
	stm, _ := db.Prepare("UPdate enwords set mp3src=? where word=?")
	stm.Exec(mp3src, dc)
	stm.Close()
	//	p(mp3src)

}
func query() {
	i := 0
	zs:=""
	//方式1 query
	rows, _ := db.Query("select word from enwords where locate('v.', translation)&& not locate('adv.', translation)")
	// select * from enwords where locate('v.', translation)&& not locate('adv.', translation)
	defer rows.Close()
	for rows.Next() {
		var word string
		if err := rows.Scan(&word); err != nil {
			log.Fatal(err)
		}
		i++
		fmt.Println(word)
		zs+=word+"\n"
	}
		W_file("word.txt",zs)
	fmt.Println(i)
}

func main() {
	query()
>>>>>>> c4889c8c845dff14412cc2b43cfc1da5f856db94
}
