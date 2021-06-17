package main

import (
	"log"
	"os"
	"os/signal"
)

func main() {
	var ch = make(chan os.Signal)
	signal.Notify(ch)
	for s:=range ch{
		log.Println(s)
	}
}
