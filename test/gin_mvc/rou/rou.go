package rou
// 导入gin包
import (
.	"../roufunc"
	"../cors"
	"github.com/gin-gonic/gin"
)

// 入口函数
func init() {
    // 初始化一个http服务对象
	r := gin.Default()
// 要在路由组之前全局使用「跨域中间件」, 否则OPTIONS会返回404
    r.Use(cors.Cors())
// 路由定义post请求, url路径为：/user/login, 绑定
r.Static("/assets", `test\gin_mvc\static`)
	
r.GET("/", Index)
r.GET("/html", Html)

r.Run(":8230")

}