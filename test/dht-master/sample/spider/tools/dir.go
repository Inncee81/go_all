package tools 

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
	"path"
	"path/filepath"
	
)

func Get(s string) string {

	resp, err := http.Get(s)
	if err != nil {
		// handle error
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		// handle error
	}

	fmt.Println(string(body))
	return string(body)
}

//<GO语言获取目录列表用 ioutil.ReadDir()，遍历目录用 filepath.Walk()>

//获取指定目录下的所有文件，不进入下一级目录搜索，可以匹配后缀过滤。
func ListDir(dirPth string, suffix string) (files []string, err error) {
	files = make([]string, 0, 10)

	dir, err := ioutil.ReadDir(dirPth)
	if err != nil {
		return nil, err
	}

	PthSep := string(os.PathSeparator)
	suffix = strings.ToUpper(suffix) //忽略后缀匹配的大小写

	for _, fi := range dir {
		if fi.IsDir() { // 忽略目录
			continue
		}
		if strings.HasSuffix(strings.ToUpper(fi.Name()), suffix) { //匹配文件
			files = append(files, dirPth+PthSep+fi.Name())
		}
	}

	return files, err
}

//遍历获取指定目录及所有子目录下的所有文件，可以匹配后缀过滤。
func AllListDir(dirPth, suffix string) (files []string, err error) {
	files = make([]string, 0, 30)
	suffix = strings.ToUpper(suffix) //忽略后缀匹配的大小写

	err = filepath.Walk(dirPth, func(filename string, fi os.FileInfo, err error) error { //遍历目录

		if fi.IsDir() { // 忽略目录
			return nil
		}

		if strings.HasSuffix(strings.ToUpper(fi.Name()), suffix) {
			files = append(files, filename)
		}

		return nil
	})

	return files, err
}

//分离目录文件名后缀文件名函数
	/*
value	D:\yun\Go_项目生成\zz2\zz2_5\zz2_5.exe
n=0 D:\yun\Go_项目生成\zz2\zz2_5\
n=1 zz2_5.exe
n=2 D:\yun\Go_项目生成\zz2\zz2_5\zz2_5
n=3 .exe
n=4 zz2_5
	 */
func Paths(value string, n int) string {
	var files string
	suffix := path.Ext(value)
	path_no_suffix := strings.TrimSuffix(value, suffix)
	path, _ := filepath.Split(value)
	filename := filepath.Base(value)
	filenameOnly := strings.TrimSuffix(filename, suffix)
	if n == 0 {
		files = path
	} else if n == 1 {
		files = filename

	} else if n == 2 {
		files = path_no_suffix
	} else if n == 3 {
		files = suffix
	} else if n == 4 {
		files = filenameOnly	
	} else {
		files = ""
	}


	return files
}


func A_file(filepath, str_content string) {
	//创建或打开一个文件,追加写入
	fd, _ := os.OpenFile(filepath, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0644)
	buf := []byte(str_content + "\n")
	fd.Write(buf)
	fd.Close()
}

//重新创建一个文件,并写入
func W_file(userFile, str_content string) {
	fd, err := os.Create(userFile)
	if err != nil {
		fmt.Println(userFile, err)
		return
	}
	defer fd.Close()
	fd.WriteString(str_content + "\n")
}

//读文件打印成字符串
func R_file(filepath string) string {

	f, err := os.Open(filepath)
	if err != nil {
		fmt.Println("read file fail", err)
		return ""
	}
	defer f.Close()

	fd, err := ioutil.ReadAll(f)
	if err != nil {
		fmt.Println("read to fd fail", err)
		return ""
	}

	return string(fd)
}

//是否目录或文件
func Ispath(path string) bool {
	_, err := os.Stat(path)
	return err == nil || os.IsExist(err)
}

//创建多级目录
func Mkdir(path string) {
	if !Ispath(path) {
		//创建多级目录 os.Mkdir 单目录
		os.MkdirAll(path, os.ModePerm)
	}
}

//windows 目录\替换成 /  linux
func Ren(path string) string {
	return strings.Replace(path, "\\", "/", -1)
}

