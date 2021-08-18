package main

import (
	"GoDemo/Echo"
	"math/rand"
	"reflect"
	"strconv"
	"strings"
	"time"
)

func main() {
	list := TestDataGen(new(T), 3)
	Echo.JsonPretty(list)
}

type T struct {
	Name string `test:"n1,n2"`
	Age  int    `test:"3,5"`
}

func TestDataGen(x interface{}, count int) (list []interface{}) {
	rand.Seed(time.Now().Unix())
	rv := reflect.ValueOf(x).Elem()
	rt := rv.Type()
	for i := 0; i < count; i++ {
		item := reflect.New(rt)
		itemElem := item.Elem()
		for i := 0; i < itemElem.NumField(); i++ {
			tag, ok := itemElem.Type().Field(i).Tag.Lookup("test")
			if !ok {
				continue
			}
			valList := strings.Split(tag, ",")
			valRand := valList[rand.Intn(len(valList))]
			switch item.Elem().Field(i).Type().Kind() {
			case reflect.String:
				item.Elem().Field(i).SetString(valRand)
			case reflect.Int:
				val, _ := strconv.Atoi(valRand)
				item.Elem().Field(i).SetInt(int64(val))
			}
		}
		list = append(list, item.Interface())
	}
	return
}
