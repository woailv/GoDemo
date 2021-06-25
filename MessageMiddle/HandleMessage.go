package main

import (
	"GoDemo/Echo"
	"fmt"
	"time"
)

var id = 0

func ReceiverMessage(themeId string, content string) {
	subIdList := themeId2SubIdList[themeId]
	message := &Message{
		Id:      fmt.Sprintf("messageId.%d", id),
		Time:    time.Now().Unix(),
		Content: content,
	}
	err := MessageSave(message)
	if err != nil {
		Echo.Json("保存消息失败:", err)
		return
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
		sub := SubGetById(mss.SubId)
		if sub == nil {
			Echo.Json("订阅中断:", mss.SubId)
			continue
		}
		message := MessageGetById(mss.MessageId)
		if message == nil {
			Echo.Json("消息丢失:", mss.MessageId)
		}
		err := sub.SendMessage(message)
		if err != nil {
			Echo.Json("消息发送失败:", mss.Id)
		}
	}
}
