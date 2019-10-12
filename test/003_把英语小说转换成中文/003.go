//
package main

import (
	"bufio"
	"fmt"
	"io"
	"net/url"
	"os"
	"strconv"
	"strings"

	. "../../private"
)

var p = fmt.Println

func fy(path string) {
	Mkdir("转换完成/cje/cn/")
	Mkdir("转换完成/cje/json/")
	Mkdir("转换完成/cje/encn/")
	q := R_file(path)
	mypath := Paths(path, 4)
	p(mypath)

	qq := url.QueryEscape(q)
	p(qq)
	get := Get(`https://translate.google.cn/translate_a/single?client=gtx&sl=en&tl=zh&dt=t&q=` + qq)
	W_file("del.txt", get)
	cn, json, encn := rl_file("del.txt")
	W_file("转换完成/cje/cn/"+mypath+".txt", cn)
	W_file("转换完成/cje/json/"+mypath+".json", json)
	W_file("转换完成/cje/encn/"+mypath+".txt", encn)
	//rl_file("color.txt")
	p(cn)

}
func main() {
	listdir()
}

//显示查找所有目录下的文件
func listdir() {
	Mkdir("需要转换")
	values, err := AllListDir("需要转换", ".txt")
	if err == nil {
		for _, value := range values {
			p(value) //需要转换\Harry Potter And The Chamber Of Secrets.txt
			vh := Ren(value)
			name := Paths(vh, 4) //Harry Potter And The Chamber Of Secrets 文件名
			pdir := "转换完成/" + name + "_合并/"
			Mkdir(pdir)
			rl(vh)
			listdir1()
			W_file(pdir+name+"_cn.txt", listdir2(name, "cn", ".txt"))
			W_file(pdir+name+"_encn.txt", listdir2(name, "encn", ".txt"))
			W_file(pdir+name+"_json.json", listdir2(name, "json", ".json"))
		}

	}
}
func listdir1() {
	Mkdir("转换完成")

	values, err := AllListDir("转换完成", ".txt")
	if err == nil {
		for _, value := range values {
			p(value) // 转换完成\encn\Harry Potter And The Chamber Of Secrets31.txt
			vh := Ren(value)

			p(Paths(vh, 4)) //Harry Potter And The Chamber Of Secrets31
			fy(vh)

		}
	}
}
func listdir2(name, dir, hz string) string {
	Mkdir("转换完成/cje")
	num := 0
	str := ""
	hbpath := "转换完成/cje/" + dir + "/"
	values, err := AllListDir(hbpath, hz)
	if err == nil {
		for _, value := range values {
			ffile := Ren(name + strconv.Itoa(num) + hz)
			rfile := R_file(hbpath + ffile)
			str += rfile + "\n"

			p(value) // 转换完成\encn\Harry Potter And The Chamber Of Secrets31.txt
			//vh := Ren(value)
			//	p(Paths(vh, 4)) //Harry Potter And The Chamber Of Secrets31
			//	fy(vh)
			num++
		}
		return str
	}
	return ""
}

//传递一个文件.txt 分割并写入多个小txt文件
func rl(path string) {
	my := Paths(path, 4)
	p(my + "\n")
	mydir := "转换完成/" + my + "/"
	Mkdir(mydir)
	s := ""
	znum := 0
	num := 0
	var pp *int
	var ss *string
	fi, err := os.Open(path)
	if err != nil {
		fmt.Printf("Error: %s\n", err)
	}
	defer fi.Close()

	br := bufio.NewReader(fi)
	for {
		a, _, c := br.ReadLine()
		stra := string(a)
		if c == io.EOF {
			break
		}
		if znum < 12000 {
			znum += len(stra)
			s += stra + "\n"
			pp = &znum
			ss = &s

		} else if znum >= 12000 {
			num++

			W_file(mydir+my+strconv.Itoa(num)+".txt", s)
			*pp = 0
			*ss = ""
		}
		if len(s) > 1 {
			W_file(mydir+my+strconv.Itoa(num)+".txt", s)
		}

	}
}

//传递一个文件.txt 返回中文,中英json ,中英对照 string
func rl_file(path string) (string, string, string) {
	json := ""
	astr := ""
	encn := ""
	fi, err := os.Open(path)
	if err != nil {
		fmt.Printf("Error: %s\n", err)
		return "N", "N", "N"
	}
	defer fi.Close()

	br := bufio.NewReader(fi)
	for {
		a, _, c := br.ReadLine()
		stra := string(a)
		if c == io.EOF {
			break
		}
		if Rec(stra) {
			straa := strings.Replace(stra, `[[[`, `,[`, 1)
			saa := strings.Replace(straa, `,null,"en"]`, ``, 1)
			sa := strings.Replace(saa, `\n`, "", -1)

			s := strings.Split(sa, `,null,`)
			p(s[0])
			if len(s[0]) > 0 {
				ress := strings.Replace(s[0], `","`, `"&_&"`, -1)
				chen := strings.Split(ress, `&_&`)
				yen := chen[1]
				qe := strings.Replace(yen, `"`, ``, -1)
				qen := strings.Replace(qe, `\`, ``, -1)
				p(qen)
				//p(fress[0])
				zh := strings.Split(chen[0], `[`)
				yzh := zh[1]
				qzh := strings.Replace(yzh, `"`, ``, -1)

				p(qzh)

				yesstr := `{"zh":` + yzh + `, "en":` + yen + `},`
				json += yesstr
				//		astr += fress[0] + "\n"
				p(yesstr)
				astr += qzh + "\n"
				encn += qzh + "\n" + qen + "\n"
			}

		}

	}
	return astr, json, encn
}

func Restr(s string) bool {
	//	fmt.Println(strings.Contains(s, `//`)) //true
	return strings.Contains(s, `//`)

}

//]如果是1返回false
func Rec(s string) bool {
	//	fmt.Println(strings.Contains(s, `//`)) //true
	nlen := len(s)
	//p(nlen)
	if nlen == 1 || nlen == 0 {
		return false
	} else {

		return true
	}

}
