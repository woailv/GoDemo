package main

import "sync"

type Message struct {
	Id      string
	Time    int64
	Content string // 方便搜索
}

var id2MessageMap = map[string]*Message{}

var id2MessageMapLock = &sync.Mutex{}

func MessageSave(message *Message) error {
	LockMessageFn(func() {
		id2MessageMap[message.Id] = message
	})
	return nil
}

func LockMessageFn(fn func()) {
	LockFn(id2MessageMapLock, fn)
}

func MessageGetById(id string) *Message {
	var message *Message
	LockMessageFn(func() {
		message = id2MessageMap[id]
	})
	return message
}
