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

	//-----------插入一条用户数据
	//定义结构体&数据
	uu := &User{Username: "tizi366", Password: "123456", CreateTime: time.Now().Unix()}
	if err := db.Create(uu).Error; err != nil {
		fmt.Println("插入失败", err)
		return
	}

	//----------获取插入记录的Id 使用了数据库连接池技术,下面代码需要添加必须使用数据库事务,否则在高并发场景下，可能会查询不到自增id，或者查询到错误的id。
	var id []int
	db.Raw("select LAST_INSERT_ID() as id").Pluck("id", &id)
	//因为Pluck函数返回的是一列值，返回结果是slice类型，我们这里只有一个值，所以取第一个值即可。
	fmt.Println(id[0])

	//执行数据库查询操作

	u := User{}
	db.First(&u)
	fmt.Println("返回第一条记录",u)

	//等价于：SELECT * FROM `foods`   LIMIT 1
	u = User{}
	db.Take(&u,237)
	// 查询ID为237的记录，只支持主键



	fmt.Println("db.Take(&u,237)",u)

	u = User{}
	db.Last(&u)
	fmt.Println("返回最后一条记录",u)
	//因为Find返回的是数组，所以定义一个商品数组用来接收结果
	
	var users []User
	//等价于：SELECT * FROM `foods`
	db.Find(&users,"username=?","hahahaha")
	fmt.Println(users)

	//查询一列值 Pluck("id", &titles),数组
	var titles []string
	//这里Model函数是为了绑定一个模型实例，可以从里面提取表名。
	db.Model(User{}).Pluck("id", &titles)
	//Pluck提取了title字段，保存到titles变量
	fmt.Println(titles)


	u = User{}
	//链式操作，先查询，然后检测查询结果
	if db.Where("id = ?", 1178).Take(&u).RecordNotFound() {
		fmt.Println("查询不到数据")
	}
	fmt.Println(u)

	foo := User{}
	db.Model(&foo).Where("id = ?", 77).Update("Username", "gai7777")
	//先查询一条记录, 保存在模型变量food
	//等价于: SELECT * FROM `foods`  WHERE (id = '2') LIMIT 1
	db.Where("id = ?", 78).Take(&foo)

	//修改food模型的值
	foo.Username = "hahahaha"
	foo.Password = "mimimi"

	//等价于: UPDATE `foods` SET `title` = '可乐', `type` = '0', `price` = '100', `stock` = '26', `create_time` = '2018-11-06 11:12:04'  WHERE `foods`.`id` = '2'
	//保存模型变量的值 相当于根据主键id，更新所有模型字段值。
	db.Save(&foo)

	//改多条
	uc := User{Username: "修改name", CreateTime: time.Now().Unix()}
	db.Model(User{}).Where("id = ?", 90).Update(uc)

	//改一条
	db.Model(User{}).Where("id = ?", 78).Update("username", "haha365")

	//db.Model(User{}).Update("Username","gdgdgd")
	//因为User{}的id为空，gorm库就不会以id作为条件，where语句就是空的
//gorm.Expr 更新数据

	db.Model(&User{}).Where("id = ?", 80).Update("createtime", gorm.Expr("createtime + 100"))



// 根据Where条件删除数据
//用法：db.Where(条件表达式).Delete(空模型变量指针 


db.Where("id = ?", 240).Delete(&User{})

}
