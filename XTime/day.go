package XTime

import (
	"fmt"
	"time"
)

func DayGetStartEndList(startDate, endDate string) [][2]string {
	location, err := time.LoadLocation("Asia/Shanghai")
	if err != nil {
		panic(err)
	}
	start, err := time.ParseInLocation("2006-01-02", startDate, location)
	if err != nil {
		panic(err)
	}
	end, err := time.ParseInLocation("2006-01-02", endDate, location)
	if err != nil {
		panic(err)
	}
	if start.After(end) {
		panic("开始时间不能小于结束时间")
	}
	xs := DayGetStartEndListByTime(start, end)
	result := [][2]string{}
	for _, x := range xs {
		result = append(result, [2]string{x[0].Format("2006-01-02"), x[1].Format("2006-01-02")})
	}
	return result
}

func DayGetStartEndListByTime(start, end time.Time) [][2]time.Time {
	if start.Equal(end) || start.After(end) {
		return nil
	}
	fmt.Println("start:", start)
	return append([][2]time.Time{{start, start.Add(time.Hour * 24)}}, DayGetStartEndListByTime(start.Add(time.Hour*24), end)...)
}

func DayStartEndGetDayList(startDate, endDate string) []string {
	location, err := time.LoadLocation("Asia/Shanghai")
	if err != nil {
		panic(err)
	}
	start, err := time.ParseInLocation("2006-01-02", startDate, location)
	if err != nil {
		panic(err)
	}
	end, err := time.ParseInLocation("2006-01-02", endDate, location)
	if err != nil {
		panic(err)
	}
	if start.After(end) {
		panic("开始时间不能小于结束时间")
	}
	result := []string{}
	tms := DayStartEndGetDayListByTime(start, end)
	for _, tm := range tms {
		result = append(result, tm.Format("2006-01-02"))
	}
	return result
}

func DayStartEndGetDayListByTime(start, end time.Time) []time.Time {
	if start.After(end) {
		return nil
	}
	fmt.Println("start:", start)
	return append([]time.Time{start}, DayStartEndGetDayListByTime(start.Add(time.Hour*24), end)...)
}
