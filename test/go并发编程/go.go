package main

import (
	"fmt"
	"time"
)

// 定义累加函数，负责将s数组所有值相加， sum函数还接受一个int类型的channel参数c
func sum(s []int, c chan int) {
	sum := 0
	// 累加数组s的所有值
	for _, v := range s {
		sum += v
	}

	// 将计算的结果发生到channel中
	c <- sum
}

func main() {
	println("start：：：：：：：：：")

	s := []int{7, 2, 8, -9, 4, 0, 4, -4, 2}

	for i := 0; i < 1500000; i++ {
		s = append(s, i)
	}

	tstart := time.Now().Unix()
	var start, end int

	lens := len(s)
	//println("任务量",lens)
	bf :=50000    //任务量等于2
	ls := lens / bf //并发数
	// println("并发数",ls)

	// 定义int类型的channel,缓冲队列大小是10
	c := make(chan int)
	for index := 0; index < ls; {
		start = index * bf
		if index < ls-1 {
			end = (index + 1) * bf
		} else {
			end = lens
		}
		// fmt.Println(s[start:end])
		go sum(s[start:end], c)
		index++
	}

	zn := 0
	for index := 0; index < ls; index++ {
		cc := <-c
		zn += cc
	}
	fmt.Println(zn)

	tend := time.Now().Unix()
	fmt.Printf("执行消耗的时间为:%v秒", tend-tstart)

	tst := time.Now().Unix()
	zvn := 0
	for _, v := range s {

		zvn += v
	}
	fmt.Println(zvn)

	ten := time.Now().Unix()
	fmt.Printf("执行消耗的时间为:%v秒", ten-tst)

}
