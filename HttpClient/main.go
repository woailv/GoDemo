package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

func main() {
	resp, err := http.PostForm("http://127.0.0.1:20008/7jssux2gw7", url.Values{
		"name":[]string{"haha"},
		"age":[]string{"11"},
	})
	if err!=nil {
		panic(err)
	}
	buf, err := ioutil.ReadAll(resp.Body)
	if err!=nil {
		panic(err)
	}
	fmt.Println(string(buf))
}
