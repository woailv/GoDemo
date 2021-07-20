package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"
)

func main() {
	PostJson()
}

type GooglePubSubMsg struct {
	Message struct {
		Attributes struct {
			Key string `json:"key"`
		} `json:"attributes"`
		Data      string `json:"data"`
		MessageID string `json:"messageId"`
	} `json:"message"`
	Subscription string `json:"subscription"`
}

func PostJson() {
	i := 0
	for {
		i++
		bs, _ := json.Marshal(GooglePubSubMsg{
			Message: struct {
				Attributes struct {
					Key string `json:"key"`
				} `json:"attributes"`
				Data      string `json:"data"`
				MessageID string `json:"messageId"`
			}{MessageID: strconv.Itoa(i)},
			Subscription: "",
		})
		path := "http://127.0.0.1:8080/2bbqbx"
		path = "http://127.0.0.1:8080/r6rgxf"
		resp, err := http.Post(path, "application-json", bytes.NewReader(bs))
		if err != nil {
			panic(err)
		}
		defer resp.Body.Close()
		fmt.Println(resp.StatusCode)
		bs, err = ioutil.ReadAll(resp.Body)
		if err != nil {
			panic(err)
		}
		fmt.Println(string(bs))
	}
}

func PostForm() {
	resp, err := http.PostForm("http://127.0.0.1:8080/hzujgr", url.Values{
		"name": []string{"haha"},
		"age":  []string{"11"},
	})
	if err != nil {
		panic(err)
	}
	buf, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(buf))
}
