package main

import (
	"net/http"
	"net/url"
)

func main() {
	resp, err:=http.PostForm("http://127.0.0.1:9500", url.Values{
		"name":[]string{"haha"},
	})
	if err!=nil{
		panic(err)
	}
	defer resp.Body.Close()
}
