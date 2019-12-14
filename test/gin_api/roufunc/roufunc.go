package roufunc
import (
 "github.com/gin-gonic/gin"
_ "time"
. "../model"
"../tools"
"fmt"
)
//获取DB初始化
 var db=tools.GetDB()

/*
// 用户注册
func Uregister(c *gin.Context) {
        // 获取post请求参数
	username := c.PostForm("username")
	password := c.PostForm("password")
var msg string
var tf int
 kong := User{};
//查询注册账号
us:=User{}
db.First(&us,"username=?",username)
//查询不到,便插入账号和密码
  if us==kong{
  uu := &User{Username:username , Password: password, CreateTime: time.Now().Unix()}
  db.Create(uu)
  msg="账号注册成功"
tf=1
  }else{
    msg="账号或密码已存在"
    tf=0
  }

	c.JSON(200,gin.H{
		"username": username,
		"msg":msg,
		"code":tf,
	})
}

//用户登录
func Ulogin(c *gin.Context) {
	username := c.PostForm("username")
	password := c.PostForm("password")
var msg string
var tf int
 kong := User{};
//查询登录账号
us:=User{}
db.First(&us,"username=?&&password=?",username,password)
//查询不到,表示账号和密码错误
  if us==kong{
    msg="账号或密码错误"
    tf=0
  }else{
  msg="账号密码正确"
tf=1
  }

	c.JSON(200,gin.H{
		"username": username,
		"msg":msg,
    "code":tf,
	})
}

*/

func Getlist(c *gin.Context) { 

bt := c.Param("bt")

fmt.Println("-----------",bt)
  ab :=[] Bt_list{}
db.Where("name like ?","%"+bt+"%").Find(&ab)
   c.JSON(200,ab)




}