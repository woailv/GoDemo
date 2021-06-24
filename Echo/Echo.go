package Echo

import (
	"encoding/json"
	"fmt"
	"reflect"
)

func Json(i ...interface{}) {
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
	fmt.Println(out...)
}
