
package main
import (
   .  "./Concurrence"
    "time"
    "fmt"
    "runtime"
)

//定义一个实现Job接口的数据
type Score struct {
    Num int
}
//定义对数据的处理
func (s *Score) Do() {
    fmt.Println("num:", s.Num)
    time.Sleep(1 * 1 * time.Second)
}
 
func main() {
    num := 100 * 100 * 2
    // debug.SetMaxThreads(num + 1000) //设置最大线程数
    // 注册工作池，传入任务
    // 参数1 worker并发个数
    p := NewWorkerPool(num)
    p.Run()
 
    //写入一亿条数据
    datanum := 100 * 100 * 10
    go func() {
        for i := 1; i <= datanum; i++ {
            sc := &Score{Num: i}
            p.JobQueue <- sc //数据传进去会被自动执行Do()方法，具体对数据的处理自己在Do()方法中定义
        }
    }()
 
//循环打印输出当前进程的Goroutine 个数
    for {
        fmt.Println("runtime.NumGoroutine() :", runtime.NumGoroutine())
        time.Sleep(2 * time.Second)
    }
 



}