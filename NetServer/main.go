package main

import (
	"fmt"
	"net"
)

func main() {
	l, err := net.Listen("tcp", "127.0.0.1:9500")
	if err != nil {
		panic(err)
	}
HERE:
	conn, err := l.Accept()
	if err != nil {
		panic(err)
	}
	for true {
		buf := make([]byte, 1024*1024)
		n, err := conn.Read(buf)
		if err != nil {
			fmt.Println("read error:", err)
			goto HERE
		}
		fmt.Println("data:", string(buf[:n]))
	}
}
