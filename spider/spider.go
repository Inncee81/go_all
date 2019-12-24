package main

import (
	"encoding/hex"
	"encoding/json"
	"fmt"
	"net/http"
	_ "net/http/pprof"
    "./cs"
	. "./model"
	"./tools"
	"github.com/shiyanhui/dht"
)



//. "../m	id, _ = G.NextID()odel"
//"../tools"

//获取DB初始化
var db = tools.GetDB()


type file struct {
	Path   []interface{} `json:"path"`
	Length int           `json:"length"`
}

type bitTorrent struct {
	InfoHash string `json:"infohash"`
	Name     string `json:"name"`
	Time     string `json:"time"`
	Files    []file `json:"files,omitempty"`
	Length   int    `json:"length,omitempty"`
	Hashpath string `json:"hashpath,omitempty"`
}

func tf(infohash string) bool {
	u := Bt_list{}
	i := Bt_list{}
	//自动生成sql： SELECT * FROM `users`  WHERE (username = 'tizi365') LIMIT 1
	db.Where(" infohash = ?", infohash).Take(&u)
	if u == i {
		return false
	}
	return true

}

func mysql(name, files, infohash string) {

	uu := &Bt_list{Name: name, Files: files, Infohash: infohash}
	//定义结构体&数据
	if err := db.Create(uu).Error; err != nil {
		fmt.Println("插入失败", err)
		return
	}

}

func main() {
    
	fmt.Println(tf("abc"))

	var zlen int
	var znum int
    G, _ := cs.NewGUID(1)
	var genid int64

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
				Time:     tools.Date("Y-m-d"),
			}

			if v, ok := info["files"]; ok {
				dx := 0
				files := v.([]interface{})
				bt.Files = make([]file, len(files))

				for i, item := range files {
					f := item.(map[string]interface{})
					bt.Files[i] = file{
						Path:   f["path"].([]interface{}),
						Length: f["length"].(int),
					}
					dx += f["length"].(int)
				}
				bt.Length = dx
				zlen = dx
				znum = len(files)
				dx = 0

			} else if _, ok := info["length"]; ok {
				bt.Length = info["length"].(int)
				zlen = bt.Length
				znum = 1
			}
         genid, _ = G.NextID()
		gid := fmt.Sprintf("%v", genid)
			bt.Hashpath = gid
			data, err := json.Marshal(bt)
			if err == nil {

				infohash := fmt.Sprintf("%s", hex.EncodeToString(resp.InfoHash))

				if !tf(infohash) {

					d := make(map[string]interface{})
					fmt.Printf("%s\n", info["name"])
					json := fmt.Sprintf("%s", data)

					fmt.Println("bt-----", bt)
					d["info"] = info
					d["infohash"] = infohash
					d["files"] = bt.Files
					d["time"] = tools.Date("Y-m-d")
					d["zlen"] = zlen
					d["znum"] = znum
					if len(bt.Files) > 1 {
						d["f"] = true
					} else {
						d["f"] = false
					}
					fmt.Println("bt.Files", bt.Files)

					path := `spider\cili\info.html`
					sdatas := tools.Tpl(d, path)
					genpath := "spider/html/"
					filename := bt.Hashpath + ".html"
					tools.W_file(genpath+filename, sdatas)
					mysql(info["name"].(string), json, infohash)

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
