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

type ServerOption struct {
	acceptData func(c *Client, data []byte)
}

type server struct {
	addr      string
	listener  net.Listener
	ClientMap sync.Map
	log       *log.Logger
	option    *ServerOption
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

func NewServer(proxyAddr string, option ...func(option *ServerOption)) *server {
	srv := &server{
		addr:   proxyAddr,
		log:    log.New(os.Stderr, "tcpProxy ", log.LstdFlags|log.Lshortfile),
		option: &ServerOption{},
	}
	for _, f := range option {
		f(srv.option)
	}
	if srv.option.acceptData == nil {
		panic("srv.option.acceptData can't be nil")
	}
	return srv
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
		acceptData: srv.option.acceptData,
	}
	wg, f := waitFunc()
	f(c.readLoop)
	f(c.ioLoop)
	srv.ClientMap.Store(c.remoteAddr, c)
	wg.Wait()
	srv.ClientMap.Delete(c.remoteAddr)
	c.Exist()
	srv.log.Println("conn offline:", c.remoteAddr)
}
