package main

import (
	"fmt"
)

func main() {
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
