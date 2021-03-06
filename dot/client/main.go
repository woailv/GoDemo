package main

import (
	"GoDemo/dot"
	"log"
	"strconv"
	"time"
)

func main() {
	client, err := dot.Dial(":8080")
	if err != nil {
		panic(err)
	}
	wait, f := dot.WaitFunc()
	f(func() {
		for i := 1; ; i++ {
			err := client.Write([]byte(strconv.Itoa(i)))
			if err != nil {
				log.Println("write error:", err)
				return
			}
			time.Sleep(time.Second * 3)
			if i == 2 {
				//client.Exist()
				//return
			}
		}
	})
	f(func() {
		client.ReadServerLoop(func(data []byte) {
			log.Println("data:", string(data))
		})
	})
	wait.Wait()
	client.Exist()
}
