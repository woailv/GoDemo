package MessageMiddle

import (
	"fmt"
	"testing"
	"time"
)

func InitData() {
	theme := &Theme{
		Id:     "theme1",
		Name:   "themeName1",
		Method: "GET",
		Path:   "/a",
	}
	_ = ThemeSave(theme)
	sub := &Sub{
		Id:      "sub1",
		ThemeId: theme.Id,
		Name:    "subName1",
		Url:     "/a/sub",
		Method:  "GET",
	}
	_ = SubSave(sub)
	sub2 := &Sub{
		Id:      "sub12",
		ThemeId: theme.Id,
		Name:    "subName2",
		Url:     "/a/sub2",
		Method:  "GET",
	}
	_ = SubSave(sub2)
}

func TestA(t *testing.T) {
	InitData()
	go SendMessage()
	for i := 0; i < 10; i++ {
		ReceiverMessage("theme1", fmt.Sprintf("this is a message:%d", i))
		time.Sleep(time.Second * 3)
	}
}
