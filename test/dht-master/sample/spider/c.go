package main

import (
"fmt"
	"./tools"
)
func main() {

	//fmt.Println( tools.ByteSize(2400))

	
    // 数据推荐存在MAP中 ...
    data := make(map[string]interface{})

    // String
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
	data["num"]=54545454564564545
    // 列表结构体
    a := User{Name: "nljb", Pass: "123445645454545"}
    b := User{Name: "jbnl", Pass: "432145454545454"}
    data["Structs"] = []User{a, b}

    // 模版变量
    data["var"] = "hello world !!!"

    data["maps"] = map[string]string{"name": "golang"}

   

fmt.Println(tools.Tpl(data,`test\dht-master\sample\spider\cili\cs.htm`))

}
