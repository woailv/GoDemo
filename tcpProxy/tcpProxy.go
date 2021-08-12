package tcpProxy

import (
	"fmt"
	"log"
	"net"
	"os"
	"sync"
)

type TCPProxy interface {
	Run() error
	Exist()
}

var _ TCPProxy = (*tcpProxy)(nil)

type tcpProxy struct {
	proxyAddr      string
	registerAddr   string
	proxyListen    net.Listener
	proxyConnMap   sync.Map
	registerListen net.Listener
	existFlag      int32
	log            *log.Logger
	errCh          chan error
}

func (tp *tcpProxy) Exist() {
	tp.proxyListen.Close()
	tp.registerListen.Close()
	// 清理客户端连接 不清理会一直存在(就算关闭了listener)
	tp.ItemProxyConnMap(func(c *Conn) {
		c.Exist()
	})
}

func (tp *tcpProxy) ItemProxyConnMap(f func(c *Conn)) {
	tp.proxyConnMap.Range(func(key, value interface{}) bool {
		f(value.(*Conn))
		return true
	})
}

func NewTCPProxy(proxyAddr, registerAddr string) *tcpProxy {
	return &tcpProxy{
		proxyAddr:    proxyAddr,
		registerAddr: registerAddr,
		log:          log.New(os.Stderr, "tcpProxy ", log.LstdFlags|log.Lshortfile),
		errCh:        make(chan error),
	}
}

func (tp *tcpProxy) Run() error {
	wg, f := waitFunc()
	var err error
	f(func() {
		tp.runProxy()
	})
	err = <-tp.errCh
	if err != nil {
		return err
	}
	f(func() {
		tp.runRegister()
	})
	err = <-tp.errCh
	if err != nil {
		return err
	}
	wg.Wait()
	tp.log.Println("end")
	return nil // TODO 错误处理
}

func (tp *tcpProxy) runProxy() {
	listen, err := net.Listen("tcp", tp.proxyAddr)
	if err != nil {
		tp.errCh <- err
		return
	}
	tp.proxyListen = listen
	tp.errCh <- nil
	tp.run(listen, tp.handleProxyConn)
}

func (tp *tcpProxy) run(listen net.Listener, handle func(conn net.Conn)) {
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

func (tp *tcpProxy) handleProxyConn(conn net.Conn) {
	c := &Conn{
		tp:         tp,
		conn:       conn,
		remoteAddr: conn.RemoteAddr().String(),
		readCh:     make(chan []byte),
		writeCh:    make(chan []byte),
		errCh:      make(chan error, 1),
		log:        log.New(os.Stderr, "clientConn ", log.LstdFlags|log.Lshortfile),
	}
	tp.proxyConnMap.Store(c.remoteAddr, c)
	wg, f := waitFunc()
	f(c.readLoop)
	f(c.ioLoop)
	wg.Wait()
	c.Exist()
	tp.log.Println("conn offline:", c.remoteAddr)
}

func (tp *tcpProxy) runRegister() {
	listen, err := net.Listen("tcp", tp.registerAddr)
	if err != nil {
		tp.errCh <- err
		return
	}
	tp.registerListen = listen
	tp.errCh <- nil
	tp.run(listen, tp.handleRegisterConn)
}

func (tp *tcpProxy) handleRegisterConn(conn net.Conn) {
	fmt.Println("bibi")
}
