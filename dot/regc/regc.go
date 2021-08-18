package regc

import (
	"GoDemo/dot"
	"encoding/json"
	"fmt"
	"log"
	"os"
)

type regc struct {
	addr   string
	server dot.Server
	log    *log.Logger
}

const Online = "Online"
const Offline = "Offline"
const ExistAddrList = "ExistAddrList"

func NewRegc(addr string) *regc {
	rc := &regc{
		addr: addr,
		log:  log.New(os.Stderr, "regc ", log.LstdFlags|log.Lshortfile),
	}
	rc.server = dot.NewServer(addr, func(option *dot.ServerOption) {
		option.ReadData = func(c *dot.Client, data []byte) {
		}
		option.OnClientOnline = func(c *dot.Client) {
			data, _ := json.Marshal(rc.server.GetClientAddrList())
			err := c.Write([]byte(fmt.Sprintf("%s,%s", ExistAddrList, string(data))))
			if err != nil {
				panic(err)
			}
			rc.server.ClientMapRange(func(exist *dot.Client) {
				err := exist.Write([]byte(fmt.Sprintf("%s,%s", Online, c.GetRemoteAddr())))
				if err != nil {
					rc.log.Println("write error", err)
				}
			})
		}
		option.OnClientOffline = func(c *dot.Client) {
			rc.server.ClientMapRange(func(exist *dot.Client) {
				err := exist.Write([]byte(fmt.Sprintf("%s,%s", Offline, c.GetRemoteAddr())))
				if err != nil {
					rc.log.Println("write error", err)
				}
			})
		}
	})
	return rc
}

func (rc *regc) Run() {
	var err error
	w, f := dot.WaitFunc()
	f(func() {
		err = rc.server.Run()
		if err != nil {
			panic(err)
		}
	})
	w.Wait()
}
