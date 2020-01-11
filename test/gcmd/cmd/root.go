package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "gcmd",                 //应用名称
	Short: "生成cobray手脚架的golang应用", //应用短说明
	Long: `
gcmd init 生成一个cobray项目
gcmd add  添加命令参数`,
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		//	fmt.Println(err)
		os.Exit(1)
	}
}
