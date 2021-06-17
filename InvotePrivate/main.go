package main

import (
	"GoDemo/InvotePrivate/b"
	"fmt"
)

func main() {
	s := b.Greet("world")
	fmt.Println(s)
	s = b.Hi("world")
	fmt.Println(s)
}
