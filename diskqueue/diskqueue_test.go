package diskqueue

import (
	"GoDemo/Err"
	"log"
	"strconv"
	"testing"
	"time"
)

func TestDiskQueue_Close(t *testing.T) {
	var err error
	dq := New("d1", "./", 1024, 1, 1024, 3, time.Second, func(lvl LogLevel, f string, args ...interface{}) {

	})
	//err = dq.Empty()
	Err.IfPanic(err)
	go func() {
		i := 0
		for false {
			//err := dq.Put([]byte(time.Now().Format(time.RFC3339)))
			err := dq.Put([]byte(strconv.Itoa(i)))
			Err.IfPanic(err)
			time.Sleep(time.Second)
			i++
			if i == 3 {
				break
			}
		}
	}()
	time.Sleep(time.Second)
	n := 0
	for data := range dq.ReadChan() {
		n++
		log.Printf("%s", data)
		if n == 2 {
			err := dq.PersistMetaReadData()
			if err != nil {
				panic(err)
			}
		}
	}
	Err.IfPanic(err)
	time.Sleep(time.Second * 3)
}
