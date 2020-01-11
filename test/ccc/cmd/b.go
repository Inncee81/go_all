package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
)

// bCmd表示b命令
var bCmd = &cobra.Command{
	Use:   "b", //使用
	Short: "b的命令的简要说明",
	Long: "b命令的较长描述，可能包含示例和使用命令的用法",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("b命令执行中...")
		
		fmt.Println(args)
	},
}

func init() {
	rootCmd.AddCommand(bCmd)
}
