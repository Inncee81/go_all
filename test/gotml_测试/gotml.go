package main
import (
	"fmt"
	"github.com/flosch/pongo2"
)

type ts struct {
Name string	
Cap string
}

func main() {
	
 
    // 数据推荐存在MAP中 ...
    data := make(map[string]interface{})

    // String
    data["Name"] = "nljb"
    data["Email"] = "nljb@qq.com"

    // 布尔
    data["True"] = true
    data["False"] = false

    // 整型
    data["Year"] = 35

    // 结构体
    type User struct {
        Name string
        Pass string
    }
    data["User"] = User{Name: "nljb", Pass: "1234"}

    // 列表
    data["List"] = []int{1, 2, 35, 4, 5, 6, 7, 8}

    // 列表结构体
    a := User{Name: "nljb", Pass: "1234"}
    b := User{Name: "jbnl", Pass: "4321"}
    data["Structs"] = []User{a, b}

    // 模版变量
    data["var"] = "hello world !!!"

    data["maps"] = map[string]string{"name": "golang"}

    // 执行
  // Compile the template first (i. e. creating the AST)
tpl, err := pongo2.FromString(`
My first name is {{ Structs.0.Name }}. My last name is {{ Structs.1.Name}}.
{%if True%}true{% endif %}
{%if True and False %}true and false{% endif %}
{%if True or False %}true or false{% endif %}
{% if False %}
   ... display 1
{% elif True %}
   ... display 2
{% else %}
   ... display 3
{% endif %}

{# 这是注释 for的用法 可以嵌套使用 {% for %} 标签 #}
  {% for p in Structs %}
            <li>{{ p.Name }}</li>
            <li>{{ p.Pass }}</li>
		{% endfor %}

`)
if err != nil {
	panic(err)
}
// Now you can render the template with the given 
// pongo2.Context how often you want to.
out, err := tpl.Execute(data)
if err != nil {
	panic(err)
}
fmt.Println(out) // Output: Hello Florian!

}