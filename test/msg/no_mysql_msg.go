// WebChat project main.go
package main

import (
	"encoding/json"
	"fmt"
_	"io/ioutil"
	"net/http"
	"net/url"
	_ "strconv"
	"strings"
	"time"

	_ "github.com/gin-gonic/gin"
	"golang.org/x/net/websocket"
)

var p = fmt.Println

//全局信息
var datas Datas
var users map[*websocket.Conn]string
var jsond Jsond

func main() {
	fmt.Println("启动时间")
	fmt.Println(time.Now())

	//初始化
	datas = Datas{}
	users = make(map[*websocket.Conn]string)
	//绑定效果页面
	http.HandleFunc("/", h_index)
	//绑定socket方法
	http.Handle("/webSocket", websocket.Handler(h_webSocket))
	//开始监听
	http.ListenAndServe(":1094", nil)

}
func Gurl(urltest string) string {
	encodeurl := url.QueryEscape(urltest)
	fmt.Println(encodeurl)
	return encodeurl
}

func h_index(w http.ResponseWriter, r *http.Request) {

}

func h_webSocket(ws *websocket.Conn) {

	var userMsg UserMsg
	var data string

	for {

		fmt.Println(users[ws])

		//判断是否重复连接
		if _, ok := users[ws]; !ok {
			users[ws] = "匿名"
			p("是否匿名:", users[ws])
		}
		userMsgsLen := len(datas.UserMsgs)
		fmt.Println("UserMsgs", userMsgsLen, "users长度：", len(users))

		//有消息时，全部分发送数据 userMsgsLen > 0
		if userMsgsLen > 0 {
			b, errMarshl := json.Marshal(datas)
			if errMarshl != nil {
				fmt.Println("全局消息内容异常...")
				break
			}
			for country := range users {
				fmt.Println(country, "是", users[country])
				num := jsond.Msg

				if num != "" {

					if users[country] == jsond.UserName {
						p("country:", country)
						start := time.Now().Unix()
						n := 60
						p(n)
						i := 0
						p(string(data))
						//关闭测试for循环
						//	for i < 200 {
						//	errMarshl = websocket.Message.Send(country, string(data))
						//消息推送分发线程并发无限制(关闭B、不存储消息、只推送消息、速度最快）
						//go websocket.Message.Send(country, string(data))
						go websocket.Message.Send(country, string(data))
						i++

						///* B 消息插入数据库,限制开4100并发
						yunum := i % 4100
						if yunum != 0 && n >= 69 {
						} else if yunum == 0 || n <= 68 {
							n := 0
							//缓冲线程完成后,继续线程并发
							n++
						}
						//*/
						p("线程i=", i)
						//	} //关闭测试for循环}
						end := time.Now().Unix()
						fmt.Printf("执行消耗的时间为:%v秒", end-start)

					}
					if users[country] == jsond.ToID {
						errMarshl = websocket.Message.Send(country, string(data))
					}
				} else {
					if users[country] == jsond.UserName {
					}
				}
				if errMarshl != nil {
					//移除出错的链接
					delete(users, country)
					fmt.Println("发送出错...")
					break

				}
			}

			//全部分发送数据
			for key, _ := range users {
				fmt.Println("jsond-form", jsond.UserName)
				fmt.Println("jsond-toid", jsond.ToID)

				fmt.Println(key)
				fmt.Println("key-------------")

				fmt.Println(users)

				fmt.Println("users-------------")
				fmt.Println("msg" + string(b))

				//errMarshl = websocket.Message.Send(key, string(b))

			}

			datas.UserMsgs = make([]UserMsg, 0)
		}

		fmt.Println("开始解析数据...")
		err := websocket.Message.Receive(ws, &data)
		fmt.Println("data：", data)
		json.Unmarshal([]byte(data), &jsond)

		if err != nil {
			//移除出错的链接
			delete(users, ws)
			p("users:", users)
			p("ws:", ws)
			fmt.Println("接收出错...用户下线")

			break
		}

		data = strings.Replace(data, "\n", "", 0)
		err = json.Unmarshal([]byte(data), &userMsg)
		if err != nil {
			fmt.Println("解析数据异常...")
			break
		}
		fmt.Println("请求数据类型：", userMsg.DataType)

		switch userMsg.DataType {
		case "send":
			//赋值对应的昵称到ws
			if _, ok := users[ws]; ok {
				users[ws] = userMsg.UserName

				//清除连接人昵称信息
				datas.UserDatas = make([]UserData, 0)
				//重新加载当前在线连接人
				for _, item := range users {

					userData := UserData{UserName: item}
					datas.UserDatas = append(datas.UserDatas, userData)
				}
			}
			datas.UserMsgs = append(datas.UserMsgs, userMsg)
		}
	}

}

type Jsond struct {
	Cmdcode  string
	UserName string
	ToID     string
	Msg      string
	DataType string
	Time     string
	Token    string
	Type     int
}
type UserMsg struct {
	UserName string
	Msg      string
	DataType string
}

type UserData struct {
	UserName string
}

type Datas struct {
	UserMsgs  []UserMsg
	UserDatas []UserData
}
