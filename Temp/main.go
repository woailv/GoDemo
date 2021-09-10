package main

import "fmt"

func main() {
	f("123")
	f(1)
	f(1.1)
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
