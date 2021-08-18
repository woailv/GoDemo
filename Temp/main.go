package main

import (
	"fmt"
	"log"
	"reflect"
	"time"
)

func main() {
	t := T{Name: "haha"}
	rv := reflect.ValueOf(t)
	var ok bool
	rvf := rv.MethodByName("NameView1")
	ok = rvf.IsValid()
	//ok = rvf.IsNil()
	//ok = rvf.IsZero()
	log.Println(ok)
	//rvf.Call([]reflect.Value{})
}

type T struct {
	Name string
	Age  int
}

func (t *T) NameView() {
	fmt.Println(t.Name)
}

func ViewTime() {
	tm := time.Unix(0, 1628834133330*1000000)
	fmt.Println(tm)
}
