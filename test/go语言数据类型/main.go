package main

// 导入fmt包
import "fmt"

// 在函数外面定义全局变量
var i int = 0

// 入口函数main
func main() {
	// 给变量i赋值
	i = 100

	// 定义string变量
	var s string
	// 给变量s赋值
	s = "hello"

	// 打印变量值
	fmt.Println(s)

	// 短变量声明用法
	yes := false
	str := "www.tizi365.com"
	code := 200

	// 赋值
	yes = true
	str = "https://www.tizi365.com"
	code = 301

	// 打印变量值
	fmt.Println(yes)
	fmt.Println(str)
	fmt.Println(code)
	fmt.Println(45 << 1)

}
