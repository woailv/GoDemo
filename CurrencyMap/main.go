package main

import (
	"GoDemo/Echo"
	"time"
)

func main() {
	m := map[int]int{
		1: 1,
		2: 2,
		3: 3,
		4: 4,
		5: 5,
	}
	for i := 0; i < 5; i++ {
		if i == 3 {
			go func() {
				for false {
					m[3] = 3
					time.Sleep(time.Millisecond)
				}
			}()
		}
		go func(k int) {
			for true {
				v := languageMapGetByTextAndKind("123", LanguageEN)
				Echo.Json(v)
				time.Sleep(time.Millisecond)
			}
		}(i)
	}
	select {}
}

const LanguageEN = "en"
const LanguageFA = "fa"

// textToKindToText
var languageMap = map[string]map[string]string{
	"123": {
		LanguageEN: "123",
		LanguageFA: "123",
	},
}

func languageMapGetByTextAndKind(text, kind string) string {
	m := languageMap[text]
	if m == nil {
		return text
	}
	v, ok := m[kind]
	if !ok {
		return text
	}
	return v
}
