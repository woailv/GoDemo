package main

import (
	"fmt"
	"sync"
)

type T struct {
	Name string
}

func main() {
	var once sync.Once
	var t *T
	var f = func() *T {
		once.Do(func() {
			t = &T{
				Name: "haha",
			}
		})
		return t
	}
	t0 := f()
	fmt.Println(t0)
	t1 := f()
	fmt.Println(t1)
}
