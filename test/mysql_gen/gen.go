package main

import (
	_"encoding/json"
	"fmt"
	_"time"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	
)

//在这里User类型可以代表mysql users表
type User struct {
	//	Id int64 `gorm:"AUTO_INCREMENT"`
	ID         int    `gorm:"column:id; PRIMARY_KEY"`
	Username   string `gorm:"column:username"`
	Password   string `gorm:"column:password"`
	CreateTime int64  `gorm:"column:createtime" json:"-"`
}






//设置表名，可以通过给struct类型定义 TableName函数，返回当前struct绑定的mysql表名是什么
func (u User) TableName() string {
	//绑定MYSQL表名为users
	return "users"
}

type Usei struct {
	Reason string
	Msg    string
	Data   []User
}

var u User


func main() {
	//配置MySQL连接参数
	username := "root"  //账号
	password := "root"  //密码
	host := "127.0.0.1" //数据库地址，可以是Ip或者域名
	port := 3306        //数据库端口
	Dbname := "cms"     //数据库名
	timeout := "10s"    //连接超时，10秒
	//拼接下dsn参数, dsn格式可以参考上面的语法，这里使用Sprintf动态拼接dsn参数，因为一般数据库连接参数，我们都是保存在配置文件里面，需要从配置文件加载参数，然后拼接dsn。
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local&timeout=%s", username, password, host, port, Dbname, timeout)
	//连接MYSQL, 获得DB类型实例，用于后面的数据库读写操作。
	db, err := gorm.Open("mysql", dsn)
	if err != nil {
		fmt.Println("连接数据库失败,oerror=" + err.Error())
	}
	//延时关闭数据库连接
	defer db.Close()
db.LogMode(true)
var s []string	
	
//获取插入记录的Id

//因为Pluck函数返回的是一列值，返回结果是slice类型，我们这里只有一个值，所以取第一个值即可。
db.Raw("show tables;").Pluck("*",&s)
fmt.Println(s)

// Scan
type Result struct {
  COLUMN_NAME string
  DATA_TYPE string
  }

//var r []string
var r []Result
for _, v := range s {
	
	db.Raw("SELECT COLUMN_NAME,DATA_TYPE FROM INFORMATION_SCHEMA.COLUMNS WHERE TABLE_NAME='?'",v).Scan(&r)
	//fmt.Println(r)
}
var rr []Result
db.Raw("SELECT COLUMN_NAME,DATA_TYPE FROM INFORMATION_SCHEMA.COLUMNS WHERE TABLE_NAME='users'").Scan(&rr)
	fmt.Println(rr)
}

