package dot

import (
	"log"
	"net"
	"sync"
	"sync/atomic"
	"time"
)

type Client struct {
	conn       net.Conn
	remoteAddr string
	writeMu    sync.Mutex
	errCh      chan error
	existFlat  int32
	log        *log.Logger
	acceptData func(c *Client, data []byte)
}

func (c *Client) Exist() {
	if !atomic.CompareAndSwapInt32(&c.existFlat, 0, 1) {
		return
	}
	c.log.Println("conn exist:", c.remoteAddr)
	_ = c.conn.Close()
}

func (c *Client) readLoop() {
	for {
		data := make([]byte, 1024)
		n, err := c.conn.Read(data)
		if err != nil {
			c.errCh <- err
			c.log.Println("read loop end")
			return
		}
		c.acceptData(c, data[:n])
	}
}

func (c *Client) Write(data []byte) error {
	c.writeMu.Lock()
	defer c.writeMu.Unlock()
	_, err := c.conn.Write(data)
	if err != nil {
		c.errCh <- err
	}
	return err
}

func (c *Client) heartBeatLoop() {
	tk := time.NewTicker(time.Second * 3)
	for {
		select {
		case <-c.errCh:
			goto end
		case <-tk.C:
			if err := c.Write([]byte(".")); err != nil {
				goto end
			}
		}
	}
end:
	tk.Stop()
	c.log.Println("hear beat end")
}
