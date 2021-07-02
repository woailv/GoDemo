package main

import (
	"GoDemo/MessageMiddle"
	"encoding/json"
	"fmt"
	"testing"
)

func Test_handle(t *testing.T) {
	bs, _ := json.Marshal(MessageMiddle.Theme{
		Id:     "123",
		Name:   "name1",
		Method: "",
		Path:   "",
	})
	r, err := handle("theme/save", "POST", bs)
	fmt.Println(r, err)
	r, err = handle("theme/123", "GET", nil)
	fmt.Println(r, err)
}
