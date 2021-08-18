package main

import (
	"GoDemo/dot"
	"GoDemo/dot/acache"
)

func main() {
	wg, f := dot.WaitFunc()
	f(func() {
		httpAddr := ":8091"
		addr := ":8081"
		addrList := []string{
			":8081",
			":8082",
		}
		ac := acache.NewACache(httpAddr, addr, addrList)
		err := ac.Run()
		if err != nil {
			panic(err)
		}
	})
	f(func() {
		httpAddr := ":8092"
		addr := ":8082"
		addrList := []string{
			":8081",
			":8082",
		}
		ac := acache.NewACache(httpAddr, addr, addrList)
		err := ac.Run()
		if err != nil {
			panic(err)
		}
	})
	wg.Wait()

}
