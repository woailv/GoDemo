package main

import (
	"fmt"
	"testing"
	"time"
)

func TestA(t *testing.T) {
	go SendMessage()
	for i := 0; ; i++ {
		ReceiverMessage("1", fmt.Sprintf("this is a message:%d", i))
		time.Sleep(time.Second * 3)
	}
}
