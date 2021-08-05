package main

import (
	"testing"

	"github.com/M-Quadra/kameria"
)

func TestMin(t *testing.T) {
	if kameria.Math.Min.Ints(1, 5) != 1 {
		t.Fail()
		return
	}

	if kameria.Math.Min.Strings("AMD", "Intel") != "AMD" {
		t.Fail()
		return
	}
}
