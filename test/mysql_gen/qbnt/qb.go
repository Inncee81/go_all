package qbnt

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

var db = &sql.DB{}

func init() {
	db, _ = sql.Open("mysql", "root:root@/cms")
}

var p = fmt.Println

type NameType struct {
  Cname string //表名
  Ctype string //表类型
  Comment string //表注释
  }
 
 //mysql类型转换成 golang类型
func Vs(s string) (string) {
if s=="bool"{
	return s
	}else if s=="varchar" || s=="char" || s=="text" || s=="longtext" || s=="tinytext" {
	return "string"
	}else if s=="date" || s=="time" || s=="datetime" {
	return "time.Time"
	}else if s=="int" || s=="tinyint" || s=="smallint" || s=="integer" || s=="mediumint" {
	return "int"
	}else if s=="bigint" {
	return "int64"
	}else if s=="double" ||s=="float" {
	return "float64"
}else{
	return s
}  
	}

func Nnts(s string) ([]NameType){
	var cnts []NameType
	//获取
			rows, _ := db.Query("SELECT COLUMN_NAME,DATA_TYPE,COLUMN_COMMENT FROM INFORMATION_SCHEMA.COLUMNS WHERE TABLE_NAME=?",s)
	defer rows.Close()
	for rows.Next() {
		var COLUMN_NAME,DATA_TYPE,COLUMN_COMMENT string
		if err := rows.Scan(&COLUMN_NAME,&DATA_TYPE ,&COLUMN_COMMENT); err != nil {
			p(err)
		}
	
dtype:=Vs(DATA_TYPE)


			 cnts = append(cnts, NameType{COLUMN_NAME,dtype,COLUMN_COMMENT})
	}
	return cnts
}