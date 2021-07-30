package diskq

import (
	"GoDemo/Err"
	"fmt"
	"log"
	"strconv"
	"testing"
	"time"
)

func TestDiskQueue_Close(t *testing.T) {
	var err error
	dq := New("d1", "./", 15, 1, 1024, 1, time.Second, func(lvl LogLevel, f string, args ...interface{}) {
		//fmt.Println(lvl, fmt.Sprintf(f, args...))
	})
	start := time.Now()
	//err = dq.Empty()
	Err.IfPanic(err)
	go func() {
		for i := 0; false; i++ {
			//err := dq.Put([]byte(time.Now().Format(time.RFC3339)))
			err := dq.Put([]byte(strconv.Itoa(i)))
			Err.IfPanic(err)
			time.Sleep(time.Second / 3)
			if i == 9 {
				break
			}
		}
	}()
	time.Sleep(time.Second)
	var dataList [][]byte
	for data := range dq.ReadChan() {
		dataList = append(dataList, data)
		if len(dq.ReadChan()) == 0 || len(dataList) == 2 {
			fmt.Println("dataList len:", len(dataList))
			var is []string
			for _, data := range dataList {
				is = append(is, string(data))
			}
			fmt.Println("is:", is)
			dq.HandForward(dataList)
			dataList = [][]byte{}
			Err.IfPanic(err)
		}
		time.Sleep(time.Second)
	}
	err = dq.Close()
	Err.IfPanic(err)
	tm := start.Sub(time.Now())
	log.Println(tm)
}
