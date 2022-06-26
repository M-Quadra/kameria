package kmath

import (
	"math"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSoftmax(t *testing.T) {
	{ // float64
		arr := []float64{1, 2, 3}
		wanna := []float64{0.0900, 0.2447, 0.6652}

		arr = Softmax(arr...)
		assert.Equal(t, len(arr), len(wanna))

		for i, v := range arr {
			assert.LessOrEqual(t, math.Abs(v-wanna[i]), 1e-4)
		}
	}

	{ // float32
		arr := []float32{1, 2, 3}
		wanna := []float32{0.0900, 0.2447, 0.6652}

		arr = Softmax(arr...)
		assert.Equal(t, len(arr), len(wanna))

		for i, v := range arr {
			assert.LessOrEqual(t, math.Abs(float64(v-wanna[i])), 1e-4)
		}
	}
}

func TestLogSoftmax(t *testing.T) {
	{ // float64
		arr := []float64{1, 2, 3}
		wanna := []float64{-2.4076, -1.4076, -0.4076}

		arr = LogSoftmax(arr...)
		assert.Equal(t, len(arr), len(wanna))

		for i, v := range arr {
			assert.LessOrEqual(t, math.Abs(v-wanna[i]), 1e-4)
		}
	}

	{ // float32
		arr := []float32{1, 2, 3}
		wanna := []float32{-2.4076, -1.4076, -0.4076}

		arr = LogSoftmax(arr...)
		assert.Equal(t, len(arr), len(wanna))

		for i, v := range arr {
			assert.LessOrEqual(t, math.Abs(float64(v-wanna[i])), 1e-4)
		}
	}
}
