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


if pag == "1" {
var titles []string

db.Where("name like ?", "%"+bt+"%").Model(Bt_list{}).Pluck("id", &titles).Count(&total)

	
	fmt.Println(total)	
} 
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
