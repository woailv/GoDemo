package main

import (
	"log"
	"reflect"
)

func main() {
	var list []int
	for i := 1; i <= 8; i++ {
		list = append(list, i)
	}
	ListStep(list, 3, func(s, e int) {
		for i := s; i < e; i++ {
			log.Println(list[i])
		}
		log.Println("==========")
	})
}

// 回调使用方式 for i := s; i < e; i++ {}
func ListStep(list interface{}, step int, f func(s, e int)) {
	s := 0
	e := step
	rv := reflect.ValueOf(list)
	for {
		if s >= rv.Len() {
			break
		}
		if e > rv.Len() {
			e = rv.Len()
		}
		f(s, e)
		s += step
		e += step
	}
}
