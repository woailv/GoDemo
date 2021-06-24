package main

type Sub struct {
	Id            string
	ThemeId       string
	Name          string
	Url           string
	MaxRetryTimes int
}

var subId2SubMap = map[string]*Sub{
	"sub1": {
		Id:            "sub1",
		ThemeId:       "a",
		Name:          "sub1name",
		Url:           "http://127.0.0.1:9001",
		MaxRetryTimes: 3,
	},
	"sub2": {
		Id:            "sub2",
		ThemeId:       "a",
		Name:          "sub2name",
		Url:           "http://127.0.0.1:9002",
		MaxRetryTimes: 3,
	},
	"sub3": {
		Id:            "sub3",
		ThemeId:       "b",
		Name:          "sub3name",
		Url:           "http://127.0.0.1:9003",
		MaxRetryTimes: 3,
	},
	"sub4": {
		Id:            "sub4",
		ThemeId:       "b",
		Name:          "sub4name",
		Url:           "http://127.0.0.1:9004",
		MaxRetryTimes: 3,
	},
}

func SubGetById(id string) *Sub {
	return subId2SubMap[id]
}
