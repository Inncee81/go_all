package main

import (
	"fmt"
	"strings"
	"github.com/atotto/clipboard"
)

func main() {

	arr := strings.Split(`C:\Program Files (x86)\NetSarang\Xshell 6\;C:\WINDOWS\system32;C:\WINDOWS;C:\WINDOWS\System32\Wbem;C:\WINDOWS\System32\WindowsPowerShell\v1.0\;C:\Program Files\IDM Computer Solutions\UltraEdit;C:\Program Files\Chez Scheme 9.5.2\bin\a6nt;C:\Go\bin;C:\Program Files\Git\bin;C:\Program Files\nodejs\;C:\Program Files\Calibre2\;C:\Users\DELL\Desktop\exe;C:\phpStudy\PHPTutorial\Apache\bin;C:\Users\DELL\AppData\Local\Programs\Python\Python37\Scripts\;C:\Users\DELL\AppData\Local\Programs\Python\Python37\;C:\Users\DELL\AppData\Local\Microsoft\WindowsApps;C:\Users\DELL\AppData\Local\Programs\Fiddler;C:\Users\DELL\AppData\Local\atom\bin;C:\Users\DELL\go\bin;C:\Users\DELL\Desktop\myexe;D:\IntelliJ IDEA Community Edition 2019.1.2\bin;D:\GoLand 2019.1.1\bin;C:\Users\DELL\Desktop\nim\nim-0.19.4\bin;C:\Users\DELL\.nimble\bin;C:\Users\DELL\Desktop\nim\nim-0.19.4\dist\mingw64\bin;C:\Users\DELL\AppData\Roaming\Composer\vendor\bin;C:\Users\DELL\AppData\Local\GitHubDesktop\bin;C:\Users\DELL\AppData\Roaming\npm;C:\Users\DELL\Desktop\path;C:\Users\DELL\AppData\Local\Programs\Python\Python37\Scripts\;C:\Users\DELL\AppData\Local\Programs\Python\Python37\;C:\Users\DELL\AppData\Local\Microsoft\WindowsApps;C:\Users\DELL\AppData\Local\Programs\Fiddler;C:\Users\DELL\AppData\Local\atom\bin;C:\Users\DELL\go\bin;C:\Users\DELL\Desktop\myexe;D:\IntelliJ IDEA Community Edition 2019.1.2\bin;D:\GoLand 2019.1.1\bin;C:\Users\DELL\Desktop\nim\nim-0.19.4\bin;C:\Users\DELL\.nimble\bin;C:\Users\DELL\Desktop\nim\nim-0.19.4\dist\mingw64\bin;C:\Users\DELL\AppData\Roaming\Composer\vendor\bin;C:\Users\DELL\AppData\Local\GitHubDesktop\bin;C:\Users\DELL\AppData\Roaming\npm;C:\Users\DELL\Desktop\p;C:\Users\DELL\AppData\Local\Programs\Python\Python37\Scripts\;C:\Users\DELL\AppData\Local\Programs\Python\Python37\;C:\Users\DELL\AppData\Local\Microsoft\WindowsApps;C:\Users\DELL\AppData\Local\Programs\Fiddler;C:\Users\DELL\AppData\Local\atom\bin;C:\Users\DELL\go\bin;C:\Users\DELL\Desktop\myexe;D:\IntelliJ IDEA Community Edition 2019.1.2\bin;D:\GoLand 2019.1.1\bin;C:\Users\DELL\Desktop\nim\nim-0.19.4\bin;C:\Users\DELL\.nimble\bin;C:\Users\DELL\Desktop\nim\nim-0.19.4\dist\mingw64\bin;C:\Users\DELL\AppData\Roaming\Composer\vendor\bin;C:\Users\DELL\AppData\Local\GitHubDesktop\bin;C:\Users\DELL\AppData\Roaming\npm;C:\Program Files (x86)\NetSarang\Xshell 6\;C:\WINDOWS\system32;C:\WINDOWS;C:\WINDOWS\System32\Wbem;C:\WINDOWS\System32\WindowsPowerShell\v1.0\;C:\Program Files\IDM Computer Solutions\UltraEdit;C:\Program Files\Chez Scheme 9.5.2\bin\a6nt;C:\Go\bin;C:\Program Files\Git\bin;C:\Program Files\nodejs\;C:\Program Files\Calibre2\;C:\Users\DELL\Desktop\exe;C:\phpStudy\PHPTutorial\Apache\bin;C:\Users\DELL\AppData\Local\Programs\Python\Python37\Scripts\;C:\Users\DELL\AppData\Local\Programs\Python\Python37\;C:\Users\DELL\AppData\Local\Microsoft\WindowsApps;C:\Users\DELL\AppData\Local\Programs\Fiddler;C:\Users\DELL\AppData\Local\atom\bin;C:\Users\DELL\go\bin;C:\Users\DELL\Desktop\myexe;D:\IntelliJ IDEA Community Edition 2019.1.2\bin;D:\GoLand 2019.1.1\bin;C:\Users\DELL\Desktop\nim\nim-0.19.4\bin;C:\Users\DELL\.nimble\bin;C:\Users\DELL\Desktop\nim\nim-0.19.4\dist\mingw64\bin;C:\Users\DELL\AppData\Roaming\Composer\vendor\bin;C:\Users\DELL\AppData\Local\GitHubDesktop\bin;C:\Users\DELL\AppData\Roaming\npm;C:\Users\DELL\Deskto;C:\Users\DELL\Desktop\c;C:\Users\DELL\Desktop\d`, ";")
	re := RemoveRepeatedElement(arr)
	cc := ""
	for i := 0; i < len(re); i++ {
		fmt.Println(re[i])
		cc += re[i] + ";"
	}
	// 复制内容到剪切板
	clipboard.WriteAll(cc)
	fmt.Println(cc)

}

//去除重复字符串
func RemoveRepeatedElement(arr []string) (newArr []string) {
	newArr = make([]string, 0)
	for i := 0; i < len(arr); i++ {
		repeat := false
		for j := i + 1; j < len(arr); j++ {
			if arr[i] == arr[j] {
				repeat = true
				break
			}
		}
		if !repeat {
			newArr = append(newArr, arr[i])
		}
	}
	return newArr
}
