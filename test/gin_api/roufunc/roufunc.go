package roufunc

import (
	"fmt"
	"strconv"
	_ "time"

	. "../model"
	"../tools"
	jwt "github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

//获取DB初始化
var db = tools.GetDB()

type Gjson struct {
	Btlist []Bt_list
	Len    int
	Pag    string
}

type Tjson struct {
	Code int `json:"code"`
	Data struct {
		Token string `json:"token"`
	} `json:"data"`
	Msg string `json:"msg"`
}

var gj Gjson
var tj Tjson

func Getlist(c *gin.Context) {

	bt := c.Param("bt")
	pag := c.Param("pag")
	total := 0
	fmt.Println("-----------", bt)

	if pag == "1" {
		var titles []string

		db.Where("name like ?", "%"+bt+"%").Model(Bt_list{}).Pluck("id", &titles).Count(&total)
		fmt.Println(total)
	}
	abp := []Bt_list{}
	n, _ := strconv.Atoi(pag)
	npag := (n - 1) * 12
	spag := strconv.Itoa(npag)
	db.Where("name like ?", "%"+bt+"%").Limit("12").Offset(spag).Find(&abp)
	gj.Btlist = abp
	gj.Len = total
	gj.Pag = pag
	c.JSON(200, gj)

}

func CheckToken(c *gin.Context) {
	username := c.Param("name")
	password := c.Param("pass")
	var us User
	db.Select("id").Where(User{Username: username, Password: password}).First(&us)
	if us.ID > 0 {
		tj.Code = 200
		//假数据
		tj.Data.Token = "yJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VybmFtZSI6InRlc3QiLCJwYXNzd29yZCI6InRlc3QxMjM0NTYiLCJleHAiOjE1MTg3MjAwMzcsImlzcyI6Imdpbi1ibG9nIn0.-kK0V9E06qTHOzupQM_gHXAGDB3EJtJS4H5TTCyWwW8"
		tj.Msg = "ok"
	} else {
		tj.Code = 400
		tj.Data.Token = ""
		tj.Msg = "授权失败"
	}
	c.JSON(200, tj)

}

func RegToken(c *gin.Context) {
	username := c.Param("name")
	password := c.Param("pass")
	var us User
	us = User{Username: username, Password: password}
	db.Select("id").Where(User{Username: username, Password: password}).First(&us)
	if us.ID > 0 {
	} else {
		GenerateToken(us)
		c.STRING(200,us)
	}
	c.JSON(200, tj)
	c.STRING(200,us)

}

func GenerateToken(user User) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"username": user.Username,
		//"exp":      time.Now().Add(time.Hour * 2).Unix(),// 可以添加过期时间
	})

	return token.SignedString([]byte("secret")) //对应的字符串请自行生成，最后足够使用加密后的字符串
}
