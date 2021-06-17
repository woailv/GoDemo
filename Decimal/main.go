package main

import (
	"fmt"
	"strconv"
)

func main() {
	var a float64 = 3.99
	a2 := a * 10 / 100
	fmt.Println(a2)
	fmt.Println(Decimal2(a2))
	fmt.Println(DollarToCent(a2))
}

func Decimal2(val float64) float64 {
	s := strconv.FormatFloat(val, 'f', 2, 64)
	r, e := strconv.ParseFloat(s, 64)
	if e != nil {
		panic(e)
	}
	return r
}

func DollarToCent(a float64) int {
	a = a * 100
	s := strconv.FormatFloat(a, 'f', 0, 64)
	r, _ := strconv.Atoi(s)
	return r
}
