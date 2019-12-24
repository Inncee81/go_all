package model
	
//在这里User类型可以代表mysql users表
type Bt_list struct {
	ID         int    `gorm:"primary_key"`
	Name   string `gorm:"column:name"`
	Files   string `gorm:"column:files"`
	Infohash  string `gorm:"column:Infohash"`
}

//设置表名，可以通过给struct类型定义 TableName函数，返回当前struct绑定的mysql表名是什么
func (Bt_list) TableName() string {
	//绑定MYSQL表名为users
	return "bt_list"
}
