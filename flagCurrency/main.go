package main

import (
	"log"
	"sync/atomic"
)

func main() {
	var a int32
	ok := atomic.CompareAndSwapInt32(&a, 0, 1)
	log.Println(ok)
	log.Println(a)
	ok = atomic.CompareAndSwapInt32(&a, 0, 1)
	log.Println(ok)
}
