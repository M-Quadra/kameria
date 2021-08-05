package main

import (
	"testing"

	"github.com/M-Quadra/kameria"
)

func TestMax(t *testing.T) {
	if kameria.Math.Max.Ints(1, 3) != 3 {
		t.Fail()
		return
	}
}
