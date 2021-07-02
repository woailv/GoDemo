package main

import (
	"GoDemo/MessageMiddle"
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"reflect"
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
	var handle interface{}
	var param interface{}
	switch request.URL.Path {
	case "themeAdd":
		param = &MessageMiddle.Theme{}
		handle = MessageMiddle.ThemeSave
	case "subAdd":
		param = &MessageMiddle.Sub{}
		handle = MessageMiddle.SubSave
	case "themeList":
	default:
		return nil, errors.New("not found")
	}
	if param != nil {
		if request.Method != "GET" {
			err = json.Unmarshal(bList, param)
			if err != nil {
				return nil, errors.New("param error")
			}
		} else {
			panic("TODO")
		}
	}
	vf := reflect.ValueOf(handle)
	var callParam []reflect.Value
	switch vf.Type().NumIn() {
	case 0:
	case 1:
		callParam = append(callParam, reflect.ValueOf(param))
	default:
		panic("TODO")
	}
	vfResult := vf.Call(callParam)

	if len(vfResult) == 1 {
		return nil, vfResult[0].Interface().(error)
	}
	if len(vfResult) == 2 {
		return vfResult[0].Interface(), vfResult[1].Interface().(error)
	}
	panic("TODO")
}
