//
package main

import (
	"fmt"
	"os"

	_ "../../private"
	"github.com/urfave/cli"
)

var p = fmt.Println

func main() {
	app := cli.NewApp()
	app.Name = "igen"                 //项目名称
	app.Author = "huvip"              //作者名称
	app.Version = "1.0.0 2019.10.3"   //版本号
	app.Copyright = "@Copyright 2019" //版权保护
	app.Usage = "快速使用go命令行生成项目手脚架"       //说明
	cli.HelpFlag = cli.BoolFlag{      //修改系统默认
		Name:  "help",
		Usage: `igen -d 项目名称`,
	}
	cli.VersionFlag = cli.BoolFlag{ //修改系统默认
		Name:  "version, v",
		Usage: "显示版本号",
	}
	app.Flags = []cli.Flag{
		cli.StringFlag{Name: "d", Value: "go-itemname", Usage: "指定目录地址与项目名称"},
	}
	app.Action = func(c *cli.Context) error {
		if c.NumFlags() == 0 {
			cli.ShowAppHelp(c)
			return nil
		}
		//
		dd := c.String("d")
		if dd != "" {

		}

		return nil
	}

	var err error
	err = app.Run(os.Args)
	if err != nil {
		//	p(err)
	}
}
