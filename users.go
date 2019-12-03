package model
type Users struct {

Id int `json:"id" gorm:"column:id"` // 自增ID
    
Username string `json:"username" gorm:"column:username"` // 账号
    
Password string `json:"password" gorm:"column:password"` // 密码
    
Createtime int `json:"createtime" gorm:"column:createtime"` // 创建时间
    
} 
//设置表名，可以通过给struct类型定义 TableName函数，返回当前struct绑定的mysql表名
func (Users) TableName() string {
	//绑定MYSQL表名为users
	return "users"
}

