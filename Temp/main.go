package main

import (
	"fmt"
	"strconv"
	"time"
)

func main() {
	a, _ := strconv.ParseInt("1625192931284", 10, 64)
	fmt.Println(a)
	fmt.Println(time.Unix(0, a*1000000))
	//timePrint()
}

func timePrint() {
	s := []int64{
		1624992851,
	}
	fmt.Println(len(s[1:]))
	for _, i := range s {
		tm := time.Unix(i, 0)
		fmt.Println(tm)
	}
	fmt.Println(len("com.security.xvpn.z35kb.unblock.fullplan.monthly_def"))
}
