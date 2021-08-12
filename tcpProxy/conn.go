package tcpProxy

import (
	"fmt"
	"log"
	"net"
	"sync"
	"sync/atomic"
	"time"
)

type Conn struct {
	conn      net.Conn
	tp        *tcpProxy
	readCh    chan []byte
	writeCh   chan []byte
	writeMu   sync.Mutex
	errCh     chan error
	existFlat int32
	log       *log.Logger
}

func (c *Conn) Exist() {
	if !atomic.CompareAndSwapInt32(&c.existFlat, 0, 1) {
		return
	}
	c.log.Println("conn exist:", c.conn.RemoteAddr().String())
	c.conn.Close()
}

func (c *Conn) readLoop() {
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

func (c *Conn) write(data []byte) {
	c.writeMu.Lock()
	defer c.writeMu.Unlock()
	_, err := c.conn.Write(data)
	if err != nil {
		c.errCh <- err
	}
}

func (c *Conn) ioLoop() {
	tk := time.NewTicker(time.Second * 1)
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
	c.log.Println("end")
}
