package main

import (
	"math"
	"testing"

	kmath "github.com/M-Quadra/kameria/k-math"
	"github.com/stretchr/testify/assert"
)

func TestMean(t *testing.T) {
	{ // int
		arr := []int{1, 2, 3, 4, 5}
		wanna := 3.
		assert.True(t, math.Abs(kmath.Mean(arr)-wanna) < 1e-8)

		arr = nil
		assert.True(t, math.Abs(kmath.Mean(arr)) < 1e-8)
	}

	{ // int64
		arr := []int64{1, 2, 3, 4, 5}
		wanna := 3.
		assert.True(t, math.Abs(kmath.Mean(arr)-wanna) < 1e-8)

		arr = nil
		assert.True(t, math.Abs(kmath.Mean(arr)) < 1e-8)
	}

	{ // float32
		arr := []float32{1, 2, 3, 4, 5}
		wanna := 3.
		assert.True(t, math.Abs(kmath.Mean(arr)-wanna) < 1e-8)

		arr = nil
		assert.True(t, math.Abs(kmath.Mean(arr)) < 1e-8)
	}

	{ // float64
		arr := []float64{1, 2, 3, 4, 5}
		wanna := 3.
		assert.True(t, math.Abs(kmath.Mean(arr)-wanna) < 1e-8)

		arr = nil
		assert.True(t, math.Abs(kmath.Mean(arr)) < 1e-8)
	}
}
