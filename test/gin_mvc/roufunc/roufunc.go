package roufunc
import (
 "github.com/gin-gonic/gin"
 
_"../model"
//数据库读取数据
"../tools"
)
//获取DB初始化
var db=tools.GetDB()

//路由跳转mvc
func Index(c *gin.Context) {
 data:=make(map[string]interface{})
    data["array"]=  tools.Array(12, "Steve", "Mark", "David",true)
    data["Name"] = "nljb"
    data["Email"] = "nljb@qq.com"

    // 布尔
    data["True"] = true
    data["False"] = false

    // 整型
    data["Year"] = 35

    // 结构体
    type User struct {
        Name string
        Pass string
    }
    data["User"] = User{Name: "nljb", Pass: "1234"}

    // 列表
    data["List"] = []int{1, 2, 35, 4, 5, 6, 7, 8}

    // 列表结构体
    a := User{Name: "nljb", Pass: "1234"}
    b := User{Name: "jbnl", Pass: "4321"}
    data["Structs"] = []User{a, b}

    // 模版变量
    data["var"] = "hello world !!!"

	data["maps"] = map[string]string{"name": "golang"}
	
 sdata:=tools.Tpl(data,`test\gin_mvc\tpl\index.htm`)
 c.Header("Content-Type", "text/html; charset=utf-8")
 c.String(200,sdata)
}
