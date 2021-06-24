package main

import (
	"GoDemo/Echo"
	"time"
)

func main() {
	m := map[int]int{
		1: 1,
		2: 2,
		3: 3,
		4: 4,
		5: 5,
	}
	for i := 0; i < 5; i++ {
		if i == 3 {
			go func() {
				for true {
					m[3] = 3
					time.Sleep(time.Millisecond)
				}
			}()
		}
		go func(k int) {
			for true {
				n := m[k]
				Echo.Json(n)
				time.Sleep(time.Millisecond)
			}
		}(i)
	}
	select {}
}
