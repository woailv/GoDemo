package main

import (
	"io"
	"log"
	"os"
)

func main() {
	f, err := os.OpenFile("./temp.txt", os.O_RDWR|os.O_CREATE, 0600)
	if err != nil {
		panic(err)
	}
	defer f.Close()
	data := make([]byte, 1)
	for i := 0; i < 5; i++ {
		n, err := io.ReadFull(f, data)
		if err != nil {
			panic(err)
		}
		log.Println("i:", i, ",n:", n, ",data:", string(data))
	}
}
