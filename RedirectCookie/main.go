package main

import (
	"fmt"
	"net/http"
)
// redirect 可以设置cookie成功
func main() {
	http.HandleFunc("/r", func(writer http.ResponseWriter, request *http.Request) {
		fmt.Fprintf(writer, "123")
	})
	http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		cookie := &http.Cookie{
			Name:     "name3",
			Value:    "a",
			Path:     "/",
			MaxAge:   300,
		}
		http.SetCookie(writer, cookie)
		http.Redirect(writer, request, "http://www.baidu.com", 302)
	})
	port := ":8081"
	fmt.Println(fmt.Sprintf("http://127.0.0.1%s", port))
	fmt.Println(fmt.Sprintf("http://127.0.0.1%s/r", port))
	if err := http.ListenAndServe(port, nil); err != nil {
		panic(err)
	}
}
