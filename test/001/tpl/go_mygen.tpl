//
package main

import (
	"fmt"
	"strconv"

	. "./src/dir"
)

var p = fmt.Println

func main() {
	//加载tpl模版到字符串
	go_main := R_file("tpl/go_main.tpl")
	go_src_dir := R_file("tpl/go_src_dir.tpl")
	go_mygen := R_file("tpl/go_mygen.tpl")
	//生成项目的路径
	spath := `C:\Users\tjqm4\Documents\我的坚果云\go语言项目库`
	//替换\ 到 /
	path := Ren(spath)
	//建立文件夹
	Mkdir(path)

	//循环生成目录
	for i := 1; i < 10; i++ {
		ii := strconv.Itoa(i)
		iii := "/zz" + ii
		Mkdir(path + "/zz" + ii)
		for a := 1; a < 6; a++ {
			ai := strconv.Itoa(a)
			mypath := path + iii + "/" + iii + "_" + ai
			myfile := iii + "_" + ai + ".go"
			Mkdir(mypath)
			W_file(mypath+myfile, go_main)
		}
	}

	//建立目录
	Mkdir(path + "/src")
	Mkdir(path + "/bin")
	Mkdir(path + "/tpl")
	Mkdir(path + "/src/dir")
	//读取文件模版并写入到文件中
	W_file(path+"/tpl/go_main.tpl", go_main)
	W_file(path+"/tpl/go_src_dir.tpl", go_src_dir)
	W_file(path+"/tpl/go_mygen.tpl", go_mygen)
	W_file(path+"/mygen.go", go_mygen)
	W_file(path+"/src/dir/dir.go", go_src_dir)

}






