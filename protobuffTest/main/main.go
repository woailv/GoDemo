package main

import (
	"GoDemo/UseTime"
	"encoding/json"
	"github.com/golang/protobuf/proto"
)

//protoc --go_out=. *.proto

func main() {
	p := &Person{
		Name:  "100000000",
		Id:    1,
		Email: "111111111111111111111",
	}
	var data []byte
	var err error
	var times = 10000 * 10000
	loop(times, func() {
		data, err = proto.Marshal(p)
		if err != nil {
			panic(err)
		}
		err = proto.Unmarshal(data, p)
		if err != nil {
			panic(err)
		}
	})
	//fmt.Println("proto date len:", len(data))
	loop(times, func() {
		data, err = json.Marshal(p)
		if err != nil {
			panic(err)
		}
		err = json.Unmarshal(data, p)
		if err != nil {
			panic(err)
		}
	})
	//fmt.Println("json date len:", len(data))
}

func loop(times int, f func()) {
	UseTime.View(func() {
		for i := 0; i < times; i++ {
			f()
		}
	})
}
