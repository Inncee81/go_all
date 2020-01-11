package main

import (
	"github.com/atotto/clipboard"
)

func main() {
	// 复制内容到剪切板
	clipboard.WriteAll(`复制这段内容到剪切板`)

	// 读取剪切板中的内容到字符串
	content, err := clipboard.ReadAll()
	if err != nil {
		panic(err)
	}
	
	println(content)
}