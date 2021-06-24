package main

import (
	"strconv"
)

type Theme struct {
	Id     string // also is message table name
	Name   string
	Method string
	Path   string // 忽略域名/IP的路径 接收消息使用
}

var themeList = []Theme{}

var path2ThemeMap = map[string]*Theme{
	"/a": {
		Id:     "a",
		Name:   "nameA",
		Method: "GET",
		Path:   "/a",
	},
	"/b": {
		Id:     "b",
		Name:   "nameB",
		Method: "GET",
		Path:   "/b",
	},
}

var themeId2SubIdList = map[string][]string{
	"a": {"sub1", "sub2"},
	"b": {"sub3", "sub4"},
}

func ThemeGetByPath(path string) *Theme {
	return path2ThemeMap[path]
}

func ThemeSave(theme *Theme) error {
	if theme.Id == "" {
		theme.Id = strconv.Itoa(len(themeList))
		themeList = append(themeList, *theme)
	} else {
		for i := 0; i < len(themeList); i++ {
			if themeList[i].Id == theme.Id {
				themeList[i] = *theme
				break
			}
			if i == len(themeList)-1 {
				return ErrNotFund
			}
		}
	}
	return nil
}
