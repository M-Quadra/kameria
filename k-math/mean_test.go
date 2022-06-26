package kmath

import (
	"math"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMean(t *testing.T) {
	{ // int
		arr := []int{1, 2, 3, 4, 5}
		wanna := 3.
		assert.LessOrEqual(t, math.Abs(Mean(arr)-wanna), 1e-8)

		arr = nil
		assert.LessOrEqual(t, math.Abs(Mean(arr)), 1e-8)
	}

	{ // int64
		arr := []int64{1, 2, 3, 4, 5}
		wanna := 3.
		assert.LessOrEqual(t, math.Abs(Mean(arr)-wanna), 1e-8)

		arr = nil
		assert.LessOrEqual(t, math.Abs(Mean(arr)), 1e-8)
	}

	{ // float32
		arr := []float32{1, 2, 3, 4, 5}
		wanna := 3.
		assert.LessOrEqual(t, math.Abs(Mean(arr)-wanna), 1e-8)

		arr = nil
		assert.LessOrEqual(t, math.Abs(Mean(arr)), 1e-8)
	}

	{ // float64
		arr := []float64{1, 2, 3, 4, 5}
		wanna := 3.
		assert.LessOrEqual(t, math.Abs(Mean(arr)-wanna), 1e-8)

		arr = nil
		assert.LessOrEqual(t, math.Abs(Mean(arr)), 1e-8)
	}
}
