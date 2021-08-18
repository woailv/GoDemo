package dot

import (
	"fmt"
	"testing"
	"time"
)

func Test_tcpProxy_Run(t *testing.T) {
	tp := NewServer(":8080", func(option *ServerOption) {
		option.ReadData = func(c *Client, data []byte) {
			c.log.Println(string(data))
		}
	})
	go func() {
		i := 0
		for {
			time.Sleep(time.Second * 3)
			tp.ClientMapRange(func(c *Client) {
				_ = c.Write([]byte(fmt.Sprintf("hi:%d", i)))
				i++
			})
		}
	}()
	err := tp.Run()
	if err != nil {
		panic(err)
	}
}
