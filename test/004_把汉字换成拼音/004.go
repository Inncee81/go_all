
import (
	"fmt"

	. "package main
../../private"
	"github.com/chain-zhang/pinyin"
)

func main() {
	q := R_file("翻译中文.txt")
	str, err := pinyin.New(q).Split("").Mode(pinyin.InitialsInCapitals).Convert()
	if err != nil {
		// 错误处理
	} else {
		fmt.Println(str+"\n")
	}

	str, err = pinyin.New(q).Split(" ").Mode(pinyin.WithoutTone).Convert()
	if err != nil {
		// 错误处理
	} else {
		fmt.Println(str+"\n")
	}

	str, err = pinyin.New(q).Split(" ").Mode(pinyin.Tone).Convert()
	if err != nil {
		// 错误处理
	} else {
		fmt.Println(str+"\n")
		W_file("py",str)
	}

	str, err = pinyin.New(q).Convert()
	if err != nil {
		// 错误处理
	} else {
		fmt.Println(str+"\n")
	}
}
