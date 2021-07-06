package main

import (
	"GoDemo/MessageMiddle"
	"errors"
	"io/ioutil"
	"log"
	"net/http"
	"reflect"
	"strings"
)

func webHandle(request *http.Request) (interface{}, error) {
	if strings.Contains(request.URL.Path, ".ico") {
		return nil, nil
	}
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
	pathElemList := strings.Split(path[1:], "/")
	if len(pathElemList) < 2 {
		log.Println(pathElemList)
		panic("TODO")
	}
	switch pathElemList[0] {
	case "theme":
		switch pathElemList[1] {
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
	switch vf.Type().NumIn() {
	case 0:
	case 1:
		switch vf.Type().In(0).Kind() {
		case reflect.Ptr, reflect.Struct, reflect.Map:
		case reflect.String: //id从path param中获取
			callParam = append(callParam, reflect.ValueOf(pathElemList[len(pathElemList)-1]))
		default:
			panic("TODO")
		}
	default:
		panic("参数个数不匹配")
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
