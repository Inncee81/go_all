package cmd

import (
	"fmt"
	"os"

	"../tools"
	"github.com/spf13/cobra"
)

// addCmd表示add命令
var addCmd = &cobra.Command{
	Use:   "add", //使用
	Short: "gcmd add  添加命令参数明",
	Long:  `gcmd add  添加命令参数`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("add命令执行中...")
		data := make(map[string]interface{})

		data["addname"] = args[0]
		a, _ := os.Getwd()
		ac := a + "/cmd/"
		c := tools.Tpls(data, `package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
)

// {{addname}}Cmd表示{{addname}}命令
var {{addname}}Cmd = &cobra.Command{
	Use:   "{{addname}}", //使用
	Short: "{{addname}}的命令的简要说明",
	Long: "{{addname}}命令的较长描述，可能包含示例和使用命令的用法",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("{{addname}}命令执行中...")
		
		fmt.Println(args)
	},
}

func init() {
	rootCmd.AddCommand({{addname}}Cmd)
}`)
		mainFile := ac + string(args[0]) + `.go`
		tools.W_file(mainFile, c)
		fmt.Println(args)
	},
}

func init() {
	rootCmd.AddCommand(addCmd)
}
