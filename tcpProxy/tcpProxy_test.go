package tcpProxy

import (
	"testing"
	"time"
)

func Test_tcpProxy_Run(t *testing.T) {
	tp := NewTCPProxy(":8080")
	go func() {
		for {
			time.Sleep(time.Second * 5)
			tp.ItemProxyConnMap(func(c *Conn) {
				c.Exist()
				tp.log.Println("conn addr:", c.conn.RemoteAddr())
			})
		}
	}()
	err := tp.Run()
	if err != nil {
		panic(err)
	}
}
