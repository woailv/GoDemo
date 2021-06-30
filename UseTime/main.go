package UseTime

import (
	"fmt"
	"time"
)

func View(f func()) {
	start := time.Now()
	f()
	useTime := time.Now().Sub(start)
	fmt.Println("use time:", useTime)
}
