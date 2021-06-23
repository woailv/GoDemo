package main

import (
	"GoDemo/Err"
	"fmt"
	"net/http"
	"strings"
)

type Theme struct {
	Id        string
	Name      string
	Method    string
	Url       string // domain/id
	TableName string // default == Id
}

type Message struct {
	Id   string
	Time int64
	Body string // 方便搜索
}

type Sub struct {
	Id            string
	ThemeId       string
	Name          string
	Url           string
	MaxRetryTimes int
}

const (
	MessageSubStatus1 = 1 // 待发送
	MessageSubStatus2 = 2 // 队列中
	MessageSubStatus3 = 3 // 成功
	MessageSubStatus4 = 4 // 失败
)

type MessageSubStatus struct {
	Id         string
	MessageId  string
	SubId      string
	Status     int8
	RetryTimes int
}

var idToThemeMap = map[string]*Theme{
	"1": {Id: "1", Name: "a", Method: "GET", Url: "/1", TableName: "1"},
}

func main() {
	port := ":8080"
	fmt.Println("http://127.0.0.1" + port)
	http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		path := request.URL.Path
		if strings.Contains(path, ".ico") {
			return
		}
		themeId := path[1:]
		fmt.Println("theme:", themeId)
		theme := idToThemeMap[themeId]
		if theme == nil {
			return
		}
		if theme.Method != request.Method {
			return
		}

	})
	err := http.ListenAndServe(port, nil)
	Err.IfPanic(err)
}
