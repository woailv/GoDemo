package main

import (
	"GoDemo/Echo"
	"GoDemo/Err"
	"fmt"
	"net/http"
	"strings"
	"time"
)

type Theme struct {
	Id        string
	Name      string
	Method    string
	Url       string // domain/id
	TableName string // default == Id
}

type Message struct {
	Id      string
	Time    int64
	Content string // 方便搜索
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

var id2ThemeMap = map[string]*Theme{
	"1": {Id: "1", Name: "a", Method: "GET", Url: "/1", TableName: "1"},
}

var themeId2SubIdList = map[string][]string{
	"1": {"11", "12", "13"},
}

var mssChan = make(chan *MessageSubStatus, 100)

var mssList = []*MessageSubStatus{}

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
		theme := id2ThemeMap[themeId]
		if theme == nil {
			return
		}
		if theme.Method != request.Method {
			return
		}
		ReceiverMessage(themeId, "this is a message")
	})
	go SendMessage()
	err := http.ListenAndServe(port, nil)
	Err.IfPanic(err)
}

func ReceiverMessage(themeId string, content string) {
	subIdList := themeId2SubIdList[themeId]
	message := Message{
		Id:      "message1",
		Time:    time.Now().Unix(),
		Content: content,
	}
	for _, subId := range subIdList {
		mss := &MessageSubStatus{
			Id:         "mssId1",
			MessageId:  message.Id,
			SubId:      subId,
			Status:     MessageSubStatus2,
			RetryTimes: 0,
		}
		mssList = append(mssList, mss)
		Echo.Json("receive:", mssList)
		mssChan <- mss
	}
}

func SendMessage() {
	for mss := range mssChan {
		for i := range mssList {
			if mssList[i].Id == mss.Id {
				mssList[i].Status = MessageSubStatus3
			}
		}
		Echo.Json("send:", mssList)
	}
}
