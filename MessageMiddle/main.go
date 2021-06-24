package main

import (
	"GoDemo/Err"
	"fmt"
	"net/http"
	"strings"
)

func main() {
	port := ":8080"
	fmt.Println("http://127.0.0.1" + port)
	http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		path := request.URL.Path
		if strings.Contains(path, ".ico") {
			return
		}
		themeId := path[1:]
		fmt.Println("theme:", themeId)
		theme := id2ThemeMap[themeId]
		if theme == nil {
			return
		}
		if theme.Method != request.Method {
			return
		}
		ReceiverMessage(themeId, "this is a message")
	})
	go SendMessage()
	err := http.ListenAndServe(port, nil)
	Err.IfPanic(err)
}
