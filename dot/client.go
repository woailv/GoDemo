package dot

import (
	"fmt"
	"log"
	"net"
	"sync"
	"sync/atomic"
	"time"
)

type Client struct {
	conn       net.Conn
	remoteAddr string
	readCh     chan []byte
	writeCh    chan []byte
	writeMu    sync.Mutex
	errCh      chan error
	existFlat  int32
	log        *log.Logger
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
			err := fmt.Errorf("read error:%s", err)
			c.errCh <- err
			return
		}
		c.log.Println("read data:", string(data[:n]))
	}
}

func (c *Client) write(data []byte) {
	c.writeMu.Lock()
	defer c.writeMu.Unlock()
	_, err := c.conn.Write(data)
	if err != nil {
		c.errCh <- err
	}
}

func (c *Client) ioLoop() {
	tk := time.NewTicker(time.Second * 3)
	for {
		select {
		case err := <-c.errCh:
			c.log.Println("error:", err)
			goto end
		case data := <-c.writeCh:
			c.log.Println("write ch data:", string(data))
			c.write(data)
		case <-tk.C:
			c.write([]byte("."))
		}
	}
end:
	tk.Stop()
	c.Exist()
	c.log.Println("io loop end")
}
