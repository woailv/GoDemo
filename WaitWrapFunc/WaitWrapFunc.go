package main

import (
	"fmt"
	"log"
	"sync"
	"time"
)

func main() {
	wg, f := WaitWrapFunc()
	defer wg.Wait()
	f(func() {
		log.Println("1")
		time.Sleep(time.Second)
		w1, f1 := WaitWrapFunc()
		f1(func() {
			log.Println("i 1")
			time.Sleep(time.Second * 8)
		}, "i1 start", "i1 end")
		w1.Wait()
	}, "1start", "1end")
	f(func() {
		log.Println("2")
		time.Sleep(time.Second * 3)
	}, "2start", "2end")
	//wg.Wait()
	fmt.Println("end")
}

// tag.len == 1 => start; tag.len ==2 => tag[0] = start, tag[1] = end
func WaitWrapFunc() (*sync.WaitGroup, func(f func(), tag ...string)) {
	wg := &sync.WaitGroup{}
	return wg, func(f func(), tag ...string) {
		wg.Add(1)
		go func() {
			defer wg.Done()
			if len(tag) > 0 {
				if tag[0] != "" {
					log.Println(tag[0])
				}
			}
			f()
			if len(tag) > 1 {
				if tag[1] != "" {
					log.Println(tag[1])
				}
			}
		}()
	}
}
