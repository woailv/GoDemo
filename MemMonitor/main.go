package main

import (
	"fmt"
	_ "net/http/pprof"
	"strings"
)

func main() {
	a:=strings.TrimPrefix("abc", "ab")
	fmt.Println(a)
}
