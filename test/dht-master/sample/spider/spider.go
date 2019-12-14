package main

import (
	"encoding/hex"
	"encoding/json"
	"fmt"
	"github.com/shiyanhui/dht"
	"net/http"
	_ "net/http/pprof"
.	"./model"
	"./tools"
)


//. "../model"
//"../tools"

//获取DB初始化
 var db=tools.GetDB()



type file struct {
	Path   []interface{} `json:"path"`
	Length int           `json:"length"`
}

type bitTorrent struct {
	InfoHash string `json:"infohash"`
	Name     string `json:"name"`
	Files    []file `json:"files,omitempty"`
	Length   int    `json:"length,omitempty"`
}

func tf(infohash string)(bool){
u := Bt_list{}
i := Bt_list{}
	//自动生成sql： SELECT * FROM `users`  WHERE (username = 'tizi365') LIMIT 1
db.Where(" infohash = ?", infohash).Take(&u)
if u == i {
return false
}
return true

}
func cs()  {
		 
			 data :=make(map[string]interface{})

//{"infohash":"4dcba1163b9214731bf25090218394c48e372c8b","name":"[ThZu.Cc]stepsiblingscaught_my_step_sisters_gift_1920","files":[{"path":["[ThZu.Cc]stepsiblingscaught_my_step_sisters_gift_1920.mp4"],"length":2499086931},{"path":["thzkk.com.url"],"length":125},{"path":["桃花族地址发布器.chm"],"length":17100}]}

data["cs"]="这是测试"
	path :=`test\dht-master\sample\spider\cili\slist.html`
	sdatas := tools.Tpl(data,path)
	tools. W_file(`test\dht-master\sample\spider\html\a.html`, sdatas)
}

func mysql(name,files,infohash string){

		uu := &Bt_list{Name: name, Files :files,Infohash:infohash }
	//定义结构体&数据
	if err := db.Create(uu).Error; err != nil {
		fmt.Println("插入失败", err)
		return
	}

}


func main() {
fmt.Println(tf("abc"))


	go func() {
		http.ListenAndServe(":6060", nil)
	}()

	w := dht.NewWire(65536, 1024, 256)
	go func() {
		for resp := range w.Response() {
			metadata, err := dht.Decode(resp.MetadataInfo)
			if err != nil {
				continue
			}
			info := metadata.(map[string]interface{})

			if _, ok := info["name"]; !ok {
				continue
			}

			bt := bitTorrent{
				InfoHash: hex.EncodeToString(resp.InfoHash),
				Name:     info["name"].(string),
			}

			if v, ok := info["files"]; ok {
				files := v.([]interface{})
				bt.Files = make([]file, len(files))

				for i, item := range files {
					f := item.(map[string]interface{})
					bt.Files[i] = file{
						Path:   f["path"].([]interface{}),
						Length: f["length"].(int),
					}
				}
			} else if _, ok := info["length"]; ok {
				bt.Length = info["length"].(int)
			}

			data, err := json.Marshal(bt)
			if err == nil {
		
			infohash := fmt.Sprintf("%s",hex.EncodeToString(resp.InfoHash))

if (! tf(infohash)){

 			d :=make(map[string]interface{})
			 fmt.Printf("%s\n",info["name"])
			 json := fmt.Sprintf("%s",data)
			 
			 fmt.Println("bt-----",bt)
	d["info"]=info
	d["infohash"]=infohash
	if info["files"]=="" {
		d["f"]=true
	}else{
		d["f"]=false
		
	}



//{"infohash":"4dcba1163b9214731bf25090218394c48e372c8b","name":"[ThZu.Cc]stepsiblingscaught_my_step_sisters_gift_1920","files":[{"path":["[ThZu.Cc]stepsiblingscaught_my_step_sisters_gift_1920.mp4"],"length":2499086931},{"path":["thzkk.com.url"],"length":125},{"path":["桃花族地址发布器.chm"],"length":17100}]}

	path :=`test\dht-master\sample\spider\cili\slist.html`
	sdatas := tools.Tpl(d,path)
	tools. W_file(`test\dht-master\sample\spider\html\a.html`, sdatas)


				mysql(info["name"].(string),json,infohash)
			
}
				
			}
		}
	}()
	go w.Run()

	config := dht.NewCrawlConfig()
	config.OnAnnouncePeer = func(infoHash, ip string, port int) {
		w.Request([]byte(infoHash), ip, port)
	}
	d := dht.New(config)

	d.Run()
}
