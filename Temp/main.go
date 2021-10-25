package main

import (
	"fmt"
	"strings"
)

func main() {
	str := `[EnableVerboseS5]jcqwah8dnr all ping lost 
hv updated successfully 1.2728m ago (2021-10-25 08:09:13.265633 +00)`
	str = `[EnableVerboseS5]jcqwah8dnr all ping lost 
hv updated successfully 1.5840m ago (2021-10-25 07:19:58.227921 +00)
Network info:
0.0.0.0/0 -> 192.168.1.1 wlan0`
	fmt.Println(strings.Contains(str, "all ping lost"))
}

type T struct {
	Name string
}

func f(i interface{}) {
	if i == "123" {
		fmt.Println("string", i)
	}
	if i == 1 {
		fmt.Println("int", i)
	}
	if i == 1.1 {

	}
}
