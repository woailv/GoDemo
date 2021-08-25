package main

import (
	"fmt"
	"log"
	"sort"
	"time"
)

func main() {
	var dayList []string
	now := time.Now()
	// 近几天的日期
	for i := 0; i <= 3; i++ {
		dayList = append(dayList, now.AddDate(0, 0, -i).Format("2006-01-02"))
	}
	log.Println("day list:", dayList)
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
