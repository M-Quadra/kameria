package main

import (
	"testing"

	"github.com/M-Quadra/kameria"
	kmath "github.com/M-Quadra/kameria/k-math"
)

func TestSumInts(t *testing.T) {
	x, y := []int{1, 2, 3}, 6

	if kmath.Sum.Ints(x) != y {
		t.Fail()
		return
	}

	if kameria.Math.Sum(x) != y {
		t.Fail()
		return
	}
}

func TestSumInt64s(t *testing.T) {
	x, y := []int64{1, 2, 3}, int64(6)

	if kmath.Sum.Int64s(x) != y {
		t.Fail()
		return
	}

	if kameria.Math.Sum(x) != y {
		t.Fail()
		return
	}
}

func TestSumFloat32s(t *testing.T) {
	x, y := []float32{1, 2, 3}, float32(6)

	if kmath.Sum.Float32s(x) != y {
		t.Fail()
		return
	}

	if kameria.Math.Sum(x) != y {
		t.Fail()
		return
	}
}

func TestSumFloat64s(t *testing.T) {
	x, y := []float64{1, 2, 3}, float64(6)

	if kmath.Sum.Float64s(x) != y {
		t.Fail()
		return
	}

	if kameria.Math.Sum(x) != y {
		t.Fail()
		return
	}
}
