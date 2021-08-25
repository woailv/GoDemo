package dot

import (
	"errors"
	"log"
	"testing"
	"time"
)

func TestWaitErr_Wait(t *testing.T) {
	we := NewWaitErr()
	go func() {
		we.Err(e(errors.New("1"), time.Second*3))
	}()
	go func() {
		we.Err(e(errors.New("2"), time.Second*2))
	}()
	e := we.Wait()
	log.Println(e)
}

func e(err error, duration time.Duration) error {
	time.Sleep(duration)
	return err
}
