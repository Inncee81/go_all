package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
)

// modelCmd表示model命令
var modelCmd = &cobra.Command{
	Use:   "model", //使用
	Short: "model的命令的简要说明",
	Long: "model命令的较长描述，可能包含示例和使用命令的用法",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("model命令执行中...")
		
		fmt.Println(args)
	},
}

func init() {
	rootCmd.AddCommand(modelCmd)
}
