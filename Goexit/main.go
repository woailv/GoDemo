package main

import (
	"fmt"
	"runtime"
)

func main() {
	//f1()
	ch := make(chan struct{})
	go func() {
		defer func() {
			fmt.Println("goroutine start")
			ch <- struct{}{}
		}()
		f2()
		fmt.Println("goroutine end")
	}()
	<-ch
	fmt.Println("main end")
}

func f1() {
	runtime.Goexit()
}

func f2() {
	defer func() {
		fmt.Println("goroutine f2 start")
	}()
	runtime.Goexit()
	fmt.Println("goroutine f2 end")
}
