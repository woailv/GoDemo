package dot

import (
	"log"
	"net"
	"os"
	"sync"
)

type Server interface {
	Run() error
	Exist()
}

var _ Server = (*server)(nil)

type server struct {
	addr      string
	listener  net.Listener
	ClientMap sync.Map
	existFlag int32
	log       *log.Logger
}

func (srv *server) Exist() {
	_ = srv.listener.Close()
	// 清理客户端连接 不清理会一直存在(就算关闭了listener)
	srv.ClientMapRange(func(c *Client) {
		c.Exist()
	})
}

func (srv *server) ClientMapRange(f func(c *Client)) {
	srv.log.Println("client map range start")
	srv.ClientMap.Range(func(key, value interface{}) bool {
		f(value.(*Client))
		return true
	})
	srv.log.Println("client map range end")
}

func NewServer(proxyAddr string) *server {
	return &server{
		addr: proxyAddr,
		log:  log.New(os.Stderr, "tcpProxy ", log.LstdFlags|log.Lshortfile),
	}
}

func (srv *server) Run() error {
	listen, err := net.Listen("tcp", srv.addr)
	if err != nil {
		return err
	}
	srv.listener = listen
	wg, f := waitFunc()
	for {
		conn, err := listen.Accept()
		if err != nil {
			srv.log.Println(err)
			break
		}
		f(func() {
			srv.handleConn(conn)
		})
	}
	wg.Wait()
	srv.Exist()
	return nil // TODO handle error
}

func (srv *server) handleConn(conn net.Conn) {
	c := &Client{
		conn:       conn,
		remoteAddr: conn.RemoteAddr().String(),
		readCh:     make(chan []byte),
		writeCh:    make(chan []byte),
		errCh:      make(chan error, 1),
		log:        log.New(os.Stderr, "clientConn ", log.LstdFlags|log.Lshortfile),
	}
	wg, f := waitFunc()
	f(c.readLoop)
	f(c.ioLoop)
	c.Enter()
	srv.ClientMap.Store(c.remoteAddr, c)
	wg.Wait()
	srv.ClientMap.Delete(c.remoteAddr)
	c.Exist()
	srv.log.Println("conn offline:", c.remoteAddr)
}
