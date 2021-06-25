package MessageMiddle

import "sync"

var mssChanLen = 5

var mssChan = make(chan *SubMessage, mssChanLen)

var id2SubMessageMap = map[string]*SubMessage{}

var id2SubMessageMapLock = &sync.Mutex{}

func LockSubMessageFn(fn func()) {
	LockFn(id2SubMessageMapLock, fn)
}

const (
	SubMessageStatus1 = 1 // 待发送
	SubMessageStatus2 = 2 // 队列中
	SubMessageStatus3 = 3 // 成功
	SubMessageStatus4 = 4 // 失败
)

type SubMessage struct {
	Id         string
	MessageId  string
	SubId      string
	Status     int8
	RetryTimes int
}

func SubMessageSave(mss *SubMessage) error {
	LockSubMessageFn(func() {
		id2SubMessageMap[mss.Id] = mss
	})
	return nil
}
