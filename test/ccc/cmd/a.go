package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
)

// aCmd表示a命令
var aCmd = &cobra.Command{
	Use:   "a", //使用
	Short: "a的命令的简要说明",
	Long: "a命令的较长描述，可能包含示例和使用命令的用法",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("a命令执行中...")
		
		fmt.Println(args)
	},
}

func init() {
	rootCmd.AddCommand(aCmd)
}
