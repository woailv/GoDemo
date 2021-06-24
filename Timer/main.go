package main

import (
	"GoDemo/Echo"
	"time"
)

func main() {
	task()
}

var ch = make(chan struct{})

func task() {
	timer := time.NewTimer(time.Second * 2)
	<-timer.C
	i := 0
	for {
		timer.Reset(time.Second * 1)
		select {
		case <-timer.C:
			Echo.Json("<-timer.C")
		case <-ch:
			Echo.Json("<-ch")
			//default:
			//	Echo.Json("default")
		}
		Echo.Json(i)
		i++
		if i == 3 {
			go func() {
				ch <- struct{}{}
			}()
		}
		if i == 10 {
			break
		}
	}
}
