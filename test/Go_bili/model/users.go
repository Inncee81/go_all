package model
import (
	"time"
)	
//在这里User类型可以代表mysql users表
type User struct {
	ID         int    `gorm:"primary_key"`
	Username   string `gorm:"column:username"`
	Password   string `gorm:"column:password"`
	CreateTime int64  `gorm:"column:createtime" json:"-"`
}

type WechatCreateLog struct {
	ID     int       `json:"id" gorm:"column:id"`
	Name   string    `json:"name" gorm:"column:name"`
	Owner  string    `json:"owner" gorm:"column:owner"`
	ChatID string    `json:"chatid" gorm:"column:chatid"`
	Cuser  string    `json:"cuser" gorm:"column:cuser"`
	Ctime  time.Time `form:"ctime" json:"ctime" gorm:"column:ctime"`
	Status int       `form:"status" json:"status" gorm:"column:status"` //1创建 2修改 3同步中 4同步完成 5同步失败
}


//设置表名，可以通过给struct类型定义 TableName函数，返回当前struct绑定的mysql表名是什么
func (User) TableName() string {
	//绑定MYSQL表名为users
	return "users"
}
