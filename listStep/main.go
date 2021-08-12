package main

import (
	"log"
	"reflect"
)

func main() {
	//listStepTest()
}

func listStepTest() {
	var list []int
	for i := 1; i <= 8; i++ {
		list = append(list, i)
	}
	ListStep(list, 3, func(s, e int) bool {
		for i := s; i < e; i++ {
			log.Println(list[i])
		}
		log.Println("==========")
		return false
	})
}

// 回调使用方式 for i := s; i < e; i++ {}
func ListStep(list interface{}, step int, f func(s, e int) bool) {
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
		if !f(s, e) {
			break
		}
		s += step
		e += step
	}
}
