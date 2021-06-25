package main

import "sync"

func LockFn(mu *sync.Mutex, fn func()) {
	mu.Lock()
	fn()
	mu.Unlock()
}
