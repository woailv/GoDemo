package tcpProxy

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
	connMap   sync.Map
	existFlag int32
	log       *log.Logger
	errCh     chan error
}

func (tp *server) Exist() {
	tp.listener.Close()
	// 清理客户端连接 不清理会一直存在(就算关闭了listener)
	tp.ItemProxyConnMap(func(c *Conn) {
		c.Exist()
	})
}

func (tp *server) ItemProxyConnMap(f func(c *Conn)) {
	tp.connMap.Range(func(key, value interface{}) bool {
		f(value.(*Conn))
		return true
	})
}

func NewTCPProxy(proxyAddr string) *server {
	return &server{
		addr:  proxyAddr,
		log:   log.New(os.Stderr, "tcpProxy ", log.LstdFlags|log.Lshortfile),
		errCh: make(chan error),
	}
}

func (tp *server) Run() error {
	wg, f := waitFunc()
	var err error
	f(func() {
		tp.runProxy()
	})
	err = <-tp.errCh
	if err != nil {
		return err
	}
	wg.Wait()
	tp.log.Println("end")
	return nil // TODO 错误处理
}

func (tp *server) runProxy() {
	listen, err := net.Listen("tcp", tp.addr)
	if err != nil {
		tp.errCh <- err
		return
	}
	tp.listener = listen
	tp.errCh <- nil
	tp.run(listen, tp.handleProxyConn)
}

func (tp *server) run(listen net.Listener, handle func(conn net.Conn)) {
	wg, f := waitFunc()
	for {
		conn, err := listen.Accept()
		if err != nil {
			tp.log.Println(err)
			break
		}
		f(func() {
			handle(conn)
		})
	}
	wg.Wait()
	tp.Exist()
}

func (tp *server) handleProxyConn(conn net.Conn) {
	c := &Conn{
		tp:         tp,
		conn:       conn,
		remoteAddr: conn.RemoteAddr().String(),
		readCh:     make(chan []byte),
		writeCh:    make(chan []byte),
		errCh:      make(chan error, 1),
		log:        log.New(os.Stderr, "clientConn ", log.LstdFlags|log.Lshortfile),
	}
	tp.connMap.Store(c.remoteAddr, c)
	wg, f := waitFunc()
	f(c.readLoop)
	f(c.ioLoop)
	wg.Wait()
	c.Exist()
	tp.log.Println("conn offline:", c.remoteAddr)
}
