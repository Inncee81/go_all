package cmd

import (
	"os"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "gginmvc",   //应用名称
	Short: "gginmvc应用", //应用短说明
	Long: "gginmvc应用程序 功能说明： 有较长描述 可能包含示例和使用命令的用法",
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		//	fmt.Println(err)
		os.Exit(1)
	}
}
                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                 
                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                              