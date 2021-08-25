// list 实现sql group by

package main

import (
	"GoDemo/Echo"
	"reflect"
	"strings"
)

func main() {
	result := ListGroupBy([]T{
		{
			A: "a",
			B: "b",
			C: "c",
			N: 1,
			S: 2,
		},
		{
			A: "a",
			B: "b",
			C: "c",
			N: 1,
			S: 2,
		},
		{
			A: "a",
			B: "b1",
			C: "c",
			N: 1,
			S: 2,
		},
	}, "A,B", "N")
	Echo.Json(result)
}

type T struct {
	A string
	B string
	C string
	N int
	S int
}

func ListGroupBy(list interface{}, groupBy string, sel string) map[string]map[string]int64 {
	byList := strings.Split(groupBy, ",")
	if len(byList) == 0 {
		panic("by list can't be zero")
	}
	selList := strings.Split(sel, ",")
	lrv := reflect.ValueOf(list)
	result := map[string]map[string]int64{}
	for i := 0; i < lrv.Len(); i++ {
		item := lrv.Index(i)
		for ; item.Kind() == reflect.Ptr; item.Elem() {
		}
		var byValList []string
		for _, by := range byList {
			fv := item.FieldByName(by)
			if !fv.IsValid() {
				panic("select field can't find")
			}
			byValList = append(byValList, fv.String())
		}
		var byVal = strings.Join(byValList, ",")
		for k := 0; k < item.NumField(); k++ {
			name := item.Type().Field(k).Name
			for _, sel := range selList {
				if name == sel {
					v := item.Field(k)
					if result[byVal] == nil {
						result[byVal] = map[string]int64{}
					}
					result[byVal][sel] += v.Int()
				}
			}
		}
		if result[byVal] == nil {
			result[byVal] = map[string]int64{}
		}
	}
	return result
}
