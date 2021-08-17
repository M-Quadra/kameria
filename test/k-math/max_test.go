package main

import (
	"testing"
	"time"

	"github.com/M-Quadra/kameria"
	kmath "github.com/M-Quadra/kameria/k-math"
)

func TestMaxInts(t *testing.T) {
	if kmath.Max.Ints(1, 3) != 3 {
		t.Fail()
		return
	}

	if kameria.Math.Max(1, 3) != 3 {
		t.Fail()
		return
	}
}

func TestMaxInt64s(t *testing.T) {
	if kmath.Max.Int64s(5, 7) != 7 {
		t.Fail()
		return
	}

	if kameria.Math.Max(int64(5), int64(7)) != int64(7) {
		t.Fail()
		return
	}
	if kameria.Math.Max(int64(5), int64(7)) == 7 { // default int, not int64
		t.Fail()
		return
	}
}

func TestMaxStrings(t *testing.T) {
	a, b := "M$", "MS"
	if a > b {
		t.Fail()
		return
	}

	if kmath.Max.Strings(a, b) != b {
		t.Fail()
		return
	}

	if kameria.Math.Max(a, b) != b {
		t.Fail()
		return
	}
}

func TestMaxFloat32s(t *testing.T) {
	if kmath.Max.Float32s(1.1, 1.2) != 1.2 {
		t.Fail()
		return
	}

	if kameria.Math.Max(float32(1.1), float32(1.2)) != float32(1.2) {
		t.Fail()
		return
	}
	if kameria.Math.Max(float32(1.1), float32(1.2)) == 1.2 { // default float64
		t.Fail()
		return
	}
}

func TestMaxFloat64s(t *testing.T) {
	if kmath.Max.Float64s(1.1, 1.2) != 1.2 {
		t.Fail()
		return
	}

	if kameria.Math.Max(1.1, 1.2) != 1.2 {
		t.Fail()
		return
	}
}

func TestMaxTimes(t *testing.T) {
	a := time.Now()
	b := a.Add(time.Hour)
	if !kmath.Max.Times(a, b).Equal(b) {
		t.Fail()
		return
	}

	if !kameria.Math.Max(a, b).(time.Time).Equal(b) {
		t.Fail()
		return
	}
}
