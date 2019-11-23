package model
	
//在这里User类型可以代表mysql users表
type User struct {
	ID         int    `gorm:"primary_key"`
	Username   string `gorm:"column:username"`
	Password   string `gorm:"column:password"`
	CreateTime int64  `gorm:"column:createtime" json:"-"`
}

//设置表名，可以通过给struct类型定义 TableName函数，返回当前struct绑定的mysql表名是什么
func (User) TableName() string {
	//绑定MYSQL表名为users
	return "users"
}
