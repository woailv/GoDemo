package MessageMiddle

import "sync"

var mssChanLen = 5

var mssChan = make(chan *SubMessage, mssChanLen)

var id2SubMessageMap = map[string]*SubMessage{}

var id2SubMessageMapLock = &sync.Mutex{}

func LockSubMessageFn(fn func()) {
	LockFn(id2SubMessageMapLock, fn)
}

var subId2SubMessageId2NilMap = map[string]map[string]struct{}{}

var subId2SubMessageId2NilMapLock = &sync.Mutex{}

func LockSubId2SubMessageId2NilFn(fn func()) {
	LockFn(subId2SubMessageId2NilMapLock, fn)
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

func SubMessageSave(subMessage *SubMessage) error {
	LockSubMessageFn(func() {
		id2SubMessageMap[subMessage.Id] = subMessage
	})
	LockSubId2SubMessageId2NilFn(func() {
		if subId2SubMessageId2NilMap[subMessage.SubId] == nil {
			subId2SubMessageId2NilMap[subMessage.SubId] = map[string]struct{}{}
		}
		subId2SubMessageId2NilMap[subMessage.SubId][subMessage.Id] = struct{}{}
	})
	return nil
}

func SubGetAllSubMessageIdsBySubId(subId string) []string {
	subMessageIds := []string{}
	LockSubId2SubMessageId2NilFn(func() {
		m := subId2SubMessageId2NilMap[subId]
		for subMessageId := range m {
			subMessageIds = append(subMessageIds, subMessageId)
		}
	})
	return subMessageIds
}

func SubMessageGetById(id string) *SubMessage {
	var subMessage *SubMessage
	LockSubMessageFn(func() {
		subMessage = id2SubMessageMap[id]
	})
	return subMessage
}

func SubDeleteAllSubMessageIdsBySubId(subId ...string) {
	for _, id := range subId {
		subMessageIds := SubGetAllSubMessageIdsBySubId(id)
		for _, subMessageId := range subMessageIds {
			subMessage := SubMessageGetById(subMessageId)
			MessageDelete(subMessage.MessageId)
			LockSubMessageFn(func() {
				delete(id2SubMessageMap, subMessageId)
			})
		}
		LockSubId2SubMessageId2NilFn(func() {
			delete(subId2SubMessageId2NilMap, id)
		})
	}
}
