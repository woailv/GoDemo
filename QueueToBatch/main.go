package main

import (
	"fmt"
	"math"
	"time"
)

// 从管道中接收数据的打包方案
func main() {
	ch := make(chan int, 100)
	go func() {
		for i := 0; i < math.MaxInt64; i++ {
			ch <- i
			//time.Sleep(time.Second / 2)
		}
	}()
	var is []int
	for i := range ch {
		is = append(is, i)
		if len(ch) == 0 || len(is) == 5 {
			fmt.Println(is)
			time.Sleep(time.Second)
			is = []int{}
		}
	}
}
