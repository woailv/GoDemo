package main

import (
	"fmt"
	"net/http"
)

var (
	username = "abc"
	password = "123"
)

func main() {
	handler := http.HandlerFunc(handleRequest)
	http.Handle("/example", handler)
	http.ListenAndServe(":8080", nil)
}

func handleRequest(w http.ResponseWriter, r *http.Request) {
	u, p, ok := r.BasicAuth()
	if !ok {
		w.Header().Set("WWW-Authenticate", `Basic realm="Dotcoo User Login"`)
		w.WriteHeader(http.StatusUnauthorized)
		return
	}
	if u != username {
		fmt.Printf("Username provided is correct: %s\n", u)
		w.WriteHeader(401)
		return
	}
	if p != password {
		fmt.Printf("Password provided is correct: %s\n", u)
		w.WriteHeader(401)
		return
	}
	fmt.Printf("Username: %s\n", u)
	fmt.Printf("Password: %s\n", p)
	w.WriteHeader(200)
	return
}
