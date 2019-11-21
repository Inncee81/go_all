// a project main.go
package main

import (
	"fmt"
)

func main() {
	// 定义int类型指针变量p
	var p *int
	// 初始化并且定义变量i
	i := 42
	p = &i
	// 将i变量的地址，赋值给p变量， 这个时候指针p指向变量i
	p = &i

	// 打印指针p，指向的值，这里输出42
	fmt.Println(*p)
	// 将100 赋值给指针p指向的存储空间，相当于赋值给变量i
	*p = 100
	fmt.Println(*p)
}
