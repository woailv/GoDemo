package main

import (
	"GoDemo/Err"
	"GoDemo/MessageMiddle"
	"encoding/json"
	"fmt"
	"net/http"
)

func main() {
	msgServer()
	webServer()
}

func webServer() {
	webPort := ":8081"
	fmt.Println("http://127.0.0.1" + webPort)
	go MessageMiddle.SendMessage()
	webMux := http.NewServeMux()
	webMux.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		msg := "ok"
		data, err := webHandle(request)
		if err != nil {
			msg = err.Error()
		}
		bs, err := json.Marshal(data)
		if err == nil {
			msg = string(bs)
		}
		fmt.Fprintf(writer, msg)
	})
	err := http.ListenAndServe(webPort, webMux)
	Err.IfPanic(err)
}

func msgServer() {
	msgPort := ":8080"
	fmt.Println("http://127.0.0.1" + msgPort)
	go MessageMiddle.SendMessage()
	msgMux := http.NewServeMux()
	msgMux.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		msg := "ok"
		err := messageHandle(request)
		if err != nil {
			msg = err.Error()
		}
		fmt.Fprintf(writer, msg)
	})
	go func() {
		err := http.ListenAndServe(msgPort, msgMux)
		Err.IfPanic(err)
	}()
}
