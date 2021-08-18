package main

import (
	"GoDemo/dot"
	"log"
)

func main() {
	client, err := dot.Dial(":8050")
	if err != nil {
		panic(err)
	}
	client.ReadServerLoop(func(data []byte) {
		log.Println(string(data))
	})
}
