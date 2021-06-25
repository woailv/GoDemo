//获取变量近似实际占用内存空间
package main

func main() {
	_ = map[string]int{}
}

func GetValMem(val interface{}) int {
	panic("TODO")
}
