package diskq

import (
	"GoDemo/Err"
	"fmt"
	"strconv"
	"testing"
	"time"
)

func TestDiskQueue_Close(t *testing.T) {
	var err error
	dqWrite := New("d1", "./", 15, 1, 1024, 1, time.Second, func(lvl LogLevel, f string, args ...interface{}) {
		//fmt.Println(lvl, fmt.Sprintf(f, args...))
	})
	dqRead := New("d1", "./", 15, 1, 1024, 1, time.Second, func(lvl LogLevel, f string, args ...interface{}) {
		//fmt.Println(lvl, fmt.Sprintf(f, args...))
	})

	//err = dq.Empty()
	Err.IfPanic(err)
	go func() {
		for i := 0; true; i++ {
			//err := dq.Put([]byte(time.Now().Format(time.RFC3339)))
			err := dqWrite.Put([]byte(strconv.Itoa(i)))
			Err.IfPanic(err)
			time.Sleep(time.Second / 3)
			if i == 9 {
				break
			}
		}
	}()
	time.Sleep(time.Second)
	var dataList [][]byte
	for data := range dqRead.ReadChan() {
		dataList = append(dataList, data)
		if len(dqWrite.ReadChan()) == 0 || len(dataList) == 2 {
			var is []string
			for _, data := range dataList {
				is = append(is, string(data))
			}
			fmt.Println("is:", is)
			dqWrite.HandForward(dataList)
			dataList = [][]byte{}
			Err.IfPanic(err)
		}
		time.Sleep(time.Second)
	}
	err = dqWrite.Close()
}
