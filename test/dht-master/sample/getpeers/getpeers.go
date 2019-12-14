package main

import (
	"fmt"
	"log"
	"time"

	"github.com/shiyanhui/dht"
)

func main() {
	d := dht.New(nil)
	d.OnGetPeersResponse = func(infoHash string, peer *dht.Peer) {
		fmt.Printf("GOT PEER: <%s:%d>\n", peer.IP, peer.Port)
	}

	go func() {
		for {
			// ubuntu-14.04.2-desktop-amd64.iso
			err := d.GetPeers("7de62b3296e13e174377cfbdcad40e8a2cce32ba")
			if err != nil && err != dht.ErrNotReady {
				log.Fatal(err)
			}

			if err == dht.ErrNotReady {
				time.Sleep(time.Second * 1)
				continue
			}

			break
		}
	}()

	d.Run()
}
