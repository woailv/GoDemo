package dot

import (
	"testing"
	"time"
)

func Test_tcpProxy_Run(t *testing.T) {
	tp := NewServer(":8080", func(option *ServerOption) {
		option.acceptData = func(c *Client, data []byte) {
			c.log.Println(string(data))
		}
	})
	go func() {
		for {
			time.Sleep(time.Second * 5)
			tp.ClientMapRange(func(c *Client) {
				tp.log.Println("conn addr:", c.conn.RemoteAddr())
				//c.Exist()
			})
		}
	}()
	err := tp.Run()
	if err != nil {
		panic(err)
	}
}
