package MessageMiddle

import "sync"

var mssChanLen = 5

var mssChan = make(chan *MessageSubStatus, mssChanLen)

var id2MessageSubStatusMap = map[string]*MessageSubStatus{}

var id2MessageSubStatusMapLock = &sync.Mutex{}

func LockMessageSubStatusFn(fn func()) {
	LockFn(id2MessageSubStatusMapLock, fn)
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

func MessageSubStatusSave(mss *MessageSubStatus) error {
	LockMessageSubStatusFn(func() {
		id2MessageSubStatusMap[mss.Id] = mss
	})
	return nil
}
