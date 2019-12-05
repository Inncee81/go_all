package roufunc
import (
 "github.com/gin-gonic/gin"
 
_"../model"
//数据库读取数据
"../tools"
)
//获取DB初始化
var db=tools.GetDB()

//路由跳转mvc
func Index(c *gin.Context) {
	c.Set("template", `../tpl/index.htm`)
	c.Set("data", map[string]interface{}{"message": "Hello World!"})
}
