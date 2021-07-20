package main

import "log"

func main() {
	ch := make(chan int, 100)
	ch <- 1
	ch <- 2
	ch <- 3
	for i := range ch {
		log.Println(i, len(ch))
		if len(ch) == 0 {
			break
		}
	}
}
