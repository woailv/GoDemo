package main

import (
	"fmt"
	"log"
	"os"
	"sync"
	"time"
)

func main() {
	wg := sync.WaitGroup{}
	wg.Add(1)
	go func() {
		defer wg.Done()
		time.Sleep(time.Second * 2)
		f, err := os.OpenFile("temp.txt", os.O_RDWR|os.O_CREATE, os.ModePerm)
		if err != nil {
			panic(err)
		}
		defer f.Close()
		i := 0
		for ; ; i++ {
			_, err := f.Write([]byte(fmt.Sprintf("%d\n", i)))
			if err != nil {
				panic(err)
			}
			time.Sleep(time.Second)
		}
	}()
	for i := 0; i < 2; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			f, err := os.OpenFile("temp.txt", os.O_RDONLY|os.O_CREATE, os.ModePerm)
			if err != nil {
				panic(err)
			}
			defer f.Close()
			for {
				var line string
				_, err := fmt.Fscanf(f, "%s\n", &line)
				if err != nil {
					log.Println(i, "error:", err)
					time.Sleep(time.Second)
					continue
				}
				log.Println(i, "data:", line)
				time.Sleep(time.Second / 2)
			}
		}(i)
	}
	wg.Wait()
}
