package DistributedCache

import (
	"log"
	"sync"
	"testing"
	"time"
)

func TestGroup_Do(t *testing.T) {
	g := Group{}
	n := 0
	wg := sync.WaitGroup{}
	for i := 0; i < 3; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			r, err := g.Do("a", func() (interface{}, error) {
				n++
				time.Sleep(time.Second * 2)
				return 123, nil
			})
			log.Println(r, err)
		}()
	}
	wg.Wait()
	log.Println(n)
}
