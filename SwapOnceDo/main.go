package main

import (
	"fmt"
	"log"
	"sync"
	"sync/atomic"
	"time"
)

func main() {
	wg := sync.WaitGroup{}
	for i := 0; i < 1; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			f1()
			log.Printf("i:%d,r:%d", i, r)
		}(i)
	}
	go func() {
		for {
			time.Sleep(time.Second)
		}
	}()
	wg.Wait()
	fmt.Println("end")
}

var a int32
var r int

func f1() {
	if !atomic.CompareAndSwapInt32(&a, 0, 1) {
		return
	}
	f1()
}
func b() {
	time.Sleep(time.Second * 2)
	r = 100
}

// 递归调用有阻塞风险
var once sync.Once

func f2() {
	log.Println("haha")
	once.Do(func() {
		f2()
	})
}
