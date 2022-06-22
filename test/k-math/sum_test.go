package main

import (
	"math"
	"testing"

	kmath "github.com/M-Quadra/kameria/k-math"
	"github.com/stretchr/testify/assert"
)

func TestSum(t *testing.T) {
	{ // int
		x, y := []int{1, 2, 3}, 6
		assert.Equal(t, kmath.Sum(x), y)

		x = nil
		assert.Equal(t, kmath.Sum(x), 0)
	}

	{ // int64
		x, y := []int64{1, 2, 3}, int64(6)
		assert.Equal(t, kmath.Sum(x), y)

		x = nil
		assert.Equal(t, kmath.Sum(x), int64(0))
	}

	{ // float32
		x, y := []float32{1, 2, 3}, float32(6)
		assert.True(t, math.Abs(float64(kmath.Sum(x)-y)) < 1e-8)

		x = nil
		assert.True(t, math.Abs(float64(kmath.Sum(x))) < 1e-8)
	}

	{ // float64
		x, y := []float64{1, 2, 3}, float64(6)
		assert.True(t, math.Abs(kmath.Sum(x)-y) < 1e-8)

		x = nil
		assert.True(t, math.Abs(kmath.Sum(x)) < 1e-8)
	}
}
