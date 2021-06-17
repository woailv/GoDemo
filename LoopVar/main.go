package main

import (
	alloc "GoDemo/AllocTrace"
	"time"
)

func main() {
	alloc.Begin()
	for i := 0; ; i++ {
		j := i
		j++
		if i%100000==0{
			time.Sleep(1)
		}
	}
}
