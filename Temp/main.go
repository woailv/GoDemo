package main

import (
	"fmt"
	"time"
)

func main() {
	//timePrint()
	fmt.Println(byte(1))
}

func timePrint() {
	s := []int64{
		1623806954,
	}
	fmt.Println(len(s[1:]))
	for _, i := range s {
		tm := time.Unix(i, 0)
		fmt.Println(tm)
	}

	fmt.Println(len("com.security.xvpn.z35kb.unblock.fullplan.monthly_def"))
}
