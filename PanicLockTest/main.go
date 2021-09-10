package main

import (
	"fmt"
	"sync"
)

func main() {
	wg := sync.WaitGroup{}
	for i := 0; i < 5; i++ {
		n := i
		wg.Add(1)
		go func() {
			defer func() {
				if err := recover(); err != nil {
					fmt.Println("recover:", err)
				}
			}()
			f(n)
		}()
	}
	wg.Wait()
	fmt.Println("end:", "main")
}

var mu = sync.Mutex{}

func f(i int) {
	mu.Lock()
	// 产生panic后此锁🔐不会被释放
	defer mu.Unlock()
	if i == 3 {
		panic(i)
	}
	fmt.Println("end:", i)
}
