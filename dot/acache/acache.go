package acache

import (
	"GoDemo/dot"
	"encoding/json"
	"log"
	"net/http"
	"os"
	"sync"
	"time"
)

type aCache struct {
	httpAddr        string
	server          dot.Server
	allAddr         []string
	addrToClientMap sync.Map
	msgQueue        []string
	log             *log.Logger
}

func NewACache(httpAddr, tcpAddr string, allAddr []string) *aCache {
	ac := &aCache{
		httpAddr: httpAddr,
		allAddr:  allAddr,
		log:      log.New(os.Stderr, "allCache ", log.LstdFlags|log.Lshortfile),
	}
	server := dot.NewServer(tcpAddr, func(option *dot.ServerOption) {
		option.ReadData = func(c *dot.Client, data []byte) {
			ac.log.Println("read client data:", string(data))
		}
		option.OnClientOnline = func(c *dot.Client) {
			ac.addrToClientMap.Store(c.GetRemoteAddr(), c)
		}
		option.OnClientOffline = func(c *dot.Client) {
			ac.addrToClientMap.Delete(c.GetRemoteAddr())
		}
	})
	ac.server = server
	return ac
}

func (ac *aCache) Run() error {
	var err error
	wg, f := dot.WaitFunc()
	f(func() {
		err = ac.server.Run()
		if err != nil {
			panic(err)
		}
	})
	ac.log.Println("tcp server ok")
	for _, addr := range ac.allAddr {
		a := addr
		f(func() {
		HERE:
			client, err := dot.Dial(a)
			if err != nil {
				ac.log.Println("dial error:", err)
				time.Sleep(time.Second * 3)
				goto HERE
			}
			client.ReadServerLoop(func(data []byte) {
				ac.log.Println("read srv data:", string(data))
				ac.msgQueue = append(ac.msgQueue, string(data))
			})
			goto HERE
		})
	}
	f(func() {
		mux := http.NewServeMux()
		mux.HandleFunc("/msgQueue", func(writer http.ResponseWriter, request *http.Request) {
			data, _ := json.Marshal(ac.msgQueue)
			_, _ = writer.Write(data)
		})
		mux.HandleFunc("/msgPush", func(writer http.ResponseWriter, request *http.Request) {
			msg := request.URL.Query().Get("msg")
			// 直接转发到所有端口
			//ac.msgQueue = append(ac.msgQueue, msg)
			ac.server.ClientMapRange(func(c *dot.Client) {
				_ = c.Write([]byte(msg))
			})
		})
		err = http.ListenAndServe(ac.httpAddr, mux)
		if err != nil {
			panic(err)
		}
	})
	wg.Wait()
	return nil
}
