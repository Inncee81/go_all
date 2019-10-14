package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

var db = &sql.DB{}

func init() {
	//	db, _ = sql.Open("mysql", "root:root@/worddb")
	db, _ = sql.Open("mysql", "euser:tjqm4912@tcp(104.238.149.37:3306)/worddb")
}

var p = fmt.Println

func check(err error) {
	if err != nil {
		fmt.Println(err)
	}

}
func qq(s string) string {
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
		html = `{"aks": 200}`
	}
	return html
}
func main() {
	//"用户名:密码@[连接方式](主机名:端口号)/数据库名"

	var html string
	r := gin.Default()
	r.GET("/search/:q/", func(c *gin.Context) {

		q := c.Param("q")

		if q != "" {
			html = qq(q)
		}
		c.String(http.StatusOK, html)
	})
	r.Run(":81") // listen and serve on 0.0.0.0:8080
}
