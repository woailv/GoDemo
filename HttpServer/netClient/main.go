package main

import (
	alloc "GoDemo/AllocTrace"
	"fmt"
	"log"
	"net"
	"sync"
	"time"
)

func main() {
	alloc.Begin()
	wg := sync.WaitGroup{}
	for i := 0; i < 1; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			sendData()
		}()
	}
	wg.Wait()
	log.Println("end")
}

func sendData() {
	conn, err := net.Dial("tcp", "127.0.0.1:8000")
	if err != nil {
		panic(err)
	}
	_, err = conn.Write([]byte(body))
	if err != nil {
		panic(err)
	}
	go func() {
		for {
			buf := make([]byte, 1024)
			n, err := conn.Read(buf)
			if err != nil {
				panic(err)
			} else {
				log.Println("read data:", string(buf[:n]))
			}
		}
	}()
	writeTotal := 0
	i := 0
	for ; true; {
		i++
		n, err := conn.Write([]byte(fmt.Sprintf("name%d=ajjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjj%d", i, i)))
		if err != nil {
			panic(err)
		}
		writeTotal += n
		if writeTotal%1000 == 0 {
			log.Println("write total:", writeTotal)
		}
		time.Sleep(time.Millisecond * 10)
	}
	for {
		time.Sleep(time.Second)
	}
}

const body = `POST / HTTP/1.1
Host: 127.0.0.1:8000
User-Agent: Go-http-client/1.1
Content-Length: 68888888888888
Content-Type: application/x-www-form-urlencoded
Accept-Encoding: gzip

name=11
`
