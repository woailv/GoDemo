package dot

import (
	"testing"
	"time"
)

func Test_tcpProxy_Run(t *testing.T) {
	tp := NewTCPProxy(":8080")
	go func() {
		for {
			time.Sleep(time.Second * 5)
			tp.Exist()
			tp.ItemProxyConnMap(func(c *Client) {
				tp.log.Println("conn addr:", c.conn.RemoteAddr())
			})
		}
	}()
	err := tp.Run()
	if err != nil {
		panic(err)
	}
}
