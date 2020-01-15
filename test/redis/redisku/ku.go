package redisku 
import (
	"github.com/go-redis/redis"
)

var Myredis *redis.Client
func init() {
// 初始化一个新的redis client
Myredis = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379", // redis地址
		Password: "", // redis没密码，没有设置，则留空
		DB:       0,  // 使用默认数据库
	})
}

func GetRS() *redis.Client {
	return Myredis
}

