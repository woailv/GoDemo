package XTime

import (
	"fmt"
	"testing"
)

func TestDayGetStartEndListByTime(t *testing.T) {
	result := DayStartEndGetDayList("2020-01-02", "2020-01-02")
	fmt.Println(result)
}
