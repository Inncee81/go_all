package model
type Staff struct {

Id int `json:"id" gorm:"column:id"` // 
    
Name string `json:"name" gorm:"column:name"` // 员工姓名
    
Phone string `json:"phone" gorm:"column:phone"` // 员工电话
    
Email string `json:"email" gorm:"column:email"` // 员工邮箱
    
Position string `json:"position" gorm:"column:position"` // 员工职位
    
Workname string `json:"workname" gorm:"column:workname"` // 公司名称
    
Dsf float64 `json:"dsf" gorm:"column:dsf"` // 
    
} 
//设置表名，可以通过给struct类型定义 TableName函数，返回当前struct绑定的mysql表名
func (Staff) TableName() string {
	//绑定MYSQL表名为staff
	return "staff"
}

