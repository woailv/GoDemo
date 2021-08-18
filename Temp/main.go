package main

import (
	"fmt"
	"time"
)

func main() {
	ViewTime()
}

type T struct {
	Name string
	Age  int
}

func ViewTime() {
	tm := time.Unix(0, 1628834133330*1000000)
	fmt.Println(tm)
}
