//获取变量近似实际占用内存空间
package main

import (
	"fmt"
	"reflect"
)

type T struct {
	Name string
	Ages []int
	m    map[string]string
}

func main() {
	a := []T{{Name: "1243", Ages: []int{1, 2, 3}, m: map[string]string{"S": "a"}}}
	result := GetValMem(a)
	fmt.Println("result:", result)
}

func GetValMem(i interface{}) int {
	return GetRValMem(reflect.ValueOf(i))
}

func GetRValMem(rv reflect.Value) int {
	tp := rv.Type()
	switch tp.Kind() {
	case reflect.Bool:
		return 1
	case reflect.Int:
		return 8
	case reflect.Int8:
		return 1
	case reflect.Int16:
		return 2
	case reflect.Int32:
		return 4
	case reflect.Int64:
		return 8
	case reflect.Uint:
		return 8
	case reflect.Uint8:
		return 1
	case reflect.Uint16:
		return 2
	case reflect.Uint32:
		return 4
	case reflect.Uint64:
		return 8
	case reflect.Uintptr:
		return 8
	case reflect.Float32:
		return 4
	case reflect.Float64:
		return 8
	case reflect.Complex64:
		return 8
	case reflect.Complex128:
		return 16
	case reflect.Array, reflect.Slice:
		result := 24
		for i := 0; i < rv.Len(); i++ {
			result += GetRValMem(rv.Index(i))
		}
		return result
	case reflect.Chan:
		return 8
	case reflect.Func, reflect.Interface, reflect.UnsafePointer:
		return 0
	case reflect.Map:
		result := 8
		keys := rv.MapKeys()
		for _, key := range keys {
			result += GetRValMem(key)
			result += GetRValMem(rv.MapIndex(key))
		}
		return result
	case reflect.Ptr:
		rv.Elem()
		return 8 + GetRValMem(rv.Elem())
	case reflect.String:
		return 16 + len(rv.String())
	case reflect.Struct:
		result := 0
		for i := 0; i < rv.NumField(); i++ {
			result += GetRValMem(rv.Field(i))
		}
		return result
	default:
		fmt.Println("kind next:", tp.Kind())
	}
	panic("TODO")
}
