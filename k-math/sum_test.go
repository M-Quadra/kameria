package kmath

import (
	"math"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSum(t *testing.T) {
	{ // int
		x, y := []int{1, 2, 3}, 6
		assert.Equal(t, Sum(x), y)

		x = nil
		assert.Equal(t, Sum(x), 0)
	}

	{ // int64
		x, y := []int64{1, 2, 3}, int64(6)
		assert.Equal(t, Sum(x), y)

		x = nil
		assert.Equal(t, Sum(x), int64(0))
	}

	{ // float32
		x, y := []float32{1, 2, 3}, float32(6)
		assert.True(t, math.Abs(float64(Sum(x)-y)) < 1e-8)

		x = nil
		assert.True(t, math.Abs(float64(Sum(x))) < 1e-8)
	}

	{ // float64
		x, y := []float64{1, 2, 3}, float64(6)
		assert.True(t, math.Abs(Sum(x)-y) < 1e-8)

		x = nil
		assert.True(t, math.Abs(Sum(x)) < 1e-8)
	}
}
