//
package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

var p = fmt.Println

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

func main() {
	Get("http://dict.youdao.com/dictvoice?audio=a")
}
