package regc

import (
	"testing"
)

func Test_regc_Run(t *testing.T) {
	rc := NewRegc(":8050")
	rc.Run()
}
