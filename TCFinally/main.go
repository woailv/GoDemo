package main

import (
	"errors"
	"log"
	"reflect"
)

func main() {
	TCFinally(func() {
		log.Println("1")
		panic(errors.New("haha"))
	},
		func(err error) {
			log.Println("error:", err)
		},
		func() {
			log.Println("finally")
		},
	)
}

func TCFinally(t func(), cf ...interface{}) {
	for i := range cf {
		f := cf[len(cf)-1-i]
		rv := reflect.ValueOf(f)
		switch rv.Type().NumIn() {
		case 0:
			defer func() {
				rv.Call([]reflect.Value{})
			}()
		case 1:
			defer func() {
				if err := recover(); err != nil {
					var e error
					switch x := err.(type) {
					case string:
						e = errors.New(x)
					case error:
						e = x
					}
					rv.Call([]reflect.Value{reflect.ValueOf(e)})
				}
			}()
		}
	}
	t()
}
