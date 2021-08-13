package dot

import (
	"testing"
	"time"
)

func Test_tcpProxy_Run(t *testing.T) {
	tp := NewServer(":8080")
	go func() {
		for {
			time.Sleep(time.Second * 5)
			tp.ClientMapRange(func(c *Client) {
				tp.log.Println("conn addr:", c.conn.RemoteAddr())
				c.Exist()
			})
		}
	}()
	err := tp.Run()
	if err != nil {
		panic(err)
	}
}
