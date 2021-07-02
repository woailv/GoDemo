package Fn

func WithErr(f func()) func() error {
	return func() error {
		f()
		return nil
	}
}

type PipeType = func(f ...func() error) error

func Pipe(f ...func() error) error {
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

// PipeWithEnd 正常结束不返回错误
func PipeWithEnd(endFn func(end ...bool) bool) PipeType {
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

func PipeWithSuccess(success func(), pipe PipeType) PipeType {
	return func(f ...func() error) error {
		err := pipe(f...)
		if err == nil {
			success()
		}
		return err
	}
}

func PipeWithFailed(failed func(err error), pipe PipeType) PipeType {
	return func(f ...func() error) error {
		err := pipe(f...)
		if err != nil {
			failed(err)
		}
		return err
	}
}

func PipeWithResult(success func(), failed func(err error), pipe PipeType) PipeType {
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
