package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

var p = fmt.Println

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
	r.GET("/allmsg/:formid/:toid/:token", func(c *gin.Context) {

		formid := c.Param("formid")
		toid := c.Param("toid")
		token := c.Param("token")

		if formid != "" && token != "" {

			Ks := ""
			var tf bool

			//方式1 query 查询数据
			rows, _ := db.Query("SELECT id,t_msg,add_time,f_id FROM webim_msg  where (f_id=? and t_id=?) or (f_id=? and t_id=?) ORDER BY `id` ASC ", formid, toid, toid, formid)
			defer rows.Close()
			for rows.Next() {
				var id, t_msg, add_time, f_id string
				if err := rows.Scan(&id, &t_msg, &add_time, &f_id); err != nil {
					log.Fatal(err)
				}
				my := `{"msgid":"` + id + `","msg":"` + t_msg + `","time":"` + add_time + `","form":"` + f_id + `"},`
				Ks += my
				//fmt.Println(Ks)
				tf = true
			}
			if tf {
				josna := `{"cmd": "allmsg","aks": 200,"msg": [`
				end := "]}"
				myjosn := josna + Ks[:len(Ks)-1] + end
				//fmt.Println(myjosn)
				//str := strings.Replace(myjosn, " ", "", -1)
				// 去除换行符
				//strn := strings.Replace(myjosn, "\n", "\r", -1)
				//print(str)
				html = "myjson(" + myjosn + ")"
				//	return myjosn
			} else {
				html = `myjson({"aks":100})`
			}

		} else {
			html = `myjson({"aks":300})`
		}
		c.String(http.StatusOK, html)
	})
	r.Run(":8888") // listen and serve on 0.0.0.0:8080
}
