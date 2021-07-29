package main

import (
	"fmt"
	"log"
	"sync"
)

func main() {
	wg := sync.WaitGroup{}
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			defer func() {
				if err := recover(); err != nil {
					log.Println("recover:", err)
				}
			}()
			do(i)
		}(i)
	}
	wg.Wait()
	fmt.Println("end")
}

func do(i int) {
	if i == 3 {
		panic(i)
	}
}
