package MessageMiddle

import (
	"fmt"
	"testing"
	"time"
)

func InitData() {

}

func TestA(t *testing.T) {
	InitData()
	go SendMessage()
	for i := 0; ; i++ {
		ReceiverMessage("a", fmt.Sprintf("this is a message:%d", i))
		time.Sleep(time.Second * 3)
	}
}
