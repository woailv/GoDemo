package main

import (
	"fmt"
	"time"
)

func main() {
	ticker := time.NewTicker(time.Second * 2)
	is := []int{}
	select {
	case <-ticker.C:

	default:
		for i := 0; i < 100; i++ {
			is = append(is, i)
			select {
			case <-ticker.C:
				fmt.Println("is t", is)
				is = []int{}
			default:
			}
			if len(is) == 10 {
				fmt.Println("is i", is)
				is = []int{}
				time.Sleep(time.Second)
			}
		}
	}
	fmt.Println("is i", is)
}
