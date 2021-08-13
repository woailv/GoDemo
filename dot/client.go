package dot

import (
	"log"
	"net"
	"os"
	"sync"
	"sync/atomic"
)

type Client struct {
	conn       net.Conn
	writeMu    sync.Mutex
	existFlat  int32
	log        *log.Logger
	acceptData func(c *Client, data []byte)
}

func (c *Client) Exist() {
	if !atomic.CompareAndSwapInt32(&c.existFlat, 0, 1) {
		return
	}
	c.log.Println("conn exist:", c.conn.RemoteAddr())
	_ = c.conn.Close()
}

func (c *Client) ReadLoop() {
	for {
		data := make([]byte, 1024)
		n, err := c.conn.Read(data)
		if err != nil {
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
	return err
}

func Dial(addr string, acceptData func(c *Client, data []byte)) (*Client, error) {
	conn, err := net.Dial("tcp", addr)
	if err != nil {
		return nil, err
	}
	c := &Client{
		conn:       conn,
		log:        log.New(os.Stderr, "clientConn ", log.LstdFlags|log.Lshortfile),
		acceptData: acceptData,
	}
	return c, nil
}
