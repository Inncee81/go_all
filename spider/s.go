package main
import (
	"fmt"

	"./cs"

)



func main() {
	G, _ := cs.NewGUID(1)
	var genid int64

    for i := 0; i < 10000; i++ {

          genid, _ = G.NextID()

				gid := fmt.Sprintf("%v", genid)
		fmt.Println(gid)
	}
}
