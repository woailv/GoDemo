package main

var mssChanLen = 5

var mssChan = make(chan *MessageSubStatus, mssChanLen)

var mssList = []*MessageSubStatus{}

var mssStatus1List = []*MessageSubStatus{}

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
