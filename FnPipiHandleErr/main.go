package main

import (
	"errors"
	"fmt"
)

func main() {
	//test1()
	test2()
}

func test2() {
	endFn := EndFn()
	fn := FnPipeWithEnd(endFn)
	fn = FnPipeWithResult(func() {
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
	_ = FnPipe(func() error {
		i++
		return nil
	}, func() error {
		i *= 2
		return nil
	}, FnWithErr(func() {
		fmt.Println(i)
	}), func() error {
		return errors.New("test error")
	}, FnWithErr(func() {
		fmt.Println("end")
	}))
}

func FnWithErr(f func()) func() error {
	return func() error {
		f()
		return nil
	}
}

type FnPipeType = func(f ...func() error) error

func FnPipe(f ...func() error) error {
	for i := range f {
		if err := f[i](); err != nil {
			return err
		}
	}
	return nil
}

func EndFn() func(end ...bool) bool {
	flag := false
	return func(end ...bool) bool {
		if len(end) == 1 && end[0] == true {
			flag = true
		}
		return flag
	}
}

// FnPipeWithEnd 正常结束不返回错误
func FnPipeWithEnd(endFn func(end ...bool) bool) FnPipeType {
	return func(f ...func() error) error {
		for i := range f {
			if err := f[i](); err != nil {
				return err
			}
			if endFn() {
				return nil // 正常终止无错误
			}
		}
		return nil
	}
}

func FnPipeWithSuccess(success func(), pipe FnPipeType) FnPipeType {
	return func(f ...func() error) error {
		err := pipe(f...)
		if err == nil {
			success()
		}
		return err
	}
}

func FnPipeWithFailed(failed func(err error), pipe FnPipeType) FnPipeType {
	return func(f ...func() error) error {
		err := pipe(f...)
		if err != nil {
			failed(err)
		}
		return err
	}
}

func FnPipeWithResult(success func(), failed func(err error), pipe FnPipeType) FnPipeType {
	return func(f ...func() error) error {
		err := pipe(f...)
		if err != nil {
			failed(err)
		} else {
			success()
		}
		return err
	}
}
