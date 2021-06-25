package MessageMiddle

import (
	"fmt"
	"testing"
	"time"
)

func TestA(t *testing.T) {
	go SendMessage()
	for i := 0; ; i++ {
		ReceiverMessage("a", fmt.Sprintf("this is a message:%d", i))
		time.Sleep(time.Second * 3)
	}
}
