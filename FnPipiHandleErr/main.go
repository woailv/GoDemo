package main

import (
	"GoDemo/Fn"
	"errors"
	"fmt"
)

func main() {
	//test1()
	test2()
}

func test2() {
	endFn := Fn.EndFn()
	fn := Fn.PipeWithEnd(endFn)
	fn = Fn.PipeWithResult(func() {
		fmt.Println("ok")
	}, func(err error) {
		fmt.Println("error:", err)
	}, fn)

	err := fn(
		func() error {
			fmt.Println(1)
			return nil
		},
		func() error {
			fmt.Println(2)
			//endFn(true)
			return nil
		},
		func() error {
			fmt.Println(3)
			return errors.New("error3")
		},
	)
	fmt.Println("end error:", err)
}

func test1() {
	var i = 0
	_ = Fn.Pipe(func() error {
		i++
		return nil
	}, func() error {
		i *= 2
		return nil
	}, Fn.WithErr(func() {
		fmt.Println(i)
	}), func() error {
		return errors.New("test error")
	}, Fn.WithErr(func() {
		fmt.Println("end")
	}))
}
