package Echo

import (
	"encoding/json"
	"fmt"
	"reflect"
	"strings"
)

func Json(i ...interface{}) {
	out := _out(i...)
	fmt.Println(out...)
}

func JsonPretty(i ...interface{}) {
	out := _outPretty(i...)
	fmt.Println(out...)
}

func _outPretty(i ...interface{}) []interface{} {
	out := []interface{}{}
	for _, item := range i {
		tk := reflect.TypeOf(item).Kind()
		if tk == reflect.Ptr || tk == reflect.Struct || tk == reflect.Slice || tk == reflect.Map {
			bs, err := json.MarshalIndent(item, "", "  ")
			if err != nil {
				panic(err)
			}
			out = append(out, string(bs))
		} else {
			out = append(out, item)
		}
	}
	return out
}

func _out(i ...interface{}) []interface{} {
	out := []interface{}{}
	for _, item := range i {
		tk := reflect.TypeOf(item).Kind()
		if tk == reflect.Ptr || tk == reflect.Struct || tk == reflect.Slice || tk == reflect.Map {
			bs, err := json.Marshal(item)
			if err != nil {
				panic(err)
			}
			out = append(out, string(bs))
		} else {
			out = append(out, item)
		}
	}
	return out
}

func JsonLnf(format string, i ...interface{}) {
	if !strings.Contains(format, "\n") {
		format += "\n"
	}
	out := _out(i...)
	fmt.Printf(format, out...)
}
