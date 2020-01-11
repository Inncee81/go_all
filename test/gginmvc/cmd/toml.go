package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"../config"
)

// tomlCmd表示toml命令
var cf config.TomlConfig

var tomlCmd = &cobra.Command{
	Use:   "toml", //使用
	Short: "toml的命令的简要说明",
	Long: "toml命令的较长描述，可能包含示例和使用命令的用法",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("toml命令执行中...")
		cf=config.Tconf()
		fmt.Println(args)
		fmt.Printf("Title: %s\n", cf.Title)
		for _, v := range cf.Rou.Roufunc {
		fmt.Println(v)	
		}
	},
}

func init() {
	rootCmd.AddCommand(tomlCmd)
}
