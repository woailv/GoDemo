package main

import (
	"errors"
	"log"
	"reflect"
)

func main() {
	TCFinally(
		func() {
			log.Println("1")
			panic(errors.New("haha"))
		},
		//func(err error) {
		//	log.Println("error:", err)
		//},
		func() {
			log.Println("finally")
		},
	)
}

/*
执行函数
panic: haha

执行函数+catch函数
1
error: haha

执行函数+catch函数+finally函数
1
error: haha
finally

执行函数+finally函数
1
finally
panic: haha

https://www.cnblogs.com/pcheng/p/10968841.html
在try catch中慎用返回值 有点绕
*/

// 也可以使用对象实现将函数绑定到对象上
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
