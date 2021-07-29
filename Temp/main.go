package main

import (
	"GoDemo/Err"
	"GoDemo/diskqueue"
	"log"
	"time"
)

func main() {
	//dqTest()
	log.Println("2006-01-02 15:04:05"[:10])
}

func dqTest() {
	var err error
	dq := diskqueue.New("d1", "", 1024, 1, 1024, 3, time.Second, func(lvl diskqueue.LogLevel, f string, args ...interface{}) {

	})
	//err = dq.Empty()
	Err.IfPanic(err)
	go func() {
		i := 0
		for {
			err := dq.Put([]byte(time.Now().Format(time.RFC3339)))
			Err.IfPanic(err)
			time.Sleep(time.Second)
			i++
			if i == 3 {
				break
			}
		}
	}()
	time.Sleep(time.Second)
	for data := range dq.ReadChan() {
		log.Printf("%s", data)
	}
	Err.IfPanic(err)
}
