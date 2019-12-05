package rou
// 导入gin包
import (
//.	"../roufunc"
	"../cors"
	"github.com/gin-gonic/gin"
	"github.com/ngerakines/ginpongo2"
)

// 入口函数
func init() {
    // 初始化一个http服务对象
	r := gin.Default()
// 要在路由组之前全局使用「跨域中间件」, 否则OPTIONS会返回404
	r.Use(cors.Cors())
	r.Use(ginpongo2.Pongo2())
// 路由定义post请求, url路径为：/user/login, 绑定

r.GET("/", func(c *gin.Context) {
		c.Set("template", "index.html")
		c.Set("data", map[string]interface{}{"message": "Hello World!"})
	})
r.Run(":8525")

}
