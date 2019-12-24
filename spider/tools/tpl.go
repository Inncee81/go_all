package tools
import (
  "github.com/flosch/pongo2"
 "strconv"
   
)

func init() {
    pongo2.RegisterFilter("scream", Scream)
}

func Scream(in *pongo2.Value, param *pongo2.Value) (out *pongo2.Value, err *pongo2.Error) {

	s :=in.String()
is , _ :=strconv.ParseInt(s, 10, 64)
	ss :=ByteSize(is)
  

    return pongo2.AsValue(ss), nil
}


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
