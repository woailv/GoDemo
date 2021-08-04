package flowerGUID

import (
	"log"
	"sync"
	"testing"
)

func Test_guidFactory_NewGUID(t *testing.T) {
	f := NewGUIDFactory(1)
	m := map[string]struct{}{}
	wg := sync.WaitGroup{}
	for i := 0; i < 10000; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			uid, err := f.NewGUID()
			if err != nil {
				panic(err)
			}
			h := uid.Hex()
			id := string(h[:])
			m[id] = struct{}{}
		}()
	}
	wg.Wait()
	log.Println(len(m))
}
