package rou
// 导入gin包
import (
.	"../roufunc"
	"github.com/gin-gonic/gin"
)

// 入口函数
func init() {
    // 初始化一个http服务对象
	r := gin.Default()

// 路由定义post请求, url路径为：/user/login, 绑定
r.POST("/user/login", Ulogin)
r.POST("/user/register", Uregister)
	
r.GET("/users/:id", GetUser)

r.Run(":8091")

}
