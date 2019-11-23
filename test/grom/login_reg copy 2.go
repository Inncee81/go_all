package main

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

//在这里User类型可以代表mysql users表
type User struct {
	//Id	int	`json:" - "` //自增ID
	Username   string `gorm:"column:username"`
	Password   string `gorm:"column:password"`
	CreateTime int64  `gorm:"column:time" json:"-"`
}

type Usei struct {
	Reason string
	Msg    string
	Data   []User
}

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

	//定义一个用户，并初始化数据
	u := User{
		Username:   "tizi366",
		Password:   "123456",
		CreateTime: time.Now().Unix(),
	}

	//插入一条用户数据
	//下面代码会自动生成SQL语句：INSERT INTO `users` (`username`,`password`,`createtime`) VALUES ('tizi365','123456','1540824823')
	if err := db.Create(u).Error; err != nil {
		fmt.Println("插入失败", err)
		return
	}

	//查询并返回第一条数据
	//定义需要保存数据的struct变量
	u = User{}
	//自动生成sql： SELECT * FROM `users`  WHERE (username = 'tizi365') LIMIT 1
	isNotFound := db.Where("username = ?", "tizi365").First(&u).RecordNotFound()
	if isNotFound {
		fmt.Println("找不到记录")
		return
	}
	//打印查询到的数据
	fmt.Println(u.Username, u.Password)

	//因为Find返回的是数组，所以定义一个商品数组用来接收结果
	var us []User

	//等价于：SELECT * FROM `foods`
	isFound := db.Where("username=?", "tizi365").Find(&us).RecordNotFound()
	if isFound {
		fmt.Println("找不到记录")
		return
	}
	fmt.Println(us)

	o := Usei{}
	o.Reason = "yes"
	// 通过UserInfo结构体初始化User字段
	o.Data = us
	// 将struct对象转换成json
	b, err := json.Marshal(o)
	if err != nil {
		// 转换失败，错误处理
		fmt.Println("error:", err)
	}

	// 因为Marshal返回的是字节数组，所以这里需要转换成字符串
	fmt.Println(string(b))

}

//更新
//自动生成Sql: UPDATE `users` SET `password` = '654321'  WHERE (username = 'tizi365')
//db.Model(User{}).Where("username = ?", "tizi365").Update("password", "654321")

//删除
//自动生成Sql： DELETE FROM `users`  WHERE (username = 'tizi365')
//db.Where("username = ?", "tizi365").Delete(User{})
