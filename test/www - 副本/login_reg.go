package main

import (
	"database/sql"
	"fmt"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

var db = &sql.DB{}

func init() {
	db, _ = sql.Open("mysql", "root:root@/cms")
}

var p = fmt.Println

func main() {

	r := gin.Default()

	r.POST("/post", func(c *gin.Context) {
		username := c.PostForm("username")
		password := c.PostForm("password")

		if qq(username, password) {
			p("true")
		} else {
			p("false")
		}
		c.JSON(200, gin.H{
			"status":   "posted",
			"username": username,
			"password": password,
		})
	})
	r.Run(":8083")
}

func qq(username, password string) bool {
	var tf bool
	//方式1 query 查询数据
	rows, _ := db.Query("SELECT * FROM user where username=? and password=?", username, password)
	defer rows.Close()
	for rows.Next() {
		tf = true
	}
	return tf
}
