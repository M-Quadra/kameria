package kameria

import (
	"fmt"
	"testing"
)

func TestInt2String(t *testing.T) {
	fmt.Println("[kameria] Int2String Test...")
	v1 := int(3990)
	v2 := int8(8)
	v3 := int16(16)
	v4 := int32(32)
	v5 := int64(64)

	if Int2String(v1, 10) != "3990" {
		t.Fail()
	}
	if Int2String(v2, 10) != "8" {
		t.Fail()
	}
	if Int2String(v3, 10) != "16" {
		t.Fail()
	}
	if Int2String(v4, 10) != "32" {
		t.Fail()
	}
	if Int2String(v5, 16) != "40" {
		t.Fail()
	}
}
