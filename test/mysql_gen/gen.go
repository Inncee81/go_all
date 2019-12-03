package main

import (
	_"encoding/json"
	"fmt"
	_"time"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"./qbnt"
	"github.com/flosch/pongo2"
	. "../../private"
)


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
var s []string	
	
//获取插入记录的Id
// 数据推荐存在MAP中 ...
    data := make(map[string]interface{})


//因为Pluck函数返回的是一列值，返回结果是slice类型，我们这里只有一个值，所以取第一个值即可。
db.Raw("show tables;").Pluck("*",&s)
fmt.Println(s)

for _, v := range s {
qbs:=qbnt.Nnts(v)

data["qbs"] = qbs
data["tname"] =v 
tpl, err := pongo2.FromFile(`test\mysql_gen\m.htm`)
if err != nil {
	panic(err)
}
out, err := tpl.Execute(data)
if err != nil {
	panic(err)
}
fmt.Println(out) // Output: Hello Florian!

//	fmt.Println(qbs)
W_file(`test\mysql_gen\genapi\` + v + ".go", out )

}


}
