package main

import (
	"GoDemo/MessageMiddle"
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
)

func webHandle(request *http.Request) (interface{}, error) {
	if request.Header.Get("Content-Length") > "9999" {
		return nil, errors.New("content too large")
	}
	bList, err := ioutil.ReadAll(request.Body)
	if err != nil {
		return nil, errors.New("read body error")
	}
	request.Body.Close()
	switch request.URL.Path {
	case "themeAdd":
		theme := &MessageMiddle.Theme{}
		json.Unmarshal(bList, theme)
		MessageMiddle.ThemeSave(theme)
	}
	return nil, errors.New("not found")
}
