package dot

import (
	"encoding/binary"
	"fmt"
	"io"
	"log"
	"net"
	"os"
	"sync"
	"sync/atomic"
)

type Client struct {
	conn      net.Conn
	writeMu   sync.Mutex
	existFlat int32
	log       *log.Logger
	readData  func(reader io.Reader) ([]byte, error)
}

func (c *Client) GetRemoteAddr() string {
	return c.conn.RemoteAddr().String()
}

func (c *Client) Exist() {
	if !atomic.CompareAndSwapInt32(&c.existFlat, 0, 1) {
		return
	}
	c.log.Println("conn exist:", c.conn.RemoteAddr())
	_ = c.conn.Close()
}

func (c *Client) ReadClientLoop(acceptData func(c *Client, data []byte)) {
	rd := c.readData
	if rd == nil {
		rd = ReadData
	}
	for {
		data, err := rd(c.conn)
		if err != nil {
			c.log.Println("read loop end")
			return
		}
		acceptData(c, data)
	}
}

func (c *Client) ReadServerLoop(acceptData func(data []byte)) {
	rd := c.readData
	if rd == nil {
		rd = ReadData
	}
	for {
		data, err := rd(c.conn)
		if err != nil {
			c.log.Println("read loop end")
			return
		}
		acceptData(data)
	}
}

func ReadData(reader io.Reader) ([]byte, error) {
	var msgSize int32
	err := binary.Read(reader, binary.BigEndian, &msgSize)
	if err != nil {
		return nil, err
	}
	if msgSize < 0 {
		return nil, fmt.Errorf("response msg size is negative: %v", msgSize)
	}
	buf := make([]byte, msgSize)
	_, err = io.ReadFull(reader, buf)
	if err != nil {
		return nil, err
	}
	return buf, nil
}

func (c *Client) Write(data []byte) error {
	c.writeMu.Lock()
	defer c.writeMu.Unlock()
	err := binary.Write(c.conn, binary.BigEndian, int32(len(data)))
	if err != nil {
		return err
	}
	_, err = c.conn.Write(data)
	if err != nil {
		return err
	}
	return err
}

func Dial(addr string) (*Client, error) {
	conn, err := net.Dial("tcp", addr)
	if err != nil {
		return nil, err
	}
	c := &Client{
		conn: conn,
		log:  log.New(os.Stderr, "clientConn ", log.LstdFlags|log.Lshortfile),
	}
	return c, nil
}
