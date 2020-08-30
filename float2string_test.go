package kameria

import (
	"strings"
	"testing"
)

func TestFloat2String(t *testing.T) {
	str := Float2String(1.2)
	if str != "1.2" {
		t.Fail()
	}

	str = Float2String(float32(2.2))
	if !strings.HasPrefix(str, "2.20000") {
		t.Fail()
	}
}
