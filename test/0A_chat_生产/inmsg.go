package main

import (
	"database/sql"
	"fmt"
	"net/http"
	_"time"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

func check(err error) {
	if err != nil {
		fmt.Println(err)
	}

}

func main() {
	//"用户名:密码@[连接方式](主机名:端口号)/数据库名"
	db, err := sql.Open("mysql", "hvp:huvip4912@(127.0.0.1:3306)/msg")
	check(err)

	var html string

	r := gin.Default()
	r.GET("/inmsg/:formid/:toid/:time/:msg", func(c *gin.Context) {

		formid := c.Param("formid")
		toid := c.Param("toid")
		msg := c.Param("msg")
		add_time := c.Param("time")
		if msg != "" {
			//add_time = time.Now().Format("2006-01-02 15:04:05")
			//add_time = time.Now().Unix()
			db.Exec("INSERT INTO webim_msg (f_id,t_id,t_msg,add_time) VALUES(?,?,?,?)", formid, toid, msg, add_time)
			html = `myjson({"aks":200})`
		} else {
			html = `myjson({"aks":300})`
		}
		c.String(http.StatusOK, html)
	})
	r.Run(":11111") // listen and serve on 0.0.0.0:8080
	//insert

}
