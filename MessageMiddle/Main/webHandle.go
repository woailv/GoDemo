package main

import (
	"GoDemo/MessageMiddle"
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"reflect"
	"strings"
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
	method := request.Method
	path := request.URL.Path
	return handle(path, method, bList)
}

func handle(path string, method string, bList []byte) (interface{}, error) {
	var handle interface{}
	actionDesc := strings.Split(path, "/")
	if len(actionDesc) != 2 {
		panic("TODO")
	}
	switch actionDesc[0] {
	case "theme":
		switch actionDesc[1] {
		case "save":
			handle = MessageMiddle.ThemeSave
		case "list":
		default:
			handle = MessageMiddle.ThemeGetById
		}
	default:
		return nil, errors.New("not found")
	}
	vf := reflect.ValueOf(handle)
	var callParam []reflect.Value
	for i := 0; i < vf.Type().NumIn(); i++ {
		// struct 从body url 中获取数据, base type 从url中获取数据
		switch vf.Type().In(i).Kind() {
		case reflect.Ptr:
			if method != "GET" {
				value := reflect.New(vf.Type().In(i).Elem())
				err := json.Unmarshal(bList, value.Interface())
				if err != nil {
					return nil, errors.New("param error")
				}
				callParam = append(callParam, value)
			} else {
				panic("TODO")
			}
		case reflect.Struct:
		case reflect.Map:
		case reflect.String: //id从path param中获取
			callParam = append(callParam, reflect.ValueOf(actionDesc[1]))
		default:
			panic("TODO")
		}
	}
	vfResult := vf.Call(callParam)

	if len(vfResult) == 1 {
		err, ok := vfResult[0].Interface().(error)
		if ok {
			return nil, err
		}
		return vfResult[0].Interface(), nil
	}
	if len(vfResult) == 2 {
		return vfResult[0].Interface(), vfResult[1].Interface().(error)
	}
	panic("TODO")
}
