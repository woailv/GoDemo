package MessageMiddle

import (
	"sync"
)

type Theme struct {
	Id     string // also is message table name
	Name   string
	Method string
	Path   string // 忽略域名/IP的路径 接收消息使用
}

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

var path2ThemeMapLock = &sync.Mutex{}

func LockThemePathFn(fn func()) {
	LockFn(path2ThemeMapLock, fn)
}

func ThemeGetByPath(path string) *Theme {
	var theme *Theme
	LockThemePathFn(func() {
		theme = path2ThemeMap[path]
	})
	return theme
}

func ThemeSave(theme *Theme) error {
	LockThemePathFn(func() {
		path2ThemeMap[theme.Path] = theme
	})
	return nil
}
