package dot

import (
	"log"
	"os"
	"sync"
)

var waitLog = log.New(os.Stderr, "wait ", log.LstdFlags|log.Lshortfile)

func waitFunc() (*sync.WaitGroup, func(f func(), tag ...string)) {
	wg := &sync.WaitGroup{}
	return wg, func(f func(), tag ...string) {
		wg.Add(1)
		go func() {
			defer wg.Done()
			if len(tag) > 0 {
				if tag[0] != "" {
					waitLog.Println(tag[0])
				}
			}
			f()
			if len(tag) > 1 {
				if tag[1] != "" {
					waitLog.Println(tag[1])
				}
			}
		}()
	}
}
