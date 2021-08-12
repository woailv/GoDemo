package main

import (
	log "fmt"
	"net"
	"strconv"
	"sync"
	"time"
)

func main() {
	conn, err := net.Dial("tcp", ":8080")
	if err != nil {
		panic(err)
	}
	i := 0
	go func() {
		for {
			data := make([]byte, 1024)
			n, err := conn.Read(data)
			if err != nil {
				panic(err)
			}
			data = data[:n]
			log.Println("read data:", string(data))
			err = write(conn, 0)
			if err != nil {
				panic(err)
			}
		}
	}()
	for {
		i++
		time.Sleep(time.Second * 3)
		if err := write(conn, i); err != nil {
			panic(err)
		}
		if i == 3 {
			continue
		}
	}
}

var writeMu sync.Mutex

func write(conn net.Conn, i int) error {
	writeMu.Lock()
	defer writeMu.Unlock()
	_, err := conn.Write([]byte(strconv.Itoa(i)))
	return err
}
