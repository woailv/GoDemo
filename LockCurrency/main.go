package main

import (
	"GoDemo/UseTime"
	"fmt"
	"sync"
	"time"
)

var flag bool
var flagMu sync.Mutex

func GetFlag() bool {
	flagMu.Lock()
	defer flagMu.Unlock()
	return flag
}

func main() {
	n := 100000000
	fmt.Println("调用次数:", n)
	UseTime.View(func() {
		wg := sync.WaitGroup{}
		for i := 0; i < n; i++ {
			wg.Add(1)
			go func() {
				_ = GetFlag()
				wg.Done()
			}()
		}
		wg.Wait()
	})
	time.Sleep(time.Second)
	UseTime.View(func() {
		wg := sync.WaitGroup{}
		for i := 0; i < n; i++ {
			wg.Add(1)
			go func() {
				_ = flag
				wg.Done()
			}()
		}
		wg.Wait()
	})
}

/*
调用次数: 10000000
lock use time: 2.832996103s
use time: 3.2379206s

调用次数: 100000000
lock use time: 30.592190783s
use time: 44.803630706s
*/
