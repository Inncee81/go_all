package main

import (
	"fmt"
	"time"
)

func main() {
	var a int
	var b, c, d int
	var bb, cc, dd string
	bbb, ccc, ddd := 111, 222, 333
	fmt.Println("数字", a, b, c, d)
	fmt.Println("字符", bb, cc, dd)
	if bb == "" && cc == "" && dd == "" {
		fmt.Println("bb cc dd is not")
	}
	fmt.Println(bbb, ccc, ddd)

	s := "packagevarfunc"
	ss := []byte(s)
	ss[0] = 'H'
	s1 := string(ss)
	fmt.Println(s1)
	cs := Cstr("huvip", 0, 'H')
	fmt.Println(cs)
	fmt.Println(Cstr("myhuvi", 0, 'M'))
	start := time.Now()
	end := time.Now()
	fmt.Println("time:", end.Sub(start).Seconds())
	for i:=0;i<100;i++{
	fmt.Println(Fstr(` 1.png
Notes to Foreigners Applying for Residence Permit for Study and Extension
read_image.png
Notes to Foreigners Applying for Private Residence Permit and Extension (Innovation and Entrepreneurship)
2.png
Notes to Foreigners Applying for Private Residence Permit and Extension
2.png
Application of Foreign Talents for a Residence Permit for Work (Annotated with "Talents")
6.png
How to apply for work residence permits and for extension
1.png
Interim measures for the participation of foreigners employed in China in social insurance
More Services >>
NewsEvents
1
Nov
Exhibition on overseas Chinese kicks off in Beijing An exhibition opened at the National Museum of China to mark the contribution that overseas Chinese made to the motherland.
1
Nov
Beijing launches public project to help citizens quit smoking The health authority in Beijing has launched a program to recruit 500 smokers this year to help them quit smoking for free through outpatient services and a dedicated hotline.
1
Nov
5G commercialization launched in Beijing China has launched its 5G commercialization at an annual international Postal and Telecom expo in Beijing.More >>
Get Involved
 Survey on New Version of the eBeijing Website Survey on New Version of the eBeijing Website
We are always trying to improve the service. Now we would like to hear how do you like what we did.

More >>
 Governmental Website Survey.jpg eBeijing Governmental Website Survey
What is your expectations for governmental website? What kinds of content should eBeijing have? What services are you interested in?`))
	}	

}
//字符串反转
func Fstr(s string) string {
	flen := len(s) - 1
	jitem := ""
	for flen >= 0 {
		jitem += string(s[flen])
		flen--
	}
	return jitem
}

//字符串修改某个字符
func Cstr(s string, i int, b byte) string {
	ss := []byte(s)
	ss[i] = b
	s1 := string(ss)
	return s1

}
