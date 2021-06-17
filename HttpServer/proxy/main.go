package main

import (
	alloc "GoDemo/AllocTrace"
	"io"
	"log"
	"net"
	"net/http"
	_ "net/http/pprof"
)

func main() {
	alloc.Begin()
	http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		request.Cookies()
		host := request.URL.Host
		log.Println("host:", host)
		cookie := request.Header.Get("Cookie")
		log.Println("cookie:", cookie)
		key1 := request.FormValue("name")
		log.Println("key1:", key1)
		proxy(writer, request, "127.0.0.1:8001")
	})
	err := http.ListenAndServe(":8000", nil)
	if err != nil {
		panic(err)
	}
}

func proxy(w http.ResponseWriter, req *http.Request, nextAddr string) {
	tran := &http.Transport{
		Dial: func(network, addr string) (net.Conn, error) {
			return net.Dial("tcp", nextAddr)
		},
	}
	req.URL.Scheme = "http"
	req.URL.Host = req.Host
	resp, err := tran.RoundTrip(req)
	if err != nil {
		println("uvmhmyyzvw", err.Error())
		w.WriteHeader(500)
		_, _ = w.Write([]byte("5ahnnr7vyb"))
		return
	}
	for key, thisValueList := range resp.Header {
		if key == "Cache-Control" {
			w.Header().Del("Cache-Control")
		}
		for _, thisValue := range thisValueList {
			w.Header().Add(key, thisValue)
		}
	}
	w.WriteHeader(resp.StatusCode)
	_, _ = io.Copy(w, resp.Body)
	_ = resp.Body.Close()
}
