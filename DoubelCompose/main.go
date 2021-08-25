package main

import (
	"fmt"
	"sort"
	"time"
)

func main() {
	var dayList []string
	now := time.Now()
	for i := 1; i <= 5; i++ {
		dayList = append(dayList, now.AddDate(0, 0, -i).Format("2006-01-02"))
	}
	res := StringsDoubleCompose(dayList)
	for _, item := range res {
		fmt.Println(item)
	}
}

func StringsDoubleCompose(strList []string) [][]string {
	var result [][]string
	for i := 0; i < len(strList); i++ {
		for k := i; k < len(strList); k++ {
			result = append(result, []string{strList[i], strList[k]})
		}
	}
	for i := range result {
		sort.Strings(result[i])
	}
	return result
}
