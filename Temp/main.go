package main

import (
	"fmt"
	"sort"
)

func main() {
	xs := []int{1, 2, 3}
	sort.Slice(xs, func(i, j int) bool {
		return xs[i] > xs[j]
	})
	fmt.Println(xs)
}
