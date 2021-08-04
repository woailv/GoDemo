package main

import (
	"fmt"
	"reflect"
	"time"
)

func main() {
	ch := make(chan int, 1)
	go func() {
		a := <-ch
		fmt.Println(a)
	}()
	select {
	case ch <- 1:
	default:
		fmt.Println("def")
	}
	time.Sleep(time.Second)
	fmt.Println("end")
}

func TryCatchFinally(t func(), cf ...interface{}) {
	if len(cf) > 0 {
		defer func() {
			if err := recover(); err != nil {
				rv := reflect.ValueOf(cf[0])
				rv.Call([]reflect.Value{reflect.ValueOf(err)})
			}
		}()
	}
	if len(cf) > 1 {
		defer func() {
			rv := reflect.ValueOf(cf[1])
			rv.Call([]reflect.Value{})
		}()
	}
	t()
}
