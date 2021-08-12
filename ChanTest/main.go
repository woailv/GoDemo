package main

import "log"

// 向已关闭的chan发送数据不会报异常会一直阻塞到那里 重复关闭chan会报异常
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
