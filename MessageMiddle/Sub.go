package MessageMiddle

import (
	"GoDemo/Echo"
	"sync"
)

type Sub struct {
	Id            string
	ThemeId       string
	Name          string
	Url           string
	Method        string
	MaxRetryTimes int
}

func (sub *Sub) SendMessage(message *Message) error {
	Echo.JsonLnf("theme id:%s, sub id:%s, message id:%s", sub.ThemeId, sub.Id, message.Id)
	return nil
}

var id2SubMap = map[string]*Sub{}

var subId2SubMapLock = &sync.Mutex{}

func LockSubFn(fn func()) {
	LockFn(subId2SubMapLock, fn)
}

func SubSave(sub *Sub) error {
	LockSubFn(func() {
		id2SubMap[sub.Id] = sub
	})
	_ = ThemeAddSub(sub.ThemeId, sub.Id)
	return nil
}

func SubGetById(id string) *Sub {
	var sub *Sub
	LockSubFn(func() {
		sub = id2SubMap[id]
	})
	return sub
}
