package main

import (
	"fmt"
	"strings"
	"time"
)

func main() {
	//a, _ := strconv.ParseInt("1625192931284", 10, 64)
	//fmt.Println(a)
	//fmt.Println(time.Unix(0, a*1000000))
	//timePrint()
	//fmt.Println(len("K1TimeSourceKindOriginalTransactionIdTransactionIdFromKindToSubOrderTrackStatus"))
	//m := map[string][]string{}
	//s := m["a"]
	//s = append(s, "1")
	//m["a"] = s
	//fmt.Println(m)
	//fmt.Println("2006-01-02"[:10])
	//type T struct {
	//	A string
	//}
	//m := map[string]*T{}
	//a := m["a"]
	//if a == nil {
	//	a = &T{}
	//	m["a"] = a
	//}
	//fmt.Println(m)
	//resp, err := http.PostForm("http://127.0.0.1:8080/b/", nil)
	//Err.IfPanic(err)
	//bs, err := ioutil.ReadAll(resp.Body)
	//Err.IfPanic(err)
	//log.Println(string(bs))
	//ch := make(chan int,1)
	//ch <- 1
	//fmt.Println(len(ch))

	//UseTime.View(func() {
	//	m := sync.Map{}
	//	//w := sync.WaitGroup{}
	//	for i := 0; i < 10000000; i++ {
	//		//w.Add(1)
	//		//go func() {
	//			m.Store(i, i)
	//		//	w.Done()
	//		//}()
	//	}
	//	//w.Wait()
	//	m.Range(func(key, value interface{}) bool {
	//		return true
	//	})
	//})

	//timePrint()

	fmt.Println(strings.SplitN("1,1,1,1", ",", 2))
}

func timePrint() {
	s := []int64{
		1626685332210,
		1626685332211,
		1626685332212,
		1626685332213,
		1626685332214,
		1626685332215,
		1626685332216,
	}
	fmt.Println(len(s[1:]))
	for _, i := range s {
		tm := time.Unix(0, i*1000000)
		fmt.Println(tm)
	}
}
