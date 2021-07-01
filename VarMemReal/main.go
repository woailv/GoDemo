//获取变量近似实际占用内存空间
package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

type T struct {
	Name string
	Ages []int
	m    map[string]string
}

func main() {
	//f1()
	//GenValMonitor("./VarMemReal/Gvar.go")
	r := MonitorGenVar("a")
	fmt.Println(r)
}

// GenValMonitor TODO
func GenValMonitor(path string) error {
	bs, err := ioutil.ReadFile(path)
	if err != nil {
		panic(err)
	}
	rows := strings.Split(string(bs), "\n")
	for _, row := range rows {
		if len(row) < 3 || row[:3] != "var" {
			continue
		}
		cols := strings.Split(row, " ")
		fmt.Println(cols[1])
	}
	return nil
}

func f1() {
	a := []T{{Name: "1243", Ages: []int{1, 2, 3}, m: map[string]string{"S": "a"}}}
	result := GetValMem(a)
	fmt.Println("result:", result)
	//result: 134
	result = GetValMem([]byte{1})
	fmt.Println("result:", result)
}
