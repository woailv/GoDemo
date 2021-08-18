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
	ClientMapRange(f func(c *Client))
	GetClientAddrList() []string
}

var _ Server = (*server)(nil)

type ServerOption struct {
	ReadData        func(c *Client, data []byte)
	OnClientOnline  func(c *Client)
	OnClientOffline func(c *Client)
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

func (srv *server) GetClientAddrList() (list []string) {
	srv.ClientMapRange(func(c *Client) {
		list = append(list, c.GetRemoteAddr())
	})
	return
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
	if srv.option.ReadData == nil {
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
	wg, f := WaitFunc()
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
		conn: conn,
		log:  log.New(os.Stderr, "clientConn ", log.LstdFlags|log.Lshortfile),
	}
	wg, f := WaitFunc()
	f(func() {
		c.ReadClientLoop(srv.option.ReadData)
	})
	srv.ClientMap.Store(c.conn.RemoteAddr(), c)
	srv.option.OnClientOnline(c)
	wg.Wait()
	srv.ClientMap.Delete(c.conn.RemoteAddr())
	srv.option.OnClientOffline(c)
	c.Exist()
	srv.log.Println("conn offline:", c.conn.RemoteAddr())
}
