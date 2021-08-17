package main

import (
	"testing"
	"time"

	"github.com/M-Quadra/kameria"
	kmath "github.com/M-Quadra/kameria/k-math"
)

func TestMinInts(t *testing.T) {
	if kmath.Min.Ints(0, 2) != 0 {
		t.Fail()
		return
	}

	if kameria.Math.Min(0, 2) != 0 {
		t.Fail()
		return
	}
}

func TestMinInt64s(t *testing.T) {
	if kmath.Min.Int64s(4, 6) != 4 {
		t.Fail()
		return
	}

	if kameria.Math.Min(int64(4), int64(6)) != int64(4) {
		t.Fail()
		return
	}
	if kameria.Math.Min(int64(4), int64(6)) == 4 { // default int, not int64
		t.Fail()
		return
	}
}

func TestMinStrings(t *testing.T) {
	a, b := "AMD", "Intel"
	if a > b {
		t.Fail()
		return
	}

	if kmath.Min.Strings(a, b) != a {
		t.Fail()
		return
	}

	if kameria.Math.Min(a, b) != a {
		t.Fail()
		return
	}
}

func TestMinFloat32s(t *testing.T) {
	a, b := float32(1.8), float32(2.0)
	if a > b {
		t.Fail()
		return
	}

	if kmath.Min.Float32s(a, b) != a {
		t.Fail()
		return
	}

	if kameria.Math.Min(a, b) != a {
		t.Fail()
		return
	}
	if kameria.Math.Min(a, b) == 2.0 { // default float64
		t.Fail()
		return
	}
}

func TestMinFloat64s(t *testing.T) {
	a, b := 2.2, 2.4
	if a > b {
		t.Fail()
		return
	}

	if kmath.Min.Float64s(a, b) != a {
		t.Fail()
		return
	}
	if kameria.Math.Min(a, b) != a {
		t.Fail()
		return
	}
}

func TestMinTimes(t *testing.T) {
	a := time.Now()
	b := a.Add(time.Hour)
	if !kmath.Min.Times(a, b).Equal(a) {
		t.Fail()
		return
	}

	if !kameria.Math.Min(a, b).(time.Time).Equal(a) {
		t.Fail()
		return
	}
}
