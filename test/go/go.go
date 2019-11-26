package main

import (
    "fmt"
)

// 定义计算斐波拉契数列的函数，通过channel参数c返回每一步的计算结果
func fibonacci(n int)  {
var	x,y int64
    x, y = 0, 1
    for i := 0; i < n; i++ {
		// 返回当前计算结果
		
		x, y = y, x+y
		fmt.Println(x)
    }
    
    // 通过close关闭channel
}

func main() {
    // 定义int类型的channel,缓冲队列大小是10
    
    // 创建一个协程，开始计算数列
    fibonacci(100)
    
    // 通过range关键词，循环遍历channel变量c，从c中读取数据
    // 如果channel没有数据，就阻塞循环，直到channel中有数据
    // 如果channel关闭，则退出循环
}