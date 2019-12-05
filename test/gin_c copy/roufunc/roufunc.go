package roufunc
import (
 "github.com/gin-gonic/gin"
_ "time"
_ "../model"
"../tools"
)
//获取DB初始化
var db=tools.GetDB()


func GetUser(c *gin.Context) { 

	c.Set("template", `test\gin_c copy\mtpl\index.htm`)
		c.Set("data", map[string]interface{}{"message": "Hello World!"})


}