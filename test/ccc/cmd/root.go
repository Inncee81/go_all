package cmd

import (
	"os"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "ccc",   //应用名称
	Short: "ccc应用", //应用短说明
	Long: "ccc应用程序 功能说明： 有较长描述 可能包含示例和使用命令的用法",
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		//	fmt.Println(err)
		os.Exit(1)
	}
}
