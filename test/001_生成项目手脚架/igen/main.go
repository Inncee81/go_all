//
package main

import (
	"fmt"
	"os"
	"strconv"

	. "./private"
	"github.com/urfave/cli"
)

var p = fmt.Println

func Mygen(s string) {
	if Ispath("tpl/go_main.go") {
		p("没有这个文件")
	} else {
		Mkdir("tpl")
		W_file("tpl/go_main.tpl", `//
package main

import (
	"fmt"

	_ "../../private"
)

var p = fmt.Println

func main() {
	
}
`)
//加载tpl模版到字符串
	go_main := R_file("tpl/go_main.tpl")
	//生成项目的路径
	spath := s
	//替换\ 到 /
	path := Ren(spath)
	//建立文件夹
	Mkdir(path)

	//循环生成目录
	for i := 1; i < 10; i++ {
		ii := strconv.Itoa(i)
		mypath := path + "/test/00" + ii + "/"
		myfile := "00" + ii + ".go"
		Mkdir(mypath)
		W_file(mypath+myfile, go_main)
	}

	//建立目录
	Mkdir(path + "/private")
	Mkdir(path + "/public")
	Mkdir(path + "/testGen")
	//读取文件模版并写入到文件中
	W_file(path+"/main.go", go_main)
	//	W_file(path+"/private/dir.go", go_src_dir)

	}

	

}

func igen(s string) {
	pathname := s
	Mygen(pathname)
	values, err := AllListDir("tpl", ".go")
	if err == nil {
		for _, value := range values {
			p(value)
			go_tpl := R_file(value)
			W_file(pathname+"/private/"+Paths(value, 1), go_tpl)

		}
	}
}
func main() {

	app := cli.NewApp()
	app.Name = "igen"                 //项目名称
	app.Author = "huvip"              //作者名称
	app.Version = "1.0.0 2019.10.3"   //版本号
	app.Copyright = "@Copyright 2019" //版权保护
	app.Usage = "快速使用go命令行生成项目手脚架"    //说明
	cli.HelpFlag = cli.BoolFlag{      //修改系统默认
		Name:"help",
		Usage:`igen -d 项目名称
	igen -d C:\项目名称`,
	}
	cli.VersionFlag = cli.BoolFlag{ //修改系统默认
		Name:  "version, v",
		Usage: "显示版本号",
	}
	app.Flags = []cli.Flag{
		cli.StringFlag{Name: "d", Value: "", Usage: "指定目录地址与项目名称"},
	}
	app.Action = func(c *cli.Context) error {
		if c.NumFlags() == 0 {
			cli.ShowAppHelp(c)
			return nil
		}
		//
		dd := c.String("d")
		if dd != "" {

			igen(dd)
		}

		return nil
	}

	var err error
	err = app.Run(os.Args)
	if err != nil {
		//	p(err)
	}

}
