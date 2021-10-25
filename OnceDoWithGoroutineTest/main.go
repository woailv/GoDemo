package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var l []int
	l = append(l, 1)
	fmt.Println(l == nil)
	//return
	w := sync.WaitGroup{}
	for i := 0; i < 10; i++ {
		w.Add(1)
		go func(i int) {
			defer w.Done()
			if i%3 == 0 {
				time.Sleep(time.Second * time.Duration(i))
			}
			f()
		}(i)
	}
	w.Wait()
	fmt.Println("end")
}

var once = sync.Once{}

//var m map[string]int
var l []int

func f() {
	once.Do(func() {
		if l == nil {
			fmt.Println("hahaha")
			l = []int{}
		}
		go func() {
			for i := 1; ; i++ {
				time.Sleep(time.Second * 1)
				l = append(l, i)
			}
		}()
	})
	fmt.Println(l)
}
