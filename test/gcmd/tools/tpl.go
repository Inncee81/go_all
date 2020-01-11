package tools
import (
  "github.com/flosch/pongo2"
)




//传入数据data 模版地址path 返回string数据
func Tpl( data map[string]interface{},src string) (string){
tpl, err := pongo2.FromFile(src)
if err != nil {
	panic(err)
}
out, err := tpl.Execute(data)
if err != nil {
	panic(err)
}
return out 
}

func Tpls( data map[string]interface{},src string) (string){
tpl, err := pongo2.FromString(src)
if err != nil {
	panic(err)
}
out, err := tpl.Execute(data)
if err != nil {
	panic(err)
}
return out 
}