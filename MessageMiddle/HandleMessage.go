package main

import (
	"GoDemo/Echo"
	"fmt"
	"time"
)

var id2ThemeMap = map[string]*Theme{
	"1": {Id: "1", Name: "a", Method: "GET", Url: "/1"},
}

var themeId2SubIdList = map[string][]string{
	"1": {"a", "b"},
}

var mssChan = make(chan *MessageSubStatus, 100)

var mssList = []*MessageSubStatus{}

var mssId = 0

func ReceiverMessage(themeId string, content string) {
	subIdList := themeId2SubIdList[themeId]
	message := Message{
		Id:      "message1",
		Time:    time.Now().Unix(),
		Content: content,
	}
	for _, subId := range subIdList {
		mss := &MessageSubStatus{
			Id:         fmt.Sprintf("mssId.%d", mssId),
			MessageId:  message.Id,
			SubId:      subId,
			Status:     MessageSubStatus2,
			RetryTimes: 0,
		}
		mssList = append(mssList, mss)
		Echo.Json("receive:", mssList)
		mssChan <- mss
		mssId++
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
