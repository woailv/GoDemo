package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	m := sync.Map{}
	for i := 0; i < 1000; i++ {
		m.Store(i, i)
	}
	w := sync.WaitGroup{}
	w.Add(2)
	f := func() {
		m.Range(func(key, value interface{}) bool {
			fmt.Println(key)
			time.Sleep(time.Second * 3)
			return true
		})
		w.Done()
	}
	go f()
	go f()
	w.Wait()
	fmt.Println("end")
}
