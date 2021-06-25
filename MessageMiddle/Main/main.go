package Main

import (
	"GoDemo/Err"
	"GoDemo/MessageMiddle"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

func main() {
	port := ":8080"
	fmt.Println("http://127.0.0.1" + port)
	http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		if strings.Contains(request.URL.Path, ".ico") {
			return
		}
		theme := MessageMiddle.ThemeGetByPath(request.URL.Path)
		if theme == nil {
			return
		}
		if theme.Method != request.Method {
			return
		}
		content := ""
		if request.Method == "GET" {
			content = request.URL.RawQuery
		} else if request.Method == "POST" {
			if request.Header.Get("Content-Length") > "9999" {
				return
			}
			bs, err := ioutil.ReadAll(request.Body)
			if err != nil {
				return
			}
			request.Body.Close()
			content = string(bs)
		}
		MessageMiddle.ReceiverMessage(theme.Id, content)
	})
	go MessageMiddle.SendMessage()
	err := http.ListenAndServe(port, nil)
	Err.IfPanic(err)
}
