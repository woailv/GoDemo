package dot

import (
	"fmt"
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
		for i := 0; ; i++ {
			time.Sleep(time.Second * 5)
			tp.ClientMapRange(func(c *Client) {
				tp.log.Println("conn addr:", c.conn.RemoteAddr())
				err := c.Write([]byte(fmt.Sprintf("hello:%d", i)))
				if err != nil {
					c.Exist()
				}
				c.Exist()
			})
		}
	}()
	err := tp.Run()
	if err != nil {
		panic(err)
	}
}
