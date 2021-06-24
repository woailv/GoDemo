package Echo

import (
	"fmt"
	"testing"
)

func TestJson(t *testing.T) {
	Json("a", map[string]string{"b": "2"})
	fmt.Println("a", map[string]string{"b": "2"})
}
