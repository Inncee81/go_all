package roufunc

import (
	"fmt"
	"strconv"
	 "time"

	. "../model"
	"../tools"
	jwt "github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

//获取DB初始化
var db = tools.GetDB()

//获取 redis 初始化
var client=tools.GetRS()

type Gjson struct {
	Btlist []Bt_list
	Len    int
	Pag    string
}

type GGjson struct {
	Code int `json:"code"`
	Data struct {
		Gjsons Gjson
	} `json:"data"`
	Msg string `json:"msg"`
}



type Tjson struct {
	Code int `json:"code"`
	Data struct {
		Token string `json:"token"`
	} `json:"data"`
	Msg string `json:"msg"`
}


type Rjson struct {
	Code int `json:"code"`
		Msg string `json:"msg"`
}


var gj Gjson
var ggjson GGjson
var tj Tjson
var rj Rjson
 
func Getlist(c *gin.Context) {




	bt := c.Param("bt")
	pag := c.Param("pag")
	username := c.Param("username")
	token := c.PostForm("token")
	isapi :=RegToken(username,token)
	if isapi{
		
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

	ggjson.Code=200
	ggjson.Msg="数据请求成功"
	ggjson.Data.Gjsons=gj

	}else{
	ggjson.Code=400
	ggjson.Msg="无权限访问"
	ggjson.Data.Gjsons=Gjson{}
	}
	
	c.JSON(200, ggjson)

}


//登录时输入账号和密码 返回生成的唯一token
func LoginToken(c *gin.Context) {
	db.LogMode(true)
	username := c.PostForm("username")
	password := c.PostForm("password")

	fmt.Println(username,password)
	us := User{}
	isNotFound := db.Debug().Where("username=? and password=?", username,password).First(&us).RecordNotFound()
	if isNotFound {
		tj.Code = 400
		tj.Data.Token = ""
		tj.Msg = "登录失败"
	}else{
		tj.Code = 200
		tj.Data.Token,_ = GenerateToken(username)
		tj.Msg = "登录成功"

		

	}
	c.JSON(200, tj)
}

//用户注册该用户不存在就注册
func Register(c *gin.Context) {
	username := c.PostForm("username")
	password := c.PostForm("password")
if username=="" || password=="" {
	rj.Msg ="注册用户名或密码不能为空"
	rj.Code=402
}
 if Isreg(username)  {
	rj.Msg ="该用户已存在"
	rj.Code=401
} else {
//定义一个用户，并初始化数据
	u := User{
		Username:username,
		Password:password,
		CreateTime:time.Now().Unix(),
	}

	//插入一条用户数据
		if err := db.Create(&u).Error; err != nil {
		fmt.Println("插入失败", err)
		rj.Msg="数据插入失败"
		rj.Code=400
	}else{
		rj.Msg="用户注册成功"
		rj.Code=200
	}
	
	}
c.JSON(200,rj) 

}

//是否注册
func Isreg(username string) (bool) {
u := User{}
isFound := db.Where("username = ?", username).First(&u).RecordNotFound()
	if !isFound {
		return true
	}else{
		return false
		
	}
}


//api权限 用户名和token 
func RegToken(username,token string) (bool) {
  // 获取post请求参数
	G,_:=GenerateToken(username)
	if G==token {
		return true
	}
return false
}


//SignWxToken 生成token,uid用户id，expireSec过期秒数
func GenerateToken(username string) (tokenStr string, err error) {
   token :=jwt.NewWithClaims(jwt.SigningMethodHS256,jwt.MapClaims{
"username":username,
})
tokenStr,err =token.SignedString([]byte("key"))
    return tokenStr, err
}
