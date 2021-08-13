package main

import (
	"fmt"
	"time"
)

func main() {
	ViewTime()
}

func ViewTime() {
	tm := time.Unix(0, 1628834133330*1000000)
	fmt.Println(tm)
}
