package main

import (
	"fmt"
	"strconv"
)

func main() {
	var amount  float64 = 100.125
	val := strconv.FormatFloat(amount, 'f', 5, 64)
	fmt.Println(val)
	val = fmt.Sprintf("%.2f", amount)
	fmt.Println(val)
}
