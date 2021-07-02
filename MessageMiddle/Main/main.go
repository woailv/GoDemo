package main

import (
	"GoDemo/Err"
	"GoDemo/MessageMiddle"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

func main() {
	port := ":8080"
	fmt.Println("http://127.0.0.1" + port)
	http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		msg := "ok"
		err := messageHandle(request)
		if err != nil {
			msg = err.Error()
		}
		fmt.Fprintf(writer, msg)
	})
	go MessageMiddle.SendMessage()
	err := http.ListenAndServe(port, nil)
	Err.IfPanic(err)
}

func messageHandle(request *http.Request) error {
	if strings.Contains(request.URL.Path, ".ico") {
		return nil
	}
	var theme *MessageMiddle.Theme
	content := ""
	theme = MessageMiddle.ThemeGetByPath(request.URL.Path)
	if theme == nil {
		return errors.New("theme not found")
	}
	if theme.Method != request.Method {
		return errors.New("http method error")
	}
	if request.Method == "GET" {
		content = request.URL.RawQuery
	} else if request.Method == "POST" {
		if request.Header.Get("Content-Length") > "9999" {
			return errors.New("content too large")
		}
		bs, err := ioutil.ReadAll(request.Body)
		if err != nil {
			return errors.New("read body error")
		}
		request.Body.Close()
		content = string(bs)
	}
	MessageMiddle.ReceiverMessage(theme.Id, content)
	return nil
}
