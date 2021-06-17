package alloc

import (
	"fmt"
	"log"
	"math"
	"runtime"
	"strconv"
	"sync"
	"time"
)

type debugAlloc struct {
	Flag     bool
	MaxAlloc float64
	mu       sync.Mutex
}

func (d *debugAlloc) Reset() {
	d.MaxAlloc = 0
}
func (d *debugAlloc) Start() {
	d.mu.Lock()
	defer d.mu.Unlock()
	if d.Flag {
		return
	}
	go func() {
		for {
			stat := &runtime.MemStats{}
			runtime.ReadMemStats(stat)
			max := math.Max(d.MaxAlloc, float64(stat.Alloc))
			if Decimal(max) > Decimal(d.MaxAlloc) {
				d.MaxAlloc = max
				log.Println("当前内存峰值:", d.GetAllocView())
			} else {
				if time.Now().Second()%4 == 0 || time.Now().Second()%3 == 0 {
					log.Println("当前内存峰值:", d.GetAllocView(), ";当前内存:", getAllocView(float64(stat.Alloc)))
				}
			}
			time.Sleep(time.Second * 2)
		}
	}()
	d.Flag = true
}

func (d *debugAlloc) GetAllocView() string {
	return fmt.Sprintf("%.2fM", d.MaxAlloc/1024/1024)
}

func getAllocView(v float64) string {
	return fmt.Sprintf("%.2fM", v/1024/1024)
}

var dAlloc = &debugAlloc{}

func Decimal(value float64) int {
	value, _ = strconv.ParseFloat(fmt.Sprintf("%.2f", value), 64)
	return int(value / 1024 / 1024 * 100)
}

func Begin() {
	dAlloc.Start()
}
