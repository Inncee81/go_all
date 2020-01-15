package main
import (
	"fmt"
	"./redisku"
	"time"
)
var client=redisku.GetRS()
func main() {
// 初始化一个新的redis client
fmt.Println(client)					//六秒
err := client.Set("key1", "value", (6 * time.Second)).Err()

// 检测设置是否成功
if err != nil {
	panic(err)
}

val, err := client.Get("key").Result()
// 检测，查询是否出错
if err != nil {
	fmt.Println("已过期需要重新登录")
}else{
	
}
fmt.Println("key", val)

}