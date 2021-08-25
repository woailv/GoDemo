package main

import (
	"fmt"
	"time"
)

func main() {
	var dayList []string
	now := time.Now()
	for i := 0; i <= 3; i++ {
		dayList = append(dayList, now.AddDate(0, 0, -i).Format("2006-01-02"))
	}
	fmt.Println(dayList[len(dayList)-1], dayList[0])
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
