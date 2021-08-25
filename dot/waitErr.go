package dot

import "sync"

type WaitErr struct {
	once  *sync.Once
	errCh chan error
}

func NewWaitErr() *WaitErr {
	return &WaitErr{
		once:  &sync.Once{},
		errCh: make(chan error),
	}
}

func (we *WaitErr) Err(err error) {
	we.once.Do(func() {
		if err != nil {
		}
		we.errCh <- err
	})
}

func (we *WaitErr) Wait() error {
	return <-we.errCh
}
