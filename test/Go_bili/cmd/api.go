package main

//导入tools包
import (
  "fmt"
  "time"
	. "../model"
	"../tools"
)

func main() {
	//获取DB
	db := tools.GetDB()
  // 读取
  var us User
 // var uss [] User
  kong := User{};
//查询登录账号
  db.First(&us,"username=?","tizi366")
  fmt.Println("查询name为tizi366的us",us)
//查询不到,便插入账号和密码
  if us==kong{
  uu := &User{Username: "tizi366", Password: "123456", CreateTime: time.Now().Unix()}
  db.Create(uu)
  }else{
    fmt.Println("账号或密码已存在")
  }
 //db.Find(&uss,"username=? && password=?","tizi366","123456")
  //fmt.Println("查询name为tizi366的us",uss)
}


