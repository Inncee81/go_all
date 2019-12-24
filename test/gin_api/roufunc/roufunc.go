package roufunc

import (
	"strconv"
	"fmt"
	_ "time"

	. "../model"
	"../tools"
	"github.com/gin-gonic/gin"
)

//获取DB初始化
var db = tools.GetDB()

type Gjson struct {
	Btlist []Bt_list
	Len    int
	Pag    string
}

var gj Gjson

func Getlist(c *gin.Context) {

	bt := c.Param("bt")
	pag := c.Param("pag")
	total := 0
	fmt.Println("-----------", bt)
	ab := []Bt_list{}
	db.Where("name like ?", "%"+bt+"%").Find(&ab).Count(&total)
	//等价于: SELECT count(*) FROM `foods`
	//这里也需要通过model设置模型，让gorm可以提取模型对应的表名
	fmt.Println(total)

	abp := []Bt_list{}
	n,_:=strconv.Atoi(pag)	
	npag:=(n-1)*12
	spag :=strconv.Itoa(npag)
	db.Where("name like ?", "%"+bt+"%").Limit("12").Offset(spag).Find(&abp)
	gj.Btlist = abp
	gj.Len = total
	gj.Pag = pag
	c.JSON(200, gj)

}
