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

var path2ThemeMap = map[string]*Theme{}

var path2ThemeMapLock = &sync.Mutex{}

func LockThemePathFn(fn func()) {
	LockFn(path2ThemeMapLock, fn)
}

var id2Theme = map[string]*Theme{}

var id2ThemeLock = &sync.Mutex{}

func LockThemeFn(fn func()) {
	LockFn(id2ThemeLock, fn)
}

func ThemeGetByPath(path string) *Theme {
	var theme *Theme
	LockThemePathFn(func() {
		theme = path2ThemeMap[path]
	})
	return theme
}

func ThemeGetById(id string) *Theme {
	var theme *Theme
	LockThemeFn(func() {
		theme = id2Theme[id]
	})
	return theme
}

func ThemeSave(theme *Theme) error {
	LockThemePathFn(func() {
		path2ThemeMap[theme.Path] = theme
	})
	LockThemeFn(func() {
		id2Theme[theme.Id] = theme
	})
	return nil
}

func ThemeDelete(id string) error {
	theme := ThemeGetById(id)
	LockThemePathFn(func() {
		delete(path2ThemeMap, theme.Path)
	})
	LockThemeFn(func() {
		delete(id2Theme, id)
	})
	// ...
	return nil
}
