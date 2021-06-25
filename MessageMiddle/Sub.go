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
	MaxRetryTimes int
}

func (sub *Sub) SendMessage(message *Message) error {
	Echo.Json("sub send message:", message)
	return nil
}

var subId2SubMap = map[string]*Sub{}

var subId2SubMapLock = &sync.Mutex{}

func LockSubFn(fn func()) {
	LockFn(subId2SubMapLock, fn)
}

func SubSave(sub *Sub) error {
	LockSubFn(func() {
		subId2SubMap[sub.Id] = sub
	})
	return nil
}

func SubGetById(id string) *Sub {
	var sub *Sub
	LockSubFn(func() {
		sub = subId2SubMap[id]
	})
	return sub
}
