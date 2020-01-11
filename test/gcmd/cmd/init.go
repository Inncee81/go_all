package cmd

import (
	"fmt"

	"github.com/spf13/cobra"

	"os"

	"../tools"
)

// cCmd表示c命令
var initCmd = &cobra.Command{
	Use:   "init", //使用
	Short: "gcmd init  生成一个cobray项目",
	Long:  `gcmd init  生成一个cobray项目`,
	Run: func(cmd *cobra.Command, args []string) {
		data := make(map[string]interface{})

		data["name"] = args[0]
		c := tools.Tpls(data, `package cmd

import (
	"os"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "{{name}}",   //应用名称
	Short: "{{name}}应用", //应用短说明
	Long: "{{name}}应用程序 功能说明： 有较长描述 可能包含示例和使用命令的用法",
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		//	fmt.Println(err)
		os.Exit(1)
	}
}`)
		a, _ := os.Getwd()
		ar := fmt.Sprintf("%s", args[0])
		a = a + `/` + ar
		tools.Mkdir(a)
		tools.Mkdir(a + "/cmd")
		userFile := a + `/cmd/root.go`
		mainFile := a+`/main.go` 
		tools.W_file(userFile, c)
		tools.W_file(mainFile, `package main

import "./cmd"

func main() {
	cmd.Execute()
}`)

		fmt.Println(data["name"])
		fmt.Println(args)
	},
}

func init() {
	rootCmd.AddCommand(initCmd)
}
