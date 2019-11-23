package main

//导入tools包
import (
	"fmt"
	"time"

	. "../model"
	"../tools"
	"github.com/jinzhu/gorm"
)

func main() {
	//获取DB
	db := tools.GetDB()


// 创建
  uu := &User{Username: "tizi366", Password: "123456", CreateTime: time.Now().Unix()}
  db.Create(uu)


  // 读取
  var us User
  var uss [] User

  db.First(&us, 235) // 查询id为1的us
  fmt.Println("查询id为235的us",us)

  db.First(&us, "id = ?", 236) // 查询code为l1212的us
  fmt.Println("查询id为236的us",us)

  db.Find(&uss,"username=?","tizi366")

  fmt.Println("查询name为tizi366的us",uss)

 	//改表的多条列数据
	uc := User{Username: "修改235name", CreateTime: time.Now().Unix()}
	db.Model(User{}).Where("id = ?", 235).Update(uc)

	//改表的一条列数据
	db.Model(User{}).Where("id = ?", 236).Update("username", "改名username236")

   //gorm.Expr 更新数据

	db.Model(&User{}).Where("id = ?",237).Update("createtime", gorm.Expr("createtime + 100"))


  // 删除 - 删除us
  db.Delete(User{},"id=?",243)
  db.Delete(User{},242)
}


