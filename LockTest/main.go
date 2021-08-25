package main

import (
	"GoDemo/UseTime"
	"fmt"
	"log"
	"sync"
	"time"
)

var l = sync.RWMutex{}

var m = map[int]int{}

func main() {
	test2()
}

func test1() {
	UseTime.View(func() {
		w := sync.WaitGroup{}
		for i := 0; i < 100000; i++ {
			w.Add(1)
			go func() {
				defer w.Done()
				for k := 0; k < 50; k++ {
					f(k)
				}
			}()
		}
		w.Wait()
		fmt.Println("rw lock")
	})
	UseTime.View(func() {
		w := sync.WaitGroup{}
		for i := 0; i < 100000; i++ {
			w.Add(1)
			go func() {
				defer w.Done()
				for k := 0; k < 50; k++ {
					f2(k)
				}
			}()
		}
		w.Wait()
		fmt.Println("lock")
	})
}

func f(k int) {
	l.RLock()
	_, ok := m[k]
	if ok {
		l.RUnlock()
		return
	}
	l.RUnlock()

	l.Lock()
	m[k] = k
	l.Unlock()
	return
}
func f2(k int) {
	l.Lock()
	_, ok := m[k]
	if ok {
		l.Unlock()
		return
	}
	m[k] = k
	l.Unlock()
	return
}

// 在写的时候拿不到读锁
func test2() {
	w := sync.WaitGroup{}
	w.Add(2)
	go func() {
		time.Sleep(time.Second * 1)
		f3_1()
		w.Done()
	}()
	go func() {
		f3_2()
		w.Done()
	}()
	w.Wait()
}

func f3_1() {
	log.Println("f3_1_1")
	l.RLock()
	log.Println("f3_1_2")
	time.Sleep(time.Second * 2)
	l.RUnlock()
	return
}

func f3_2() {
	log.Println("f3_2_1")
	l.Lock()
	time.Sleep(time.Second * 5)
	log.Println("f3_2_2")
	l.Unlock()
	return
}
