package main

import (
	"math"
	"testing"

	"github.com/M-Quadra/kameria"
	kmath "github.com/M-Quadra/kameria/k-math"
)

func TestMeanInts(t *testing.T) {
	ary := []int{1, 2, 3, 4, 5}
	wanna := 3.

	if math.Abs(kmath.Mean.Ints(ary)-wanna) > 1e-8 {
		t.Fail()
		return
	}

	if math.Abs(kameria.Math.Mean(ary)-wanna) > 1e-8 {
		t.Fail()
		return
	}
}

func TestMeanInt64s(t *testing.T) {
	ary := []int64{1, 2, 3, 4, 5}
	wanna := 3.

	if math.Abs(kmath.Mean.Int64s(ary)-wanna) > 1e-8 {
		t.Fail()
		return
	}

	if math.Abs(kameria.Math.Mean(ary)-wanna) > 1e-8 {
		t.Fail()
		return
	}
}

func TestMeanFloat32s(t *testing.T) {
	ary := []float32{1, 2, 3, 4, 5}
	wanna := 3.

	if math.Abs(kmath.Mean.Float32s(ary)-wanna) > 1e-8 {
		t.Fail()
		return
	}

	if math.Abs(kameria.Math.Mean(ary)-wanna) > 1e-8 {
		t.Fail()
		return
	}
}

func TestMeanFloat64s(t *testing.T) {
	ary := []float64{1, 2, 3, 4, 5}
	wanna := 3.

	if math.Abs(kmath.Mean.Float64s(ary)-wanna) > 1e-8 {
		t.Fail()
		return
	}

	if math.Abs(kameria.Math.Mean(ary)-wanna) > 1e-8 {
		t.Fail()
		return
	}
}
