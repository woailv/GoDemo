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
	switch request.URL.Path {
	case "themeAdd":
		handle = MessageMiddle.ThemeSave
	case "subAdd":
		handle = MessageMiddle.SubSave
	case "themeList":
	default:
		return nil, errors.New("not found")
	}
	vf := reflect.ValueOf(handle)
	var callParam []reflect.Value
	for i := 0; i < vf.Type().NumIn(); i++ {
		// struct 从body url 中获取数据, base type 从url中获取数据
		switch vf.Type().In(i).Kind() {
		case reflect.Ptr:
			if request.Method != "GET" {
				err = json.Unmarshal(bList, reflect.New(reflect.TypeOf(vf.Type().In(i).Elem())).Interface())
				if err != nil {
					return nil, errors.New("param error")
				}
			} else {
				panic("TODO")
			}
		case reflect.Struct:
		case reflect.Map:
		default:
			panic("TODO")
		}
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
