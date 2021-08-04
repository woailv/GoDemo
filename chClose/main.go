package main

import "fmt"

func main() {
	ch := make(chan int)
	go func() {
		ch <- 2
		close(ch)
	}()
	a := <-ch
	b := <-ch
	b, ok := <-ch
	fmt.Println("a:", a)
	fmt.Println("b:", b)
	fmt.Println("ok:", ok)
	fmt.Println("end")
}
