package main

import (
	"fmt"
	"time"
)

var id = 0

func ReceiverMessage(themeId string, content string) {
	subIdList := themeId2SubIdList[themeId]
	message := Message{
		Id:      fmt.Sprintf("messageId.%d", id),
		Time:    time.Now().Unix(),
		Content: content,
	}
	for _, subId := range subIdList {
		mss := &MessageSubStatus{
			Id:         fmt.Sprintf("mssId.%d", id),
			MessageId:  message.Id,
			SubId:      subId,
			Status:     MessageSubStatus1,
			RetryTimes: 0,
		}
		mssList = append(mssList, mss)
		if len(mssChan) < mssChanLen {
			mss.Status = MessageSubStatus2
			mssChan <- mss
		} else {
			mssStatus1List = append(mssStatus1List, mss)
		}
		id++
	}
}

func SendMessage() {
	for mss := range mssChan {
		panic("TODO")
		_ = subId2SubMap[mss.SubId]
	}
}
